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
    static fm0(p0, globalState) {
      return new BigNumber((((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (_dafny.MultiSet.fromElements(_dafny.Set.fromElements(true))).Elements) {
          let _0_v0 = _compr_0;
          if ((_dafny.MultiSet.fromElements(_dafny.Set.fromElements(true))).contains(_0_v0)) {
            _coll0.push([_0_v0,false]);
          }
        }
        return _coll0;
      }()).Merge(function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of (_dafny.MultiSet.fromElements(_dafny.Set.fromElements(false))).Elements) {
          let _1_v1 = _compr_1;
          if ((_dafny.MultiSet.fromElements(_dafny.Set.fromElements(false))).contains(_1_v1)) {
            _coll1.push([_1_v1,false]);
          }
        }
        return _coll1;
      }())).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(false, !(!(true))),!(false)))).length);
    };
    static fm1(globalState) {
      return (_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(false, false, false, false, !(false))));
    };
    static fm6(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("brkfsxkwu"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(235)), function (_2_i0) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      }));
    };
    static fm7(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.UnicodeFromString("ykyytibe")).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(new BigNumber(749))).length)));
    };
    static fm8(p0, p1, globalState) {
      return ((function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of (_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(972)), _dafny.MultiSet.fromElements(new BigNumber(169)))).Elements) {
          let _3_v0 = _compr_2;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(972)), _dafny.MultiSet.fromElements(new BigNumber(169))), _3_v0)) {
            _coll2.push([_3_v0,_dafny.Seq.UnicodeFromString("d")]);
          }
        }
        return _coll2;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe((_module.D19.create_DC39(_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(971), new BigNumber(664))).cardinality())))).dtor_cf60,_dafny.Seq.UnicodeFromString("w")))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber(-308)),_dafny.Seq.UnicodeFromString("hkclbdj")));
    };
    static fm11(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(783)), _dafny.Set.fromElements(new BigNumber(290))), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(-640))), _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(-318)))));
    };
    static fm17(p0, p1, p2, p3, globalState) {
      return !_dafny.areEqual(_dafny.Seq.of(!(true)), _dafny.Seq.of(!(false)));
    };
    static fm20(p0, p1, globalState) {
      return (_dafny.MultiSet.fromElements(function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of (_dafny.Set.fromElements(new BigNumber(-761))).Elements) {
          let _4_v0 = _compr_3;
          if ((_dafny.Set.fromElements(new BigNumber(-761))).contains(_4_v0)) {
            _coll3.push([(_4_v0).plus(new BigNumber(949)),new BigNumber((_dafny.Seq.UnicodeFromString("tvghivqib")).length)]);
          }
        }
        return _coll3;
      }(), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-361),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(510)), function (_5_i0) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      })).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true, !(false))).length),!(false))).length),new BigNumber(-874)))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(587),new BigNumber((_dafny.Seq.UnicodeFromString("qbojidmhb")).length))), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-55),(_dafny.ZERO).minus(new BigNumber(-384)))))));
    };
    static fm21(globalState) {
      return !(!(!(true) || (!((new BigNumber((_dafny.Seq.of(false, true, !(true))).length)).isLessThan(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of (_dafny.MultiSet.fromElements(_dafny.Seq.of(false, true))).Elements) {
          let _6_v0 = _compr_4;
          if ((_dafny.MultiSet.fromElements(_dafny.Seq.of(false, true))).contains(_6_v0)) {
            _coll4.push([_6_v0,true]);
          }
        }
        return _coll4;
      }()).length),!(true))).length))))));
    };
    static fm22(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(344), new BigNumber(378), new BigNumber((_dafny.Seq.UnicodeFromString("utdhwcb")).length))).cardinality()),true)).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(960),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-518)),false)));
    };
    static fm25(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber(-48))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-685)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(465), new BigNumber((_dafny.Seq.of(!(false))).length))).cardinality()))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber(-478))));
    };
    static fm26(p0, p1, p2, p3, globalState) {
      return (((!(false)) ? ((_module.D28.create_DC59(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('e'.codePointAt(0)),false))).dtor_cf91) : (function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of (_dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0)))).Elements) {
          let _7_v0 = _compr_5;
          if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0))), _7_v0)) {
            _coll5.push([_7_v0,!(false)]);
          }
        }
        return _coll5;
      }()))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),true));
    };
    static fm27(p0, p1, p2, p3, globalState) {
      return !(!(((true) ? (true) : (false))) || ((_dafny.MultiSet.FromArray(_dafny.Seq.of(false, true))).IsProperSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.of(true)))));
    };
    static fm28(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("axni"), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(330)), function (_8_i0) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(293)), function (_9_i1) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      })));
    };
    static fm29(globalState) {
      return _module.D3.create_DC3(!(true));
    };
    static fm30(p0, globalState) {
      return (_dafny.MultiSet.fromElements(new BigNumber(409))).Intersect(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("bm")).length), new BigNumber(-783)));
    };
    static fm31(p0, p1, globalState) {
      return !(_dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(214), new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-799)))).length)), _dafny.Seq.of(new BigNumber(928))), _dafny.Seq.of(new BigNumber(-772))));
    };
    static fm32(p0, p1, p2, p3, globalState) {
      let _source0 = _module.D6.create_DC10();
      if (_source0.is_DC10) {
        return _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(733));
      } else {
        let _10___mcc_h0 = (_source0).cf14;
        let _11_cf14 = _10___mcc_h0;
        return _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.UnicodeFromString("olpebd")).length));
      }
    };
    static fm34(p0, p1, p2, p3, globalState) {
      return new _dafny.CodePoint('c'.codePointAt(0));
    };
    static fm35(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(638)), function (_12_i0) {
        return new BigNumber(((_module.D29.create_DC63(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-409),_dafny.Seq.UnicodeFromString("vyixwxt")))).dtor_cf100).length);
      }), _dafny.Seq.of(new BigNumber((function () {
        let _coll6 = new _dafny.Map();
        for (const _compr_6 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),true)).Keys.Elements) {
          let _13_v0 = _compr_6;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),true)).contains(_13_v0)) {
            _coll6.push([_13_v0,false]);
          }
        }
        return _coll6;
      }()).length), new BigNumber(800), new BigNumber(95), new BigNumber(-392)));
    };
    static fm38(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("oihg")).length)), new BigNumber(-997), new BigNumber((_dafny.Seq.UnicodeFromString("gekldkc")).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(995)), function (_14_i0) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("fjogb")).length);
      }));
    };
    static fm39(globalState) {
      return new _dafny.CodePoint('r'.codePointAt(0));
    };
    static fm40(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,!(false)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(!(false),!(false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,true)));
    };
    static fm41(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.Map.Empty.slice().updateUnsafe(false,false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(false,false)));
    };
    static fm42(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(33)), function (_15_i0) {
        return ((false) ? (new _dafny.CodePoint('u'.codePointAt(0))) : (new _dafny.CodePoint('r'.codePointAt(0))));
      });
    };
    static fm43(p0, globalState) {
      return !((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-928),false)).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
        let _coll7 = new _dafny.Map();
        for (const _compr_7 of _dafny.IntegerRange(new BigNumber(280), new BigNumber(197))) {
          let _16_v0 = _compr_7;
          if (((new BigNumber(280)).isLessThanOrEqualTo(_16_v0)) && ((_16_v0).isLessThan(new BigNumber(197)))) {
            _coll7.push([_module.__default.safeModuloInt(_16_v0, new BigNumber((_dafny.Seq.UnicodeFromString("hcqe")).length)),new _dafny.CodePoint('j'.codePointAt(0))]);
          }
        }
        return _coll7;
      }()).length),false))) || ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('a'.codePointAt(0)))).IsProperSubsetOf(_dafny.MultiSet.fromElements(new _dafny.CodePoint('m'.codePointAt(0)))));
    };
    static fm44(p0, p1, globalState) {
      return (_dafny.Set.fromElements(false, true)).Difference((_dafny.Set.fromElements(true, false, true)).Difference(_dafny.Set.fromElements(false, false)));
    };
    static fm45(p0, globalState) {
      return _dafny.Set.fromElements(new _dafny.CodePoint('e'.codePointAt(0)));
    };
    static fm46(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('u'.codePointAt(0)),new BigNumber((_dafny.Seq.of(false)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),new BigNumber(822)))).Merge((function () {
        let _coll8 = new _dafny.Map();
        for (const _compr_8 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),false)).Keys.Elements) {
          let _17_v0 = _compr_8;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),false)).contains(_17_v0)) {
            _coll8.push([_17_v0,new BigNumber(732)]);
          }
        }
        return _coll8;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))));
    };
    static fm47(p0, globalState) {
      return _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of(true)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(111)),new BigNumber((_dafny.Set.fromElements(!(false))).length))).length));
    };
    static fm48(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(!(true),true),new BigNumber(209))).Merge(function () {
        let _coll9 = new _dafny.Map();
        for (const _compr_9 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(216)), function (_18_i0) {
          return _dafny.Map.Empty.slice().updateUnsafe(false,false);
        })).Elements) {
          let _19_v0 = _compr_9;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(216)), function (_18_i0) {
            return _dafny.Map.Empty.slice().updateUnsafe(false,false);
          }), _19_v0)) {
            _coll9.push([_19_v0,new BigNumber(519)]);
          }
        }
        return _coll9;
      }())).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,true),new BigNumber(-941))).Merge(function () {
        let _coll10 = new _dafny.Map();
        for (const _compr_10 of (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,true))).Elements) {
          let _20_v1 = _compr_10;
          if ((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,true))).contains(_20_v1)) {
            _coll10.push([_20_v1,new BigNumber((_dafny.Seq.UnicodeFromString("fratralci")).length)]);
          }
        }
        return _coll10;
      }()));
    };
    static fm49(p0, p1, p2, p3, globalState) {
      return (function () {
        let _coll11 = new _dafny.Set();
        for (const _compr_11 of _dafny.IntegerRange(new BigNumber(753), new BigNumber(841))) {
          let _21_v0 = _compr_11;
          if (((new BigNumber(753)).isLessThanOrEqualTo(_21_v0)) && ((_21_v0).isLessThan(new BigNumber(841)))) {
            _coll11.add((_21_v0).multipliedBy(new BigNumber((_dafny.Set.fromElements(new BigNumber(-474))).length)));
          }
        }
        return _coll11;
      }()).Difference((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,!(false))).length), new BigNumber((_dafny.Seq.UnicodeFromString("yattoed")).length), new BigNumber(-345), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(445)), function (_22_i0) {
        return new BigNumber(756);
      })).length), new BigNumber(-986))).Union(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("gx")).length))));
    };
    static fm50(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(_dafny.MultiSet.fromElements(true))).Difference(_dafny.Set.fromElements(_dafny.MultiSet.fromElements(!(true)), _dafny.MultiSet.fromElements(true)))).Difference(function () {
        let _coll12 = new _dafny.Set();
        for (const _compr_12 of (_dafny.Seq.of(_dafny.MultiSet.fromElements(false, true, true, false, true), _dafny.MultiSet.fromElements(!(false)))).Elements) {
          let _23_v0 = _compr_12;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.MultiSet.fromElements(false, true, true, false, true), _dafny.MultiSet.fromElements(!(false))), _23_v0)) {
            _coll12.add(_23_v0);
          }
        }
        return _coll12;
      }());
    };
    static fm51(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-37))).cardinality())),((false) ? (false) : (true)));
    };
    static fm52(p0, p1, p2, p3, globalState) {
      return (_dafny.Seq.of(true, false));
    };
    static fm53(p0, globalState) {
      return _module.D3.create_DC3(false);
    };
    static fm54(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(20)), function (_24_i0) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      })).length))), ((false) ? (_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber(994)))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(353)), function (_25_i1) {
        return _dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((_dafny.Set.fromElements(new BigNumber(970))).length));
      }))));
    };
    static fm55(p0, p1, globalState) {
      return _module.D16.create_DC32((_dafny.MultiSet.fromElements(true, true)).Intersect(_dafny.MultiSet.fromElements(true, true, false, true)));
    };
    static fm56(p0, p1, globalState) {
      return _module.D14.create_DC27(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(false,true))),new BigNumber((_dafny.Seq.UnicodeFromString("fpjbrs")).length)));
    };
    static fm57(globalState) {
      return (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,!(!(false))))).Intersect(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(!(!(!(false))),true), _dafny.Map.Empty.slice().updateUnsafe(false,false), _dafny.Map.Empty.slice().updateUnsafe(false,false)));
    };
    static fm58(p0, p1, p2, globalState) {
      if (false) {
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("uua")).length),_module.D4.create_DC7(_dafny.Seq.of(true), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(320),_dafny.Set.fromElements(false))).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(581),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(!(true), false)).length),false)).length),new _dafny.CodePoint('y'.codePointAt(0)))).length))).length))).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-772),_module.D4.create_DC7(_dafny.Seq.of(false, !(false)), new BigNumber(953))));
      } else if (!(false)) {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(504),_module.D4.create_DC7(_dafny.Seq.of(!(true)), new BigNumber((_dafny.Seq.UnicodeFromString("lakodemt")).length)));
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("dmiwuxmt")).length),_module.D4.create_DC7(_dafny.Seq.of(true), new BigNumber(-934)));
      }
    };
    static fm59(globalState) {
      return _module.D4.create_DC7(_dafny.Seq.of(true), new BigNumber((_dafny.Set.fromElements(!(false), !(true), false, false)).length));
    };
    static fm60(p0, p1, p2, globalState) {
      return _module.D11.create_DC20();
    };
    static fm61(p0, p1, globalState) {
      return ((function () {
        let _coll13 = new _dafny.Map();
        for (const _compr_13 of (_dafny.Seq.of(_dafny.Seq.UnicodeFromString("yp"), _dafny.Seq.UnicodeFromString("lt"), _dafny.Seq.UnicodeFromString("kkt"))).Elements) {
          let _26_v0 = _compr_13;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("yp"), _dafny.Seq.UnicodeFromString("lt"), _dafny.Seq.UnicodeFromString("kkt")), _26_v0)) {
            _coll13.push([_26_v0,true]);
          }
        }
        return _coll13;
      }()).Merge(function () {
        let _coll14 = new _dafny.Map();
        for (const _compr_14 of (_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(863)), function (_27_i0) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("mh"))).Elements) {
          let _28_v1 = _compr_14;
          if ((_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(863)), function (_27_i0) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          }), _dafny.Seq.UnicodeFromString("mh"))).contains(_28_v1)) {
            _coll14.push([_28_v1,!(false)]);
          }
        }
        return _coll14;
      }())).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("issb"),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(373)), function (_29_i1) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      }),true)));
    };
    static fm62(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_module.D3.create_DC3(false), _module.D3.create_DC3(false)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(926)), function (_30_i0) {
        return _module.D3.create_DC3(false);
      })), _dafny.Seq.Create(_module.__default.abs(new BigNumber(928)), function (_31_i1) {
        return _module.D3.create_DC3(false);
      }));
    };
    static fm63(globalState) {
      return _module.D19.create_DC39((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(584), new BigNumber(((_module.D28.create_DC60(new BigNumber(40), _dafny.Seq.UnicodeFromString("x"), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(377), new BigNumber(135), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false, true)).cardinality())))).length), new BigNumber((function () {
  let _coll15 = new _dafny.Set();
  for (const _compr_15 of _dafny.IntegerRange(new BigNumber(504), new BigNumber(758))) {
    let _32_v0 = _compr_15;
    if (((new BigNumber(504)).isLessThanOrEqualTo(_32_v0)) && ((_32_v0).isLessThan(new BigNumber(758)))) {
      _coll15.add((_32_v0).multipliedBy(new BigNumber((_dafny.Seq.of(true)).length)));
    }
  }
  return _coll15;
}()).length), new BigNumber((_dafny.Seq.UnicodeFromString("kfsij")).length)))).dtor_cf93).length)))).Difference(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(339)), function (_33_i0) {
  return new _dafny.CodePoint('o'.codePointAt(0));
})).length))));
    };
    static fm64(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('m'.codePointAt(0)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('e'.codePointAt(0))))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('b'.codePointAt(0)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('i'.codePointAt(0)))));
    };
    static fm65(p0, globalState) {
      return _module.D3.create_DC4((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('c'.codePointAt(0)),false)).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('c'.codePointAt(0)),!(!(false)))), new BigNumber((_dafny.Set.fromElements(false, !(!(false)))).length), _module.__default.safeModuloInt(new BigNumber(445), new BigNumber(979)), (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(137),!(true))).length)).isLessThan(new BigNumber((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(true, !(true)), _dafny.MultiSet.fromElements(true, false))).cardinality())), (_dafny.MultiSet.fromElements(false, !(false), !(true), true, false)).IsProperSubsetOf(_dafny.MultiSet.fromElements(!(false), (_module.D16.create_DC33(false, false, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(287)), function (_34_i0) {
  return new _dafny.CodePoint('s'.codePointAt(0));
})).length), new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()))).length), new BigNumber(262))).length), false)).dtor_cf48)));
    };
    static fm66(p0, p1, p2, p3, globalState) {
      return _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(false, !(!(false)))).length)),new BigNumber(-355)));
    };
    static fm67(p0, globalState) {
      return _module.D7.create_DC12();
    };
    static fm68(p0, p1, globalState) {
      return _module.D6.create_DC9(new _dafny.CodePoint('h'.codePointAt(0)));
    };
    static fm69(p0, globalState) {
      return (function () {
        let _coll16 = new _dafny.Map();
        for (const _compr_16 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("vioord"),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(763)), function (_35_i0) {
          return (_module.D6.create_DC9(new _dafny.CodePoint('k'.codePointAt(0)))).dtor_cf14;
        })).length))).Keys.Elements) {
          let _36_v0 = _compr_16;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("vioord"),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(763)), function (_35_i0) {
            return (_module.D6.create_DC9(new _dafny.CodePoint('k'.codePointAt(0)))).dtor_cf14;
          })).length))).contains(_36_v0)) {
            _coll16.push([_36_v0,new BigNumber((_dafny.Seq.UnicodeFromString("pgrs")).length)]);
          }
        }
        return _coll16;
      }()).Merge((_module.D13.create_DC25(function () {
  let _coll17 = new _dafny.Map();
  for (const _compr_17 of (_dafny.Map.Empty.slice().updateUnsafe((_module.D12.create_DC22(_dafny.Seq.UnicodeFromString("oixmk"))).dtor_cf29,new _dafny.CodePoint('m'.codePointAt(0)))).Keys.Elements) {
    let _37_v1 = _compr_17;
    if ((_dafny.Map.Empty.slice().updateUnsafe((_module.D12.create_DC22(_dafny.Seq.UnicodeFromString("oixmk"))).dtor_cf29,new _dafny.CodePoint('m'.codePointAt(0)))).contains(_37_v1)) {
      _coll17.push([_37_v1,new BigNumber((_dafny.Seq.UnicodeFromString("illhgq")).length)]);
    }
  }
  return _coll17;
}())).dtor_cf38);
    };
    static fm70(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(862),new BigNumber((_dafny.Set.fromElements(true)).length));
    };
    static fm71(p0, p1, p2, p3, globalState) {
      return _module.D16.create_DC33(false, false, (_module.D25.create_DC52(new BigNumber(791))).dtor_cf79, !(false));
    };
    static m16(p0, p1, p2, p3, globalState) {
      let r0 = [];
      let r1 = _dafny.Seq.UnicodeFromString("");
      let r2 = _dafny.ZERO;
      let _38_i0;
      _38_i0 = _dafny.ZERO;
      L0: {
        while ((p1).isLessThan(p1)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_38_i0)) {
              break L0;
            }
            _38_i0 = (_38_i0).plus(_dafny.ONE);
            let _39_v0;
            _39_v0 = _dafny.Map.Empty.slice().updateUnsafe(p3,true);
            let _40_v1;
            _40_v1 = _dafny.MultiSet.fromElements(p3, (((_39_v0).contains(p3)) ? ((_39_v0).get(p3)) : (p3)), p3, p3, p3);
            let _41_v2;
            _41_v2 = _dafny.Seq.of(true, true);
            let _42_v3;
            _42_v3 = _dafny.Set.fromElements(!(p3));
            let _43_v4;
            _43_v4 = _dafny.MultiSet.fromElements(p1);
            let _44_v5;
            _44_v5 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm39(globalState),p3);
            let _45_v6;
            let _nw0 = Array((new BigNumber(8)).toNumber());
            _nw0[(_dafny.ZERO).toNumber()] = p1;
            _nw0[(_dafny.ONE).toNumber()] = p1;
            _nw0[(new BigNumber(2)).toNumber()] = new BigNumber((_43_v4).cardinality());
            _nw0[(new BigNumber(3)).toNumber()] = p1;
            _nw0[(new BigNumber(4)).toNumber()] = p1;
            _nw0[(new BigNumber(5)).toNumber()] = p1;
            _nw0[(new BigNumber(6)).toNumber()] = p1;
            _nw0[(new BigNumber(7)).toNumber()] = new BigNumber((_44_v5).length);
            _45_v6 = _nw0;
            let _46_v7;
            _46_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_42_v3).length),_45_v6);
            let _47_v8;
            let _nw1 = Array((new BigNumber(16)).toNumber());
            _nw1[(_dafny.ZERO).toNumber()] = p3;
            _nw1[(_dafny.ONE).toNumber()] = (_40_v1).IsProperSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.of(p3)));
            _nw1[(new BigNumber(2)).toNumber()] = p3;
            _nw1[(new BigNumber(3)).toNumber()] = p3;
            _nw1[(new BigNumber(4)).toNumber()] = !(p3);
            _nw1[(new BigNumber(5)).toNumber()] = (_41_v2)[_module.__default.safeIndex(new BigNumber(-168), new BigNumber((_41_v2).length))];
            _nw1[(new BigNumber(6)).toNumber()] = p3;
            _nw1[(new BigNumber(7)).toNumber()] = p3;
            _nw1[(new BigNumber(8)).toNumber()] = p3;
            _nw1[(new BigNumber(9)).toNumber()] = p3;
            _nw1[(new BigNumber(10)).toNumber()] = p3;
            _nw1[(new BigNumber(11)).toNumber()] = p3;
            _nw1[(new BigNumber(12)).toNumber()] = p3;
            _nw1[(new BigNumber(13)).toNumber()] = !(_46_v7).contains(p1);
            _nw1[(new BigNumber(14)).toNumber()] = !(p3) || (p3);
            _nw1[(new BigNumber(15)).toNumber()] = p3;
            _47_v8 = _nw1;
            let _index0 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_47_v8).length));
            (_47_v8)[_index0] = p3;
            let _index1 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_47_v8).length));
            let _rhs0 = false;
            let _rhs1 = (_dafny.ZERO).minus((_dafny.ZERO).minus(p1));
            let _lhs0 = _47_v8;
            let _lhs1 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_47_v8).length));
            let _lhs2 = globalState;
            _lhs0[_lhs1] = _rhs0;
            _lhs2.f7 = _rhs1;
            let _48_v9;
            _48_v9 = new _dafny.CodePoint('f'.codePointAt(0));
            _48_v9 = _48_v9;
            let _49_v10;
            let _nw2 = Array((new BigNumber(6)).toNumber()).fill([]);
            _49_v10 = _nw2;
            let _50_v11;
            _50_v11 = _dafny.Map.Empty.slice().updateUnsafe(p1,p3);
            let _51_v12;
            _51_v12 = _dafny.Seq.UnicodeFromString("ouvug");
            let _52_v13;
            let _nw3 = new _module.C3();
            _nw3.__ctor(new BigNumber((_50_v11).length), _51_v12, !((_47_v8)[_module.__default.safeIndex(new BigNumber(485), new BigNumber((_47_v8).length))]), p3);
            _52_v13 = _nw3;
            let _53_v14;
            let _nw4 = Array((new BigNumber(7)).toNumber());
            _nw4[(_dafny.ZERO).toNumber()] = _52_v13;
            _nw4[(_dafny.ONE).toNumber()] = _52_v13;
            _nw4[(new BigNumber(2)).toNumber()] = _52_v13;
            _nw4[(new BigNumber(3)).toNumber()] = _52_v13;
            _nw4[(new BigNumber(4)).toNumber()] = _52_v13;
            _nw4[(new BigNumber(5)).toNumber()] = _52_v13;
            _nw4[(new BigNumber(6)).toNumber()] = _52_v13;
            _53_v14 = _nw4;
            let _index2 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_49_v10).length));
            (_49_v10)[_index2] = _53_v14;
            let _index3 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_49_v10).length));
            (_49_v10)[_index3] = _53_v14;
            let _54_v15;
            _54_v15 = _module.D3.create_DC3(false);
            let _55_v16;
            let _nw5 = new _module.C2();
            _nw5.__ctor(_54_v15, (_47_v8)[_module.__default.safeIndex(new BigNumber(485), new BigNumber((_47_v8).length))], p3);
            _55_v16 = _nw5;
            let _56_v17;
            _56_v17 = _dafny.Seq.of(_55_v16);
            let _index4 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_47_v8).length));
            (_47_v8)[_index4] = (((p1).isLessThan(new BigNumber((_56_v17).length))) ? (!(((_module.__default.fm21(globalState)) ? (false) : (!((_47_v8)[_module.__default.safeIndex(new BigNumber(485), new BigNumber((_47_v8).length))]))))) : (false));
          }
        }
      }
      let _57_v18;
      let _nw6 = Array((new BigNumber(16)).toNumber()).fill([]);
      _57_v18 = _nw6;
      _57_v18 = _57_v18;
      let _58_v19;
      let _init0 = function (_59_i1) {
        return false;
      };
      let _nw7 = Array((new BigNumber(23)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw7.length); _i0_0++) {
        _nw7[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _58_v19 = _nw7;
      let _60_v20;
      _60_v20 = _dafny.Seq.of(_58_v19);
      _60_v20 = _dafny.Seq.of(_58_v19, _58_v19, _58_v19, _58_v19, _58_v19);
      let _61_v21;
      let _nw8 = new _module.C13();
      _nw8.__ctor();
      _61_v21 = _nw8;
      let _62_v22;
      _62_v22 = false;
      _62_v22 = _module.__default.fm17(_62_v22, p1, new BigNumber((_dafny.Seq.Concat(p2, _dafny.Seq.of(_module.__default.fm0(!(p3), globalState)))).length), p1, globalState);
      let _63_v23;
      _63_v23 = _dafny.Seq.UnicodeFromString("vfxq");
      let _64_v24;
      let _nw9 = new _module.C3();
      _nw9.__ctor(p1, _63_v23, true, (_module.D3.create_DC3(!(true))).dtor_cf3);
      _64_v24 = _nw9;
      let _65_v25;
      let _nw10 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Set.Empty);
      _65_v25 = _nw10;
      let _66_v26;
      let _nw11 = new _module.C0();
      _nw11.__ctor(_65_v25, p1);
      _66_v26 = _nw11;
      let _67_v27;
      _67_v27 = _dafny.Map.Empty.slice().updateUnsafe(_64_v24,_66_v26);
      let _68_v28;
      let _nw12 = new _module.C12();
      _nw12.__ctor(!(_67_v27).contains(_64_v24), _66_v26.f27);
      _68_v28 = _nw12;
      r0 = _58_v19;
      r1 = _dafny.Seq.update(_63_v23, _module.__default.safeIndex(p1, new BigNumber((_63_v23).length)), p0);
      let _69_v29;
      _69_v29 = _dafny.MultiSet.fromElements(p3);
      let _70_v30;
      _70_v30 = _dafny.Map.Empty.slice().updateUnsafe(p1,_69_v29);
      r2 = new BigNumber((((_69_v29).Union((((_70_v30).contains(p1)) ? ((_70_v30).get(p1)) : (_69_v29)))).Difference(_69_v29)).cardinality());
      return [r0, r1, r2];
    }
    static Main(__noArgsParameter) {
      let _71_v0;
      let _init1 = function (_72_i0) {
        return true;
      };
      let _nw13 = Array((new BigNumber(11)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw13.length); _i0_1++) {
        _nw13[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _71_v0 = _nw13;
      let _73_v1;
      let _init2 = function (_74_i1) {
        return (_74_i1).multipliedBy(new BigNumber(616));
      };
      let _nw14 = Array((new BigNumber(18)).toNumber());
      for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw14.length); _i0_2++) {
        _nw14[_i0_2] = _init2(new BigNumber(_i0_2));
      }
      _73_v1 = _nw14;
      let _75_v2;
      _75_v2 = _dafny.MultiSet.fromElements(_73_v1, _73_v1, _73_v1);
      let _76_globalState;
      let _nw15 = new _module.GlobalState();
      _nw15.__ctor(new BigNumber(150), new BigNumber(-360), new BigNumber(288), new BigNumber(359), true, new BigNumber(773), false, new BigNumber(4), new BigNumber(763), new BigNumber(290), true, _71_v0, _75_v2, new BigNumber(105));
      _76_globalState = _nw15;
      let _77_v3;
      _77_v3 = true;
      let _78_v4;
      let _nw16 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.of());
      _78_v4 = _nw16;
      let _79_v5;
      _79_v5 = new BigNumber(63);
      let _index5 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_78_v4).length));
      (_78_v4)[_index5] = _dafny.Seq.Concat(_dafny.Seq.of(_79_v5), _dafny.Seq.of(_79_v5, _79_v5));
      let _80_v6;
      _80_v6 = _dafny.Map.Empty.slice().updateUnsafe(_79_v5,_77_v3);
      let _81_v7;
      _81_v7 = _dafny.Map.Empty.slice().updateUnsafe(_77_v3,_77_v3);
      let _82_v8;
      _82_v8 = _dafny.MultiSet.fromElements(new BigNumber((_81_v7).length), _79_v5);
      let _83_v9;
      _83_v9 = _dafny.Seq.of((_dafny.ZERO).minus(_79_v5), new BigNumber((_82_v8).cardinality()), (_dafny.ZERO).minus(_79_v5), _79_v5, _79_v5);
      let _84_v10;
      _84_v10 = new _dafny.CodePoint('c'.codePointAt(0));
      let _85_v11;
      _85_v11 = _dafny.Map.Empty.slice().updateUnsafe(_79_v5,_84_v10);
      let _index6 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_78_v4).length));
      let _rhs2 = (((_80_v6).contains((_dafny.ZERO).minus(_79_v5))) ? ((_80_v6).get((_dafny.ZERO).minus(_79_v5))) : (_77_v3));
      let _rhs3 = (_79_v5).minus(_79_v5);
      let _rhs4 = _77_v3;
      let _rhs5 = _dafny.Seq.Concat(_83_v9, _dafny.Seq.update(_dafny.Seq.of(new BigNumber((_85_v11).length)), _module.__default.safeIndex(_79_v5, new BigNumber((_dafny.Seq.of(new BigNumber((_85_v11).length))).length)), _79_v5));
      let _rhs6 = ((_77_v3) ? (!(new BigNumber((_dafny.MultiSet.FromArray(_83_v9)).cardinality())).isEqualTo(_79_v5)) : (_77_v3));
      let _lhs3 = _76_globalState;
      let _lhs4 = _78_v4;
      let _lhs5 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_78_v4).length));
      _77_v3 = _rhs2;
      _lhs3.f7 = _rhs3;
      _77_v3 = _rhs4;
      _lhs4[_lhs5] = _rhs5;
      _77_v3 = _rhs6;
      let _86_v12;
      _86_v12 = _dafny.Seq.of(_77_v3, _77_v3, _77_v3);
      let _87_v13;
      _87_v13 = _dafny.Seq.of(_86_v12);
      _79_v5 = new BigNumber((_dafny.Seq.Concat(_87_v13, _dafny.Seq.of(_86_v12))).length);
      let _88_v14;
      _88_v14 = _dafny.Seq.of(_73_v1);
      let _89_v15;
      _89_v15 = _dafny.Set.fromElements(_77_v3, _77_v3);
      let _90_v16;
      _90_v16 = _dafny.Map.Empty.slice().updateUnsafe(_73_v1,_89_v15);
      let _91_v17;
      _91_v17 = _dafny.MultiSet.fromElements((_90_v16).contains((_88_v14)[_module.__default.safeIndex(_79_v5, new BigNumber((_88_v14).length))]), !(_77_v3), false);
      let _92_v18;
      _92_v18 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_77_v3, _76_globalState),_91_v17);
      let _93_v19;
      _93_v19 = _dafny.Seq.of(_91_v17, _dafny.MultiSet.fromElements(_77_v3, _77_v3), _91_v17);
      _91_v17 = ((((_92_v18).contains(new BigNumber(-359))) ? ((_92_v18).get(new BigNumber(-359))) : ((_93_v19)[_module.__default.safeIndex(_79_v5, new BigNumber((_93_v19).length))]))).Difference((_91_v17).Union(_module.__default.fm1(_76_globalState)));
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_73_v1).length))) {
        let _94_i2 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_94_i2)) && ((_94_i2).isLessThan(new BigNumber((_73_v1).length))))) {
          (_73_v1)[(_94_i2)] = _module.__default.safeDivisionInt(_94_i2, _79_v5);
        }
      }
      if (false) {
        let _95_v20;
        let _nw17 = new _module.C13();
        _nw17.__ctor();
        _95_v20 = _nw17;
        _77_v3 = (_module.__default.fm31(_77_v3, _77_v3, _76_globalState)) && (_77_v3);
        let _96_v21;
        let _nw18 = new _module.C5();
        _nw18.__ctor(_77_v3, _77_v3);
        _96_v21 = _nw18;
        _96_v21 = ((!(_77_v3)) ? (_96_v21) : (_96_v21));
        let _97_v22;
        _97_v22 = _dafny.Map.Empty.slice().updateUnsafe(_79_v5,_79_v5);
        let _98_v23;
        let _nw19 = new _module.C10();
        _nw19.__ctor(_77_v3, _77_v3, new BigNumber((_97_v22).length));
        _98_v23 = _nw19;
        _79_v5 = (_79_v5).minus(_79_v5);
      } else {
        let _index7 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length));
        (_71_v0)[_index7] = !(false);
        let _index8 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length));
        let _rhs7 = _84_v10;
        let _rhs8 = !(!(false) || (_77_v3));
        let _rhs9 = !(true);
        let _lhs6 = _71_v0;
        let _lhs7 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length));
        _84_v10 = _rhs7;
        _77_v3 = _rhs8;
        _lhs6[_lhs7] = _rhs9;
        let _99_v24;
        let _nw20 = Array((new BigNumber(13)).toNumber()).fill(_module.D16.Default());
        _99_v24 = _nw20;
        let _index9 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_99_v24).length));
        (_99_v24)[_index9] = _module.D16.create_DC33(_77_v3, false, _79_v5, _77_v3);
        let _100_v25;
        _100_v25 = _dafny.Map.Empty.slice().updateUnsafe(_79_v5,_79_v5);
        let _index10 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_99_v24).length));
        (_99_v24)[_index10] = _module.__default.fm71(_100_v25, _module.__default.fm52(!(_77_v3), _79_v5, _79_v5, _79_v5, _76_globalState), new BigNumber(741), (_89_v15).Union(_dafny.Set.fromElements(!(_77_v3))), _76_globalState);
        if ((((_81_v7).contains((_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))])) ? ((_81_v7).get((_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))])) : ((_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))]))) {
          let _101_v26;
          let _nw21 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.of());
          _101_v26 = _nw21;
          let _102_v27;
          _102_v27 = _dafny.Seq.UnicodeFromString("mxwae");
          let _103_v28;
          _103_v28 = _dafny.Map.Empty.slice().updateUnsafe(_102_v27,new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(916)), ((_104_v10) => function (_105_i3) {
            return _104_v10;
          })(_84_v10)), _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("lwfis")).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(916)), ((_106_v10) => function (_107_i3) {
            return _106_v10;
          })(_84_v10))).length)), _84_v10)).length));
          let _108_v29;
          _108_v29 = _module.D13.create_DC25(_103_v28);
          let _109_v30;
          _109_v30 = _dafny.Seq.of(_module.D13.create_DC25(_103_v28), _108_v29);
          let _index11 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_101_v26).length));
          (_101_v26)[_index11] = _109_v30;
          let _index12 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_101_v26).length));
          (_101_v26)[_index12] = _109_v30;
          _77_v3 = false;
          let _110_v31;
          let _init3 = function (_111_i4) {
            return true;
          };
          let _nw22 = Array((new BigNumber(9)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw22.length); _i0_3++) {
            _nw22[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _110_v31 = _nw22;
          _110_v31 = _110_v31;
          _103_v28 = (_103_v28).update(_102_v27, _79_v5);
          let _112_v32;
          let _nw23 = Array((new BigNumber(28)).toNumber()).fill([]);
          _112_v32 = _nw23;
          let _113_v33;
          _113_v33 = _module.D25.create_DC51(_112_v32);
          _112_v32 = (((_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))]) ? (_112_v32) : ((_113_v33).dtor_cf78));
        } else {
          let _index13 = _module.__default.safeIndex(new BigNumber(794), new BigNumber((_73_v1).length));
          (_73_v1)[_index13] = _79_v5;
          let _index14 = _module.__default.safeIndex(new BigNumber(794), new BigNumber((_73_v1).length));
          let _rhs10 = (_79_v5).minus((_dafny.ZERO).minus(_79_v5));
          let _rhs11 = (_79_v5).isLessThanOrEqualTo(new BigNumber(931));
          let _rhs12 = _81_v7;
          let _lhs8 = _73_v1;
          let _lhs9 = _module.__default.safeIndex(new BigNumber(794), new BigNumber((_73_v1).length));
          _lhs8[_lhs9] = _rhs10;
          _77_v3 = _rhs11;
          _81_v7 = _rhs12;
          let _114_v34;
          let _nw24 = new _module.C13();
          _nw24.__ctor();
          _114_v34 = _nw24;
          _114_v34 = _114_v34;
          _79_v5 = _module.__default.safeDivisionInt(_79_v5, _79_v5);
          let _115_v35;
          let _nw25 = Array((new BigNumber(29)).toNumber());
          _nw25[(_dafny.ZERO).toNumber()] = _77_v3;
          _nw25[(_dafny.ONE).toNumber()] = !(_77_v3);
          _nw25[(new BigNumber(2)).toNumber()] = (_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))];
          _nw25[(new BigNumber(3)).toNumber()] = _77_v3;
          _nw25[(new BigNumber(4)).toNumber()] = (_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))];
          _nw25[(new BigNumber(5)).toNumber()] = _77_v3;
          _nw25[(new BigNumber(6)).toNumber()] = (_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))];
          _nw25[(new BigNumber(7)).toNumber()] = (_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))];
          _nw25[(new BigNumber(8)).toNumber()] = (_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))];
          _nw25[(new BigNumber(9)).toNumber()] = _77_v3;
          _nw25[(new BigNumber(10)).toNumber()] = false;
          _nw25[(new BigNumber(11)).toNumber()] = _77_v3;
          _nw25[(new BigNumber(12)).toNumber()] = (_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))];
          _nw25[(new BigNumber(13)).toNumber()] = (_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))];
          _nw25[(new BigNumber(14)).toNumber()] = (_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))];
          _nw25[(new BigNumber(15)).toNumber()] = (_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))];
          _nw25[(new BigNumber(16)).toNumber()] = _module.__default.fm31(_77_v3, _77_v3, _76_globalState);
          _nw25[(new BigNumber(17)).toNumber()] = _module.__default.fm43((_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))], _76_globalState);
          _nw25[(new BigNumber(18)).toNumber()] = _77_v3;
          _nw25[(new BigNumber(19)).toNumber()] = !((_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))]);
          _nw25[(new BigNumber(20)).toNumber()] = true;
          _nw25[(new BigNumber(21)).toNumber()] = _77_v3;
          _nw25[(new BigNumber(22)).toNumber()] = (_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))];
          _nw25[(new BigNumber(23)).toNumber()] = _77_v3;
          _nw25[(new BigNumber(24)).toNumber()] = (_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))];
          _nw25[(new BigNumber(25)).toNumber()] = (_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))];
          _nw25[(new BigNumber(26)).toNumber()] = (_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))];
          _nw25[(new BigNumber(27)).toNumber()] = _77_v3;
          _nw25[(new BigNumber(28)).toNumber()] = _77_v3;
          _115_v35 = _nw25;
          let _116_v36;
          _116_v36 = _dafny.Seq.UnicodeFromString("s");
          let _117_v37;
          _117_v37 = _dafny.Map.Empty.slice().updateUnsafe(_115_v35,_116_v36);
          let _118_v38;
          _118_v38 = _module.D12.create_DC22(_116_v36);
          _117_v37 = (_117_v37).update(_115_v35, _dafny.Seq.Concat((_118_v38).dtor_cf29, _dafny.Seq.Create(_module.__default.abs(new BigNumber(15)), ((_119_v10) => function (_120_i5) {
            return _119_v10;
          })(_84_v10))));
          let _121_v39;
          let _nw26 = Array((new BigNumber(6)).toNumber()).fill([]);
          _121_v39 = _nw26;
          let _122_v40;
          let _nw27 = Array((new BigNumber(3)).toNumber());
          _122_v40 = _nw27;
          let _index15 = _module.__default.safeIndex(new BigNumber(503), new BigNumber((_121_v39).length));
          (_121_v39)[_index15] = _122_v40;
          let _index16 = _module.__default.safeIndex(new BigNumber(503), new BigNumber((_121_v39).length));
          (_121_v39)[_index16] = _122_v40;
        }
        let _123_v41;
        _123_v41 = _dafny.Seq.UnicodeFromString("wfsifht");
        if (!(_dafny.Seq.IsPrefixOf(_dafny.Seq.update((((_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))]) ? (_123_v41) : (_123_v41)), _module.__default.safeIndex(_79_v5, new BigNumber(((((_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))]) ? (_123_v41) : (_123_v41))).length)), _84_v10), _123_v41))) {
          let _124_v42;
          let _nw28 = new _module.C13();
          _nw28.__ctor();
          _124_v42 = _nw28;
          let _125_v43;
          _125_v43 = _dafny.MultiSet.fromElements(_87_v13, _87_v13, _87_v13);
          let _126_v44;
          _126_v44 = _dafny.Seq.of(_77_v3, _77_v3);
          let _127_v45;
          _127_v45 = _module.D4.create_DC7(_126_v44, _79_v5);
          (_76_globalState).f2 = (((_125_v43).contains(_dafny.Seq.of(_86_v12))) ? ((_125_v43).get(_dafny.Seq.of(_86_v12))) : ((_dafny.ZERO).minus(((_127_v45).dtor_cf12).multipliedBy(_79_v5))));
          _80_v6 = (_80_v6).update(_module.__default.fm0((_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))], _76_globalState), _dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(602)), ((_128_v10) => function (_129_i6) {
            return _128_v10;
          })(_84_v10)), _123_v41));
          _80_v6 = (_80_v6).update((_79_v5).minus(new BigNumber((_83_v9).length)), _77_v3);
          let _130_v46;
          _130_v46 = _dafny.Set.fromElements(_79_v5, _79_v5, new BigNumber(442));
          let _131_v47;
          let _nw29 = Array((new BigNumber(14)).toNumber()).fill(false);
          _131_v47 = _nw29;
          let _132_v48;
          _132_v48 = _module.D9.create_DC15((_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))], (_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))], _130_v46, _73_v1, _131_v47);
          let _pat_let_tv0 = _73_v1;
          let _pat_let_tv1 = _131_v47;
          let _pat_let_tv2 = _131_v47;
          let _pat_let_tv3 = _73_v1;
          let _133_v50;
          let _nw30 = Array((new BigNumber(14)).toNumber());
          _nw30[(_dafny.ZERO).toNumber()] = _132_v48;
          _nw30[(_dafny.ONE).toNumber()] = _132_v48;
          _nw30[(new BigNumber(2)).toNumber()] = function (_pat_let0_0) {
            return function (_134_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_135_dt__update_hcf21_h0) {
                  return function (_pat_let2_0) {
                    return function (_136_dt__update_hcf22_h0) {
                      return _module.D9.create_DC15((_134_dt__update__tmp_h0).dtor_cf18, (_134_dt__update__tmp_h0).dtor_cf19, (_134_dt__update__tmp_h0).dtor_cf20, _135_dt__update_hcf21_h0, _136_dt__update_hcf22_h0);
                    }(_pat_let2_0);
                  }(_pat_let_tv1);
                }(_pat_let1_0);
              }(_pat_let_tv0);
            }(_pat_let0_0);
          }(_module.D9.create_DC15(_77_v3, (_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))], _130_v46, _73_v1, _131_v47));
          _nw30[(new BigNumber(3)).toNumber()] = _module.D9.create_DC15((_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))], _module.__default.fm43(false, _76_globalState), function () {
  let _coll18 = new _dafny.Set();
  for (const _compr_18 of (_100_v25).Keys.Elements) {
    let _137_v49 = _compr_18;
    if ((_100_v25).contains(_137_v49)) {
      _coll18.add(_module.__default.safeModuloInt(_137_v49, new BigNumber((_dafny.Set.fromElements(true, true, false, false, !(true))).length)));
    }
  }
  return _coll18;
}(), _73_v1, _131_v47);
          _nw30[(new BigNumber(4)).toNumber()] = _132_v48;
          _nw30[(new BigNumber(5)).toNumber()] = _132_v48;
          _nw30[(new BigNumber(6)).toNumber()] = _132_v48;
          _nw30[(new BigNumber(7)).toNumber()] = _132_v48;
          _nw30[(new BigNumber(8)).toNumber()] = _132_v48;
          _nw30[(new BigNumber(9)).toNumber()] = _132_v48;
          _nw30[(new BigNumber(10)).toNumber()] = function (_pat_let3_0) {
            return function (_138_dt__update__tmp_h1) {
              return function (_pat_let4_0) {
                return function (_139_dt__update_hcf22_h1) {
                  return function (_pat_let5_0) {
                    return function (_140_dt__update_hcf19_h0) {
                      return function (_pat_let6_0) {
                        return function (_141_dt__update_hcf21_h1) {
                          return _module.D9.create_DC15((_138_dt__update__tmp_h1).dtor_cf18, _140_dt__update_hcf19_h0, (_138_dt__update__tmp_h1).dtor_cf20, _141_dt__update_hcf21_h1, _139_dt__update_hcf22_h1);
                        }(_pat_let6_0);
                      }(_pat_let_tv3);
                    }(_pat_let5_0);
                  }(false);
                }(_pat_let4_0);
              }(_pat_let_tv2);
            }(_pat_let3_0);
          }(_132_v48);
          _nw30[(new BigNumber(11)).toNumber()] = _132_v48;
          _nw30[(new BigNumber(12)).toNumber()] = _132_v48;
          _nw30[(new BigNumber(13)).toNumber()] = _132_v48;
          _133_v50 = _nw30;
          let _142_v51;
          _142_v51 = _module.D18.create_DC37(_133_v50);
          _142_v51 = _142_v51;
        } else {
          _123_v41 = _dafny.Seq.Concat(_dafny.Seq.Concat(_123_v41, _123_v41), _123_v41);
          let _143_v52;
          let _nw31 = Array((new BigNumber(25)).toNumber());
          _143_v52 = _nw31;
          let _144_v53;
          let _nw32 = Array((new BigNumber(26)).toNumber());
          _nw32[(_dafny.ZERO).toNumber()] = _dafny.Set.fromElements(_77_v3);
          _nw32[(_dafny.ONE).toNumber()] = _89_v15;
          _nw32[(new BigNumber(2)).toNumber()] = _dafny.Set.fromElements(_77_v3);
          _nw32[(new BigNumber(3)).toNumber()] = _89_v15;
          _nw32[(new BigNumber(4)).toNumber()] = _dafny.Set.fromElements((_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))]);
          _nw32[(new BigNumber(5)).toNumber()] = _89_v15;
          _nw32[(new BigNumber(6)).toNumber()] = _89_v15;
          _nw32[(new BigNumber(7)).toNumber()] = _89_v15;
          _nw32[(new BigNumber(8)).toNumber()] = _dafny.Set.fromElements(_77_v3, false, _77_v3, false, !(_77_v3));
          _nw32[(new BigNumber(9)).toNumber()] = _89_v15;
          _nw32[(new BigNumber(10)).toNumber()] = _89_v15;
          _nw32[(new BigNumber(11)).toNumber()] = _89_v15;
          _nw32[(new BigNumber(12)).toNumber()] = _89_v15;
          _nw32[(new BigNumber(13)).toNumber()] = _89_v15;
          _nw32[(new BigNumber(14)).toNumber()] = _dafny.Set.fromElements(_77_v3);
          _nw32[(new BigNumber(15)).toNumber()] = _89_v15;
          _nw32[(new BigNumber(16)).toNumber()] = _89_v15;
          _nw32[(new BigNumber(17)).toNumber()] = _dafny.Set.fromElements(_77_v3);
          _nw32[(new BigNumber(18)).toNumber()] = _89_v15;
          _nw32[(new BigNumber(19)).toNumber()] = _89_v15;
          _nw32[(new BigNumber(20)).toNumber()] = _89_v15;
          _nw32[(new BigNumber(21)).toNumber()] = _89_v15;
          _nw32[(new BigNumber(22)).toNumber()] = _module.__default.fm44(_77_v3, false, _76_globalState);
          _nw32[(new BigNumber(23)).toNumber()] = _89_v15;
          _nw32[(new BigNumber(24)).toNumber()] = _89_v15;
          _nw32[(new BigNumber(25)).toNumber()] = _89_v15;
          _144_v53 = _nw32;
          let _145_v54;
          let _nw33 = new _module.C0();
          _nw33.__ctor(_144_v53, new BigNumber(((_78_v4)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_78_v4).length))]).length));
          _145_v54 = _nw33;
          let _index17 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_143_v52).length));
          (_143_v52)[_index17] = _145_v54;
          let _index18 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_143_v52).length));
          (_143_v52)[_index18] = _145_v54;
          let _146_v55;
          let _nw34 = Array((new BigNumber(24)).toNumber()).fill(false);
          _146_v55 = _nw34;
          let _147_v56;
          _147_v56 = _module.D11.create_DC19(_100_v25);
          let _148_v57;
          _148_v57 = _dafny.Map.Empty.slice().updateUnsafe(_146_v55,_147_v56);
          _148_v57 = (_148_v57).update(_146_v55, _147_v56);
          let _149_v59;
          _149_v59 = _dafny.Set.fromElements(new BigNumber((function () {
            let _coll19 = new _dafny.Set();
            for (const _compr_19 of _dafny.IntegerRange(new BigNumber(137), new BigNumber(987))) {
              let _150_v58 = _compr_19;
              if (((new BigNumber(137)).isLessThanOrEqualTo(_150_v58)) && ((_150_v58).isLessThan(new BigNumber(987)))) {
                _coll19.add((_150_v58).minus(_145_v54.f27));
              }
            }
            return _coll19;
          }()).length));
          let _151_v60;
          let _nw35 = new _module.C2();
          _nw35.__ctor(_module.D3.create_DC3(_77_v3), (_149_v59).IsDisjointFrom(_149_v59), _77_v3);
          _151_v60 = _nw35;
          let _index19 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length));
          let _index20 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length));
          let _rhs13 = _77_v3;
          let _rhs14 = _77_v3;
          let _rhs15 = !(_79_v5).isEqualTo(_79_v5);
          let _rhs16 = _77_v3;
          let _rhs17 = _151_v60;
          let _lhs10 = _71_v0;
          let _lhs11 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length));
          let _lhs12 = _71_v0;
          let _lhs13 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length));
          _77_v3 = _rhs13;
          _lhs10[_lhs11] = _rhs14;
          _77_v3 = _rhs15;
          _lhs12[_lhs13] = _rhs16;
          _151_v60 = _rhs17;
          let _152_v61;
          _152_v61 = _dafny.Map.Empty.slice().updateUnsafe(_77_v3,((_dafny.ZERO).minus(new BigNumber((_83_v9).length))).minus(new BigNumber(697)));
          let _153_v62;
          _153_v62 = _module.D12.create_DC22(_123_v41);
          let _154_v63;
          _154_v63 = _dafny.Seq.of(_153_v62, _module.D12.create_DC22(_123_v41), _153_v62);
          _79_v5 = (((_152_v61).contains(!(_152_v61).contains((_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))]))) ? ((_152_v61).get(!(_152_v61).contains((_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))]))) : (new BigNumber((((_77_v3) ? (_154_v63) : (_154_v63))).length)));
        }
        if (_77_v3) {
          let _155_v64;
          let _nw36 = new _module.C9();
          _nw36.__ctor();
          _155_v64 = _nw36;
          let _156_v65;
          _156_v65 = _dafny.Set.fromElements(_155_v64, _155_v64, _155_v64, _155_v64, _155_v64);
          let _157_v66;
          let _nw37 = new _module.C6();
          _nw37.__ctor(_79_v5, _module.__default.fm43((_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))], _76_globalState), _77_v3, (_dafny.Set.fromElements(_155_v64, _155_v64)).IsProperSubsetOf(_156_v65), ((_dafny.ZERO).minus((_dafny.ZERO).minus(_79_v5))).multipliedBy(new BigNumber((_83_v9).length)));
          _157_v66 = _nw37;
          let _158_v67;
          let _nw38 = new _module.C10();
          _nw38.__ctor((_157_v66).f20, !(((_157_v66).f20) && (true)), (_157_v66).f19);
          _158_v67 = _nw38;
          let _159_v68;
          let _nw39 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Set.Empty);
          _159_v68 = _nw39;
          let _160_v69;
          let _nw40 = new _module.C0();
          _nw40.__ctor(_159_v68, _79_v5);
          _160_v69 = _nw40;
          let _161_v70;
          _161_v70 = _dafny.Map.Empty.slice().updateUnsafe(_79_v5,_160_v69);
          let _162_v71;
          _162_v71 = _dafny.Set.fromElements(_79_v5, (_157_v66).f19, new BigNumber((_123_v41).length), new BigNumber((_123_v41).length), new BigNumber((_161_v70).length));
          let _163_v72;
          _163_v72 = _dafny.Map.Empty.slice().updateUnsafe((_157_v66).f20,_162_v71);
          let _164_v73;
          let _nw41 = new _module.C11();
          _nw41.__ctor(_160_v69.f27, (_157_v66).f20, (_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))]);
          _164_v73 = _nw41;
          let _index21 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length));
          let _rhs18 = _163_v72;
          let _rhs19 = (_157_v66).f20;
          let _rhs20 = _79_v5;
          let _rhs21 = _164_v73;
          let _rhs22 = (_81_v7).Merge(_81_v7);
          let _lhs14 = _71_v0;
          let _lhs15 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length));
          _163_v72 = _rhs18;
          _lhs14[_lhs15] = _rhs19;
          _79_v5 = _rhs20;
          _164_v73 = _rhs21;
          _81_v7 = _rhs22;
          _77_v3 = _dafny.Seq.IsProperPrefixOf(_123_v41, (_158_v67).fm9(_79_v5, _dafny.Seq.of(_162_v71, _162_v71), new _dafny.CodePoint('s'.codePointAt(0)), _160_v69.f27, _76_globalState));
          let _165_v74;
          _165_v74 = _dafny.Map.Empty.slice().updateUnsafe((_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))],new BigNumber(24));
          _165_v74 = (_165_v74).update(((_71_v0)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length))]) && ((_157_v66).f20), _module.__default.safeDivisionInt(_160_v69.f27, new BigNumber((_dafny.Set.fromElements(_160_v69.f27)).length)));
        } else {
          let _166_v75;
          let _nw42 = new _module.C13();
          _nw42.__ctor();
          _166_v75 = _nw42;
          let _167_v76;
          _167_v76 = _dafny.MultiSet.fromElements(_166_v75);
          let _168_v77;
          _168_v77 = _dafny.Seq.of(_166_v75);
          _77_v3 = (_167_v76).IsDisjointFrom(_dafny.MultiSet.FromArray(_168_v77));
          let _169_v78;
          _169_v78 = _dafny.Map.Empty.slice().updateUnsafe(_79_v5,_123_v41);
          _77_v3 = !_dafny.areEqual(_123_v41, _dafny.Seq.Concat((((_169_v78).contains(new BigNumber(964))) ? ((_169_v78).get(new BigNumber(964))) : (_123_v41)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(68)), ((_170_v10) => function (_171_i7) {
            return _170_v10;
          })(_84_v10))));
          (_76_globalState).f2 = _79_v5;
          let _index22 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_71_v0).length));
          (_71_v0)[_index22] = !(_module.__default.safeModuloInt(new BigNumber(-188), _79_v5)).isEqualTo(((_77_v3) ? (_79_v5) : (_79_v5)));
          let _172_v79;
          let _nw43 = new _module.C13();
          _nw43.__ctor();
          _172_v79 = _nw43;
        }
      }
      let _init4 = ((_173_v5) => function (_174_i8) {
        return (_174_i8).plus(_173_v5);
      })(_79_v5);
      let _nw44 = Array((new BigNumber(5)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw44.length); _i0_4++) {
        _nw44[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _73_v1 = _nw44;
      let _source1 = _module.D25.create_DC53();
      if (_source1.is_DC52) {
        let _175___mcc_h0 = (_source1).cf79;
        let _176_cf79 = _175___mcc_h0;
        _77_v3 = _77_v3;
        _77_v3 = _77_v3;
        _77_v3 = ((_dafny.ZERO).minus((((_82_v8).contains(_176_cf79)) ? ((_82_v8).get(_176_cf79)) : (_79_v5)))).isLessThan(_176_cf79);
        let _177_v80;
        _177_v80 = _dafny.Seq.UnicodeFromString("t");
        let _178_v81;
        _178_v81 = _dafny.Map.Empty.slice().updateUnsafe(_77_v3,(_dafny.ZERO).minus(new BigNumber((_177_v80).length)));
        let _179_v82;
        _179_v82 = _dafny.MultiSet.fromElements(_module.__default.fm38(_77_v3, _176_cf79, _176_cf79, _76_globalState), _dafny.Seq.of(new BigNumber((_178_v81).length), _79_v5));
        _80_v6 = (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.update(_177_v80, _module.__default.safeIndex(new BigNumber((_179_v82).cardinality()), new BigNumber((_177_v80).length)), _84_v10)).length),_77_v3)).Merge(_80_v6);
      } else if (_source1.is_DC53) {
        let _180_v83;
        let _nw45 = new _module.C7();
        _nw45.__ctor(_77_v3, ((_77_v3) ? (_77_v3) : (_77_v3)), new BigNumber(997));
        _180_v83 = _nw45;
        _180_v83 = _180_v83;
        (_180_v83).f16 = !(true) || (_77_v3);
        (_180_v83).f16 = _module.__default.fm17(_180_v83.f16, _module.__default.safeModuloInt(_79_v5, _79_v5), (_79_v5).multipliedBy(new BigNumber(159)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_79_v5,_79_v5)).length), _76_globalState);
        if (!(_77_v3)) {
          let _181_v84;
          _181_v84 = _module.D21.create_DC44(_79_v5, !(!((_180_v83).f17)));
          let _182_v85;
          _182_v85 = _dafny.Seq.UnicodeFromString("ewiw");
          let _183_v86;
          let _nw46 = new _module.C6();
          _nw46.__ctor(_79_v5, _77_v3, (_180_v83).f17, (_180_v83).f17, (_dafny.ZERO).minus(_79_v5));
          _183_v86 = _nw46;
          let _184_v87;
          _184_v87 = _dafny.Set.fromElements(_183_v86);
          let _185_v88;
          _185_v88 = _dafny.Map.Empty.slice().updateUnsafe((_180_v83).f17,new BigNumber((_184_v87).length));
          let _186_v89;
          let _187_v90;
          let _188_v91;
          let _189_v92;
          let _out0;
          let _out1;
          let _out2;
          let _out3;
          let _outcollector0 = (_180_v83).m6((_180_v83).f17, (_181_v84).dtor_cf70, _module.__default.safeDivisionInt(new BigNumber(776), _79_v5), _module.__default.fm34(new BigNumber((_182_v85).length), new BigNumber((_185_v88).length), (_180_v83).f17, new BigNumber(-499), _76_globalState), _76_globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _out2 = _outcollector0[2];
          _out3 = _outcollector0[3];
          _186_v89 = _out0;
          _187_v90 = _out1;
          _188_v91 = _out2;
          _189_v92 = _out3;
          let _190_v93;
          _190_v93 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_187_v90));
          let _191_v94;
          let _nw47 = new _module.C11();
          _nw47.__ctor(_186_v89, (_190_v93).equals(_190_v93), !(_189_v92));
          _191_v94 = _nw47;
          (_76_globalState).f2 = new BigNumber(158);
          let _192_v95;
          _192_v95 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(339)), function (_193_i9) {
            return new _dafny.CodePoint('m'.codePointAt(0));
          })).length),new BigNumber((_86_v12).length));
          let _194_v101;
          _194_v101 = _module.D11.create_DC19(_192_v95);
          let _195_v102;
          let _nw48 = Array((new BigNumber(22)).toNumber());
          _nw48[(_dafny.ZERO).toNumber()] = ((_180_v83.f16) ? (_192_v95) : (_192_v95));
          _nw48[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_91_v17).cardinality()),_183_v86.f14);
          _nw48[(new BigNumber(2)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_79_v5,new BigNumber((_192_v95).length))).Merge(_192_v95);
          _nw48[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_183_v86.f14,new BigNumber((_dafny.Seq.UnicodeFromString("fjhe")).length));
          _nw48[(new BigNumber(4)).toNumber()] = (function () {
            let _coll20 = new _dafny.Map();
            for (const _compr_20 of _dafny.IntegerRange(new BigNumber(156), new BigNumber(112))) {
              let _196_v96 = _compr_20;
              if (((new BigNumber(156)).isLessThanOrEqualTo(_196_v96)) && ((_196_v96).isLessThan(new BigNumber(112)))) {
                _coll20.push([(_196_v96).plus(_186_v89),(_dafny.ZERO).minus(_186_v89)]);
              }
            }
            return _coll20;
          }()).Merge(function () {
            let _coll21 = new _dafny.Map();
            for (const _compr_21 of ((_78_v4)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_78_v4).length))]).Elements) {
              let _197_v97 = _compr_21;
              if (_dafny.Seq.contains((_78_v4)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_78_v4).length))], _197_v97)) {
                _coll21.push([_module.__default.safeDivisionInt(_197_v97, _183_v86.f14),new BigNumber((_dafny.Seq.of(_180_v83.f16, _77_v3, (_180_v83).f17)).length)]);
              }
            }
            return _coll21;
          }());
          _nw48[(new BigNumber(5)).toNumber()] = (_192_v95).Merge(_192_v95);
          _nw48[(new BigNumber(6)).toNumber()] = function () {
            let _coll22 = new _dafny.Map();
            for (const _compr_22 of (_83_v9).Elements) {
              let _198_v98 = _compr_22;
              if (_dafny.Seq.contains(_83_v9, _198_v98)) {
                _coll22.push([(_198_v98).minus(_186_v89),_183_v86.f14]);
              }
            }
            return _coll22;
          }();
          _nw48[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_183_v86.f14,_186_v89);
          _nw48[(new BigNumber(8)).toNumber()] = (_192_v95).Merge(function () {
            let _coll23 = new _dafny.Map();
            for (const _compr_23 of ((_78_v4)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_78_v4).length))]).Elements) {
              let _199_v99 = _compr_23;
              if (_dafny.Seq.contains((_78_v4)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_78_v4).length))], _199_v99)) {
                _coll23.push([_module.__default.safeModuloInt(_199_v99, _183_v86.f14),_186_v89]);
              }
            }
            return _coll23;
          }());
          _nw48[(new BigNumber(9)).toNumber()] = (_192_v95).Merge(_192_v95);
          _nw48[(new BigNumber(10)).toNumber()] = _192_v95;
          _nw48[(new BigNumber(11)).toNumber()] = (_192_v95).Merge(_192_v95);
          _nw48[(new BigNumber(12)).toNumber()] = _192_v95;
          _nw48[(new BigNumber(13)).toNumber()] = function () {
            let _coll24 = new _dafny.Map();
            for (const _compr_24 of _dafny.IntegerRange(new BigNumber(986), new BigNumber(286))) {
              let _200_v100 = _compr_24;
              if (((new BigNumber(986)).isLessThanOrEqualTo(_200_v100)) && ((_200_v100).isLessThan(new BigNumber(286)))) {
                _coll24.push([(_200_v100).multipliedBy(new BigNumber((_192_v95).length)),_183_v86.f14]);
              }
            }
            return _coll24;
          }();
          _nw48[(new BigNumber(14)).toNumber()] = _192_v95;
          _nw48[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_186_v89,_186_v89);
          _nw48[(new BigNumber(16)).toNumber()] = (_192_v95).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_83_v9).length),_79_v5));
          _nw48[(new BigNumber(17)).toNumber()] = (_194_v101).dtor_cf27;
          _nw48[(new BigNumber(18)).toNumber()] = _192_v95;
          _nw48[(new BigNumber(19)).toNumber()] = (_180_v83).fm3(_76_globalState);
          _nw48[(new BigNumber(20)).toNumber()] = _192_v95;
          _nw48[(new BigNumber(21)).toNumber()] = _192_v95;
          _195_v102 = _nw48;
          let _201_v104;
          _201_v104 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_182_v85).length)));
          let _202_v106;
          _202_v106 = _dafny.Map.Empty.slice().updateUnsafe((_180_v83).f17,(_194_v101).dtor_cf27);
          let _203_v107;
          let _nw49 = Array((new BigNumber(15)).toNumber()).fill(_module.D9.Default());
          _203_v107 = _nw49;
          let _204_v108;
          _204_v108 = _dafny.Seq.of(_203_v107, _203_v107);
          let _205_v109;
          _205_v109 = _module.D12.create_DC24(_189_v92, _86_v12, _204_v108, (_180_v83).f17, _dafny.Seq.Create(_module.__default.abs(new BigNumber(852)), ((_206_v83) => function (_207_i10) {
  return _module.D3.create_DC3(_206_v83.f16);
})(_180_v83)));
          let _nw50 = Array((new BigNumber(26)).toNumber());
          _nw50[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_79_v5,_186_v89);
          _nw50[(_dafny.ONE).toNumber()] = _192_v95;
          _nw50[(new BigNumber(2)).toNumber()] = _192_v95;
          _nw50[(new BigNumber(3)).toNumber()] = _192_v95;
          _nw50[(new BigNumber(4)).toNumber()] = _192_v95;
          _nw50[(new BigNumber(5)).toNumber()] = _192_v95;
          _nw50[(new BigNumber(6)).toNumber()] = _192_v95;
          _nw50[(new BigNumber(7)).toNumber()] = function () {
            let _coll25 = new _dafny.Map();
            for (const _compr_25 of (_201_v104).Elements) {
              let _208_v103 = _compr_25;
              if ((_201_v104).contains(_208_v103)) {
                _coll25.push([(_208_v103).minus(_186_v89),(_dafny.ZERO).minus((_dafny.ZERO).minus(_183_v86.f14))]);
              }
            }
            return _coll25;
          }();
          _nw50[(new BigNumber(8)).toNumber()] = (_192_v95).Merge(_192_v95);
          _nw50[(new BigNumber(9)).toNumber()] = _192_v95;
          _nw50[(new BigNumber(10)).toNumber()] = _192_v95;
          _nw50[(new BigNumber(11)).toNumber()] = function () {
            let _coll26 = new _dafny.Map();
            for (const _compr_26 of _dafny.IntegerRange(new BigNumber(157), new BigNumber(-96))) {
              let _209_v105 = _compr_26;
              if (((new BigNumber(157)).isLessThanOrEqualTo(_209_v105)) && ((_209_v105).isLessThan(new BigNumber(-96)))) {
                _coll26.push([_module.__default.safeDivisionInt(_209_v105, _183_v86.f14),_183_v86.f14]);
              }
            }
            return _coll26;
          }();
          _nw50[(new BigNumber(12)).toNumber()] = _192_v95;
          _nw50[(new BigNumber(13)).toNumber()] = (_192_v95).Merge(_dafny.Map.Empty.slice().updateUnsafe(_186_v89,_186_v89));
          _nw50[(new BigNumber(14)).toNumber()] = (_192_v95).Merge(_192_v95);
          _nw50[(new BigNumber(15)).toNumber()] = _192_v95;
          _nw50[(new BigNumber(16)).toNumber()] = _192_v95;
          _nw50[(new BigNumber(17)).toNumber()] = _192_v95;
          _nw50[(new BigNumber(18)).toNumber()] = _192_v95;
          _nw50[(new BigNumber(19)).toNumber()] = _192_v95;
          _nw50[(new BigNumber(20)).toNumber()] = (_192_v95).Merge((((_202_v106).contains((_205_v109).dtor_cf36)) ? ((_202_v106).get((_205_v109).dtor_cf36)) : (_192_v95)));
          _nw50[(new BigNumber(21)).toNumber()] = (_192_v95).Merge(_192_v95);
          _nw50[(new BigNumber(22)).toNumber()] = _192_v95;
          _nw50[(new BigNumber(23)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_79_v5,new BigNumber((_dafny.Seq.UnicodeFromString("rceafxpqm")).length))).Merge(_192_v95);
          _nw50[(new BigNumber(24)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_183_v86.f14,_183_v86.f14);
          _nw50[(new BigNumber(25)).toNumber()] = _192_v95;
          _195_v102 = _nw50;
          let _index23 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_73_v1).length));
          (_73_v1)[_index23] = _79_v5;
          let _210_v110;
          _210_v110 = _dafny.Map.Empty.slice().updateUnsafe(false,_79_v5);
          let _index24 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_73_v1).length));
          let _rhs23 = (_dafny.ZERO).minus((_180_v83).fm12(new BigNumber(727), _76_globalState));
          let _rhs24 = _module.__default.fm21(_76_globalState);
          let _rhs25 = _183_v86;
          let _rhs26 = (((_192_v95).contains(_183_v86.f14)) ? ((_192_v95).get(_183_v86.f14)) : ((_dafny.ZERO).minus(_79_v5)));
          let _rhs27 = _210_v110;
          let _lhs16 = _76_globalState;
          let _lhs17 = _73_v1;
          let _lhs18 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_73_v1).length));
          _lhs16.f2 = _rhs23;
          _77_v3 = _rhs24;
          _183_v86 = _rhs25;
          _lhs17[_lhs18] = _rhs26;
          _210_v110 = _rhs27;
        } else {
          let _211_v111;
          _211_v111 = _dafny.Map.Empty.slice().updateUnsafe(_180_v83.f16,new BigNumber(-58));
          let _212_v112;
          _212_v112 = _dafny.Seq.UnicodeFromString("b");
          let _213_v113;
          let _nw51 = new _module.C3();
          _nw51.__ctor(new BigNumber(352), _212_v112, _180_v83.f16, _180_v83.f16);
          _213_v113 = _nw51;
          let _214_v114;
          _214_v114 = _dafny.Map.Empty.slice().updateUnsafe(_211_v111,_213_v113);
          let _215_v115;
          _215_v115 = _dafny.Seq.of((((_214_v114).contains(_211_v111)) ? ((_214_v114).get(_211_v111)) : (_213_v113)), _213_v113);
          (_76_globalState).f7 = new BigNumber((_215_v115).length);
          let _216_v116;
          let _nw52 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _216_v116 = _nw52;
          let _217_v117;
          let _nw53 = new _module.C4();
          _nw53.__ctor(true, (_180_v83).f17, _180_v83.f16);
          _217_v117 = _nw53;
          let _218_v118;
          _218_v118 = _dafny.Map.Empty.slice().updateUnsafe(_217_v117,_216_v116);
          _216_v116 = (((_218_v118).contains(_217_v117)) ? ((_218_v118).get(_217_v117)) : (_216_v116));
          let _rhs28 = _dafny.Seq.UnicodeFromString("lkc");
          let _rhs29 = (_86_v12)[_module.__default.safeIndex(_module.__default.safeDivisionInt((_213_v113).f22, (_213_v113).f22), new BigNumber((_86_v12).length))];
          let _rhs30 = _180_v83.f16;
          let _lhs19 = _180_v83;
          _212_v112 = _rhs28;
          _77_v3 = _rhs29;
          _lhs19.f16 = _rhs30;
          let _219_v119;
          _219_v119 = _module.D21.create_DC44(_79_v5, _180_v83.f16);
          let _220_v120;
          _220_v120 = _dafny.Map.Empty.slice().updateUnsafe((_213_v113).f22,_219_v119);
          let _221_v121;
          let _nw54 = new _module.C7();
          _nw54.__ctor(true, (((_78_v4)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_78_v4).length))])[_module.__default.safeIndex(_79_v5, new BigNumber(((_78_v4)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_78_v4).length))]).length))]).isLessThanOrEqualTo(new BigNumber((_220_v120).length)), (_213_v113).f22);
          _221_v121 = _nw54;
          let _index25 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_73_v1).length));
          (_73_v1)[_index25] = ((_dafny.ZERO).minus(_79_v5)).multipliedBy(new BigNumber((_83_v9).length));
          let _index26 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_73_v1).length));
          (_73_v1)[_index26] = (_dafny.ZERO).minus((_213_v113).f22);
        }
      } else if (_source1.is_DC51) {
        let _222___mcc_h1 = (_source1).cf78;
        let _223_cf78 = _222___mcc_h1;
        let _224_v122;
        _224_v122 = _dafny.Seq.UnicodeFromString("it");
        let _225_v123;
        let _nw55 = new _module.C8();
        _nw55.__ctor(_224_v122, _77_v3, _77_v3, new BigNumber(569));
        _225_v123 = _nw55;
        let _226_v124;
        _226_v124 = _dafny.Seq.of(_225_v123, _225_v123, _225_v123, _225_v123);
        let _227_v125;
        _227_v125 = _dafny.Set.fromElements(_225_v123, _225_v123, _225_v123, _225_v123, (_226_v124)[_module.__default.safeIndex(_79_v5, new BigNumber((_226_v124).length))]);
        let _228_v126;
        _228_v126 = _dafny.Set.fromElements(new BigNumber((_81_v7).length));
        let _229_v127;
        _229_v127 = _module.D9.create_DC15((_227_v125).contains(_225_v123), _77_v3, _228_v126, _73_v1, _71_v0);
        _229_v127 = _229_v127;
        (_76_globalState).f2 = _79_v5;
        let _230_v128;
        _230_v128 = _module.D19.create_DC39(_dafny.MultiSet.fromElements(_79_v5));
        let _231_v129;
        _231_v129 = _dafny.Map.Empty.slice().updateUnsafe(_230_v128,_77_v3);
        let _232_v130;
        _232_v130 = _module.D26.create_DC55(_231_v129);
        let _233_v131;
        let _nw56 = Array((new BigNumber(15)).toNumber());
        _nw56[(_dafny.ZERO).toNumber()] = _231_v129;
        _nw56[(_dafny.ONE).toNumber()] = (_232_v130).dtor_cf81;
        _nw56[(new BigNumber(2)).toNumber()] = _231_v129;
        _nw56[(new BigNumber(3)).toNumber()] = _231_v129;
        _nw56[(new BigNumber(4)).toNumber()] = _231_v129;
        _nw56[(new BigNumber(5)).toNumber()] = (_231_v129).Merge(_231_v129);
        _nw56[(new BigNumber(6)).toNumber()] = _231_v129;
        _nw56[(new BigNumber(7)).toNumber()] = (_231_v129).update(_230_v128, _77_v3);
        _nw56[(new BigNumber(8)).toNumber()] = (_231_v129).Merge(_dafny.Map.Empty.slice().updateUnsafe(_230_v128,_77_v3));
        _nw56[(new BigNumber(9)).toNumber()] = _231_v129;
        _nw56[(new BigNumber(10)).toNumber()] = _231_v129;
        _nw56[(new BigNumber(11)).toNumber()] = _231_v129;
        _nw56[(new BigNumber(12)).toNumber()] = _231_v129;
        _nw56[(new BigNumber(13)).toNumber()] = _231_v129;
        _nw56[(new BigNumber(14)).toNumber()] = _231_v129;
        _233_v131 = _nw56;
        let _234_v133;
        _234_v133 = _dafny.Seq.of(function () {
          let _coll27 = new _dafny.Map();
          for (const _compr_27 of (_dafny.MultiSet.fromElements(_230_v128)).Elements) {
            let _235_v132 = _compr_27;
            if ((_dafny.MultiSet.fromElements(_230_v128)).contains(_235_v132)) {
              _coll27.push([_235_v132,_77_v3]);
            }
          }
          return _coll27;
        }());
        let _index27 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_233_v131).length));
        (_233_v131)[_index27] = ((_234_v133)[_module.__default.safeIndex((_225_v123).fm12(_module.__default.fm0(_77_v3, _76_globalState), _76_globalState), new BigNumber((_234_v133).length))]).Merge(_231_v129);
        let _236_v134;
        let _nw57 = new _module.C5();
        _nw57.__ctor(((_77_v3) ? (!(_77_v3)) : (_77_v3)), (_77_v3) || (_77_v3));
        _236_v134 = _nw57;
        let _index28 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_233_v131).length));
        let _rhs31 = _231_v129;
        let _rhs32 = !(_79_v5).isEqualTo((_79_v5).plus(new BigNumber((_224_v122).length)));
        let _rhs33 = _236_v134;
        let _lhs20 = _233_v131;
        let _lhs21 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_233_v131).length));
        _lhs20[_lhs21] = _rhs31;
        _77_v3 = _rhs32;
        _236_v134 = _rhs33;
        let _rhs34 = _77_v3;
        let _rhs35 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_225_v123).fm12(_79_v5, _76_globalState), ((false) ? (new BigNumber((_dafny.Seq.of(_77_v3)).length)) : (_79_v5))));
        let _rhs36 = _77_v3;
        let _lhs22 = _76_globalState;
        _77_v3 = _rhs34;
        _lhs22.f7 = _rhs35;
        _77_v3 = _rhs36;
      } else {
        let _237___mcc_h2 = (_source1).cf80;
        let _238_cf80 = _237___mcc_h2;
        _77_v3 = ((_77_v3) ? (_77_v3) : (_77_v3));
        if (_77_v3) {
          _77_v3 = _77_v3;
          let _239_v135;
          _239_v135 = _dafny.Map.Empty.slice().updateUnsafe(!(_77_v3),_79_v5);
          let _index29 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_73_v1).length));
          (_73_v1)[_index29] = _module.__default.safeDivisionInt(_79_v5, (((_239_v135).contains(_77_v3)) ? ((_239_v135).get(_77_v3)) : (_79_v5)));
          let _index30 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_73_v1).length));
          (_73_v1)[_index30] = _79_v5;
          (_76_globalState).f2 = ((_73_v1)[_module.__default.safeIndex(new BigNumber(355), new BigNumber((_73_v1).length))]).minus(_module.__default.safeModuloInt(_79_v5, (_73_v1)[_module.__default.safeIndex(new BigNumber(355), new BigNumber((_73_v1).length))]));
          let _240_v136;
          let _init5 = ((_241_v7) => function (_242_i11) {
            return (_242_i11).plus(new BigNumber((_241_v7).length));
          })(_81_v7);
          let _nw58 = Array((new BigNumber(21)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw58.length); _i0_5++) {
            _nw58[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _240_v136 = _nw58;
          let _243_v137;
          _243_v137 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_79_v5),_240_v136);
          let _244_v138;
          _244_v138 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(_module.__default.fm0(_77_v3, _76_globalState), new BigNumber((_83_v9).length)),(((_243_v137).contains((_73_v1)[_module.__default.safeIndex(new BigNumber(355), new BigNumber((_73_v1).length))])) ? ((_243_v137).get((_73_v1)[_module.__default.safeIndex(new BigNumber(355), new BigNumber((_73_v1).length))])) : (_240_v136)));
          let _245_v139;
          let _nw59 = new _module.C6();
          _nw59.__ctor(new BigNumber(478), !(_77_v3), _77_v3, _module.__default.fm21(_76_globalState), _79_v5);
          _245_v139 = _nw59;
          let _246_v140;
          _246_v140 = _dafny.Map.Empty.slice().updateUnsafe(_245_v139,_244_v138);
          _243_v137 = (((_246_v140).contains(_245_v139)) ? ((_246_v140).get(_245_v139)) : (_243_v137));
          let _247_v141;
          let _nw60 = new _module.C13();
          _nw60.__ctor();
          _247_v141 = _nw60;
        } else {
          let _248_v142;
          _248_v142 = _dafny.Set.fromElements(_79_v5, _79_v5);
          let _249_v143;
          _249_v143 = _dafny.Map.Empty.slice().updateUnsafe(_77_v3,_248_v142);
          let _250_v144;
          let _251_v145;
          let _252_v146;
          let _out4;
          let _out5;
          let _out6;
          let _outcollector1 = _module.__default.m16(_84_v10, _module.__default.safeModuloInt(_79_v5, _79_v5), _83_v9, (new BigNumber((_dafny.Seq.update(_86_v12, _module.__default.safeIndex(_79_v5, new BigNumber((_86_v12).length)), _77_v3)).length)).isLessThan(new BigNumber((_249_v143).length)), _76_globalState);
          _out4 = _outcollector1[0];
          _out5 = _outcollector1[1];
          _out6 = _outcollector1[2];
          _250_v144 = _out4;
          _251_v145 = _out5;
          _252_v146 = _out6;
          _77_v3 = _77_v3;
          _77_v3 = _77_v3;
          _77_v3 = true;
          let _253_v147;
          _253_v147 = _dafny.Map.Empty.slice().updateUnsafe(!(_77_v3),_79_v5);
          let _254_v148;
          _254_v148 = _dafny.Map.Empty.slice().updateUnsafe(_253_v147,_module.__default.safeModuloInt(_252_v146, _252_v146));
          _254_v148 = (_254_v148).update(_253_v147, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(244)), ((_255_v3, _256_v17) => function (_257_i12) {
            return (((_256_v17).contains(_255_v3)) ? ((_256_v17).get(_255_v3)) : (new BigNumber(-741)));
          })(_77_v3, _91_v17))).length));
        }
        _77_v3 = _77_v3;
        let _258_v149;
        let _nw61 = Array((new BigNumber(17)).toNumber());
        _258_v149 = _nw61;
        let _259_v150;
        let _nw62 = new _module.C12();
        _nw62.__ctor(_77_v3, new BigNumber(827));
        _259_v150 = _nw62;
        let _index31 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_258_v149).length));
        (_258_v149)[_index31] = _259_v150;
        let _260_v151;
        _260_v151 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_79_v5),_259_v150);
        let _index32 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_258_v149).length));
        (_258_v149)[_index32] = (((_260_v151).contains(_module.__default.safeModuloInt(_79_v5, _79_v5))) ? ((_260_v151).get(_module.__default.safeModuloInt(_79_v5, _79_v5))) : (_259_v150));
      }
      let _index33 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_73_v1).length));
      (_73_v1)[_index33] = ((_77_v3) ? (_module.__default.fm0(_77_v3, _76_globalState)) : (new BigNumber(17)));
      let _index34 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_73_v1).length));
      (_73_v1)[_index34] = _79_v5;
      let _261_v152;
      _261_v152 = _module.D16.create_DC32(_91_v17);
      let _hi0 = _79_v5;
      for (let _262_i13 = new BigNumber(((_261_v152).dtor_cf46).cardinality()); _262_i13.isLessThan(_hi0); _262_i13 = _262_i13.plus(_dafny.ONE)) {
        let _263_v153;
        _263_v153 = _dafny.Seq.UnicodeFromString("rbx");
        let _264_v154;
        let _nw63 = new _module.C3();
        _nw63.__ctor(_262_i13, _263_v153, _77_v3, _77_v3);
        _264_v154 = _nw63;
        let _265_v155;
        _265_v155 = _dafny.Map.Empty.slice().updateUnsafe(_264_v154,_262_i13);
        let _266_v156;
        let _nw64 = new _module.C7();
        _nw64.__ctor(_77_v3, _77_v3, new BigNumber(932));
        _266_v156 = _nw64;
        let _267_v157;
        _267_v157 = _dafny.Set.fromElements(_266_v156);
        let _268_v158;
        _268_v158 = _dafny.Map.Empty.slice().updateUnsafe(((_73_v1)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_73_v1).length))]).plus((((_265_v155).contains(_264_v154)) ? ((_265_v155).get(_264_v154)) : (_79_v5))),(_267_v157).Difference(_267_v157));
        let _index35 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_73_v1).length));
        let _rhs37 = (_dafny.ZERO).minus(_79_v5);
        let _rhs38 = new BigNumber((_268_v158).length);
        let _lhs23 = _73_v1;
        let _lhs24 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_73_v1).length));
        let _lhs25 = _76_globalState;
        _lhs23[_lhs24] = _rhs37;
        _lhs25.f7 = _rhs38;
        let _269_v159;
        _269_v159 = _dafny.Set.fromElements(new BigNumber(-72));
        let _270_v160;
        _270_v160 = _module.D9.create_DC15(!(_77_v3), _77_v3, _269_v159, _73_v1, _71_v0);
        let _271_v161;
        _271_v161 = _dafny.Seq.of(_270_v160, _270_v160);
        let _272_v162;
        _272_v162 = _dafny.Map.Empty.slice().updateUnsafe(((_77_v3) ? (_77_v3) : (_77_v3)),_271_v161);
        _272_v162 = (_272_v162).update(_77_v3, _dafny.Seq.update(_dafny.Seq.Concat(_271_v161, _271_v161), _module.__default.safeIndex((_264_v154).f22, new BigNumber((_dafny.Seq.Concat(_271_v161, _271_v161)).length)), _270_v160));
        _71_v0 = _71_v0;
        let _273_v163;
        let _nw65 = new _module.C13();
        _nw65.__ctor();
        _273_v163 = _nw65;
      }
      let _274_v164;
      _274_v164 = _module.D15.create_DC30(_84_v10);
      _274_v164 = _274_v164;
      let _275_v165;
      _275_v165 = _module.D27.create_DC57(_dafny.Seq.of(_80_v6));
      let _index36 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_73_v1).length));
      (_73_v1)[_index36] = new BigNumber(((_275_v165).dtor_cf87).length);
      let _276_v167;
      _276_v167 = _dafny.Seq.of(function () {
        let _coll28 = new _dafny.Map();
        for (const _compr_28 of (_dafny.MultiSet.FromArray(_83_v9)).Elements) {
          let _277_v166 = _compr_28;
          if ((_dafny.MultiSet.FromArray(_83_v9)).contains(_277_v166)) {
            _coll28.push([(_277_v166).minus(_79_v5),_77_v3]);
          }
        }
        return _coll28;
      }());
      (_76_globalState).f2 = new BigNumber((_276_v167).length);
      let _index37 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_71_v0).length));
      (_71_v0)[_index37] = (_77_v3) || (_77_v3);
      let _index38 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_71_v0).length));
      (_71_v0)[_index38] = _77_v3;
      let _hi1 = (_dafny.ZERO).minus((_73_v1)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_73_v1).length))]);
      for (let _278_i14 = _module.__default.fm0(!((_71_v0)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_71_v0).length))]), _76_globalState); _278_i14.isLessThan(_hi1); _278_i14 = _278_i14.plus(_dafny.ONE)) {
        let _279_v168;
        _279_v168 = _dafny.Seq.UnicodeFromString("x");
        let _280_v169;
        let _281_v170;
        let _282_v171;
        let _out7;
        let _out8;
        let _out9;
        let _outcollector2 = _module.__default.m16(_84_v10, new BigNumber(361), _83_v9, _dafny.Seq.IsPrefixOf(_279_v168, _279_v168), _76_globalState);
        _out7 = _outcollector2[0];
        _out8 = _outcollector2[1];
        _out9 = _outcollector2[2];
        _280_v169 = _out7;
        _281_v170 = _out8;
        _282_v171 = _out9;
        (_76_globalState).f2 = _module.__default.safeModuloInt(_278_i14, _79_v5);
        let _283_v172;
        _283_v172 = _dafny.Map.Empty.slice().updateUnsafe((_73_v1)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_73_v1).length))],_79_v5);
        let _284_v173;
        let _nw66 = Array((new BigNumber(14)).toNumber()).fill(_module.D9.Default());
        _284_v173 = _nw66;
        let _285_v174;
        _285_v174 = _module.D18.create_DC37(_284_v173);
        let _source2 = ((!(_283_v172).equals(_283_v172)) ? (((_77_v3) ? (_module.D18.create_DC37(_284_v173)) : (_285_v174))) : (_285_v174));
        if (_source2.is_DC38) {
          let _286___mcc_h3 = (_source2).cf57;
          let _287___mcc_h4 = (_source2).cf58;
          let _288___mcc_h5 = (_source2).cf59;
          let _289_cf59 = _288___mcc_h5;
          let _290_cf58 = _287___mcc_h4;
          let _291_cf57 = _286___mcc_h3;
          let _292_v175;
          let _nw67 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.of());
          _292_v175 = _nw67;
          let _293_v176;
          _293_v176 = _dafny.Set.fromElements(new BigNumber(-789));
          let _rhs39 = _292_v175;
          let _rhs40 = (((_71_v0)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_71_v0).length))]) ? ((_293_v176).IsDisjointFrom(_293_v176)) : ((((_80_v6).contains(_291_cf57)) ? ((_80_v6).get(_291_cf57)) : (true))));
          _292_v175 = _rhs39;
          _77_v3 = _rhs40;
          let _294_v177;
          _294_v177 = _dafny.Map.Empty.slice().updateUnsafe((_71_v0)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_71_v0).length))],_73_v1);
          _294_v177 = (_294_v177).update((_71_v0)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_71_v0).length))], _73_v1);
          let _295_v178;
          let _nw68 = Array((new BigNumber(19)).toNumber()).fill([]);
          _295_v178 = _nw68;
          let _index39 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_295_v178).length));
          (_295_v178)[_index39] = _73_v1;
          let _index40 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_295_v178).length));
          let _rhs41 = _module.__default.fm0((!(_module.__default.fm43(_290_cf58, _76_globalState))) && (_77_v3), _76_globalState);
          let _rhs42 = _73_v1;
          let _lhs26 = _295_v178;
          let _lhs27 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_295_v178).length));
          _79_v5 = _rhs41;
          _lhs26[_lhs27] = _rhs42;
          (_76_globalState).f7 = _module.__default.fm0((_71_v0)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_71_v0).length))], _76_globalState);
        } else {
          let _296___mcc_h6 = (_source2).cf56;
          let _297_cf56 = _296___mcc_h6;
          let _298_v179;
          let _nw69 = new _module.C6();
          _nw69.__ctor(_282_v171, _77_v3, (_71_v0)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_71_v0).length))], (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_77_v3,new BigNumber(574))).length)).isLessThanOrEqualTo(new BigNumber(992)), (_73_v1)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_73_v1).length))]);
          _298_v179 = _nw69;
          _91_v17 = _91_v17;
          let _299_v180;
          let _nw70 = new _module.C13();
          _nw70.__ctor();
          _299_v180 = _nw70;
          let _300_v181;
          let _nw71 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
          _300_v181 = _nw71;
          let _301_v182;
          _301_v182 = _dafny.Map.Empty.slice().updateUnsafe(_91_v17,_77_v3);
          let _index41 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_300_v181).length));
          (_300_v181)[_index41] = _301_v182;
          let _302_v183;
          _302_v183 = _module.D16.create_DC33((_298_v179).f20, _77_v3, (_73_v1)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_73_v1).length))], _77_v3);
          let _303_v184;
          _303_v184 = _dafny.Seq.of(_302_v183);
          let _304_v185;
          let _nw72 = new _module.C9();
          _nw72.__ctor();
          _304_v185 = _nw72;
          let _305_v186;
          _305_v186 = _dafny.Map.Empty.slice().updateUnsafe(_79_v5,_304_v185);
          let _306_v187;
          _306_v187 = _dafny.Set.fromElements((((_305_v186).contains((_73_v1)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_73_v1).length))])) ? ((_305_v186).get((_73_v1)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_73_v1).length))])) : (_304_v185)));
          let _307_v188;
          _307_v188 = _dafny.Seq.of(_303_v184, _303_v184, _dafny.Seq.update(_303_v184, _module.__default.safeIndex(new BigNumber(17), new BigNumber((_303_v184).length)), _302_v183));
          let _pat_let_tv4 = _282_v171;
          let _index42 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_300_v181).length));
          let _rhs43 = _module.__default.safeModuloInt(((_73_v1)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_73_v1).length))]).plus(_282_v171), new BigNumber(-573));
          let _rhs44 = _299_v180;
          let _rhs45 = _301_v182;
          let _rhs46 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(function (_pat_let7_0) {
            return function (_308_dt__update__tmp_h2) {
              return function (_pat_let8_0) {
                return function (_309_dt__update_hcf49_h0) {
                  return _module.D16.create_DC33((_308_dt__update__tmp_h2).dtor_cf47, (_308_dt__update__tmp_h2).dtor_cf48, _309_dt__update_hcf49_h0, (_308_dt__update__tmp_h2).dtor_cf50);
                }(_pat_let8_0);
              }(_pat_let_tv4);
            }(_pat_let7_0);
          }(_302_v183)), _303_v184), (_307_v188)[_module.__default.safeIndex(new BigNumber((_279_v168).length), new BigNumber((_307_v188).length))]);
          let _rhs47 = _dafny.Set.fromElements(_304_v185, _304_v185, _304_v185);
          let _lhs28 = _76_globalState;
          let _lhs29 = _300_v181;
          let _lhs30 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_300_v181).length));
          _lhs28.f2 = _rhs43;
          _299_v180 = _rhs44;
          _lhs29[_lhs30] = _rhs45;
          _303_v184 = _rhs46;
          _306_v187 = _rhs47;
          let _310_v189;
          let _311_v190;
          let _out10;
          let _out11;
          let _outcollector3 = (_304_v185).m5(_76_globalState);
          _out10 = _outcollector3[0];
          _out11 = _outcollector3[1];
          _310_v189 = _out10;
          _311_v190 = _out11;
        }
        let _index43 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_71_v0).length));
        (_71_v0)[_index43] = !((_86_v12)[_module.__default.safeIndex(_278_i14, new BigNumber((_86_v12).length))]) || ((_71_v0)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_71_v0).length))]);
      }
      let _312_v191;
      let _init6 = ((_313_v15) => function (_314_i16) {
        return _313_v15;
      })(_89_v15);
      let _nw73 = Array((new BigNumber(13)).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw73.length); _i0_6++) {
        _nw73[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _312_v191 = _nw73;
      let _ingredients0 = [];
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_312_v191).length))) {
        let _315_i15 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_315_i15)) && ((_315_i15).isLessThan(new BigNumber((_312_v191).length))))) {
          _ingredients0.push(_dafny.Tuple.of(_312_v191, (_315_i15).toNumber(), _dafny.Set.fromElements(_77_v3, (_dafny.Set.fromElements(_79_v5, (_73_v1)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_73_v1).length))])).IsSubsetOf(_dafny.Set.fromElements((_73_v1)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_73_v1).length))])))));
        }
      }
      for (const _tup0 of _ingredients0) {
        _tup0[0][_tup0[1]] = _tup0[2];
      }
      (_76_globalState).f2 = new BigNumber(115);
      process.stdout.write(_dafny.toString((_71_v0)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v0)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v0)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v0)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v0)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v0)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v0)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v0)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v0)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v0)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v0)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_73_v1)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_73_v1)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_73_v1)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_73_v1)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_73_v1)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_75_v2).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_76_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_76_globalState.f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_76_globalState).f11)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_76_globalState).f11)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_76_globalState).f11)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_76_globalState).f11)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_76_globalState).f11)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_76_globalState).f11)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_76_globalState).f11)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_76_globalState).f11)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_76_globalState).f11)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_76_globalState).f11)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_76_globalState).f11)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_76_globalState).f12).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_77_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_78_v4)[_dafny.ZERO], _dafny.Seq.of(new BigNumber(-63), new BigNumber(2), new BigNumber(-63), new BigNumber(63), new BigNumber(63), new BigNumber(63)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_79_v5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_80_v6).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(63),true).updateUnsafe(new BigNumber(3),false).updateUnsafe(new BigNumber(-4),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_81_v7).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_82_v8).equals(_dafny.MultiSet.fromElements(_dafny.ONE, new BigNumber(63)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_83_v9, _dafny.Seq.of(new BigNumber(-63), new BigNumber(2), new BigNumber(-63), new BigNumber(63), new BigNumber(63)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_84_v10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v11).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(63),new _dafny.CodePoint('c'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_86_v12, _dafny.Seq.of(true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_87_v13, _dafny.Seq.of(_dafny.Seq.of(true, true, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_88_v14).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_89_v15).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_90_v16).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_v17).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_92_v18).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(3),_dafny.MultiSet.fromElements(true, false, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_93_v19, _dafny.Seq.of(_dafny.MultiSet.fromElements(true, false, false), _dafny.MultiSet.fromElements(true, true), _dafny.MultiSet.fromElements(true, false, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v152).dtor_cf46).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_274_v164).dtor_cf44));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_275_v165).dtor_cf87, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(63),true).updateUnsafe(new BigNumber(3),false).updateUnsafe(new BigNumber(-4),true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_276_v167, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-64),true).updateUnsafe(_dafny.ONE,true).updateUnsafe(new BigNumber(62),true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_312_v191)[_dafny.ZERO]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_312_v191)[_dafny.ONE]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_312_v191)[new BigNumber(2)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_312_v191)[new BigNumber(3)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_312_v191)[new BigNumber(4)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_312_v191)[new BigNumber(5)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_312_v191)[new BigNumber(6)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_312_v191)[new BigNumber(7)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_312_v191)[new BigNumber(8)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_312_v191)[new BigNumber(9)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_312_v191)[new BigNumber(10)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_312_v191)[new BigNumber(11)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_312_v191)[new BigNumber(12)]).equals(_dafny.Set.fromElements(true))));
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
        return other.$tag === 0 && _dafny.areEqual(this.cf0, other.cf0);
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
          return D0.Default();
        }
      };
    }
  }

  $module.D1 = class D1 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1) {
      let $dt = new D1(0);
      $dt.cf1 = cf1;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf1 === other.cf1;
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
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC2(cf2) {
      let $dt = new D2(0);
      $dt.cf2 = cf2;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get dtor_cf2() { return this.cf2; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC2" + "(" + _dafny.toString(this.cf2) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf2, other.cf2);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.ZERO;
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
    static create_DC4(cf4, cf5, cf6, cf7, cf8) {
      let $dt = new D3(0);
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC3(cf3) {
      let $dt = new D3(1);
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC5(cf9) {
      let $dt = new D3(2);
      $dt.cf9 = cf9;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf9() { return this.cf9; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC4" + "(" + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC3" + "(" + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC5" + "(" + _dafny.toString(this.cf9) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf4 === other.cf4 && _dafny.areEqual(this.cf5, other.cf5) && _dafny.areEqual(this.cf6, other.cf6) && this.cf7 === other.cf7 && this.cf8 === other.cf8;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf3 === other.cf3;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf9, other.cf9);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC4(false, _dafny.ZERO, _dafny.ZERO, false, false);
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
    static create_DC7(cf11, cf12) {
      let $dt = new D4(0);
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC6(cf10) {
      let $dt = new D4(1);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC7" + "(" + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC6" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf11, other.cf11) && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC7(_dafny.Seq.of(), _dafny.ZERO);
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
    static create_DC8(cf13) {
      let $dt = new D5(0);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC8" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf13, other.cf13);
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
    static create_DC10() {
      let $dt = new D6(0);
      return $dt;
    }
    static create_DC9(cf14) {
      let $dt = new D6(1);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC10";
      } else if (this.$tag === 1) {
        return "D6.DC9" + "(" + _dafny.toString(this.cf14) + ")";
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
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC10();
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
    static create_DC12() {
      let $dt = new D7(0);
      return $dt;
    }
    static create_DC11(cf15) {
      let $dt = new D7(1);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC12";
      } else if (this.$tag === 1) {
        return "D7.DC11" + "(" + _dafny.toString(this.cf15) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf15, other.cf15);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC12();
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
    static create_DC13(cf16) {
      let $dt = new D8(0);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC13" + "(" + _dafny.toString(this.cf16) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf16, other.cf16);
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
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC15(cf18, cf19, cf20, cf21, cf22) {
      let $dt = new D9(0);
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC16(cf23, cf24) {
      let $dt = new D9(1);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC14(cf17) {
      let $dt = new D9(2);
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC17(cf25) {
      let $dt = new D9(3);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get is_DC14() { return this.$tag === 2; }
    get is_DC17() { return this.$tag === 3; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC15" + "(" + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC16" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC14" + "(" + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 3) {
        return "D9.DC17" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf18 === other.cf18 && this.cf19 === other.cf19 && _dafny.areEqual(this.cf20, other.cf20) && this.cf21 === other.cf21 && this.cf22 === other.cf22;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf23, other.cf23) && this.cf24 === other.cf24;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf17 === other.cf17;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf25, other.cf25);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC15(false, false, _dafny.Set.Empty, [], []);
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
    static create_DC18(cf26) {
      let $dt = new D10(0);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC18" + "(" + _dafny.toString(this.cf26) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf26, other.cf26);
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
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC20() {
      let $dt = new D11(0);
      return $dt;
    }
    static create_DC19(cf27) {
      let $dt = new D11(1);
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC21(cf28) {
      let $dt = new D11(2);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC20";
      } else if (this.$tag === 1) {
        return "D11.DC19" + "(" + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC21" + "(" + _dafny.toString(this.cf28) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf27, other.cf27);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf28, other.cf28);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC20();
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
    static create_DC23(cf30, cf31, cf32) {
      let $dt = new D12(0);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC24(cf33, cf34, cf35, cf36, cf37) {
      let $dt = new D12(1);
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC22(cf29) {
      let $dt = new D12(2);
      $dt.cf29 = cf29;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get is_DC22() { return this.$tag === 2; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf29() { return this.cf29; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC23" + "(" + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC24" + "(" + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC22" + "(" + this.cf29.toVerbatimString(true) + ")";
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
        return other.$tag === 1 && this.cf33 === other.cf33 && _dafny.areEqual(this.cf34, other.cf34) && _dafny.areEqual(this.cf35, other.cf35) && this.cf36 === other.cf36 && _dafny.areEqual(this.cf37, other.cf37);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf29, other.cf29);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC23(false, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC26() {
      let $dt = new D13(0);
      return $dt;
    }
    static create_DC25(cf38) {
      let $dt = new D13(1);
      $dt.cf38 = cf38;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get dtor_cf38() { return this.cf38; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC26";
      } else if (this.$tag === 1) {
        return "D13.DC25" + "(" + _dafny.toString(this.cf38) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf38, other.cf38);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC26();
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
    static create_DC28(cf40, cf41, cf42) {
      let $dt = new D14(0);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC27(cf39) {
      let $dt = new D14(1);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC28" + "(" + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ", " + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC27" + "(" + _dafny.toString(this.cf39) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf40, other.cf40) && this.cf41 === other.cf41 && this.cf42 === other.cf42;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf39, other.cf39);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC28(_dafny.ZERO, false, null);
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
    static create_DC30(cf44) {
      let $dt = new D15(0);
      $dt.cf44 = cf44;
      return $dt;
    }
    static create_DC29(cf43) {
      let $dt = new D15(1);
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC31(cf45) {
      let $dt = new D15(2);
      $dt.cf45 = cf45;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get is_DC29() { return this.$tag === 1; }
    get is_DC31() { return this.$tag === 2; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf45() { return this.cf45; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC30" + "(" + _dafny.toString(this.cf44) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC29" + "(" + _dafny.toString(this.cf43) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC31" + "(" + _dafny.toString(this.cf45) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf44, other.cf44);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf43, other.cf43);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf45, other.cf45);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC30(new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC33(cf47, cf48, cf49, cf50) {
      let $dt = new D16(0);
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC34(cf51, cf52, cf53) {
      let $dt = new D16(1);
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      return $dt;
    }
    static create_DC32(cf46) {
      let $dt = new D16(2);
      $dt.cf46 = cf46;
      return $dt;
    }
    static create_DC35(cf54) {
      let $dt = new D16(3);
      $dt.cf54 = cf54;
      return $dt;
    }
    get is_DC33() { return this.$tag === 0; }
    get is_DC34() { return this.$tag === 1; }
    get is_DC32() { return this.$tag === 2; }
    get is_DC35() { return this.$tag === 3; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf54() { return this.cf54; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC33" + "(" + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC34" + "(" + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC32" + "(" + _dafny.toString(this.cf46) + ")";
      } else if (this.$tag === 3) {
        return "D16.DC35" + "(" + _dafny.toString(this.cf54) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf47 === other.cf47 && this.cf48 === other.cf48 && _dafny.areEqual(this.cf49, other.cf49) && this.cf50 === other.cf50;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf51 === other.cf51 && this.cf52 === other.cf52 && _dafny.areEqual(this.cf53, other.cf53);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf46, other.cf46);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf54, other.cf54);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC33(false, false, _dafny.ZERO, false);
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
    static create_DC36(cf55) {
      let $dt = new D17(0);
      $dt.cf55 = cf55;
      return $dt;
    }
    get is_DC36() { return this.$tag === 0; }
    get dtor_cf55() { return this.cf55; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC36" + "(" + _dafny.toString(this.cf55) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf55, other.cf55);
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
          return D17.Default();
        }
      };
    }
  }

  $module.D18 = class D18 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC38(cf57, cf58, cf59) {
      let $dt = new D18(0);
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      return $dt;
    }
    static create_DC37(cf56) {
      let $dt = new D18(1);
      $dt.cf56 = cf56;
      return $dt;
    }
    get is_DC38() { return this.$tag === 0; }
    get is_DC37() { return this.$tag === 1; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf56() { return this.cf56; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC38" + "(" + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC37" + "(" + _dafny.toString(this.cf56) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf57, other.cf57) && this.cf58 === other.cf58 && _dafny.areEqual(this.cf59, other.cf59);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf56 === other.cf56;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC38(_dafny.ZERO, false, _dafny.ZERO);
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
    static create_DC40(cf61, cf62, cf63) {
      let $dt = new D19(0);
      $dt.cf61 = cf61;
      $dt.cf62 = cf62;
      $dt.cf63 = cf63;
      return $dt;
    }
    static create_DC39(cf60) {
      let $dt = new D19(1);
      $dt.cf60 = cf60;
      return $dt;
    }
    get is_DC40() { return this.$tag === 0; }
    get is_DC39() { return this.$tag === 1; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf60() { return this.cf60; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC40" + "(" + _dafny.toString(this.cf61) + ", " + _dafny.toString(this.cf62) + ", " + _dafny.toString(this.cf63) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC39" + "(" + _dafny.toString(this.cf60) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf61, other.cf61) && this.cf62 === other.cf62 && _dafny.areEqual(this.cf63, other.cf63);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf60, other.cf60);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC40(_dafny.ZERO, false, _dafny.ZERO);
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
    static create_DC42(cf65, cf66, cf67) {
      let $dt = new D20(0);
      $dt.cf65 = cf65;
      $dt.cf66 = cf66;
      $dt.cf67 = cf67;
      return $dt;
    }
    static create_DC41(cf64) {
      let $dt = new D20(1);
      $dt.cf64 = cf64;
      return $dt;
    }
    get is_DC42() { return this.$tag === 0; }
    get is_DC41() { return this.$tag === 1; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf64() { return this.cf64; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC42" + "(" + _dafny.toString(this.cf65) + ", " + _dafny.toString(this.cf66) + ", " + _dafny.toString(this.cf67) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC41" + "(" + _dafny.toString(this.cf64) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf65, other.cf65) && this.cf66 === other.cf66 && this.cf67 === other.cf67;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf64, other.cf64);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC42(_dafny.ZERO, false, null);
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
    static create_DC44(cf69, cf70) {
      let $dt = new D21(0);
      $dt.cf69 = cf69;
      $dt.cf70 = cf70;
      return $dt;
    }
    static create_DC43(cf68) {
      let $dt = new D21(1);
      $dt.cf68 = cf68;
      return $dt;
    }
    static create_DC45(cf71) {
      let $dt = new D21(2);
      $dt.cf71 = cf71;
      return $dt;
    }
    get is_DC44() { return this.$tag === 0; }
    get is_DC43() { return this.$tag === 1; }
    get is_DC45() { return this.$tag === 2; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf71() { return this.cf71; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC44" + "(" + _dafny.toString(this.cf69) + ", " + _dafny.toString(this.cf70) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC43" + "(" + _dafny.toString(this.cf68) + ")";
      } else if (this.$tag === 2) {
        return "D21.DC45" + "(" + _dafny.toString(this.cf71) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf69, other.cf69) && this.cf70 === other.cf70;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf68, other.cf68);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf71, other.cf71);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC44(_dafny.ZERO, false);
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
    static create_DC46(cf72) {
      let $dt = new D22(0);
      $dt.cf72 = cf72;
      return $dt;
    }
    get is_DC46() { return this.$tag === 0; }
    get dtor_cf72() { return this.cf72; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC46" + "(" + _dafny.toString(this.cf72) + ")";
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
      return _dafny.Seq.of();
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
    static create_DC48(cf74, cf75) {
      let $dt = new D23(0);
      $dt.cf74 = cf74;
      $dt.cf75 = cf75;
      return $dt;
    }
    static create_DC47(cf73) {
      let $dt = new D23(1);
      $dt.cf73 = cf73;
      return $dt;
    }
    static create_DC49(cf76) {
      let $dt = new D23(2);
      $dt.cf76 = cf76;
      return $dt;
    }
    get is_DC48() { return this.$tag === 0; }
    get is_DC47() { return this.$tag === 1; }
    get is_DC49() { return this.$tag === 2; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf76() { return this.cf76; }
    toString() {
      if (this.$tag === 0) {
        return "D23.DC48" + "(" + _dafny.toString(this.cf74) + ", " + _dafny.toString(this.cf75) + ")";
      } else if (this.$tag === 1) {
        return "D23.DC47" + "(" + _dafny.toString(this.cf73) + ")";
      } else if (this.$tag === 2) {
        return "D23.DC49" + "(" + _dafny.toString(this.cf76) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf74 === other.cf74 && _dafny.areEqual(this.cf75, other.cf75);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf73, other.cf73);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf76, other.cf76);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D23.create_DC48(false, _dafny.Seq.of());
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
    static create_DC50(cf77) {
      let $dt = new D24(0);
      $dt.cf77 = cf77;
      return $dt;
    }
    get is_DC50() { return this.$tag === 0; }
    get dtor_cf77() { return this.cf77; }
    toString() {
      if (this.$tag === 0) {
        return "D24.DC50" + "(" + _dafny.toString(this.cf77) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf77, other.cf77);
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
    static create_DC52(cf79) {
      let $dt = new D25(0);
      $dt.cf79 = cf79;
      return $dt;
    }
    static create_DC53() {
      let $dt = new D25(1);
      return $dt;
    }
    static create_DC51(cf78) {
      let $dt = new D25(2);
      $dt.cf78 = cf78;
      return $dt;
    }
    static create_DC54(cf80) {
      let $dt = new D25(3);
      $dt.cf80 = cf80;
      return $dt;
    }
    get is_DC52() { return this.$tag === 0; }
    get is_DC53() { return this.$tag === 1; }
    get is_DC51() { return this.$tag === 2; }
    get is_DC54() { return this.$tag === 3; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf80() { return this.cf80; }
    toString() {
      if (this.$tag === 0) {
        return "D25.DC52" + "(" + _dafny.toString(this.cf79) + ")";
      } else if (this.$tag === 1) {
        return "D25.DC53";
      } else if (this.$tag === 2) {
        return "D25.DC51" + "(" + _dafny.toString(this.cf78) + ")";
      } else if (this.$tag === 3) {
        return "D25.DC54" + "(" + _dafny.toString(this.cf80) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf79, other.cf79);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf78 === other.cf78;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf80, other.cf80);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D25.create_DC52(_dafny.ZERO);
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
    static create_DC56(cf82, cf83, cf84, cf85, cf86) {
      let $dt = new D26(0);
      $dt.cf82 = cf82;
      $dt.cf83 = cf83;
      $dt.cf84 = cf84;
      $dt.cf85 = cf85;
      $dt.cf86 = cf86;
      return $dt;
    }
    static create_DC55(cf81) {
      let $dt = new D26(1);
      $dt.cf81 = cf81;
      return $dt;
    }
    get is_DC56() { return this.$tag === 0; }
    get is_DC55() { return this.$tag === 1; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf83() { return this.cf83; }
    get dtor_cf84() { return this.cf84; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf81() { return this.cf81; }
    toString() {
      if (this.$tag === 0) {
        return "D26.DC56" + "(" + _dafny.toString(this.cf82) + ", " + _dafny.toString(this.cf83) + ", " + _dafny.toString(this.cf84) + ", " + _dafny.toString(this.cf85) + ", " + this.cf86.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D26.DC55" + "(" + _dafny.toString(this.cf81) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf82, other.cf82) && this.cf83 === other.cf83 && _dafny.areEqual(this.cf84, other.cf84) && _dafny.areEqual(this.cf85, other.cf85) && _dafny.areEqual(this.cf86, other.cf86);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf81, other.cf81);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D26.create_DC56(_dafny.ZERO, null, _dafny.Seq.of(), _dafny.Seq.of(), _dafny.Seq.UnicodeFromString(""));
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
    static create_DC58(cf88, cf89, cf90) {
      let $dt = new D27(0);
      $dt.cf88 = cf88;
      $dt.cf89 = cf89;
      $dt.cf90 = cf90;
      return $dt;
    }
    static create_DC57(cf87) {
      let $dt = new D27(1);
      $dt.cf87 = cf87;
      return $dt;
    }
    get is_DC58() { return this.$tag === 0; }
    get is_DC57() { return this.$tag === 1; }
    get dtor_cf88() { return this.cf88; }
    get dtor_cf89() { return this.cf89; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf87() { return this.cf87; }
    toString() {
      if (this.$tag === 0) {
        return "D27.DC58" + "(" + _dafny.toString(this.cf88) + ", " + _dafny.toString(this.cf89) + ", " + _dafny.toString(this.cf90) + ")";
      } else if (this.$tag === 1) {
        return "D27.DC57" + "(" + _dafny.toString(this.cf87) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf88 === other.cf88 && _dafny.areEqual(this.cf89, other.cf89) && this.cf90 === other.cf90;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf87, other.cf87);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D27.create_DC58([], _dafny.ZERO, false);
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
    static create_DC60(cf92, cf93, cf94) {
      let $dt = new D28(0);
      $dt.cf92 = cf92;
      $dt.cf93 = cf93;
      $dt.cf94 = cf94;
      return $dt;
    }
    static create_DC61(cf95, cf96, cf97, cf98) {
      let $dt = new D28(1);
      $dt.cf95 = cf95;
      $dt.cf96 = cf96;
      $dt.cf97 = cf97;
      $dt.cf98 = cf98;
      return $dt;
    }
    static create_DC59(cf91) {
      let $dt = new D28(2);
      $dt.cf91 = cf91;
      return $dt;
    }
    static create_DC62(cf99) {
      let $dt = new D28(3);
      $dt.cf99 = cf99;
      return $dt;
    }
    get is_DC60() { return this.$tag === 0; }
    get is_DC61() { return this.$tag === 1; }
    get is_DC59() { return this.$tag === 2; }
    get is_DC62() { return this.$tag === 3; }
    get dtor_cf92() { return this.cf92; }
    get dtor_cf93() { return this.cf93; }
    get dtor_cf94() { return this.cf94; }
    get dtor_cf95() { return this.cf95; }
    get dtor_cf96() { return this.cf96; }
    get dtor_cf97() { return this.cf97; }
    get dtor_cf98() { return this.cf98; }
    get dtor_cf91() { return this.cf91; }
    get dtor_cf99() { return this.cf99; }
    toString() {
      if (this.$tag === 0) {
        return "D28.DC60" + "(" + _dafny.toString(this.cf92) + ", " + this.cf93.toVerbatimString(true) + ", " + _dafny.toString(this.cf94) + ")";
      } else if (this.$tag === 1) {
        return "D28.DC61" + "(" + _dafny.toString(this.cf95) + ", " + _dafny.toString(this.cf96) + ", " + _dafny.toString(this.cf97) + ", " + _dafny.toString(this.cf98) + ")";
      } else if (this.$tag === 2) {
        return "D28.DC59" + "(" + _dafny.toString(this.cf91) + ")";
      } else if (this.$tag === 3) {
        return "D28.DC62" + "(" + _dafny.toString(this.cf99) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf92, other.cf92) && _dafny.areEqual(this.cf93, other.cf93) && _dafny.areEqual(this.cf94, other.cf94);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf95, other.cf95) && this.cf96 === other.cf96 && _dafny.areEqual(this.cf97, other.cf97) && this.cf98 === other.cf98;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf91, other.cf91);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf99, other.cf99);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D28.create_DC60(_dafny.ZERO, _dafny.Seq.UnicodeFromString(""), _dafny.Set.Empty);
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
    static create_DC64() {
      let $dt = new D29(0);
      return $dt;
    }
    static create_DC63(cf100) {
      let $dt = new D29(1);
      $dt.cf100 = cf100;
      return $dt;
    }
    static create_DC65(cf101) {
      let $dt = new D29(2);
      $dt.cf101 = cf101;
      return $dt;
    }
    get is_DC64() { return this.$tag === 0; }
    get is_DC63() { return this.$tag === 1; }
    get is_DC65() { return this.$tag === 2; }
    get dtor_cf100() { return this.cf100; }
    get dtor_cf101() { return this.cf101; }
    toString() {
      if (this.$tag === 0) {
        return "D29.DC64";
      } else if (this.$tag === 1) {
        return "D29.DC63" + "(" + _dafny.toString(this.cf100) + ")";
      } else if (this.$tag === 2) {
        return "D29.DC65" + "(" + _dafny.toString(this.cf101) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf100, other.cf100);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf101, other.cf101);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D29.create_DC64();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D29.Default();
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
      this.f2 = _dafny.ZERO;
      this.f7 = _dafny.ZERO;
      this._f0 = _dafny.ZERO;
      this._f1 = _dafny.ZERO;
      this._f3 = _dafny.ZERO;
      this._f4 = false;
      this._f5 = _dafny.ZERO;
      this._f6 = false;
      this._f8 = _dafny.ZERO;
      this._f9 = _dafny.ZERO;
      this._f10 = false;
      this._f11 = [];
      this._f12 = _dafny.MultiSet.Empty;
      this._f13 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this).f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
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
      this.f26 = [];
      this.f27 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f26, f27) {
      let _this = this;
      (_this).f26 = f26;
      (_this).f27 = f27;
      return;
    }
    fm36(p0, p1, p2, globalState) {
      let _this = this;
      return _this.f27;
    };
    fm37(p0, globalState) {
      let _this = this;
      return (_module.D9.create_DC16(_this.f27, !(false))).dtor_cf24;
    };
    m15(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _316_v0;
      let _init7 = function (_317_i0) {
        return true;
      };
      let _nw74 = Array((new BigNumber(14)).toNumber());
      for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw74.length); _i0_7++) {
        _nw74[_i0_7] = _init7(new BigNumber(_i0_7));
      }
      _316_v0 = _nw74;
      let _index44 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_316_v0).length));
      (_316_v0)[_index44] = (_this.f27).isLessThanOrEqualTo(new BigNumber(520));
      let _318_v1;
      _318_v1 = true;
      let _index45 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_316_v0).length));
      (_316_v0)[_index45] = _318_v1;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_316_v0).length))) {
        let _319_i1 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_319_i1)) && ((_319_i1).isLessThan(new BigNumber((_316_v0).length))))) {
          (_316_v0)[(_319_i1)] = !((p2).isLessThan(_this.f27));
        }
      }
      let _320_v2;
      _320_v2 = _dafny.MultiSet.fromElements(false);
      let _321_v3;
      _321_v3 = _dafny.Seq.of((_316_v0)[_module.__default.safeIndex(new BigNumber(253), new BigNumber((_316_v0).length))], _318_v1);
      let _322_v4;
      _322_v4 = new BigNumber((_321_v3).length);
      let _323_v5;
      _323_v5 = _dafny.Set.fromElements((((_320_v2).contains((_this).fm37((_316_v0)[_module.__default.safeIndex(new BigNumber(253), new BigNumber((_316_v0).length))], globalState))) ? ((_320_v2).get((_this).fm37((_316_v0)[_module.__default.safeIndex(new BigNumber(253), new BigNumber((_316_v0).length))], globalState))) : ((_322_v4))), _this.f27, p2);
      let _324_v6;
      let _init8 = ((_325_p2) => function (_326_i2) {
        return _module.__default.safeDivisionInt(_326_i2, _325_p2);
      })(p2);
      let _nw75 = Array((new BigNumber(14)).toNumber());
      for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw75.length); _i0_8++) {
        _nw75[_i0_8] = _init8(new BigNumber(_i0_8));
      }
      _324_v6 = _nw75;
      let _327_v7;
      _327_v7 = _module.D9.create_DC15(false, (_316_v0)[_module.__default.safeIndex(new BigNumber(253), new BigNumber((_316_v0).length))], _323_v5, _324_v6, _316_v0);
      _327_v7 = _327_v7;
      let _328_v8;
      _328_v8 = _module.D7.create_DC12();
      let _329_i3;
      _329_i3 = _dafny.ZERO;
      L1: {
        let _pat_let_tv5 = _323_v5;
        let _pat_let_tv6 = _318_v1;
        while (function (_source3) {
          if (_source3.is_DC12) {
            return !(_pat_let_tv5).equals(_dafny.Set.fromElements(_this.f27));
          } else {
            let _332___mcc_h0 = (_source3).cf15;
            let _333_cf15 = _332___mcc_h0;
            return _pat_let_tv6;
          }
        }(_328_v8)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_329_i3)) {
              break L1;
            }
            _329_i3 = (_329_i3).plus(_dafny.ONE);
            let _index46 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_316_v0).length));
            let _rhs48 = true;
            let _rhs49 = (_dafny.ZERO).minus(_module.__default.fm0((_316_v0)[_module.__default.safeIndex(new BigNumber(253), new BigNumber((_316_v0).length))], globalState));
            let _lhs31 = _316_v0;
            let _lhs32 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_316_v0).length));
            let _lhs33 = globalState;
            _lhs31[_lhs32] = _rhs48;
            _lhs33.f7 = _rhs49;
            let _330_v9;
            _330_v9 = _module.D9.create_DC16(new BigNumber((p3).length), (_316_v0)[_module.__default.safeIndex(new BigNumber(253), new BigNumber((_316_v0).length))]);
            let _index47 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_316_v0).length));
            (_316_v0)[_index47] = !(!(!((_330_v9).dtor_cf24)));
            _318_v1 = (p2).isLessThanOrEqualTo(_this.f27);
            let _331_v10;
            let _nw76 = Array((new BigNumber(17)).toNumber());
            _nw76[(_dafny.ZERO).toNumber()] = _316_v0;
            _nw76[(_dafny.ONE).toNumber()] = _316_v0;
            _nw76[(new BigNumber(2)).toNumber()] = _316_v0;
            _nw76[(new BigNumber(3)).toNumber()] = _316_v0;
            _nw76[(new BigNumber(4)).toNumber()] = _316_v0;
            _nw76[(new BigNumber(5)).toNumber()] = _316_v0;
            _nw76[(new BigNumber(6)).toNumber()] = _316_v0;
            _nw76[(new BigNumber(7)).toNumber()] = _316_v0;
            _nw76[(new BigNumber(8)).toNumber()] = _316_v0;
            _nw76[(new BigNumber(9)).toNumber()] = _316_v0;
            _nw76[(new BigNumber(10)).toNumber()] = _316_v0;
            _nw76[(new BigNumber(11)).toNumber()] = _316_v0;
            _nw76[(new BigNumber(12)).toNumber()] = (((_316_v0)[_module.__default.safeIndex(new BigNumber(253), new BigNumber((_316_v0).length))]) ? (_316_v0) : (_316_v0));
            _nw76[(new BigNumber(13)).toNumber()] = _316_v0;
            _nw76[(new BigNumber(14)).toNumber()] = _316_v0;
            _nw76[(new BigNumber(15)).toNumber()] = _316_v0;
            _nw76[(new BigNumber(16)).toNumber()] = _316_v0;
            _331_v10 = _nw76;
            _331_v10 = ((_318_v1) ? (_331_v10) : (((_318_v1) ? (_331_v10) : (_331_v10))));
          }
        }
      }
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_316_v0).length))) {
        let _334_i4 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_334_i4)) && ((_334_i4).isLessThan(new BigNumber((_316_v0).length))))) {
          (_316_v0)[(_334_i4)] = (_module.D9.create_DC16(_this.f27, _318_v1)).dtor_cf24;
        }
      }
      let _335_v11;
      _335_v11 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_316_v0)[_module.__default.safeIndex(new BigNumber(253), new BigNumber((_316_v0).length))]);
      let _336_i5;
      _336_i5 = _dafny.ZERO;
      L2: {
        while (!(_335_v11).contains(p2)) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_336_i5)) {
              break L2;
            }
            _336_i5 = (_336_i5).plus(_dafny.ONE);
            if (!((_316_v0)[_module.__default.safeIndex(new BigNumber(253), new BigNumber((_316_v0).length))])) {
              _323_v5 = _323_v5;
              let _337_v12;
              _337_v12 = _dafny.Map.Empty.slice().updateUnsafe(_this.f27,_316_v0);
              let _338_v13;
              let _nw77 = Array((new BigNumber(4)).toNumber());
              _nw77[(_dafny.ZERO).toNumber()] = _337_v12;
              _nw77[(_dafny.ONE).toNumber()] = _337_v12;
              _nw77[(new BigNumber(2)).toNumber()] = _337_v12;
              _nw77[(new BigNumber(3)).toNumber()] = _337_v12;
              _338_v13 = _nw77;
              let _index48 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_338_v13).length));
              (_338_v13)[_index48] = (_337_v12).Merge(_337_v12);
              let _339_v14;
              _339_v14 = _dafny.Seq.of(_this.f27, new BigNumber((p3).length));
              let _index49 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_338_v13).length));
              (_338_v13)[_index49] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_339_v14).length),_316_v0);
              _318_v1 = false;
              let _340_v15;
              _340_v15 = _dafny.Map.Empty.slice().updateUnsafe(p2,_339_v14);
              _340_v15 = (_340_v15).update(new BigNumber((_320_v2).cardinality()), _dafny.Seq.Concat((((p1).contains(_318_v1)) ? ((p1).get(_318_v1)) : (_dafny.Seq.of(_this.f27))), _339_v14));
              (globalState).f7 = new BigNumber((_323_v5).length);
            } else {
              _318_v1 = ((_320_v2).update(!(!((_316_v0)[_module.__default.safeIndex(new BigNumber(253), new BigNumber((_316_v0).length))])), _module.__default.abs(p2))).IsSubsetOf(_320_v2);
              let _341_v16;
              let _init9 = ((_342_v4) => function (_343_i6) {
                return _342_v4;
              })(_322_v4);
              let _nw78 = Array((new BigNumber(28)).toNumber());
              for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw78.length); _i0_9++) {
                _nw78[_i0_9] = _init9(new BigNumber(_i0_9));
              }
              _341_v16 = _nw78;
              let _index50 = _module.__default.safeIndex(new BigNumber(559), new BigNumber((_341_v16).length));
              (_341_v16)[_index50] = _322_v4;
              let _index51 = _module.__default.safeIndex(new BigNumber(559), new BigNumber((_341_v16).length));
              (_341_v16)[_index51] = (p2).minus(_this.f27);
              let _index52 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_316_v0).length));
              (_316_v0)[_index52] = _318_v1;
              let _344_v17;
              _344_v17 = _module.D3.create_DC4(_318_v1, _this.f27, _this.f27, (_316_v0)[_module.__default.safeIndex(new BigNumber(253), new BigNumber((_316_v0).length))], (_316_v0)[_module.__default.safeIndex(new BigNumber(253), new BigNumber((_316_v0).length))]);
              (_this).f27 = (_344_v17).dtor_cf5;
              _335_v11 = (_335_v11).update(new BigNumber(366), _318_v1);
            }
            let _345_v18;
            _345_v18 = _dafny.Seq.of(_this.f27);
            let _346_v19;
            _346_v19 = _dafny.Seq.of(_module.__default.safeModuloInt(new BigNumber((_321_v3).length), p2), _this.f27, (_345_v18)[_module.__default.safeIndex(_this.f27, new BigNumber((_345_v18).length))], new BigNumber(821), p2);
            let _347_v20;
            _347_v20 = _dafny.Map.Empty.slice().updateUnsafe(_318_v1,_318_v1);
            let _rhs50 = p2;
            let _rhs51 = _dafny.Seq.Concat(_dafny.Seq.update(_module.__default.fm38(false, p2, p2, globalState), _module.__default.safeIndex(p2, new BigNumber((_module.__default.fm38(false, p2, p2, globalState)).length)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f27,_318_v1)).length)), _346_v19);
            let _rhs52 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p2),(((_316_v0)[_module.__default.safeIndex(new BigNumber(253), new BigNumber((_316_v0).length))]) ? ((((_347_v20).contains(!((_316_v0)[_module.__default.safeIndex(new BigNumber(253), new BigNumber((_316_v0).length))]))) ? ((_347_v20).get(!((_316_v0)[_module.__default.safeIndex(new BigNumber(253), new BigNumber((_316_v0).length))]))) : (_318_v1))) : ((_316_v0)[_module.__default.safeIndex(new BigNumber(253), new BigNumber((_316_v0).length))])));
            let _lhs34 = globalState;
            _lhs34.f7 = _rhs50;
            _345_v18 = _rhs51;
            _335_v11 = _rhs52;
            let _348_v21;
            _348_v21 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_this).fm37(_318_v1, globalState),p2),p2);
            let _349_v22;
            _349_v22 = _dafny.Map.Empty.slice().updateUnsafe((_316_v0)[_module.__default.safeIndex(new BigNumber(253), new BigNumber((_316_v0).length))],(_345_v18)[_module.__default.safeIndex(p2, new BigNumber((_345_v18).length))]);
            _348_v21 = (_348_v21).update(_349_v22, _this.f27);
            let _350_v23;
            let _init10 = function (_351_i7) {
              return new _dafny.CodePoint('y'.codePointAt(0));
            };
            let _nw79 = Array((new BigNumber(12)).toNumber());
            for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw79.length); _i0_10++) {
              _nw79[_i0_10] = _init10(new BigNumber(_i0_10));
            }
            _350_v23 = _nw79;
            let _352_v24;
            _352_v24 = new _dafny.CodePoint('e'.codePointAt(0));
            let _index53 = _module.__default.safeIndex(new BigNumber(81), new BigNumber((_350_v23).length));
            (_350_v23)[_index53] = _352_v24;
            let _index54 = _module.__default.safeIndex(new BigNumber(81), new BigNumber((_350_v23).length));
            (_350_v23)[_index54] = _352_v24;
          }
        }
      }
      return;
    }
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f14 = _dafny.ZERO;
      this.f25 = false;
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
    __ctor(f25, f14) {
      let _this = this;
      (_this).f25 = f25;
      (_this)._f14 = f14;
      return;
    }
    m13(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = new _dafny.CodePoint('D'.codePointAt(0));
      let r3 = _dafny.Map.Empty;
      let _353_v0;
      let _nw80 = Array((new BigNumber(6)).toNumber()).fill(false);
      _353_v0 = _nw80;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_353_v0).length))) {
        let _354_i0 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_354_i0)) && ((_354_i0).isLessThan(new BigNumber((_353_v0).length))))) {
          (_353_v0)[(_354_i0)] = (p0) || (p0);
        }
      }
      let _355_v1;
      let _init11 = ((_356_p1) => function (_357_i1) {
        return (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f14,_dafny.Seq.of(_356_p1))).length))).Difference(_dafny.MultiSet.fromElements(_this.f14, _this.f14));
      })(p1);
      let _nw81 = Array((new BigNumber(17)).toNumber());
      for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw81.length); _i0_11++) {
        _nw81[_i0_11] = _init11(new BigNumber(_i0_11));
      }
      _355_v1 = _nw81;
      let _358_v2;
      _358_v2 = _dafny.Seq.UnicodeFromString("njtfyb");
      let _359_v3;
      _359_v3 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_358_v2).length));
      let _360_v4;
      _360_v4 = _dafny.MultiSet.fromElements(new BigNumber((_359_v3).length), _this.f14);
      let _index55 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_355_v1).length));
      (_355_v1)[_index55] = (_360_v4).Intersect(_360_v4);
      let _index56 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_355_v1).length));
      (_355_v1)[_index56] = _360_v4;
      let _361_v5;
      let _init12 = ((_362_p0, _363_p1) => function (_364_i2) {
        return _dafny.Set.fromElements(_362_p0, _363_p1);
      })(p0, p1);
      let _nw82 = Array((new BigNumber(27)).toNumber());
      for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw82.length); _i0_12++) {
        _nw82[_i0_12] = _init12(new BigNumber(_i0_12));
      }
      _361_v5 = _nw82;
      let _365_v6;
      let _nw83 = new _module.C0();
      _nw83.__ctor(_361_v5, _this.f14);
      _365_v6 = _nw83;
      let _366_v7;
      _366_v7 = _dafny.Map.Empty.slice().updateUnsafe(_358_v2,true);
      let _367_v8;
      _367_v8 = _module.D3.create_DC4(p0, new BigNumber(517), _365_v6.f27, true, (((_366_v7).contains(_358_v2)) ? ((_366_v7).get(_358_v2)) : (p1)));
      let _hi2 = _365_v6.f27;
      for (let _368_i3 = (_367_v8).dtor_cf5; _368_i3.isLessThan(_hi2); _368_i3 = _368_i3.plus(_dafny.ONE)) {
        (_this).f25 = p0;
        let _369_v9;
        _369_v9 = _dafny.Map.Empty.slice().updateUnsafe((_365_v6).fm37(!(p1), globalState),_module.__default.fm39(globalState));
        let _370_v10;
        _370_v10 = new _dafny.CodePoint('y'.codePointAt(0));
        _369_v9 = (_369_v9).update(_this.f25, _370_v10);
        let _371_v11;
        _371_v11 = _dafny.Map.Empty.slice().updateUnsafe(_this.f25,_this.f25);
        let _372_v12;
        _372_v12 = _dafny.Map.Empty.slice().updateUnsafe(p1,_371_v11);
        let _373_v13;
        _373_v13 = _dafny.Map.Empty.slice().updateUnsafe(p0,_371_v11);
        r1 = !((((_372_v12)).update(p1, _module.__default.fm40(p1, globalState))).Merge(_373_v13)).contains(true);
        (_this).f25 = p1;
      }
      let _374_v14;
      _374_v14 = new _dafny.CodePoint('o'.codePointAt(0));
      let _375_v15;
      _375_v15 = _dafny.Map.Empty.slice().updateUnsafe(_374_v14,_this.f14);
      let _376_v16;
      _376_v16 = _module.D9.create_DC16(new BigNumber(((_359_v3).update(_this.f25, new BigNumber(968))).length), true);
      let _hi3 = _this.f14;
      for (let _377_i4 = (((_375_v15).contains(_374_v14)) ? ((_375_v15).get(_374_v14)) : (_module.__default.fm0((_376_v16).dtor_cf24, globalState))); _377_i4.isLessThan(_hi3); _377_i4 = _377_i4.plus(_dafny.ONE)) {
        (_this).f25 = _this.f25;
        let _378_v17;
        _378_v17 = _dafny.Seq.of(p0, p1);
        let _379_v18;
        _379_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_378_v17).length),new BigNumber(-186));
        let _380_v19;
        _380_v19 = _dafny.Seq.of(_module.__default.safeModuloInt(_365_v6.f27, new BigNumber(-108)), new BigNumber(((_379_v18).Merge((_379_v18).update(_this.f14, (_dafny.ZERO).minus(_365_v6.f27)))).length), (_dafny.ZERO).minus(_365_v6.f27), _365_v6.f27);
        let _rhs53 = _365_v6.f27;
        let _rhs54 = _dafny.Seq.of(_365_v6.f27);
        let _rhs55 = _module.__default.fm38((_376_v16).dtor_cf24, (_dafny.ZERO).minus(_365_v6.f27), _module.__default.safeDivisionInt(_365_v6.f27, _377_i4), globalState);
        let _lhs35 = globalState;
        _lhs35.f2 = _rhs53;
        _380_v19 = _rhs54;
        _380_v19 = _rhs55;
        if (!(p1)) {
          (_this).f25 = !(p1);
          let _381_v20;
          _381_v20 = _dafny.Map.Empty.slice().updateUnsafe(_374_v14,_this.f25);
          let _382_v21;
          _382_v21 = _dafny.Seq.of(_378_v17);
          let _383_v22;
          _383_v22 = _dafny.MultiSet.fromElements(p0, p1);
          let _384_v23;
          let _nw84 = Array((new BigNumber(21)).toNumber());
          _nw84[(_dafny.ZERO).toNumber()] = _module.__default.safeDivisionInt(_this.f14, _365_v6.f27);
          _nw84[(_dafny.ONE).toNumber()] = (_365_v6.f27).plus(_module.__default.fm0(_this.f25, globalState));
          _nw84[(new BigNumber(2)).toNumber()] = (_module.__default.fm0((((_381_v20).contains(_374_v14)) ? ((_381_v20).get(_374_v14)) : (_this.f25)), globalState)).plus(new BigNumber(-340));
          _nw84[(new BigNumber(3)).toNumber()] = new BigNumber(902);
          _nw84[(new BigNumber(4)).toNumber()] = _365_v6.f27;
          _nw84[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus((((_359_v3).contains(p1)) ? ((_359_v3).get(p1)) : (_365_v6.f27)));
          _nw84[(new BigNumber(6)).toNumber()] = new BigNumber((_359_v3).length);
          _nw84[(new BigNumber(7)).toNumber()] = new BigNumber((_358_v2).length);
          _nw84[(new BigNumber(8)).toNumber()] = (new BigNumber(((_382_v21)[_module.__default.safeIndex(_365_v6.f27, new BigNumber((_382_v21).length))]).length)).minus(_365_v6.f27);
          _nw84[(new BigNumber(9)).toNumber()] = _365_v6.f27;
          _nw84[(new BigNumber(10)).toNumber()] = (_367_v8).dtor_cf5;
          _nw84[(new BigNumber(11)).toNumber()] = _365_v6.f27;
          _nw84[(new BigNumber(12)).toNumber()] = (((_379_v18).contains(_module.__default.fm0(p1, globalState))) ? ((_379_v18).get(_module.__default.fm0(p1, globalState))) : (new BigNumber((_383_v22).cardinality())));
          _nw84[(new BigNumber(13)).toNumber()] = _365_v6.f27;
          _nw84[(new BigNumber(14)).toNumber()] = new BigNumber((_380_v19).length);
          _nw84[(new BigNumber(15)).toNumber()] = _365_v6.f27;
          _nw84[(new BigNumber(16)).toNumber()] = _365_v6.f27;
          _nw84[(new BigNumber(17)).toNumber()] = _365_v6.f27;
          _nw84[(new BigNumber(18)).toNumber()] = _this.f14;
          _nw84[(new BigNumber(19)).toNumber()] = (_365_v6.f27).minus(new BigNumber(-472));
          _nw84[(new BigNumber(20)).toNumber()] = _365_v6.f27;
          _384_v23 = _nw84;
          _384_v23 = _384_v23;
          let _385_v24;
          _385_v24 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Map.Empty.slice().updateUnsafe(_this.f25,_this.f25));
          let _386_v25;
          _386_v25 = _385_v24;
          let _387_v26;
          _387_v26 = _dafny.Map.Empty.slice().updateUnsafe(_this.f25,p0);
          let _388_v27;
          let _nw85 = Array((new BigNumber(28)).toNumber());
          _nw85[(_dafny.ZERO).toNumber()] = _386_v25;
          _nw85[(_dafny.ONE).toNumber()] = _385_v24;
          _nw85[(new BigNumber(2)).toNumber()] = _module.__default.fm41(p1, p1, globalState);
          _nw85[(new BigNumber(3)).toNumber()] = ((p1) ? (_386_v25) : (_386_v25));
          _nw85[(new BigNumber(4)).toNumber()] = _386_v25;
          _nw85[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_this.f25,_387_v26);
          _nw85[(new BigNumber(6)).toNumber()] = _386_v25;
          _nw85[(new BigNumber(7)).toNumber()] = _386_v25;
          _nw85[(new BigNumber(8)).toNumber()] = _386_v25;
          _nw85[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,_387_v26);
          _nw85[(new BigNumber(10)).toNumber()] = _386_v25;
          _nw85[(new BigNumber(11)).toNumber()] = _386_v25;
          _nw85[(new BigNumber(12)).toNumber()] = _386_v25;
          _nw85[(new BigNumber(13)).toNumber()] = _386_v25;
          _nw85[(new BigNumber(14)).toNumber()] = _386_v25;
          _nw85[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,_387_v26);
          _nw85[(new BigNumber(16)).toNumber()] = _386_v25;
          _nw85[(new BigNumber(17)).toNumber()] = _module.__default.fm41(p0, p0, globalState);
          _nw85[(new BigNumber(18)).toNumber()] = _386_v25;
          _nw85[(new BigNumber(19)).toNumber()] = _385_v24;
          _nw85[(new BigNumber(20)).toNumber()] = _386_v25;
          _nw85[(new BigNumber(21)).toNumber()] = _386_v25;
          _nw85[(new BigNumber(22)).toNumber()] = _386_v25;
          _nw85[(new BigNumber(23)).toNumber()] = _386_v25;
          _nw85[(new BigNumber(24)).toNumber()] = _386_v25;
          _nw85[(new BigNumber(25)).toNumber()] = _module.__default.fm41(p0, _this.f25, globalState);
          _nw85[(new BigNumber(26)).toNumber()] = _386_v25;
          _nw85[(new BigNumber(27)).toNumber()] = ((p0) ? (_386_v25) : (_386_v25));
          _388_v27 = _nw85;
          let _index57 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_388_v27).length));
          (_388_v27)[_index57] = _385_v24;
          let _index58 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_388_v27).length));
          (_388_v27)[_index58] = _386_v25;
          (globalState).f7 = _this.f14;
          let _389_v28;
          let _nw86 = Array((new BigNumber(2)).toNumber()).fill([]);
          _389_v28 = _nw86;
          let _index59 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_389_v28).length));
          (_389_v28)[_index59] = _384_v23;
          let _index60 = _module.__default.safeIndex(new BigNumber(672), new BigNumber((_389_v28).length));
          (_389_v28)[_index60] = ((_this.f25) ? (_384_v23) : (_384_v23));
          let _390_v29;
          _390_v29 = _module.D6.create_DC9(_374_v14);
          let _391_v31;
          _391_v31 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll29 = new _dafny.Set();
            for (const _compr_29 of (_dafny.Seq.of(_365_v6.f27)).Elements) {
              let _392_v30 = _compr_29;
              if (_dafny.Seq.contains(_dafny.Seq.of(_365_v6.f27), _392_v30)) {
                _coll29.add((_392_v30).minus(new BigNumber(668)));
              }
            }
            return _coll29;
          }(),_358_v2);
          let _393_v32;
          _393_v32 = _dafny.Set.fromElements(_365_v6.f27, _365_v6.f27);
          let _394_v33;
          let _nw87 = Array((new BigNumber(24)).toNumber());
          _nw87[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_358_v2, _module.__default.safeIndex(_365_v6.f27, new BigNumber((_358_v2).length)), (_390_v29).dtor_cf14);
          _nw87[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_358_v2, _358_v2);
          _nw87[(new BigNumber(2)).toNumber()] = _358_v2;
          _nw87[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_358_v2, _358_v2);
          _nw87[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(595)), ((_395_v14) => function (_396_i5) {
            return _395_v14;
          })(_374_v14));
          _nw87[(new BigNumber(5)).toNumber()] = _358_v2;
          _nw87[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("slao");
          _nw87[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm42(p1, new BigNumber((_dafny.MultiSet.fromElements(_365_v6.f27, _365_v6.f27, _365_v6.f27)).cardinality()), _365_v6.f27, p1, globalState), _358_v2);
          _nw87[(new BigNumber(8)).toNumber()] = _dafny.Seq.update(_358_v2, _module.__default.safeIndex(_this.f14, new BigNumber((_358_v2).length)), _374_v14);
          _nw87[(new BigNumber(9)).toNumber()] = _358_v2;
          _nw87[(new BigNumber(10)).toNumber()] = _358_v2;
          _nw87[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_358_v2, _module.__default.fm42(p0, _365_v6.f27, _365_v6.f27, p1, globalState));
          _nw87[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_358_v2, _358_v2);
          _nw87[(new BigNumber(13)).toNumber()] = _358_v2;
          _nw87[(new BigNumber(14)).toNumber()] = _358_v2;
          _nw87[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_358_v2, _358_v2);
          _nw87[(new BigNumber(16)).toNumber()] = _dafny.Seq.UnicodeFromString("sub");
          _nw87[(new BigNumber(17)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(21)), ((_397_v14) => function (_398_i6) {
            return _397_v14;
          })(_374_v14)), _module.__default.safeIndex(new BigNumber(-560), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(21)), ((_399_v14) => function (_400_i6) {
            return _399_v14;
          })(_374_v14))).length)), _374_v14);
          _nw87[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_358_v2, _358_v2);
          _nw87[(new BigNumber(19)).toNumber()] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("qrtitlk"), _module.__default.safeIndex(_377_i4, new BigNumber((_dafny.Seq.UnicodeFromString("qrtitlk")).length)), _374_v14);
          _nw87[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(605)), ((_401_v14) => function (_402_i7) {
            return _401_v14;
          })(_374_v14)), _dafny.Seq.UnicodeFromString("ytmlhni"));
          _nw87[(new BigNumber(21)).toNumber()] = _358_v2;
          _nw87[(new BigNumber(22)).toNumber()] = (((_391_v31).contains(_393_v32)) ? ((_391_v31).get(_393_v32)) : (_module.__default.fm42(p0, _365_v6.f27, new BigNumber(453), false, globalState)));
          _nw87[(new BigNumber(23)).toNumber()] = _358_v2;
          _394_v33 = _nw87;
          let _index61 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_394_v33).length));
          (_394_v33)[_index61] = _dafny.Seq.UnicodeFromString("bwcpbhs");
          let _index62 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_384_v23).length));
          (_384_v23)[_index62] = new BigNumber((_358_v2).length);
          let _403_v34;
          _403_v34 = _dafny.Seq.of(_358_v2);
          let _404_v35;
          _404_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_378_v17)).length),_353_v0);
          let _index63 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_389_v28).length));
          let _index64 = _module.__default.safeIndex(new BigNumber(672), new BigNumber((_389_v28).length));
          let _index65 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_394_v33).length));
          let _index66 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_384_v23).length));
          let _rhs56 = _384_v23;
          let _rhs57 = _384_v23;
          let _rhs58 = _dafny.Seq.Concat(((p0) ? ((_403_v34)[_module.__default.safeIndex(_365_v6.f27, new BigNumber((_403_v34).length))]) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-320)), ((_405_v14) => function (_406_i8) {
            return _405_v14;
          })(_374_v14)))), _358_v2);
          let _rhs59 = (_dafny.ZERO).minus(new BigNumber(-799));
          let _rhs60 = new BigNumber((_404_v35).length);
          let _lhs36 = _389_v28;
          let _lhs37 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_389_v28).length));
          let _lhs38 = _389_v28;
          let _lhs39 = _module.__default.safeIndex(new BigNumber(672), new BigNumber((_389_v28).length));
          let _lhs40 = _394_v33;
          let _lhs41 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_394_v33).length));
          let _lhs42 = globalState;
          let _lhs43 = _384_v23;
          let _lhs44 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_384_v23).length));
          _lhs36[_lhs37] = _rhs56;
          _lhs38[_lhs39] = _rhs57;
          _lhs40[_lhs41] = _rhs58;
          _lhs42.f7 = _rhs59;
          _lhs43[_lhs44] = _rhs60;
        } else {
          (globalState).f7 = _365_v6.f27;
          let _407_v36;
          _407_v36 = _dafny.Set.fromElements(_377_i4);
          let _408_v37;
          let _init13 = function (_409_i9) {
            return (_409_i9).multipliedBy(_this.f14);
          };
          let _nw88 = Array((new BigNumber(24)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw88.length); _i0_13++) {
            _nw88[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _408_v37 = _nw88;
          let _410_v38;
          _410_v38 = _dafny.Map.Empty.slice().updateUnsafe(_377_i4,_353_v0);
          let _411_v39;
          _411_v39 = _module.D9.create_DC15(p0, p0, _407_v36, _408_v37, (((_410_v38).contains(_365_v6.f27)) ? ((_410_v38).get(_365_v6.f27)) : (_353_v0)));
          r1 = ((false) ? (_this.f25) : ((_411_v39).dtor_cf19));
          (_this).f25 = (_module.D3.create_DC3(!(p1))).dtor_cf3;
          let _412_v40;
          let _init14 = ((_413_v3, _414_v6) => function (_415_i10) {
            return _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(118)), ((_416_v3) => function (_417_i11) {
              return new BigNumber((_416_v3).length);
            })(_413_v3)), _dafny.Seq.of(_414_v6.f27));
          })(_359_v3, _365_v6);
          let _nw89 = Array((new BigNumber(8)).toNumber());
          for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw89.length); _i0_14++) {
            _nw89[_i0_14] = _init14(new BigNumber(_i0_14));
          }
          _412_v40 = _nw89;
          let _418_v41;
          _418_v41 = _dafny.Seq.of(_380_v19);
          let _index67 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_412_v40).length));
          (_412_v40)[_index67] = _418_v41;
          let _419_v42;
          let _init15 = ((_420_p0) => function (_421_i12) {
            return _dafny.Map.Empty.slice().updateUnsafe(_420_p0,_this.f25);
          })(p0);
          let _nw90 = Array((new BigNumber(6)).toNumber());
          for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw90.length); _i0_15++) {
            _nw90[_i0_15] = _init15(new BigNumber(_i0_15));
          }
          _419_v42 = _nw90;
          let _index68 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_412_v40).length));
          let _rhs61 = _418_v41;
          let _rhs62 = _419_v42;
          let _rhs63 = _408_v37;
          let _rhs64 = _365_v6.f27;
          let _rhs65 = _this.f25;
          let _lhs45 = _412_v40;
          let _lhs46 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_412_v40).length));
          let _lhs47 = globalState;
          _lhs45[_lhs46] = _rhs61;
          _419_v42 = _rhs62;
          _408_v37 = _rhs63;
          _lhs47.f7 = _rhs64;
          r1 = _rhs65;
          r2 = _374_v14;
        }
        let _422_v43;
        _422_v43 = _dafny.Map.Empty.slice().updateUnsafe(_380_v19,new BigNumber(-50));
        let _423_v44;
        _423_v44 = _dafny.Set.fromElements(p1, p1, p1, p0);
        (_this).f25 = !(_377_i4).isEqualTo((((_422_v43).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(274)), function (_425_i13) {
          return new BigNumber(349);
        }))) ? ((_422_v43).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(274)), function (_424_i13) {
          return new BigNumber(349);
        }))) : (new BigNumber((_423_v44).length))));
      }
      let _426_v45;
      _426_v45 = _dafny.Set.fromElements(_this.f14, _this.f14);
      let _427_v46;
      _427_v46 = _dafny.MultiSet.fromElements(p0, p1, p1, (_this.f14).isLessThan(new BigNumber((_426_v45).length)), p0);
      let _428_v47;
      _428_v47 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(802)), ((_429_v14) => function (_430_i14) {
        return _429_v14;
      })(_374_v14)));
      let _rhs66 = (_this.f14).multipliedBy((_dafny.ZERO).minus((((_427_v46).contains(p1)) ? ((_427_v46).get(p1)) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f25,p1)).length)))));
      let _rhs67 = ((_this.f25) ? (_374_v14) : (_374_v14));
      let _rhs68 = !(p1) || ((_428_v47).IsSubsetOf(_428_v47));
      let _rhs69 = _365_v6.f27;
      let _rhs70 = _427_v46;
      let _lhs48 = globalState;
      r0 = _rhs66;
      r2 = _rhs67;
      r1 = _rhs68;
      _lhs48.f2 = _rhs69;
      _427_v46 = _rhs70;
      r0 = _365_v6.f27;
      r1 = !(!(!(_this.f25) || (!(p0))));
      r2 = _374_v14;
      r3 = _359_v3;
      return [r0, r1, r2, r3];
    }
    m14(globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.ZERO;
      let r2 = false;
      let r3 = _dafny.ZERO;
      (globalState).f7 = (_dafny.ZERO).minus(_this.f14);
      let _431_v0;
      _431_v0 = _dafny.Seq.of(_this.f25);
      let _432_v1;
      _432_v1 = _dafny.Set.fromElements(_this.f25, _this.f25);
      let _433_v2;
      _433_v2 = _dafny.MultiSet.fromElements(_this.f14);
      let _434_v3;
      _434_v3 = _dafny.Set.fromElements(_this.f14);
      let _435_v4;
      _435_v4 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_this.f14);
      let _436_v5;
      let _nw91 = Array((new BigNumber(6)).toNumber());
      _nw91[(_dafny.ZERO).toNumber()] = new BigNumber(-4);
      _nw91[(_dafny.ONE).toNumber()] = new BigNumber((_435_v4).length);
      _nw91[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(_this.f14);
      _nw91[(new BigNumber(3)).toNumber()] = new BigNumber(225);
      _nw91[(new BigNumber(4)).toNumber()] = _this.f14;
      _nw91[(new BigNumber(5)).toNumber()] = _this.f14;
      _436_v5 = _nw91;
      let _437_v6;
      let _init16 = function (_438_i0) {
        return _this.f25;
      };
      let _nw92 = Array((new BigNumber(3)).toNumber());
      for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw92.length); _i0_16++) {
        _nw92[_i0_16] = _init16(new BigNumber(_i0_16));
      }
      _437_v6 = _nw92;
      let _439_v7;
      _439_v7 = _module.D9.create_DC15(true, _this.f25, _434_v3, _436_v5, _437_v6);
      let _440_v8;
      _440_v8 = _dafny.Seq.of(_439_v7);
      let _441_v9;
      _441_v9 = _dafny.Map.Empty.slice().updateUnsafe(_440_v8,_this.f14);
      let _442_v10;
      _442_v10 = _dafny.Seq.UnicodeFromString("bgnjpywx");
      let _443_v11;
      _443_v11 = _dafny.Map.Empty.slice().updateUnsafe(_this.f25,_this.f25);
      let _444_v12;
      _444_v12 = _dafny.Seq.of(_this.f14, _this.f14, _this.f14, _this.f14);
      let _445_v13;
      let _nw93 = Array((new BigNumber(23)).toNumber());
      _nw93[(_dafny.ZERO).toNumber()] = new BigNumber((_431_v0).length);
      _nw93[(_dafny.ONE).toNumber()] = _this.f14;
      _nw93[(new BigNumber(2)).toNumber()] = (new BigNumber((_432_v1).length)).multipliedBy(new BigNumber(105));
      _nw93[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_432_v1).length));
      _nw93[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_433_v2).cardinality()));
      _nw93[(new BigNumber(5)).toNumber()] = _this.f14;
      _nw93[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_this.f14);
      _nw93[(new BigNumber(7)).toNumber()] = _this.f14;
      _nw93[(new BigNumber(8)).toNumber()] = _this.f14;
      _nw93[(new BigNumber(9)).toNumber()] = (((_441_v9).contains(_dafny.Seq.update(_440_v8, _module.__default.safeIndex(_this.f14, new BigNumber((_440_v8).length)), _439_v7))) ? ((_441_v9).get(_dafny.Seq.update(_440_v8, _module.__default.safeIndex(_this.f14, new BigNumber((_440_v8).length)), _439_v7))) : (new BigNumber((_442_v10).length)));
      _nw93[(new BigNumber(10)).toNumber()] = (_this.f14).plus(new BigNumber(907));
      _nw93[(new BigNumber(11)).toNumber()] = _this.f14;
      _nw93[(new BigNumber(12)).toNumber()] = _this.f14;
      _nw93[(new BigNumber(13)).toNumber()] = new BigNumber((_443_v11).length);
      _nw93[(new BigNumber(14)).toNumber()] = _this.f14;
      _nw93[(new BigNumber(15)).toNumber()] = (_this.f14).minus(new BigNumber((_442_v10).length));
      _nw93[(new BigNumber(16)).toNumber()] = _this.f14;
      _nw93[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_442_v10).length)));
      _nw93[(new BigNumber(18)).toNumber()] = _this.f14;
      _nw93[(new BigNumber(19)).toNumber()] = new BigNumber((_dafny.Seq.update(_431_v0, _module.__default.safeIndex(new BigNumber(202), new BigNumber((_431_v0).length)), _this.f25)).length);
      _nw93[(new BigNumber(20)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_444_v12).length), _this.f14);
      _nw93[(new BigNumber(21)).toNumber()] = _this.f14;
      _nw93[(new BigNumber(22)).toNumber()] = new BigNumber(113);
      _445_v13 = _nw93;
      let _index69 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_436_v5).length));
      (_436_v5)[_index69] = _module.__default.safeModuloInt(_this.f14, _this.f14);
      let _index70 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_436_v5).length));
      (_436_v5)[_index70] = new BigNumber((_431_v0).length);
      _439_v7 = _439_v7;
      let _hi4 = _this.f14;
      for (let _446_i1 = (_dafny.ZERO).minus(new BigNumber((_431_v0).length)); _446_i1.isLessThan(_hi4); _446_i1 = _446_i1.plus(_dafny.ONE)) {
        let _447_v14;
        let _nw94 = Array((new BigNumber(3)).toNumber()).fill([]);
        _447_v14 = _nw94;
        let _index71 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_447_v14).length));
        (_447_v14)[_index71] = _436_v5;
        let _index72 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_447_v14).length));
        (_447_v14)[_index72] = _445_v13;
        let _init17 = function (_448_i2) {
          return (_448_i2).multipliedBy(new BigNumber(444));
        };
        let _nw95 = Array((new BigNumber(13)).toNumber());
        for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw95.length); _i0_17++) {
          _nw95[_i0_17] = _init17(new BigNumber(_i0_17));
        }
        _436_v5 = _nw95;
        let _index73 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_436_v5).length));
        (_436_v5)[_index73] = (_436_v5)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_436_v5).length))];
        (globalState).f7 = new BigNumber(548);
      }
      let _449_v15;
      _449_v15 = _dafny.Map.Empty.slice().updateUnsafe(_439_v7,_432_v1);
      let _rhs71 = new BigNumber(((_module.D11.create_DC19((_435_v4).update((_dafny.ZERO).minus(_this.f14), (_436_v5)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_436_v5).length))]))).dtor_cf27).length);
      let _rhs72 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(96)), ((_450_v5) => function (_451_i3) {
        return _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("ymnrs"), _module.__default.safeIndex((_450_v5)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_450_v5).length))], new BigNumber((_dafny.Seq.UnicodeFromString("ymnrs")).length)), new _dafny.CodePoint('r'.codePointAt(0)))).length), (_450_v5)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_450_v5).length))]);
      })(_436_v5));
      let _rhs73 = _437_v6;
      let _rhs74 = _module.__default.fm43(_this.f25, globalState);
      let _rhs75 = _module.__default.fm42(true, (_436_v5)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_436_v5).length))], (new BigNumber(-970)).minus(_this.f14), (_432_v1).IsSubsetOf((((_449_v15).contains(_module.D9.create_DC15(_this.f25, _this.f25, _434_v3, _445_v13, _437_v6))) ? ((_449_v15).get(_module.D9.create_DC15(_this.f25, _this.f25, _434_v3, _445_v13, _437_v6))) : (_432_v1))), globalState);
      let _lhs49 = _this;
      _lhs49.f14 = _rhs71;
      _444_v12 = _rhs72;
      _437_v6 = _rhs73;
      r2 = _rhs74;
      _442_v10 = _rhs75;
      (_this).f14 = _module.__default.safeModuloInt((_436_v5)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_436_v5).length))], (_436_v5)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_436_v5).length))]);
      let _452_v16;
      let _init18 = function (_453_i4) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      };
      let _nw96 = Array((new BigNumber(10)).toNumber());
      for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw96.length); _i0_18++) {
        _nw96[_i0_18] = _init18(new BigNumber(_i0_18));
      }
      _452_v16 = _nw96;
      r0 = _452_v16;
      r1 = new BigNumber(-994);
      let _454_v17;
      _454_v17 = _module.D7.create_DC12();
      r2 = function (_source4) {
        if (_source4.is_DC12) {
          if (_this.f25) {
            return _this.f25;
          } else {
            return _this.f25;
          }
        } else {
          let _455___mcc_h0 = (_source4).cf15;
          let _456_cf15 = _455___mcc_h0;
          return _this.f25;
        }
      }(_454_v17);
      let _457_v18;
      _457_v18 = _dafny.Set.fromElements(_437_v6);
      let _458_v19;
      _458_v19 = new BigNumber((_457_v18).length);
      r3 = (_dafny.ZERO).minus((_458_v19));
      return [r0, r1, r2, r3];
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f16 = false;
      this._f17 = false;
      this.f24 = _module.D3.Default();
    }
    _parentTraits() {
      return [_module.T1];
    }
    get f16() {
      let _this = this;
      return _this._f16;
    };
    set f16(value) {
      let _this = this;
      _this._f16 = value;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    __ctor(f24, f16, f17) {
      let _this = this;
      (_this).f24 = f24;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      return;
    }
    fm3(globalState) {
      let _this = this;
      return ((function () {
        let _coll30 = new _dafny.Map();
        for (const _compr_30 of _dafny.IntegerRange(new BigNumber(-864), new BigNumber(515))) {
          let _459_v0 = _compr_30;
          if (((new BigNumber(-864)).isLessThanOrEqualTo(_459_v0)) && ((_459_v0).isLessThan(new BigNumber(515)))) {
            _coll30.push([(_459_v0).multipliedBy(new BigNumber(832)),new BigNumber(330)]);
          }
        }
        return _coll30;
      }()).Merge(function () {
        let _coll31 = new _dafny.Map();
        for (const _compr_31 of _dafny.IntegerRange(new BigNumber(374), new BigNumber(74))) {
          let _460_v1 = _compr_31;
          if (((new BigNumber(374)).isLessThanOrEqualTo(_460_v1)) && ((_460_v1).isLessThan(new BigNumber(74)))) {
            _coll31.push([_module.__default.safeDivisionInt(_460_v1, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("jypfe")).length), new BigNumber(105), new BigNumber((_dafny.MultiSet.fromElements((_this).f17, _this.f16)).cardinality()), new BigNumber(562), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(640)), function (_461_i0) {
              return new _dafny.CodePoint('n'.codePointAt(0));
            })).length))).length))).length)),(_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-325)))]);
          }
        }
        return _coll31;
      }())).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-423),new BigNumber((function () {
        let _coll32 = new _dafny.Set();
        for (const _compr_32 of (function () {
          let _coll33 = new _dafny.Map();
          for (const _compr_33 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(161)), function (_462_i1) {
            return new _dafny.CodePoint('x'.codePointAt(0));
          }),(_this).f17)).Keys.Elements) {
            let _463_v2 = _compr_33;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(161)), function (_462_i1) {
              return new _dafny.CodePoint('x'.codePointAt(0));
            }),(_this).f17)).contains(_463_v2)) {
              _coll33.push([_463_v2,new BigNumber(669)]);
            }
          }
          return _coll33;
        }()).Keys.Elements) {
          let _464_v3 = _compr_32;
          if ((function () {
            let _coll34 = new _dafny.Map();
            for (const _compr_34 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(161)), function (_462_i1) {
              return new _dafny.CodePoint('x'.codePointAt(0));
            }),(_this).f17)).Keys.Elements) {
              let _463_v2 = _compr_34;
              if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(161)), function (_462_i1) {
                return new _dafny.CodePoint('x'.codePointAt(0));
              }),(_this).f17)).contains(_463_v2)) {
                _coll34.push([_463_v2,new BigNumber(669)]);
              }
            }
            return _coll34;
          }()).contains(_464_v3)) {
            _coll32.add(_464_v3);
          }
        }
        return _coll32;
      }()).length)));
    };
    fm33(globalState) {
      let _this = this;
      let _source5 = _module.D6.create_DC9(new _dafny.CodePoint('k'.codePointAt(0)));
      if (_source5.is_DC10) {
        return _dafny.Seq.of(false, (_this).f17);
      } else {
        let _465___mcc_h0 = (_source5).cf14;
        let _466_cf14 = _465___mcc_h0;
        return _dafny.Seq.of(_this.f16);
      }
    };
    m2(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _467_v0;
      let _nw97 = Array((new BigNumber(19)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _467_v0 = _nw97;
      let _468_v1;
      _468_v1 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _index74 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_467_v0).length));
      (_467_v0)[_index74] = _module.__default.fm34(p0, (((_468_v1).contains(p0)) ? ((_468_v1).get(p0)) : (p0)), p1, p0, globalState);
      let _469_v2;
      _469_v2 = new _dafny.CodePoint('x'.codePointAt(0));
      let _index75 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_467_v0).length));
      (_467_v0)[_index75] = _469_v2;
      let _470_i0;
      _470_i0 = _dafny.ZERO;
      L3: {
        while (true) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_470_i0)) {
              break L3;
            }
            _470_i0 = (_470_i0).plus(_dafny.ONE);
            let _471_v3;
            let _nw98 = new _module.C1();
            _nw98.__ctor(p1, p0);
            _471_v3 = _nw98;
            let _472_v4;
            _472_v4 = _dafny.MultiSet.fromElements(_471_v3, _471_v3, _471_v3);
            let _473_v5;
            let _474_v6;
            let _out12;
            let _out13;
            let _outcollector4 = (_this).m12((_this).f17, new BigNumber((_module.__default.fm35(p0, false, p0, (_467_v0)[_module.__default.safeIndex(new BigNumber(349), new BigNumber((_467_v0).length))], globalState)).length), _472_v4, !((_this).f17), globalState);
            _out12 = _outcollector4[0];
            _out13 = _outcollector4[1];
            _473_v5 = _out12;
            _474_v6 = _out13;
            let _475_v7;
            let _nw99 = Array((new BigNumber(24)).toNumber()).fill([]);
            _475_v7 = _nw99;
            let _476_v8;
            _476_v8 = _dafny.Seq.UnicodeFromString("kiji");
            let _477_v9;
            _477_v9 = _dafny.Map.Empty.slice().updateUnsafe(_474_v6,_471_v3.f14);
            let _478_v10;
            _478_v10 = _dafny.MultiSet.fromElements(p1);
            let _479_v11;
            _479_v11 = _dafny.Set.fromElements(p0);
            let _480_v12;
            _480_v12 = _dafny.Seq.of((((_478_v10).contains(_this.f16)) ? ((_478_v10).get(_this.f16)) : (new BigNumber(419))), new BigNumber((_479_v11).length));
            let _481_v13;
            _481_v13 = _dafny.Set.fromElements((_480_v12)[_module.__default.safeIndex(_471_v3.f14, new BigNumber((_480_v12).length))]);
            let _482_v14;
            let _nw100 = Array((new BigNumber(28)).toNumber());
            _nw100[(_dafny.ZERO).toNumber()] = p0;
            _nw100[(_dafny.ONE).toNumber()] = p0;
            _nw100[(new BigNumber(2)).toNumber()] = _471_v3.f14;
            _nw100[(new BigNumber(3)).toNumber()] = new BigNumber(748);
            _nw100[(new BigNumber(4)).toNumber()] = new BigNumber(525);
            _nw100[(new BigNumber(5)).toNumber()] = p0;
            _nw100[(new BigNumber(6)).toNumber()] = p0;
            _nw100[(new BigNumber(7)).toNumber()] = new BigNumber((_468_v1).length);
            _nw100[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_476_v8).length));
            _nw100[(new BigNumber(9)).toNumber()] = p0;
            _nw100[(new BigNumber(10)).toNumber()] = new BigNumber(724);
            _nw100[(new BigNumber(11)).toNumber()] = p0;
            _nw100[(new BigNumber(12)).toNumber()] = p0;
            _nw100[(new BigNumber(13)).toNumber()] = new BigNumber((_468_v1).length);
            _nw100[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(-858));
            _nw100[(new BigNumber(15)).toNumber()] = p0;
            _nw100[(new BigNumber(16)).toNumber()] = p0;
            _nw100[(new BigNumber(17)).toNumber()] = _471_v3.f14;
            _nw100[(new BigNumber(18)).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_477_v9))).cardinality());
            _nw100[(new BigNumber(19)).toNumber()] = new BigNumber((_dafny.Set.fromElements(_471_v3.f14, _471_v3.f14, _471_v3.f14, new BigNumber((_476_v8).length), p0)).length);
            _nw100[(new BigNumber(20)).toNumber()] = p0;
            _nw100[(new BigNumber(21)).toNumber()] = new BigNumber(-302);
            _nw100[(new BigNumber(22)).toNumber()] = _471_v3.f14;
            _nw100[(new BigNumber(23)).toNumber()] = (_dafny.ZERO).minus(p0);
            _nw100[(new BigNumber(24)).toNumber()] = _471_v3.f14;
            _nw100[(new BigNumber(25)).toNumber()] = new BigNumber((_481_v13).length);
            _nw100[(new BigNumber(26)).toNumber()] = p0;
            _nw100[(new BigNumber(27)).toNumber()] = _471_v3.f14;
            _482_v14 = _nw100;
            let _index76 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_475_v7).length));
            (_475_v7)[_index76] = _482_v14;
            let _index77 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_475_v7).length));
            let _nw101 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
            (_475_v7)[_index77] = _nw101;
            let _483_v15;
            _483_v15 = _dafny.Set.fromElements((_this.f24).dtor_cf3, !(p1));
            (_this).f16 = !((_483_v15).IsDisjointFrom(_module.__default.fm44(p1, _this.f16, globalState)));
            let _484_v16;
            let _nw102 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Set.Empty);
            _484_v16 = _nw102;
            let _485_v17;
            _485_v17 = p0;
            let _486_v18;
            _486_v18 = _dafny.Map.Empty.slice().updateUnsafe(_484_v16,(_485_v17));
            _486_v18 = (_486_v18).update(_484_v16, _module.__default.fm0((_this).f17, globalState));
          }
        }
      }
      let _487_v19;
      let _nw103 = Array((new BigNumber(6)).toNumber()).fill(false);
      _487_v19 = _nw103;
      let _488_v20;
      let _nw104 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
      _488_v20 = _nw104;
      let _489_v21;
      _489_v21 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(333)),p0);
      let _490_v22;
      _490_v22 = _dafny.Seq.of((_this).f17, p1, p1);
      let _491_v23;
      _491_v23 = _487_v19;
      let _492_v24;
      _492_v24 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,p0);
      let _493_v25;
      _493_v25 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Map.Empty.slice().updateUnsafe(_492_v24,p0));
      let _494_v26;
      _494_v26 = _490_v22;
      let _rhs76 = (_487_v19);
      let _rhs77 = _module.__default.fm43(p1, globalState);
      let _rhs78 = _488_v20;
      let _rhs79 = (((_493_v25).contains(p0)) ? ((_493_v25).get(p0)) : (_489_v21));
      let _rhs80 = (_494_v26);
      let _lhs50 = _this;
      _487_v19 = _rhs76;
      _lhs50.f16 = _rhs77;
      _488_v20 = _rhs78;
      _489_v21 = _rhs79;
      _490_v22 = _rhs80;
      let _495_v27;
      _495_v27 = _dafny.Set.fromElements(_469_v2);
      _495_v27 = _module.__default.fm45((_467_v0)[_module.__default.safeIndex(new BigNumber(349), new BigNumber((_467_v0).length))], globalState);
      r0 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(p0, p0));
      let _496_v28;
      _496_v28 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f16);
      let _497_v29;
      _497_v29 = _dafny.Seq.of(_496_v28);
      let _498_i1;
      _498_i1 = _dafny.ZERO;
      L4: {
        while (function (_source6) {
          if (_source6.is_DC12) {
            if ((_this).f17) {
              return (_this).f17;
            } else {
              return _this.f16;
            }
          } else {
            let _516___mcc_h0 = (_source6).cf15;
            let _517_cf15 = _516___mcc_h0;
            return (_this).f17;
          }
        }(_module.D7.create_DC11((((_497_v29)[_module.__default.safeIndex(p0, new BigNumber((_497_v29).length))]).update(p0, (_this).f17)).update(p0, _this.f16)))) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_498_i1)) {
              break L4;
            }
            _498_i1 = (_498_i1).plus(_dafny.ONE);
            let _499_v30;
            let _nw105 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Set.Empty);
            _499_v30 = _nw105;
            let _500_v31;
            let _nw106 = new _module.C0();
            _nw106.__ctor(((false) ? (_499_v30) : (_499_v30)), new BigNumber(-730));
            _500_v31 = _nw106;
            let _501_v32;
            _501_v32 = _dafny.Map.Empty.slice().updateUnsafe(_469_v2,(_dafny.ZERO).minus(p0));
            _501_v32 = ((_module.__default.fm46(_this.f16, false, globalState)).Merge(_501_v32)).Merge(_501_v32);
            if (p1) {
              let _502_v33;
              _502_v33 = _dafny.Seq.UnicodeFromString("ftj");
              let _503_v34;
              let _nw107 = Array((new BigNumber(9)).toNumber());
              _nw107[(_dafny.ZERO).toNumber()] = _502_v33;
              _nw107[(_dafny.ONE).toNumber()] = _502_v33;
              _nw107[(new BigNumber(2)).toNumber()] = _502_v33;
              _nw107[(new BigNumber(3)).toNumber()] = _502_v33;
              _nw107[(new BigNumber(4)).toNumber()] = _502_v33;
              _nw107[(new BigNumber(5)).toNumber()] = _502_v33;
              _nw107[(new BigNumber(6)).toNumber()] = _502_v33;
              _nw107[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_502_v33, _dafny.Seq.UnicodeFromString("utmfmeu"));
              _nw107[(new BigNumber(8)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(925)), ((_504_v2) => function (_505_i2) {
                return (_module.D6.create_DC9(_504_v2)).dtor_cf14;
              })(_469_v2));
              _503_v34 = _nw107;
              _503_v34 = _503_v34;
              let _506_v35;
              _506_v35 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,(_467_v0)[_module.__default.safeIndex(new BigNumber(349), new BigNumber((_467_v0).length))]);
              let _507_v36;
              _507_v36 = _dafny.Map.Empty.slice().updateUnsafe(_506_v35,(_500_v31.f27).plus(p0));
              _507_v36 = (_507_v36).Merge(_507_v36);
              let _508_v37;
              let _nw108 = Array((new BigNumber(4)).toNumber()).fill([]);
              _508_v37 = _nw108;
              let _index78 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_508_v37).length));
              (_508_v37)[_index78] = _503_v34;
              let _index79 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_508_v37).length));
              (_508_v37)[_index79] = _503_v34;
              let _509_v38;
              _509_v38 = _module.D3.create_DC4(!((_module.D3.create_DC3(p1)).dtor_cf3), _500_v31.f27, new BigNumber((_496_v28).length), true, true);
              _509_v38 = _509_v38;
              (_this).f16 = !((p0).isLessThanOrEqualTo(new BigNumber(-834)));
            } else {
              (_this).f16 = ((p1) ? ((_this).f17) : ((_this).f17));
              (_500_v31).f27 = (_dafny.ZERO).minus(new BigNumber(-26));
              let _510_v39;
              _510_v39 = _dafny.Seq.of(_491_v23);
              let _511_v40;
              _511_v40 = _dafny.Set.fromElements(_492_v24, _492_v24);
              let _rhs81 = ((_dafny.Set.fromElements(_492_v24, _492_v24)).Difference(_511_v40)).IsDisjointFrom(_511_v40);
              let _rhs82 = _dafny.Seq.Concat(_510_v39, _510_v39);
              let _lhs51 = _this;
              _lhs51.f16 = _rhs81;
              _510_v39 = _rhs82;
              let _512_v41;
              let _nw109 = Array((new BigNumber(9)).toNumber());
              _512_v41 = _nw109;
              _512_v41 = _512_v41;
              (globalState).f7 = ((_dafny.ZERO).minus(_500_v31.f27)).plus((new BigNumber(372)).plus((((_468_v1).contains(p0)) ? ((_468_v1).get(p0)) : (_500_v31.f27))));
            }
            let _513_v42;
            _513_v42 = _dafny.MultiSet.fromElements((_this).f17);
            let _514_v43;
            _514_v43 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,p1)).length), _500_v31.f27);
            let _515_v44;
            _515_v44 = _dafny.Map.Empty.slice().updateUnsafe(_500_v31.f27,_514_v43);
            (globalState).f7 = _module.__default.safeModuloInt(new BigNumber((_513_v42).cardinality()), ((_514_v43)[_module.__default.safeIndex(new BigNumber((_515_v44).length), new BigNumber((_514_v43).length))]).multipliedBy((_514_v43)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_514_v43).length))]));
          }
        }
      }
      r0 = _module.__default.safeModuloInt(_module.__default.fm0(p1, globalState), _module.__default.fm0(_this.f16, globalState));
      return r0;
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _518_v0;
      _518_v0 = _dafny.Seq.of(true);
      let _519_v1;
      _519_v1 = _dafny.Map.Empty.slice().updateUnsafe(p0,_518_v0);
      (_this).f16 = (new BigNumber((_519_v1).length)).isEqualTo(p3);
      let _520_v2;
      let _nw110 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
      _520_v2 = _nw110;
      _520_v2 = _520_v2;
      let _521_v3;
      let _nw111 = Array((new BigNumber(25)).toNumber());
      _nw111[(_dafny.ZERO).toNumber()] = p2;
      _nw111[(_dafny.ONE).toNumber()] = _this.f16;
      _nw111[(new BigNumber(2)).toNumber()] = (_this.f24).dtor_cf3;
      _nw111[(new BigNumber(3)).toNumber()] = (_this).f17;
      _nw111[(new BigNumber(4)).toNumber()] = true;
      _nw111[(new BigNumber(5)).toNumber()] = p2;
      _nw111[(new BigNumber(6)).toNumber()] = _this.f16;
      _nw111[(new BigNumber(7)).toNumber()] = _this.f16;
      _nw111[(new BigNumber(8)).toNumber()] = false;
      _nw111[(new BigNumber(9)).toNumber()] = p2;
      _nw111[(new BigNumber(10)).toNumber()] = _this.f16;
      _nw111[(new BigNumber(11)).toNumber()] = (_this).f17;
      _nw111[(new BigNumber(12)).toNumber()] = _this.f16;
      _nw111[(new BigNumber(13)).toNumber()] = p2;
      _nw111[(new BigNumber(14)).toNumber()] = p2;
      _nw111[(new BigNumber(15)).toNumber()] = _this.f16;
      _nw111[(new BigNumber(16)).toNumber()] = p2;
      _nw111[(new BigNumber(17)).toNumber()] = p2;
      _nw111[(new BigNumber(18)).toNumber()] = (_this).f17;
      _nw111[(new BigNumber(19)).toNumber()] = (_this).f17;
      _nw111[(new BigNumber(20)).toNumber()] = !((_this).f17);
      _nw111[(new BigNumber(21)).toNumber()] = _this.f16;
      _nw111[(new BigNumber(22)).toNumber()] = p2;
      _nw111[(new BigNumber(23)).toNumber()] = _this.f16;
      _nw111[(new BigNumber(24)).toNumber()] = p2;
      _521_v3 = _nw111;
      let _522_v4;
      let _nw112 = Array((new BigNumber(25)).toNumber());
      _nw112[(_dafny.ZERO).toNumber()] = _521_v3;
      _nw112[(_dafny.ONE).toNumber()] = _521_v3;
      _nw112[(new BigNumber(2)).toNumber()] = _521_v3;
      _nw112[(new BigNumber(3)).toNumber()] = _521_v3;
      _nw112[(new BigNumber(4)).toNumber()] = _521_v3;
      _nw112[(new BigNumber(5)).toNumber()] = _521_v3;
      _nw112[(new BigNumber(6)).toNumber()] = _521_v3;
      _nw112[(new BigNumber(7)).toNumber()] = _521_v3;
      _nw112[(new BigNumber(8)).toNumber()] = _521_v3;
      _nw112[(new BigNumber(9)).toNumber()] = _521_v3;
      _nw112[(new BigNumber(10)).toNumber()] = _521_v3;
      _nw112[(new BigNumber(11)).toNumber()] = _521_v3;
      _nw112[(new BigNumber(12)).toNumber()] = _521_v3;
      _nw112[(new BigNumber(13)).toNumber()] = _521_v3;
      _nw112[(new BigNumber(14)).toNumber()] = _521_v3;
      _nw112[(new BigNumber(15)).toNumber()] = _521_v3;
      _nw112[(new BigNumber(16)).toNumber()] = _521_v3;
      _nw112[(new BigNumber(17)).toNumber()] = _521_v3;
      _nw112[(new BigNumber(18)).toNumber()] = _521_v3;
      _nw112[(new BigNumber(19)).toNumber()] = _521_v3;
      _nw112[(new BigNumber(20)).toNumber()] = _521_v3;
      _nw112[(new BigNumber(21)).toNumber()] = _521_v3;
      _nw112[(new BigNumber(22)).toNumber()] = _521_v3;
      _nw112[(new BigNumber(23)).toNumber()] = _521_v3;
      _nw112[(new BigNumber(24)).toNumber()] = _521_v3;
      _522_v4 = _nw112;
      _522_v4 = _522_v4;
      let _hi5 = new BigNumber(71);
      for (let _523_i0 = new BigNumber(815); _523_i0.isLessThan(_hi5); _523_i0 = _523_i0.plus(_dafny.ONE)) {
        let _524_v5;
        let _nw113 = new _module.C1();
        _nw113.__ctor(p2, p3);
        _524_v5 = _nw113;
        let _525_v6;
        let _nw114 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Set.Empty);
        _525_v6 = _nw114;
        let _526_v7;
        _526_v7 = _dafny.Set.fromElements(p0);
        let _index80 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_525_v6).length));
        (_525_v6)[_index80] = _526_v7;
        let _index81 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_525_v6).length));
        (_525_v6)[_index81] = (_526_v7).Union(_dafny.Set.fromElements(p0, p0));
        (_this).f16 = _module.__default.fm43((_this).f17, globalState);
        let _index82 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((_520_v2).length));
        (_520_v2)[_index82] = (_dafny.ZERO).minus(_523_i0);
        let _527_v8;
        _527_v8 = _module.__default.safeDivisionInt(new BigNumber(590), p3);
        let _528_v9;
        _528_v9 = _dafny.Map.Empty.slice().updateUnsafe(p3,_521_v3);
        let _529_v10;
        _529_v10 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(p3,p2));
        let _530_v11;
        _530_v11 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-48),_524_v5.f25);
        let _531_v12;
        _531_v12 = _dafny.Seq.UnicodeFromString("ktd");
        let _index83 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((_520_v2).length));
        let _rhs83 = new BigNumber((_528_v9).length);
        let _rhs84 = !(!((_529_v10)[_module.__default.safeIndex(p1, new BigNumber((_529_v10).length))]).equals(_530_v11));
        let _rhs85 = new BigNumber((_531_v12).length);
        let _rhs86 = (_dafny.ZERO).minus(p1);
        let _rhs87 = _module.__default.fm47((p1).multipliedBy(p3), globalState);
        let _lhs52 = globalState;
        let _lhs53 = _524_v5;
        let _lhs54 = globalState;
        let _lhs55 = _520_v2;
        let _lhs56 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((_520_v2).length));
        _lhs52.f2 = _rhs83;
        _lhs53.f25 = _rhs84;
        _lhs54.f2 = _rhs85;
        _lhs55[_lhs56] = _rhs86;
        _527_v8 = _rhs87;
      }
      r0 = _module.__default.safeDivisionInt(p1, p3);
      let _hi6 = p3;
      for (let _532_i1 = p3; _532_i1.isLessThan(_hi6); _532_i1 = _532_i1.plus(_dafny.ONE)) {
        let _533_v13;
        let _nw115 = Array((new BigNumber(5)).toNumber());
        _nw115[(_dafny.ZERO).toNumber()] = p0;
        _nw115[(_dafny.ONE).toNumber()] = p0;
        _nw115[(new BigNumber(2)).toNumber()] = p0;
        _nw115[(new BigNumber(3)).toNumber()] = new _dafny.CodePoint('i'.codePointAt(0));
        _nw115[(new BigNumber(4)).toNumber()] = p0;
        _533_v13 = _nw115;
        let _index84 = _module.__default.safeIndex(new BigNumber(521), new BigNumber((_533_v13).length));
        (_533_v13)[_index84] = p0;
        let _index85 = _module.__default.safeIndex(new BigNumber(521), new BigNumber((_533_v13).length));
        (_533_v13)[_index85] = p0;
        let _534_v14;
        _534_v14 = _dafny.MultiSet.fromElements(p2, (_this).f17, p2);
        let _535_v15;
        _535_v15 = _dafny.Seq.of(_534_v14);
        let _index86 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_520_v2).length));
        (_520_v2)[_index86] = _532_i1;
        let _index87 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_520_v2).length));
        let _rhs88 = _dafny.Seq.Concat(_535_v15, _dafny.Seq.Concat(_535_v15, _535_v15));
        let _rhs89 = new BigNumber(321);
        let _rhs90 = (new BigNumber(648)).minus(p1);
        let _rhs91 = (p1).plus(new BigNumber(655));
        let _lhs57 = _520_v2;
        let _lhs58 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_520_v2).length));
        _535_v15 = _rhs88;
        _lhs57[_lhs58] = _rhs89;
        r0 = _rhs90;
        r0 = _rhs91;
        let _536_v16;
        _536_v16 = _dafny.Map.Empty.slice().updateUnsafe(false,p0);
        if ((((_536_v16).contains((_this).f17)) ? ((_this).f17) : ((_518_v0)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(p3, p3))).cardinality()), new BigNumber((_518_v0).length))]))) {
          (globalState).f7 = _module.__default.safeModuloInt((_520_v2)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_520_v2).length))], ((_this.f16) ? (p1) : (new BigNumber(-740))));
          let _537_v17;
          _537_v17 = _dafny.MultiSet.fromElements(p1);
          let _538_v18;
          _538_v18 = _dafny.Map.Empty.slice().updateUnsafe((p3).plus(p1),_537_v17);
          _538_v18 = _538_v18;
          let _539_v19;
          _539_v19 = _dafny.Seq.UnicodeFromString("awfmvcj");
          _539_v19 = ((!(false) || (_this.f16)) ? (_539_v19) : (_539_v19));
          (_this).f16 = (_518_v0)[_module.__default.safeIndex((_520_v2)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_520_v2).length))], new BigNumber((_518_v0).length))];
          let _540_v20;
          _540_v20 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)));
          let _541_v21;
          let _nw116 = new _module.C1();
          _nw116.__ctor(p2, p3);
          _541_v21 = _nw116;
          let _542_v22;
          _542_v22 = _dafny.MultiSet.fromElements(_541_v21);
          let _543_v23;
          _543_v23 = _dafny.Set.fromElements((_this).f17);
          let _544_v24;
          _544_v24 = _dafny.Seq.of(new BigNumber((_543_v23).length), _541_v21.f14);
          let _545_v25;
          _545_v25 = _dafny.Set.fromElements((_520_v2)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_520_v2).length))], p3);
          let _546_v26;
          let _nw117 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
          _546_v26 = _nw117;
          let _547_v27;
          _547_v27 = _module.D9.create_DC15(p2, p2, _545_v25, _546_v26, _521_v3);
          let _548_v28;
          let _549_v29;
          let _out14;
          let _out15;
          let _outcollector5 = (_this).m12(((_this.f16) ? ((_this).f17) : (_this.f16)), (_module.__default.fm0(_this.f16, globalState)).multipliedBy(new BigNumber((_540_v20).cardinality())), ((_542_v22).update(_541_v21, _module.__default.abs(new BigNumber((_544_v24).length)))).Intersect(_542_v22), (_545_v25).IsProperSubsetOf((_547_v27).dtor_cf20), globalState);
          _out14 = _outcollector5[0];
          _out15 = _outcollector5[1];
          _548_v28 = _out14;
          _549_v29 = _out15;
        } else {
          let _550_v30;
          _550_v30 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_533_v13)[_module.__default.safeIndex(new BigNumber(521), new BigNumber((_533_v13).length))]);
          let _551_v31;
          _551_v31 = _dafny.Seq.UnicodeFromString("atptjlkbt");
          let _552_v32;
          _552_v32 = _dafny.Seq.of(p1, _532_i1, new BigNumber((_551_v31).length));
          let _553_v33;
          _553_v33 = _dafny.Map.Empty.slice().updateUnsafe(_550_v30,_552_v32);
          _553_v33 = _553_v33;
          let _554_v34;
          _554_v34 = _dafny.Set.fromElements(p2);
          let _555_v35;
          let _nw118 = Array((new BigNumber(13)).toNumber());
          _nw118[(_dafny.ZERO).toNumber()] = _dafny.Set.fromElements(!((_this).f17), _this.f16, _this.f16, p2, (_this).f17);
          _nw118[(_dafny.ONE).toNumber()] = _554_v34;
          _nw118[(new BigNumber(2)).toNumber()] = _554_v34;
          _nw118[(new BigNumber(3)).toNumber()] = _dafny.Set.fromElements(p2, (_this).f17);
          _nw118[(new BigNumber(4)).toNumber()] = _554_v34;
          _nw118[(new BigNumber(5)).toNumber()] = _dafny.Set.fromElements(p2, p2);
          _nw118[(new BigNumber(6)).toNumber()] = _554_v34;
          _nw118[(new BigNumber(7)).toNumber()] = _554_v34;
          _nw118[(new BigNumber(8)).toNumber()] = _554_v34;
          _nw118[(new BigNumber(9)).toNumber()] = _554_v34;
          _nw118[(new BigNumber(10)).toNumber()] = _554_v34;
          _nw118[(new BigNumber(11)).toNumber()] = _554_v34;
          _nw118[(new BigNumber(12)).toNumber()] = _554_v34;
          _555_v35 = _nw118;
          let _556_v36;
          let _nw119 = new _module.C0();
          _nw119.__ctor(_555_v35, (_532_i1).multipliedBy(p3));
          _556_v36 = _nw119;
          let _557_v37;
          let _init19 = ((_558_v0) => function (_559_i2) {
            return (_559_i2).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_558_v0).length),_this.f16)).length));
          })(_518_v0);
          let _nw120 = Array((new BigNumber(12)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw120.length); _i0_19++) {
            _nw120[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _557_v37 = _nw120;
          _557_v37 = _557_v37;
          let _560_v38;
          _560_v38 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_520_v2)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_520_v2).length))]),p2);
          let _561_v39;
          _561_v39 = _dafny.MultiSet.fromElements(_module.__default.fm42((((_560_v38).contains(_532_i1)) ? ((_560_v38).get(_532_i1)) : (true)), _556_v36.f27, _556_v36.f27, _this.f16, globalState));
          _561_v39 = _561_v39;
          let _index88 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_521_v3).length));
          (_521_v3)[_index88] = _module.__default.fm43(_this.f16, globalState);
          let _562_v40;
          _562_v40 = _dafny.Seq.of(_521_v3, _521_v3);
          let _563_v41;
          _563_v41 = (_562_v40)[_module.__default.safeIndex((_520_v2)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_520_v2).length))], new BigNumber((_562_v40).length))];
          let _index89 = _module.__default.safeIndex(new BigNumber(521), new BigNumber((_533_v13).length));
          let _index90 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_521_v3).length));
          let _index91 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_520_v2).length));
          let _rhs92 = p0;
          let _rhs93 = p2;
          let _rhs94 = (_dafny.ZERO).minus(new BigNumber(((_module.D12.create_DC22(_dafny.Seq.UnicodeFromString("bsgrpsbcd"))).dtor_cf29).length));
          let _rhs95 = _521_v3;
          let _lhs59 = _533_v13;
          let _lhs60 = _module.__default.safeIndex(new BigNumber(521), new BigNumber((_533_v13).length));
          let _lhs61 = _521_v3;
          let _lhs62 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_521_v3).length));
          let _lhs63 = _520_v2;
          let _lhs64 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_520_v2).length));
          _lhs59[_lhs60] = _rhs92;
          _lhs61[_lhs62] = _rhs93;
          _lhs63[_lhs64] = _rhs94;
          _563_v41 = _rhs95;
        }
        (globalState).f2 = p1;
      }
      r0 = (_module.__default.fm0(p2, globalState)).minus(p3);
      return r0;
    }
    m12(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = false;
      let _564_v0;
      _564_v0 = _dafny.Map.Empty.slice().updateUnsafe(true,p0);
      let _565_v1;
      _565_v1 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_564_v0);
      let _566_v2;
      _566_v2 = _565_v1;
      _566_v2 = _566_v2;
      r1 = p0;
      if (_module.__default.fm43(p3, globalState)) {
        let _567_v3;
        let _nw121 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _567_v3 = _nw121;
        let _568_v4;
        _568_v4 = _dafny.Seq.UnicodeFromString("yktx");
        let _index92 = _module.__default.safeIndex(new BigNumber(545), new BigNumber((_567_v3).length));
        (_567_v3)[_index92] = _dafny.Seq.Concat(_568_v4, _dafny.Seq.UnicodeFromString("i"));
        let _569_v5;
        _569_v5 = new _dafny.CodePoint('s'.codePointAt(0));
        let _index93 = _module.__default.safeIndex(new BigNumber(545), new BigNumber((_567_v3).length));
        (_567_v3)[_index93] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ilx"), _dafny.Seq.Concat(_dafny.Seq.update(_568_v4, _module.__default.safeIndex(p1, new BigNumber((_568_v4).length)), _569_v5), _dafny.Seq.UnicodeFromString("t")));
        r1 = (_this).f17;
        let _570_v6;
        let _nw122 = Array((_dafny.ONE).toNumber()).fill(_dafny.Set.Empty);
        _570_v6 = _nw122;
        let _571_v7;
        let _nw123 = new _module.C0();
        _nw123.__ctor(_570_v6, p1);
        _571_v7 = _nw123;
        let _572_v9;
        _572_v9 = _dafny.Seq.of(_dafny.Seq.of(p0, !(p0), (_this).f17));
        let _573_v10;
        _573_v10 = _dafny.MultiSet.fromElements((_this).f17);
        (globalState).f7 = (new BigNumber((function () {
          let _coll35 = new _dafny.Map();
          for (const _compr_35 of (_572_v9).Elements) {
            let _574_v8 = _compr_35;
            if (_dafny.Seq.contains(_572_v9, _574_v8)) {
              _coll35.push([_574_v8,false]);
            }
          }
          return _coll35;
        }()).length)).minus((((_573_v10).contains(true)) ? ((_573_v10).get(true)) : (p1)));
        if ((_this).f17) {
          let _575_v11;
          _575_v11 = _dafny.MultiSet.fromElements(new BigNumber(-551));
          let _576_v12;
          _576_v12 = _dafny.Seq.of(_575_v11);
          let _577_v13;
          let _nw124 = Array((new BigNumber(12)).toNumber());
          _nw124[(_dafny.ZERO).toNumber()] = (_this).f17;
          _nw124[(_dafny.ONE).toNumber()] = !((_this).f17);
          _nw124[(new BigNumber(2)).toNumber()] = p3;
          _nw124[(new BigNumber(3)).toNumber()] = p3;
          _nw124[(new BigNumber(4)).toNumber()] = (_this.f24).dtor_cf3;
          _nw124[(new BigNumber(5)).toNumber()] = (_this).f17;
          _nw124[(new BigNumber(6)).toNumber()] = p0;
          _nw124[(new BigNumber(7)).toNumber()] = false;
          _nw124[(new BigNumber(8)).toNumber()] = p0;
          _nw124[(new BigNumber(9)).toNumber()] = p0;
          _nw124[(new BigNumber(10)).toNumber()] = p0;
          _nw124[(new BigNumber(11)).toNumber()] = _this.f16;
          _577_v13 = _nw124;
          let _578_v14;
          _578_v14 = _dafny.Map.Empty.slice().updateUnsafe(_577_v13,p1);
          let _579_v15;
          _579_v15 = _dafny.Set.fromElements(!((_this).f17));
          let _580_v16;
          _580_v16 = _dafny.Set.fromElements(new BigNumber((_575_v11).cardinality()), (new BigNumber((_579_v15).length)).minus(p1));
          let _581_v17;
          _581_v17 = _dafny.Map.Empty.slice().updateUnsafe(_571_v7.f27,(_571_v7.f27).plus(new BigNumber(784)));
          let _582_v18;
          _582_v18 = _dafny.Seq.of(new BigNumber(951));
          let _583_v20;
          _583_v20 = _dafny.Map.Empty.slice().updateUnsafe(_571_v7.f27,p0);
          let _rhs96 = _576_v12;
          let _rhs97 = ((_578_v14).Merge(_578_v14)).Merge(_578_v14);
          let _rhs98 = (_564_v0).update(_this.f16, (_this).f17);
          let _rhs99 = ((_580_v16).Intersect(_580_v16)).Intersect((_580_v16).Union(_580_v16));
          let _rhs100 = ((_581_v17).update((_582_v18)[_module.__default.safeIndex(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(p1,_577_v13)).update(p1, _577_v13)).length), new BigNumber((_582_v18).length))], new BigNumber(831))).Merge(function () {
            let _coll36 = new _dafny.Map();
            for (const _compr_36 of (_583_v20).Keys.Elements) {
              let _584_v19 = _compr_36;
              if ((_583_v20).contains(_584_v19)) {
                _coll36.push([_module.__default.safeDivisionInt(_584_v19, (((_573_v10).contains((_this).f17)) ? ((_573_v10).get((_this).f17)) : (new BigNumber((_568_v4).length)))),_571_v7.f27]);
              }
            }
            return _coll36;
          }());
          _576_v12 = _rhs96;
          _578_v14 = _rhs97;
          _564_v0 = _rhs98;
          _580_v16 = _rhs99;
          _581_v17 = _rhs100;
          let _585_v21;
          let _nw125 = new _module.C0();
          _nw125.__ctor(_570_v6, _571_v7.f27);
          _585_v21 = _nw125;
          (_this).f16 = !(_this.f16);
          let _586_v22;
          _586_v22 = _dafny.Map.Empty.slice().updateUnsafe((_567_v3)[_module.__default.safeIndex(new BigNumber(545), new BigNumber((_567_v3).length))],!(!(!(!(false)) || (p0))));
          r1 = (((_586_v22).contains((_567_v3)[_module.__default.safeIndex(new BigNumber(545), new BigNumber((_567_v3).length))])) ? ((_586_v22).get((_567_v3)[_module.__default.safeIndex(new BigNumber(545), new BigNumber((_567_v3).length))])) : ((_this).f17));
          let _587_v23;
          _587_v23 = _dafny.MultiSet.fromElements(_569_v5, _569_v5);
          (_this).f16 = ((_this.f16) ? (!(_587_v23).contains(_569_v5)) : (p0));
        } else {
          (globalState).f7 = p1;
          let _588_v24;
          _588_v24 = _module.D9.create_DC16(_571_v7.f27, p0);
          let _589_v25;
          _589_v25 = _dafny.Seq.of((_this).f17, _this.f16);
          let _590_v26;
          let _nw126 = Array((new BigNumber(6)).toNumber());
          _nw126[(_dafny.ZERO).toNumber()] = p1;
          _nw126[(_dafny.ONE).toNumber()] = (_588_v24).dtor_cf23;
          _nw126[(new BigNumber(2)).toNumber()] = p1;
          _nw126[(new BigNumber(3)).toNumber()] = new BigNumber((_589_v25).length);
          _nw126[(new BigNumber(4)).toNumber()] = p1;
          _nw126[(new BigNumber(5)).toNumber()] = p1;
          _590_v26 = _nw126;
          let _591_v27;
          _591_v27 = _dafny.Map.Empty.slice().updateUnsafe((p1).multipliedBy(new BigNumber((_568_v4).length)),_590_v26);
          _591_v27 = (_591_v27).update(new BigNumber((_589_v25).length), _590_v26);
          let _592_v28;
          _592_v28 = _dafny.Map.Empty.slice().updateUnsafe(true,_568_v4);
          let _593_v29;
          _593_v29 = _dafny.MultiSet.fromElements((((_592_v28).contains(p3)) ? ((_592_v28).get(p3)) : (_568_v4)), (_567_v3)[_module.__default.safeIndex(new BigNumber(545), new BigNumber((_567_v3).length))], _dafny.Seq.UnicodeFromString("imiolrx"));
          let _594_v30;
          _594_v30 = _dafny.Map.Empty.slice().updateUnsafe(_571_v7.f27,_571_v7.f27);
          let _595_v31;
          _595_v31 = _dafny.Map.Empty.slice().updateUnsafe((_593_v29).IsSubsetOf(_593_v29),_594_v30);
          _595_v31 = (_595_v31).update(_this.f16, _594_v30);
          let _596_v32;
          _596_v32 = _589_v25;
          let _597_v33;
          _597_v33 = _dafny.Set.fromElements(!(p3));
          let _598_v34;
          _598_v34 = _dafny.Set.fromElements(new BigNumber((_597_v33).length), _571_v7.f27, new BigNumber(627));
          let _599_v35;
          _599_v35 = _module.D4.create_DC7(_589_v25, new BigNumber((_598_v34).length));
          let _600_v36;
          _600_v36 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("rxnylel"),p1);
          let _601_v37;
          _601_v37 = _module.D13.create_DC25(_600_v36);
          let _602_v38;
          _602_v38 = _dafny.Map.Empty.slice().updateUnsafe(_564_v0,new BigNumber(((_601_v37).dtor_cf38).length));
          let _603_v39;
          _603_v39 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_599_v35).dtor_cf12),_602_v38);
          _603_v39 = (_603_v39).update(p1, (_module.__default.fm48(globalState)).update(_564_v0, _571_v7.f27));
          let _index94 = _module.__default.safeIndex(new BigNumber(545), new BigNumber((_567_v3).length));
          (_567_v3)[_index94] = _dafny.Seq.Concat(_dafny.Seq.Concat(_568_v4, (_567_v3)[_module.__default.safeIndex(new BigNumber(545), new BigNumber((_567_v3).length))]), _dafny.Seq.UnicodeFromString("ouxahjk"));
        }
      } else {
        let _604_v40;
        let _nw127 = new _module.C1();
        _nw127.__ctor(_this.f16, p1);
        _604_v40 = _nw127;
        let _605_v41;
        let _nw128 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _605_v41 = _nw128;
        let _606_v42;
        _606_v42 = _dafny.Seq.UnicodeFromString("k");
        let _index95 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_605_v41).length));
        (_605_v41)[_index95] = _606_v42;
        let _607_v43;
        _607_v43 = _dafny.Seq.of((_this).f17);
        let _608_v44;
        _608_v44 = _dafny.Set.fromElements(new BigNumber((_606_v42).length), _module.__default.fm0((_this).f17, globalState));
        let _609_v45;
        _609_v45 = _dafny.MultiSet.fromElements(_608_v44, _608_v44);
        let _610_v46;
        _610_v46 = _dafny.MultiSet.fromElements(_604_v40.f25, (_this).f17, _604_v40.f25);
        let _611_v47;
        let _nw129 = Array((new BigNumber(3)).toNumber());
        _nw129[(_dafny.ZERO).toNumber()] = false;
        _nw129[(_dafny.ONE).toNumber()] = (_610_v46).IsSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.update(_607_v43, _module.__default.safeIndex(new BigNumber((_609_v45).cardinality()), new BigNumber((_607_v43).length)), _this.f16)));
        _nw129[(new BigNumber(2)).toNumber()] = (p0) || (p0);
        _611_v47 = _nw129;
        let _index96 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_611_v47).length));
        (_611_v47)[_index96] = (_dafny.MultiSet.FromArray(_607_v43)).IsDisjointFrom(_610_v46);
        let _612_v48;
        _612_v48 = new _dafny.CodePoint('u'.codePointAt(0));
        let _613_v49;
        _613_v49 = _dafny.Map.Empty.slice().updateUnsafe(_612_v48,p1);
        let _614_v50;
        _614_v50 = _dafny.Seq.of(p1, p1);
        let _index97 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_605_v41).length));
        let _index98 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_611_v47).length));
        let _rhs101 = _dafny.Seq.UnicodeFromString("hbphfkp");
        let _rhs102 = (_module.__default.fm0(!(p3), globalState)).isLessThanOrEqualTo(p1);
        let _rhs103 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((((_613_v49).contains(new _dafny.CodePoint('m'.codePointAt(0)))) ? ((_613_v49).get(new _dafny.CodePoint('m'.codePointAt(0)))) : (p1)), (_614_v50)[_module.__default.safeIndex(p1, new BigNumber((_614_v50).length))]));
        let _rhs104 = new BigNumber((_606_v42).length);
        let _lhs65 = _605_v41;
        let _lhs66 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_605_v41).length));
        let _lhs67 = _611_v47;
        let _lhs68 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_611_v47).length));
        let _lhs69 = globalState;
        let _lhs70 = globalState;
        _lhs65[_lhs66] = _rhs101;
        _lhs67[_lhs68] = _rhs102;
        _lhs69.f2 = _rhs103;
        _lhs70.f7 = _rhs104;
        if (_604_v40.f25) {
          let _615_v51;
          let _init20 = ((_616_v40) => function (_617_i0) {
            return _dafny.Set.fromElements(_616_v40.f25);
          })(_604_v40);
          let _nw130 = Array((new BigNumber(2)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw130.length); _i0_20++) {
            _nw130[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _615_v51 = _nw130;
          let _618_v52;
          let _nw131 = new _module.C0();
          _nw131.__ctor(_615_v51, (_dafny.ZERO).minus(p1));
          _618_v52 = _nw131;
          let _619_v53;
          _619_v53 = _dafny.Set.fromElements(p3, true);
          let _620_v54;
          _620_v54 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm42(_604_v40.f25, p1, p1, _604_v40.f25, globalState),(_619_v53).Union(_619_v53));
          _620_v54 = (_620_v54).update((_605_v41)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_605_v41).length))], _619_v53);
          let _621_v55;
          _621_v55 = _dafny.MultiSet.fromElements(_606_v42, (_605_v41)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_605_v41).length))]);
          let _622_v56;
          _622_v56 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,p1);
          let _623_v57;
          _623_v57 = _dafny.Map.Empty.slice().updateUnsafe(!((_611_v47)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_611_v47).length))]),new _dafny.CodePoint('k'.codePointAt(0)));
          let _624_v58;
          _624_v58 = _dafny.Seq.of((_this).f17, _604_v40.f25, true);
          let _625_v59;
          _625_v59 = _dafny.Map.Empty.slice().updateUnsafe(_618_v52.f27,new BigNumber((_610_v46).cardinality()));
          let _626_v60;
          let _nw132 = Array((new BigNumber(22)).toNumber());
          _nw132[(_dafny.ZERO).toNumber()] = new BigNumber((_621_v55).cardinality());
          _nw132[(_dafny.ONE).toNumber()] = _618_v52.f27;
          _nw132[(new BigNumber(2)).toNumber()] = _618_v52.f27;
          _nw132[(new BigNumber(3)).toNumber()] = new BigNumber(((_605_v41)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_605_v41).length))]).length);
          _nw132[(new BigNumber(4)).toNumber()] = (((_622_v56).contains((_this).f17)) ? ((_622_v56).get((_this).f17)) : (p1));
          _nw132[(new BigNumber(5)).toNumber()] = (((_613_v49).contains(_612_v48)) ? ((_613_v49).get(_612_v48)) : (p1));
          _nw132[(new BigNumber(6)).toNumber()] = (new BigNumber((_dafny.Seq.update((_605_v41)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_605_v41).length))], _module.__default.safeIndex(new BigNumber((_623_v57).length), new BigNumber(((_605_v41)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_605_v41).length))]).length)), _612_v48)).length)).multipliedBy(_618_v52.f27);
          _nw132[(new BigNumber(7)).toNumber()] = new BigNumber(((_624_v58)).length);
          _nw132[(new BigNumber(8)).toNumber()] = (_618_v52.f27).multipliedBy(_618_v52.f27);
          _nw132[(new BigNumber(9)).toNumber()] = _618_v52.f27;
          _nw132[(new BigNumber(10)).toNumber()] = (p1).plus(p1);
          _nw132[(new BigNumber(11)).toNumber()] = new BigNumber(249);
          _nw132[(new BigNumber(12)).toNumber()] = _618_v52.f27;
          _nw132[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_619_v53).length));
          _nw132[(new BigNumber(14)).toNumber()] = p1;
          _nw132[(new BigNumber(15)).toNumber()] = _618_v52.f27;
          _nw132[(new BigNumber(16)).toNumber()] = new BigNumber(((_605_v41)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_605_v41).length))]).length);
          _nw132[(new BigNumber(17)).toNumber()] = p1;
          _nw132[(new BigNumber(18)).toNumber()] = ((_dafny.ZERO).minus(_618_v52.f27)).minus(_618_v52.f27);
          _nw132[(new BigNumber(19)).toNumber()] = (((_611_v47)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_611_v47).length))]) ? (new BigNumber((_dafny.MultiSet.fromElements(_625_v59, _625_v59)).cardinality())) : (_618_v52.f27));
          _nw132[(new BigNumber(20)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_618_v52.f27), p1);
          _nw132[(new BigNumber(21)).toNumber()] = new BigNumber(637);
          _626_v60 = _nw132;
          let _index99 = _module.__default.safeIndex(new BigNumber(421), new BigNumber((_626_v60).length));
          (_626_v60)[_index99] = _618_v52.f27;
          let _index100 = _module.__default.safeIndex(new BigNumber(421), new BigNumber((_626_v60).length));
          (_626_v60)[_index100] = _618_v52.f27;
          let _627_v61;
          _627_v61 = _dafny.Map.Empty.slice().updateUnsafe(p1,_module.__default.fm43((_this).f17, globalState));
          let _628_v62;
          _628_v62 = _dafny.Map.Empty.slice().updateUnsafe((((_627_v61).contains(_618_v52.f27)) ? ((_627_v61).get(_618_v52.f27)) : (_this.f16)),_626_v60);
          _628_v62 = (_628_v62).update(false, _626_v60);
          (globalState).f7 = (_626_v60)[_module.__default.safeIndex(new BigNumber(421), new BigNumber((_626_v60).length))];
        } else {
          let _629_v63;
          _629_v63 = _dafny.Set.fromElements((_this).f17, _this.f16);
          let _630_v64;
          let _nw133 = Array((new BigNumber(17)).toNumber());
          _nw133[(_dafny.ZERO).toNumber()] = p1;
          _nw133[(_dafny.ONE).toNumber()] = p1;
          _nw133[(new BigNumber(2)).toNumber()] = p1;
          _nw133[(new BigNumber(3)).toNumber()] = p1;
          _nw133[(new BigNumber(4)).toNumber()] = new BigNumber(-868);
          _nw133[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(p1);
          _nw133[(new BigNumber(6)).toNumber()] = p1;
          _nw133[(new BigNumber(7)).toNumber()] = new BigNumber((_629_v63).length);
          _nw133[(new BigNumber(8)).toNumber()] = p1;
          _nw133[(new BigNumber(9)).toNumber()] = p1;
          _nw133[(new BigNumber(10)).toNumber()] = p1;
          _nw133[(new BigNumber(11)).toNumber()] = new BigNumber((_606_v42).length);
          _nw133[(new BigNumber(12)).toNumber()] = p1;
          _nw133[(new BigNumber(13)).toNumber()] = new BigNumber((_614_v50).length);
          _nw133[(new BigNumber(14)).toNumber()] = p1;
          _nw133[(new BigNumber(15)).toNumber()] = p1;
          _nw133[(new BigNumber(16)).toNumber()] = p1;
          _630_v64 = _nw133;
          let _631_v65;
          _631_v65 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_630_v64),_612_v48);
          let _632_v66;
          _632_v66 = _dafny.MultiSet.fromElements(_630_v64);
          _631_v65 = (_631_v65).update(_632_v66, (((_607_v43)[_module.__default.safeIndex(p1, new BigNumber((_607_v43).length))]) ? (_612_v48) : (_612_v48)));
          (globalState).f7 = (p1).multipliedBy(new BigNumber((_610_v46).cardinality()));
          (globalState).f7 = ((p1).minus(p1)).plus(p1);
          let _633_v67;
          let _634_v68;
          let _635_v69;
          let _636_v70;
          let _out16;
          let _out17;
          let _out18;
          let _out19;
          let _outcollector6 = (_604_v40).m13(!(!(!(!(p0)))), p3, globalState);
          _out16 = _outcollector6[0];
          _out17 = _outcollector6[1];
          _out18 = _outcollector6[2];
          _out19 = _outcollector6[3];
          _633_v67 = _out16;
          _634_v68 = _out17;
          _635_v69 = _out18;
          _636_v70 = _out19;
          (globalState).f7 = _633_v67;
        }
        let _637_v71;
        _637_v71 = _dafny.Set.fromElements(_612_v48);
        _637_v71 = _637_v71;
        let _638_v72;
        _638_v72 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _639_v73;
        let _init21 = ((_640_p1) => function (_641_i1) {
          return (_641_i1).multipliedBy((_dafny.ZERO).minus(_640_p1));
        })(p1);
        let _nw134 = Array((new BigNumber(17)).toNumber());
        for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw134.length); _i0_21++) {
          _nw134[_i0_21] = _init21(new BigNumber(_i0_21));
        }
        _639_v73 = _nw134;
        let _642_v74;
        _642_v74 = _dafny.MultiSet.fromElements(_639_v73, _639_v73, _639_v73, _639_v73, _639_v73);
        let _643_v75;
        let _nw135 = Array((new BigNumber(11)).toNumber());
        _nw135[(_dafny.ZERO).toNumber()] = p1;
        _nw135[(_dafny.ONE).toNumber()] = p1;
        _nw135[(new BigNumber(2)).toNumber()] = new BigNumber((_638_v72).length);
        _nw135[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(p1);
        _nw135[(new BigNumber(4)).toNumber()] = new BigNumber((_642_v74).cardinality());
        _nw135[(new BigNumber(5)).toNumber()] = p1;
        _nw135[(new BigNumber(6)).toNumber()] = p1;
        _nw135[(new BigNumber(7)).toNumber()] = p1;
        _nw135[(new BigNumber(8)).toNumber()] = new BigNumber(330);
        _nw135[(new BigNumber(9)).toNumber()] = (p1).minus(p1);
        _nw135[(new BigNumber(10)).toNumber()] = new BigNumber(((_605_v41)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_605_v41).length))]).length);
        _643_v75 = _nw135;
        let _index101 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_639_v73).length));
        (_639_v73)[_index101] = p1;
        let _index102 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_639_v73).length));
        (_639_v73)[_index102] = p1;
      }
      let _644_v76;
      _644_v76 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(_module.__default.fm43(p3, globalState)),p1)).length),p3);
      let _hi7 = new BigNumber((_644_v76).length);
      for (let _645_i2 = new BigNumber(177); _645_i2.isLessThan(_hi7); _645_i2 = _645_i2.plus(_dafny.ONE)) {
        let _646_v77;
        _646_v77 = _dafny.Seq.of(p0);
        let _647_v78;
        _647_v78 = _dafny.Map.Empty.slice().updateUnsafe(_645_i2,_dafny.Seq.Concat(_646_v77, _646_v77));
        _647_v78 = ((p0) ? (_647_v78) : (_647_v78));
        let _648_v79;
        let _nw136 = new _module.C1();
        _nw136.__ctor(p3, (p1).plus(new BigNumber(879)));
        _648_v79 = _nw136;
        (globalState).f2 = p1;
        let _649_v80;
        let _init22 = ((_650_v79) => function (_651_i3) {
          return _dafny.Set.fromElements(_650_v79.f25, _this.f16);
        })(_648_v79);
        let _nw137 = Array((new BigNumber(17)).toNumber());
        for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw137.length); _i0_22++) {
          _nw137[_i0_22] = _init22(new BigNumber(_i0_22));
        }
        _649_v80 = _nw137;
        let _652_v81;
        let _nw138 = new _module.C0();
        _nw138.__ctor(_649_v80, _module.__default.safeDivisionInt(p1, _645_i2));
        _652_v81 = _nw138;
      }
      let _653_v82;
      _653_v82 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,p1);
      let _hi8 = p1;
      for (let _654_i4 = new BigNumber((_653_v82).length); _654_i4.isLessThan(_hi8); _654_i4 = _654_i4.plus(_dafny.ONE)) {
        (_this).f16 = _module.__default.fm43(p3, globalState);
        let _655_v83;
        _655_v83 = _dafny.Seq.of(p1, p1, p1);
        _655_v83 = _655_v83;
        (globalState).f7 = p1;
        let _656_v84;
        _656_v84 = _dafny.MultiSet.fromElements(!((_this).f17), _this.f16, false);
        let _657_v85;
        _657_v85 = _dafny.Map.Empty.slice().updateUnsafe((_564_v0).Merge(_564_v0),(_656_v84).IsSubsetOf(_656_v84));
        let _658_v86;
        _658_v86 = _dafny.Seq.UnicodeFromString("afrpdrn");
        if (!((((_657_v85).contains((((_this).f17) ? (_564_v0) : (_564_v0)))) ? ((_657_v85).get((((_this).f17) ? (_564_v0) : (_564_v0)))) : (_dafny.Seq.IsPrefixOf(_658_v86, _658_v86))))) {
          (globalState).f7 = p1;
          (globalState).f7 = new BigNumber((_658_v86).length);
          let _659_v87;
          _659_v87 = _dafny.Map.Empty.slice().updateUnsafe(false,_566_v2);
          let _660_v88;
          _660_v88 = _dafny.Map.Empty.slice().updateUnsafe(_659_v87,_654_i4);
          let _661_v89;
          _661_v89 = _module.D14.create_DC27(_660_v88);
          _660_v88 = ((_661_v89).dtor_cf39).Merge((_660_v88).Merge(_660_v88));
          (globalState).f2 = _654_i4;
          let _662_v90;
          _662_v90 = new _dafny.CodePoint('e'.codePointAt(0));
          let _663_v91;
          _663_v91 = _dafny.Map.Empty.slice().updateUnsafe(p3,_module.__default.fm35((_655_v83)[_module.__default.safeIndex(p1, new BigNumber((_655_v83).length))], p3, new BigNumber((_564_v0).length), _662_v90, globalState));
          let _664_v92;
          let _nw139 = Array((new BigNumber(5)).toNumber());
          _nw139[(_dafny.ZERO).toNumber()] = _654_i4;
          _nw139[(_dafny.ONE).toNumber()] = new BigNumber((_663_v91).length);
          _nw139[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(_656_v84)).cardinality());
          _nw139[(new BigNumber(3)).toNumber()] = new BigNumber(842);
          _nw139[(new BigNumber(4)).toNumber()] = (p1).minus(p1);
          _664_v92 = _nw139;
          let _index103 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_664_v92).length));
          (_664_v92)[_index103] = new BigNumber(850);
          let _index104 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_664_v92).length));
          (_664_v92)[_index104] = (_dafny.ZERO).minus(p1);
        } else {
          let _665_v93;
          let _nw140 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Map.Empty);
          _665_v93 = _nw140;
          let _666_v94;
          _666_v94 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_658_v86).length),(_dafny.ZERO).minus(_654_i4));
          let _index105 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_665_v93).length));
          (_665_v93)[_index105] = _666_v94;
          let _index106 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_665_v93).length));
          (_665_v93)[_index106] = _666_v94;
          (_this).f16 = (p3) && (p3);
          let _667_v95;
          _667_v95 = _dafny.Set.fromElements((_this).f17, p3);
          let _668_v96;
          _668_v96 = _dafny.Map.Empty.slice().updateUnsafe(p3,_667_v95);
          let _669_v97;
          _669_v97 = _dafny.Seq.of(_667_v95);
          let _670_v98;
          let _nw141 = Array((new BigNumber(26)).toNumber());
          _nw141[(_dafny.ZERO).toNumber()] = (((_668_v96).contains(_this.f16)) ? ((_668_v96).get(_this.f16)) : ((_669_v97)[_module.__default.safeIndex(_654_i4, new BigNumber((_669_v97).length))]));
          _nw141[(_dafny.ONE).toNumber()] = _667_v95;
          _nw141[(new BigNumber(2)).toNumber()] = _dafny.Set.fromElements(true);
          _nw141[(new BigNumber(3)).toNumber()] = _667_v95;
          _nw141[(new BigNumber(4)).toNumber()] = _667_v95;
          _nw141[(new BigNumber(5)).toNumber()] = _667_v95;
          _nw141[(new BigNumber(6)).toNumber()] = _667_v95;
          _nw141[(new BigNumber(7)).toNumber()] = _667_v95;
          _nw141[(new BigNumber(8)).toNumber()] = _667_v95;
          _nw141[(new BigNumber(9)).toNumber()] = _667_v95;
          _nw141[(new BigNumber(10)).toNumber()] = _dafny.Set.fromElements(p0, p0, _this.f16, _module.__default.fm43((_this).f17, globalState));
          _nw141[(new BigNumber(11)).toNumber()] = _667_v95;
          _nw141[(new BigNumber(12)).toNumber()] = _667_v95;
          _nw141[(new BigNumber(13)).toNumber()] = _667_v95;
          _nw141[(new BigNumber(14)).toNumber()] = _667_v95;
          _nw141[(new BigNumber(15)).toNumber()] = _667_v95;
          _nw141[(new BigNumber(16)).toNumber()] = _667_v95;
          _nw141[(new BigNumber(17)).toNumber()] = _667_v95;
          _nw141[(new BigNumber(18)).toNumber()] = _dafny.Set.fromElements((((_564_v0).contains(p3)) ? ((_564_v0).get(p3)) : (p0)));
          _nw141[(new BigNumber(19)).toNumber()] = _dafny.Set.fromElements(p3);
          _nw141[(new BigNumber(20)).toNumber()] = _dafny.Set.fromElements((_this).f17);
          _nw141[(new BigNumber(21)).toNumber()] = _667_v95;
          _nw141[(new BigNumber(22)).toNumber()] = _667_v95;
          _nw141[(new BigNumber(23)).toNumber()] = _667_v95;
          _nw141[(new BigNumber(24)).toNumber()] = _dafny.Set.fromElements(p0, p3);
          _nw141[(new BigNumber(25)).toNumber()] = _module.__default.fm44(_this.f16, _this.f16, globalState);
          _670_v98 = _nw141;
          let _671_v99;
          let _nw142 = new _module.C0();
          _nw142.__ctor(_670_v98, _654_i4);
          _671_v99 = _nw142;
          r1 = ((false) ? (_dafny.Seq.contains(_658_v86, new _dafny.CodePoint('w'.codePointAt(0)))) : (p3));
          let _672_v100;
          let _nw143 = Array((new BigNumber(27)).toNumber()).fill(false);
          _672_v100 = _nw143;
          let _index107 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_672_v100).length));
          (_672_v100)[_index107] = p0;
          let _index108 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_672_v100).length));
          (_672_v100)[_index108] = p3;
        }
      }
      let _673_v101;
      _673_v101 = _dafny.Set.fromElements(p1);
      r1 = (_module.__default.fm49(p1, p1, _this.f16, p1, globalState)).equals((function () {
        let _coll37 = new _dafny.Set();
        for (const _compr_37 of (_673_v101).Elements) {
          let _674_v102 = _compr_37;
          if ((_673_v101).contains(_674_v102)) {
            _coll37.add(_module.__default.safeDivisionInt(_674_v102, new BigNumber((_dafny.Seq.of(false)).length)));
          }
        }
        return _coll37;
      }()).Union(_673_v101));
      let _675_v103;
      _675_v103 = _dafny.MultiSet.fromElements(p0);
      let _676_v104;
      _676_v104 = _dafny.Map.Empty.slice().updateUnsafe(p3,_675_v103);
      r0 = (_676_v104).Merge(_676_v104);
      r1 = p0;
      return [r0, r1];
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f16 = false;
      this._f17 = false;
      this._f22 = _dafny.ZERO;
      this._f23 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T1];
    }
    get f16() {
      let _this = this;
      return _this._f16;
    };
    set f16(value) {
      let _this = this;
      _this._f16 = value;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    __ctor(f22, f23, f16, f17) {
      let _this = this;
      (_this)._f22 = f22;
      (_this)._f23 = f23;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      return;
    }
    fm3(globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f16,(_this).f17)).length),new BigNumber(((_this).f23).length))).length),(_this).f22)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_this).f22)).length),new BigNumber(391)));
    };
    m2(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _677_v0;
      _677_v0 = new _dafny.CodePoint('w'.codePointAt(0));
      (_this).f16 = function (_source7) {
        if (_source7.is_DC10) {
          if (_this.f16) {
            return (_this).f17;
          } else {
            return true;
          }
        } else {
          let _678___mcc_h0 = (_source7).cf14;
          let _679_cf14 = _678___mcc_h0;
          return (_this).f17;
        }
      }(_module.D6.create_DC9(_677_v0));
      let _680_v1;
      _680_v1 = _dafny.Seq.of((_this).f17);
      let _681_v2;
      _681_v2 = _680_v1;
      let _682_v3;
      _682_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,(_this).f22);
      let _683_v4;
      _683_v4 = _dafny.Seq.of(p0, (((_682_v3).contains((_this).f22)) ? ((_682_v3).get((_this).f22)) : (p0)), p0);
      let _source8 = _module.D4.create_DC7(_681_v2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_683_v4,_this.f16)).length));
      if (_source8.is_DC7) {
        let _684___mcc_h1 = (_source8).cf11;
        let _685___mcc_h2 = (_source8).cf12;
        let _686_cf12 = _685___mcc_h2;
        let _687_cf11 = _684___mcc_h1;
        let _688_v5;
        let _nw144 = Array((new BigNumber(9)).toNumber()).fill(false);
        _688_v5 = _nw144;
        let _689_v6;
        _689_v6 = _dafny.Map.Empty.slice().updateUnsafe(p1,_688_v5);
        _689_v6 = _689_v6;
        let _nw145 = Array((new BigNumber(21)).toNumber()).fill(false);
        _688_v5 = _nw145;
        (globalState).f2 = _686_cf12;
        let _rhs105 = (_this).f17;
        let _rhs106 = _686_cf12;
        let _lhs71 = _this;
        _lhs71.f16 = _rhs105;
        r0 = _rhs106;
      } else {
        let _690___mcc_h3 = (_source8).cf10;
        let _691_cf10 = _690___mcc_h3;
        let _692_v7;
        _692_v7 = _dafny.Map.Empty.slice().updateUnsafe((_691_cf10).Merge((_dafny.Map.Empty.slice().updateUnsafe(_this.f16,(_this).f17)).update((_this).f17, _this.f16)),(_dafny.ZERO).minus(p0));
        _692_v7 = (_692_v7).update(_691_cf10, p0);
        _677_v0 = _677_v0;
        let _693_v8;
        _693_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,(_this).f17);
        let _694_v9;
        _694_v9 = _module.D3.create_DC3((_this).f17);
        let _695_v10;
        let _nw146 = Array((new BigNumber(3)).toNumber()).fill(false);
        _695_v10 = _nw146;
        let _696_v11;
        _696_v11 = _dafny.Set.fromElements(_695_v10, _695_v10);
        let _697_v12;
        _697_v12 = _module.D3.create_DC4(false, p0, p0, p1, (_this).f17);
        let _698_v13;
        let _nw147 = Array((new BigNumber(28)).toNumber());
        _nw147[(_dafny.ZERO).toNumber()] = p1;
        _nw147[(_dafny.ONE).toNumber()] = (_this).f17;
        _nw147[(new BigNumber(2)).toNumber()] = !(p1);
        _nw147[(new BigNumber(3)).toNumber()] = (_this).f17;
        _nw147[(new BigNumber(4)).toNumber()] = (((_693_v8).contains(p0)) ? ((_693_v8).get(p0)) : ((((_693_v8).contains(p0)) ? ((_693_v8).get(p0)) : ((((_693_v8).contains(p0)) ? ((_693_v8).get(p0)) : (_this.f16))))));
        _nw147[(new BigNumber(5)).toNumber()] = ((p1) ? (_this.f16) : (_this.f16));
        _nw147[(new BigNumber(6)).toNumber()] = (_694_v9).dtor_cf3;
        _nw147[(new BigNumber(7)).toNumber()] = ((_this).f22).isLessThan(new BigNumber(536));
        _nw147[(new BigNumber(8)).toNumber()] = _this.f16;
        _nw147[(new BigNumber(9)).toNumber()] = !(_this.f16);
        _nw147[(new BigNumber(10)).toNumber()] = true;
        _nw147[(new BigNumber(11)).toNumber()] = p1;
        _nw147[(new BigNumber(12)).toNumber()] = (new BigNumber((_dafny.Set.fromElements(!(_this.f16))).length)).isLessThanOrEqualTo((_dafny.ZERO).minus(p0));
        _nw147[(new BigNumber(13)).toNumber()] = _this.f16;
        _nw147[(new BigNumber(14)).toNumber()] = p1;
        _nw147[(new BigNumber(15)).toNumber()] = !(p1);
        _nw147[(new BigNumber(16)).toNumber()] = false;
        _nw147[(new BigNumber(17)).toNumber()] = (_680_v1)[_module.__default.safeIndex(p0, new BigNumber((_680_v1).length))];
        _nw147[(new BigNumber(18)).toNumber()] = p1;
        _nw147[(new BigNumber(19)).toNumber()] = _this.f16;
        _nw147[(new BigNumber(20)).toNumber()] = p1;
        _nw147[(new BigNumber(21)).toNumber()] = (_696_v11).IsDisjointFrom(_696_v11);
        _nw147[(new BigNumber(22)).toNumber()] = !((_this).f17);
        _nw147[(new BigNumber(23)).toNumber()] = _this.f16;
        _nw147[(new BigNumber(24)).toNumber()] = _this.f16;
        _nw147[(new BigNumber(25)).toNumber()] = (_this).f17;
        _nw147[(new BigNumber(26)).toNumber()] = true;
        _nw147[(new BigNumber(27)).toNumber()] = ((_697_v12).dtor_cf6).isLessThan(new BigNumber((_dafny.Seq.UnicodeFromString("rxvxscggb")).length));
        _698_v13 = _nw147;
        let _index109 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_698_v13).length));
        (_698_v13)[_index109] = _module.__default.fm31(p1, p1, globalState);
        let _699_v14;
        _699_v14 = _dafny.Set.fromElements(p0);
        let _700_v16;
        _700_v16 = (_this).f22;
        let _701_v17;
        _701_v17 = _dafny.MultiSet.fromElements(_dafny.Set.fromElements((_dafny.ZERO).minus((_700_v16)), p0), _699_v14, _699_v14, _699_v14);
        let _index110 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_698_v13).length));
        let _rhs107 = ((!(_this.f16)) ? (_677_v0) : (_677_v0));
        let _rhs108 = _module.__default.fm31((true) && (true), _this.f16, globalState);
        let _rhs109 = _module.__default.safeModuloInt(new BigNumber((_683_v4).length), _module.__default.fm0(p1, globalState));
        let _rhs110 = _dafny.Seq.of(p0, (_this).f22, p0, p0);
        let _rhs111 = _module.__default.fm31((_this).f17, (_701_v17).IsSubsetOf(_dafny.MultiSet.fromElements(_699_v14, function () {
          let _coll38 = new _dafny.Set();
          for (const _compr_38 of _dafny.IntegerRange(new BigNumber(122), new BigNumber(68))) {
            let _702_v15 = _compr_38;
            if (((new BigNumber(122)).isLessThanOrEqualTo(_702_v15)) && ((_702_v15).isLessThan(new BigNumber(68)))) {
              _coll38.add(_module.__default.safeDivisionInt(_702_v15, p0));
            }
          }
          return _coll38;
        }())), globalState);
        let _lhs72 = _this;
        let _lhs73 = globalState;
        let _lhs74 = _698_v13;
        let _lhs75 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_698_v13).length));
        _677_v0 = _rhs107;
        _lhs72.f16 = _rhs108;
        _lhs73.f7 = _rhs109;
        _683_v4 = _rhs110;
        _lhs74[_lhs75] = _rhs111;
        let _703_v18;
        _703_v18 = _dafny.Seq.of(_698_v13);
        let _704_v19;
        let _nw148 = Array((new BigNumber(13)).toNumber());
        _nw148[(_dafny.ZERO).toNumber()] = _module.__default.safeDivisionInt((_this).f22, (_dafny.ZERO).minus((_this).f22));
        _nw148[(_dafny.ONE).toNumber()] = new BigNumber(297);
        _nw148[(new BigNumber(2)).toNumber()] = ((_this).f22).minus((_this).f22);
        _nw148[(new BigNumber(3)).toNumber()] = _module.__default.fm0(_this.f16, globalState);
        _nw148[(new BigNumber(4)).toNumber()] = p0;
        _nw148[(new BigNumber(5)).toNumber()] = new BigNumber((_703_v18).length);
        _nw148[(new BigNumber(6)).toNumber()] = (_this).f22;
        _nw148[(new BigNumber(7)).toNumber()] = p0;
        _nw148[(new BigNumber(8)).toNumber()] = _module.__default.fm0((_this).f17, globalState);
        _nw148[(new BigNumber(9)).toNumber()] = (_this).f22;
        _nw148[(new BigNumber(10)).toNumber()] = p0;
        _nw148[(new BigNumber(11)).toNumber()] = new BigNumber((_680_v1).length);
        _nw148[(new BigNumber(12)).toNumber()] = p0;
        _704_v19 = _nw148;
        let _705_v20;
        _705_v20 = _dafny.Seq.UnicodeFromString("maojmxu");
        let _706_v21;
        _706_v21 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f17), _dafny.Map.Empty.slice().updateUnsafe(p0,(_698_v13)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_698_v13).length))]));
        let _707_v22;
        _707_v22 = _module.D9.create_DC15((_this).f17, (_698_v13)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_698_v13).length))], _dafny.Set.fromElements(p0), _704_v19, _698_v13);
        let _rhs112 = (_707_v22).dtor_cf21;
        let _rhs113 = (_this).f23;
        let _rhs114 = _706_v21;
        _704_v19 = _rhs112;
        _705_v20 = _rhs113;
        _706_v21 = _rhs114;
      }
      if (!((p1) && (_module.__default.fm31(true, p1, globalState)))) {
        let _708_v23;
        _708_v23 = _dafny.Set.fromElements((_this).f22);
        _708_v23 = _708_v23;
        let _709_v24;
        _709_v24 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm32((_this).f17, !(p1), _this.f16, _this.f16, globalState),new BigNumber(((_this).f23).length));
        let _710_v26;
        let _init23 = function (_711_i0) {
          return (_711_i0).minus((_this).f22);
        };
        let _nw149 = Array((new BigNumber(12)).toNumber());
        for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw149.length); _i0_23++) {
          _nw149[_i0_23] = _init23(new BigNumber(_i0_23));
        }
        _710_v26 = _nw149;
        let _712_v27;
        _712_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,_710_v26);
        let _713_v28;
        _713_v28 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,(_this).f17);
        let _714_v29;
        let _nw150 = Array((_dafny.ONE).toNumber()).fill(false);
        _714_v29 = _nw150;
        let _715_v30;
        _715_v30 = _dafny.Seq.of(_714_v29, _714_v29);
        let _716_v31;
        _716_v31 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_713_v28).length),(_715_v30)[_module.__default.safeIndex((_this).f22, new BigNumber((_715_v30).length))]);
        let _717_v32;
        _717_v32 = _dafny.Set.fromElements((_this).f17, _this.f16, !((_this).f17));
        let _718_v33;
        _718_v33 = _module.D9.create_DC15((_709_v24).contains((_dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f22)).update((_this).f17, p0)), (_this).f17, function () {
  let _coll39 = new _dafny.Set();
  for (const _compr_39 of _dafny.IntegerRange(new BigNumber(679), new BigNumber(561))) {
    let _719_v25 = _compr_39;
    if (((new BigNumber(679)).isLessThanOrEqualTo(_719_v25)) && ((_719_v25).isLessThan(new BigNumber(561)))) {
      _coll39.add((_719_v25).minus((_this).f22));
    }
  }
  return _coll39;
}(), (((_712_v27).contains((_this).f22)) ? ((_712_v27).get((_this).f22)) : (_710_v26)), (((_716_v31).contains((_dafny.ZERO).minus(new BigNumber((_717_v32).length)))) ? ((_716_v31).get((_dafny.ZERO).minus(new BigNumber((_717_v32).length)))) : (_714_v29)));
        _718_v33 = _module.D9.create_DC15(((_this).f17) === (p1), _module.__default.fm31(_this.f16, _this.f16, globalState), _708_v23, _710_v26, _714_v29);
        (globalState).f2 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(p0, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(p1, !(_this.f16))).length))));
        let _720_v34;
        let _nw151 = Array((new BigNumber(19)).toNumber()).fill([]);
        _720_v34 = _nw151;
        let _index111 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_720_v34).length));
        (_720_v34)[_index111] = _714_v29;
        let _index112 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_720_v34).length));
        (_720_v34)[_index112] = _714_v29;
        let _721_v35;
        _721_v35 = _dafny.MultiSet.fromElements((_this).f22);
        let _722_v36;
        _722_v36 = _dafny.MultiSet.fromElements(_721_v35);
        (_this).f16 = (_722_v36).IsProperSubsetOf(((_722_v36).update(_721_v35, _module.__default.abs(new BigNumber(330)))).update(_dafny.MultiSet.fromElements(new BigNumber((_680_v1).length), new BigNumber(((_this).f23).length)), _module.__default.abs((_this).f22)));
      } else {
        let _723_v37;
        let _nw152 = Array((new BigNumber(14)).toNumber()).fill(false);
        _723_v37 = _nw152;
        let _724_v38;
        _724_v38 = _723_v37;
        let _725_v39;
        _725_v39 = _dafny.Seq.of(_724_v38);
        let _726_v40;
        _726_v40 = _module.D3.create_DC3(false);
        let _727_v41;
        _727_v41 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_725_v39).length)).isLessThanOrEqualTo(p0),_726_v40);
        let _728_v42;
        _728_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,p1);
        _727_v41 = (_727_v41).update((((_728_v42).contains(p1)) ? ((_728_v42).get(p1)) : (p1)), _726_v40);
        let _729_v43;
        let _init24 = ((_730_v0) => function (_731_i1) {
          return _dafny.MultiSet.FromArray(_dafny.Seq.of(_module.D6.create_DC9(_730_v0)));
        })(_677_v0);
        let _nw153 = Array((new BigNumber(27)).toNumber());
        for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw153.length); _i0_24++) {
          _nw153[_i0_24] = _init24(new BigNumber(_i0_24));
        }
        _729_v43 = _nw153;
        let _732_v44;
        _732_v44 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,(_this).f22);
        let _733_v45;
        _733_v45 = _dafny.Map.Empty.slice().updateUnsafe(_729_v43,_732_v44);
        let _734_v46;
        _734_v46 = _dafny.Seq.of(_729_v43, _729_v43, _729_v43);
        _733_v45 = (_733_v45).update((_734_v46)[_module.__default.safeIndex((_this).f22, new BigNumber((_734_v46).length))], _module.__default.fm32(p1, false, _this.f16, p1, globalState));
        (_this).f16 = false;
        _683_v4 = _683_v4;
        let _735_v47;
        _735_v47 = _dafny.Seq.Concat(_683_v4, _683_v4);
        _735_v47 = ((_module.__default.fm31((_this).f17, _module.__default.fm31(p1, _this.f16, globalState), globalState)) ? (((p1) ? (_735_v47) : (_735_v47))) : (_735_v47));
      }
      (globalState).f2 = (p0).multipliedBy((_this).f22);
      if (_module.__default.fm31(p1, (_this).f17, globalState)) {
        r0 = _module.__default.fm0(_this.f16, globalState);
        let _736_v48;
        _736_v48 = _dafny.Set.fromElements(p1);
        (globalState).f7 = ((p1) ? (new BigNumber((_736_v48).length)) : (p0));
        (_this).f16 = !((_this).f17) || (((_this).f22).isEqualTo(new BigNumber(775)));
        let _737_v49;
        _737_v49 = _module.D4.create_DC7(_681_v2, (_this).f22);
        let _738_v50;
        let _nw154 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
        _738_v50 = _nw154;
        let _739_v51;
        _739_v51 = _dafny.Map.Empty.slice().updateUnsafe(_738_v50,p0);
        let _740_v52;
        _740_v52 = _dafny.MultiSet.fromElements((_this).f22);
        let _741_v53;
        _741_v53 = _dafny.Map.Empty.slice().updateUnsafe(_736_v48,_740_v52);
        let _742_v54;
        _742_v54 = _module.D9.create_DC16((_this).f22, p1);
        let _743_v55;
        let _nw155 = Array((new BigNumber(29)).toNumber());
        _nw155[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("sdsflqmgt")).length))).minus(new BigNumber((_736_v48).length)));
        _nw155[(_dafny.ONE).toNumber()] = (_module.D9.create_DC16((_this).f22, (_this).f17)).dtor_cf23;
        _nw155[(new BigNumber(2)).toNumber()] = (_this).f22;
        _nw155[(new BigNumber(3)).toNumber()] = (_this).f22;
        _nw155[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xsh"), (_this).f23)).length);
        _nw155[(new BigNumber(5)).toNumber()] = p0;
        _nw155[(new BigNumber(6)).toNumber()] = new BigNumber(-190);
        _nw155[(new BigNumber(7)).toNumber()] = new BigNumber(20);
        _nw155[(new BigNumber(8)).toNumber()] = (_this).f22;
        _nw155[(new BigNumber(9)).toNumber()] = (p0).multipliedBy((_dafny.ZERO).minus(p0));
        _nw155[(new BigNumber(10)).toNumber()] = (_this).f22;
        _nw155[(new BigNumber(11)).toNumber()] = p0;
        _nw155[(new BigNumber(12)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_677_v0,p1)).length);
        _nw155[(new BigNumber(13)).toNumber()] = p0;
        _nw155[(new BigNumber(14)).toNumber()] = p0;
        _nw155[(new BigNumber(15)).toNumber()] = p0;
        _nw155[(new BigNumber(16)).toNumber()] = (_this).f22;
        _nw155[(new BigNumber(17)).toNumber()] = p0;
        _nw155[(new BigNumber(18)).toNumber()] = ((_this.f16) ? (new BigNumber(817)) : (_module.__default.fm0(p1, globalState)));
        _nw155[(new BigNumber(19)).toNumber()] = (_this).f22;
        _nw155[(new BigNumber(20)).toNumber()] = (_this).f22;
        _nw155[(new BigNumber(21)).toNumber()] = p0;
        _nw155[(new BigNumber(22)).toNumber()] = (_this).f22;
        _nw155[(new BigNumber(23)).toNumber()] = (_this).f22;
        _nw155[(new BigNumber(24)).toNumber()] = new BigNumber((_683_v4).length);
        _nw155[(new BigNumber(25)).toNumber()] = (((_739_v51).contains(_738_v50)) ? ((_739_v51).get(_738_v50)) : (p0));
        _nw155[(new BigNumber(26)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_740_v52).cardinality()), (_this).f22);
        _nw155[(new BigNumber(27)).toNumber()] = new BigNumber(((((_741_v53).contains(_736_v48)) ? ((_741_v53).get(_736_v48)) : ((_dafny.MultiSet.fromElements(p0)).update((_742_v54).dtor_cf23, _module.__default.abs((_this).f22))))).cardinality());
        _nw155[(new BigNumber(28)).toNumber()] = _module.__default.safeModuloInt((_this).f22, p0);
        _743_v55 = _nw155;
        let _index113 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_743_v55).length));
        (_743_v55)[_index113] = ((_this).f22).multipliedBy(_module.__default.fm0(p1, globalState));
        let _index114 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_743_v55).length));
        let _rhs115 = _737_v49;
        let _rhs116 = (_this).f22;
        let _lhs76 = _743_v55;
        let _lhs77 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_743_v55).length));
        _737_v49 = _rhs115;
        _lhs76[_lhs77] = _rhs116;
        if (_this.f16) {
          (globalState).f7 = p0;
          let _744_v56;
          let _nw156 = new _module.C1();
          _nw156.__ctor(p1, new BigNumber(((_this).f23).length));
          _744_v56 = _nw156;
          let _745_v57;
          let _nw157 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Set.Empty);
          _745_v57 = _nw157;
          let _746_v58;
          let _nw158 = new _module.C0();
          _nw158.__ctor(_745_v57, (_this).f22);
          _746_v58 = _nw158;
          (_744_v56).f25 = _this.f16;
          let _747_v59;
          let _init25 = function (_748_i2) {
            return true;
          };
          let _nw159 = Array((new BigNumber(12)).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw159.length); _i0_25++) {
            _nw159[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _747_v59 = _nw159;
          let _index115 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_747_v59).length));
          (_747_v59)[_index115] = false;
          let _index116 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_747_v59).length));
          (_747_v59)[_index116] = _744_v56.f25;
        } else {
          let _749_v60;
          let _nw160 = Array((_dafny.ONE).toNumber()).fill(false);
          _749_v60 = _nw160;
          let _index117 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_749_v60).length));
          (_749_v60)[_index117] = _this.f16;
          let _750_v61;
          _750_v61 = _dafny.MultiSet.fromElements(_749_v60, _749_v60, _749_v60, _749_v60);
          let _index118 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_749_v60).length));
          let _rhs117 = new _dafny.CodePoint('m'.codePointAt(0));
          let _rhs118 = (_this).f17;
          let _rhs119 = !((_750_v61).Difference(_750_v61)).equals((_dafny.MultiSet.fromElements(_749_v60, _749_v60)).Difference(_750_v61));
          let _lhs78 = _this;
          let _lhs79 = _749_v60;
          let _lhs80 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_749_v60).length));
          _677_v0 = _rhs117;
          _lhs78.f16 = _rhs118;
          _lhs79[_lhs80] = _rhs119;
          let _751_v62;
          _751_v62 = _dafny.MultiSet.fromElements(_677_v0);
          let _752_v63;
          _752_v63 = _dafny.MultiSet.fromElements(_this.f16, (_this).f17);
          let _index119 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_749_v60).length));
          let _index120 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_743_v55).length));
          let _rhs120 = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_751_v62).cardinality())));
          let _rhs121 = p0;
          let _rhs122 = (_752_v63).equals(_752_v63);
          let _rhs123 = p0;
          let _rhs124 = _738_v50;
          let _lhs81 = globalState;
          let _lhs82 = _749_v60;
          let _lhs83 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_749_v60).length));
          let _lhs84 = _743_v55;
          let _lhs85 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_743_v55).length));
          _lhs81.f7 = _rhs120;
          r0 = _rhs121;
          _lhs82[_lhs83] = _rhs122;
          _lhs84[_lhs85] = _rhs123;
          _738_v50 = _rhs124;
          (globalState).f7 = (_this).f22;
          let _index121 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_749_v60).length));
          (_749_v60)[_index121] = !((_this).f17);
          let _753_v64;
          _753_v64 = _module.D3.create_DC3(p1);
          let _754_v65;
          let _nw161 = new _module.C2();
          _nw161.__ctor(_753_v64, !(p1), (_752_v63).equals(_752_v63));
          _754_v65 = _nw161;
        }
      } else {
        r0 = ((p1) ? (p0) : ((_dafny.ZERO).minus((_this).f22)));
        let _755_v66;
        let _init26 = ((_756_p0) => function (_757_i3) {
          return (_757_i3).plus(_756_p0);
        })(p0);
        let _nw162 = Array((new BigNumber(29)).toNumber());
        for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw162.length); _i0_26++) {
          _nw162[_i0_26] = _init26(new BigNumber(_i0_26));
        }
        _755_v66 = _nw162;
        let _758_v67;
        _758_v67 = _dafny.Map.Empty.slice().updateUnsafe(p1,_755_v66);
        let _759_v68;
        _759_v68 = _dafny.Set.fromElements(_this.f16, _this.f16, (_this).f17, false);
        let _760_v69;
        _760_v69 = _dafny.Map.Empty.slice().updateUnsafe((((_758_v67).contains(_this.f16)) ? ((_758_v67).get(_this.f16)) : (_755_v66)),_759_v68);
        let _761_v70;
        _761_v70 = _dafny.Map.Empty.slice().updateUnsafe(_683_v4,new BigNumber((_760_v69).length));
        let _762_v71;
        _762_v71 = _dafny.MultiSet.fromElements(p1, (_this).f17);
        let _763_v72;
        _763_v72 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_762_v71);
        let _rhs125 = _683_v4;
        let _rhs126 = (((_761_v70).contains(_dafny.Seq.of((_this).f22, p0, new BigNumber((_763_v72).length), new BigNumber(633)))) ? ((_761_v70).get(_dafny.Seq.of((_this).f22, p0, new BigNumber((_763_v72).length), new BigNumber(633)))) : ((_this).f22));
        _683_v4 = _rhs125;
        r0 = _rhs126;
        let _764_v73;
        _764_v73 = _dafny.Set.fromElements(_762_v71, _762_v71, _762_v71, _762_v71);
        (_this).m11(_module.__default.safeModuloInt(p0, new BigNumber(258)), (_this).f17, (_764_v73).Union(_module.__default.fm50((_this).f22, p1, p1, globalState)), _762_v71, globalState);
        let _765_v74;
        let _nw163 = new _module.C1();
        _nw163.__ctor(p1, _module.__default.safeDivisionInt(new BigNumber(502), p0));
        _765_v74 = _nw163;
        let _766_v75;
        _766_v75 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,(_this).f17);
        _766_v75 = (_766_v75).update(false, p1);
      }
      let _767_v76;
      _767_v76 = _dafny.MultiSet.fromElements((_this).f17, p1);
      let _768_v77;
      _768_v77 = _module.D12.create_DC23(_this.f16, (_this).f22, (_this).f22);
      if (!(((_768_v77).dtor_cf31).isLessThan(new BigNumber((_767_v76).cardinality())))) {
        let _769_v78;
        _769_v78 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f17) || ((_this).f17),_683_v4);
        _769_v78 = (_769_v78).update(!(!(p1)), _683_v4);
        let _770_v79;
        _770_v79 = _module.D6.create_DC10();
        let _771_v80;
        _771_v80 = _dafny.Map.Empty.slice().updateUnsafe(_680_v1,_770_v79);
        _771_v80 = (function () {
          let _coll40 = new _dafny.Map();
          for (const _compr_40 of (_dafny.Map.Empty.slice().updateUnsafe(_680_v1,new BigNumber((_dafny.Seq.of(p1, _this.f16, p1, (_this).f17, (_this).f17)).length))).Keys.Elements) {
            let _772_v81 = _compr_40;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_680_v1,new BigNumber((_dafny.Seq.of(p1, _this.f16, p1, (_this).f17, (_this).f17)).length))).contains(_772_v81)) {
              _coll40.push([_772_v81,_770_v79]);
            }
          }
          return _coll40;
        }()).Merge((_771_v80).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_this).f17, _this.f16),_770_v79)));
        let _773_v82;
        _773_v82 = _module.D13.create_DC26();
        let _source9 = _773_v82;
        if (_source9.is_DC26) {
          let _774_v83;
          _774_v83 = _dafny.Set.fromElements((_this).f22, p0);
          let _775_v84;
          _775_v84 = _dafny.Map.Empty.slice().updateUnsafe(p0,_774_v83);
          let _776_v85;
          _776_v85 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_680_v1, _dafny.Seq.of(!(_this.f16)))).length),_775_v84);
          _776_v85 = (_776_v85).update(p0, function () {
            let _coll41 = new _dafny.Map();
            for (const _compr_41 of _dafny.IntegerRange(new BigNumber(462), new BigNumber(717))) {
              let _777_v86 = _compr_41;
              if (((new BigNumber(462)).isLessThanOrEqualTo(_777_v86)) && ((_777_v86).isLessThan(new BigNumber(717)))) {
                _coll41.push([(_777_v86).multipliedBy((_this).f22),_774_v83]);
              }
            }
            return _coll41;
          }());
          let _778_v87;
          _778_v87 = _dafny.Set.fromElements(_767_v76);
          (_this).m11(p0, p1, (_778_v87).Difference(_778_v87), _767_v76, globalState);
          let _779_v88;
          _779_v88 = _dafny.Seq.UnicodeFromString("ampyus");
          let _780_v89;
          _780_v89 = _dafny.Map.Empty.slice().updateUnsafe(_677_v0,(_this).f17);
          let _781_v90;
          let _nw164 = Array((new BigNumber(26)).toNumber());
          _nw164[(_dafny.ZERO).toNumber()] = (_this).f22;
          _nw164[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("t")).length);
          _nw164[(new BigNumber(2)).toNumber()] = _module.__default.safeModuloInt((_this).f22, (_this).f22);
          _nw164[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(p0);
          _nw164[(new BigNumber(4)).toNumber()] = p0;
          _nw164[(new BigNumber(5)).toNumber()] = p0;
          _nw164[(new BigNumber(6)).toNumber()] = p0;
          _nw164[(new BigNumber(7)).toNumber()] = (_this).f22;
          _nw164[(new BigNumber(8)).toNumber()] = (_683_v4)[_module.__default.safeIndex((_this).f22, new BigNumber((_683_v4).length))];
          _nw164[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus((_this).f22);
          _nw164[(new BigNumber(10)).toNumber()] = (p0).minus(p0);
          _nw164[(new BigNumber(11)).toNumber()] = new BigNumber(671);
          _nw164[(new BigNumber(12)).toNumber()] = new BigNumber((_780_v89).length);
          _nw164[(new BigNumber(13)).toNumber()] = p0;
          _nw164[(new BigNumber(14)).toNumber()] = p0;
          _nw164[(new BigNumber(15)).toNumber()] = p0;
          _nw164[(new BigNumber(16)).toNumber()] = ((true) ? (_module.__default.fm0(p1, globalState)) : (_module.__default.fm0(true, globalState)));
          _nw164[(new BigNumber(17)).toNumber()] = new BigNumber((_683_v4).length);
          _nw164[(new BigNumber(18)).toNumber()] = ((_dafny.ZERO).minus(p0)).plus((_dafny.ZERO).minus(p0));
          _nw164[(new BigNumber(19)).toNumber()] = p0;
          _nw164[(new BigNumber(20)).toNumber()] = new BigNumber(26);
          _nw164[(new BigNumber(21)).toNumber()] = p0;
          _nw164[(new BigNumber(22)).toNumber()] = p0;
          _nw164[(new BigNumber(23)).toNumber()] = new BigNumber((_module.__default.fm42(_this.f16, p0, p0, true, globalState)).length);
          _nw164[(new BigNumber(24)).toNumber()] = (_this).f22;
          _nw164[(new BigNumber(25)).toNumber()] = (_this).f22;
          _781_v90 = _nw164;
          let _index122 = _module.__default.safeIndex(new BigNumber(322), new BigNumber((_781_v90).length));
          (_781_v90)[_index122] = (_this).f22;
          let _782_v91;
          _782_v91 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm35(p0, p1, (_this).f22, _677_v0, globalState),(_this).f22);
          let _783_v92;
          _783_v92 = _module.D12.create_DC22(_779_v88);
          let _index123 = _module.__default.safeIndex(new BigNumber(322), new BigNumber((_781_v90).length));
          let _rhs127 = _dafny.Seq.Concat(_779_v88, (_this).f23);
          let _rhs128 = (_this).f17;
          let _rhs129 = _module.__default.safeModuloInt(_module.__default.safeModuloInt((_this).f22, (_dafny.ZERO).minus((_this).f22)), new BigNumber(((_782_v91).Merge(_782_v91)).length));
          let _rhs130 = p0;
          let _rhs131 = _dafny.Seq.contains((_783_v92).dtor_cf29, _677_v0);
          let _lhs86 = _this;
          let _lhs87 = _781_v90;
          let _lhs88 = _module.__default.safeIndex(new BigNumber(322), new BigNumber((_781_v90).length));
          let _lhs89 = _this;
          _779_v88 = _rhs127;
          _lhs86.f16 = _rhs128;
          _lhs87[_lhs88] = _rhs129;
          r0 = _rhs130;
          _lhs89.f16 = _rhs131;
          r0 = (_this).f22;
        } else {
          let _784___mcc_h4 = (_source9).cf38;
          let _785_cf38 = _784___mcc_h4;
          let _786_v93;
          _786_v93 = _module.D3.create_DC3((_this).f17);
          let _787_v94;
          let _nw165 = new _module.C2();
          _nw165.__ctor(_786_v93, (_this).f17, p1);
          _787_v94 = _nw165;
          (_this).f16 = p1;
          (globalState).f2 = p0;
          let _788_v95;
          _788_v95 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f22);
          let _789_v96;
          _789_v96 = _dafny.Seq.of(_788_v95, (_788_v95).Merge(_788_v95), _788_v95, _788_v95);
          r0 = new BigNumber(((_789_v96)[_module.__default.safeIndex((_this).f22, new BigNumber((_789_v96).length))]).length);
        }
        if (p1) {
          (globalState).f7 = (_683_v4)[_module.__default.safeIndex((_this).f22, new BigNumber((_683_v4).length))];
          let _790_v97;
          _790_v97 = _dafny.Map.Empty.slice().updateUnsafe(!(((_this).f22).isEqualTo((_this).f22)),_769_v78);
          _790_v97 = (_790_v97).update(true, _769_v78);
          (_this).f16 = p1;
          let _791_v98;
          let _nw166 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _791_v98 = _nw166;
          let _index124 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_791_v98).length));
          (_791_v98)[_index124] = (_dafny.ZERO).minus((_this).f22);
          let _index125 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_791_v98).length));
          (_791_v98)[_index125] = (_this).f22;
          (globalState).f7 = ((_this.f16) ? ((_791_v98)[_module.__default.safeIndex(new BigNumber(403), new BigNumber((_791_v98).length))]) : (p0));
        } else {
          let _792_v99;
          _792_v99 = _dafny.Set.fromElements(_767_v76);
          (_this).m11((_this).f22, (_this).f17, _792_v99, _767_v76, globalState);
          (_this).f16 = p1;
          let _793_v100;
          _793_v100 = _module.D7.create_DC12();
          let _794_v101;
          _794_v101 = _dafny.Map.Empty.slice().updateUnsafe(!((new BigNumber((_683_v4).length)).isEqualTo(new BigNumber(402))),_793_v100);
          _794_v101 = (_794_v101).update(!((_680_v1)[_module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_680_v1).length))]) || (false), _module.D7.create_DC12());
          let _795_v102;
          let _nw167 = new _module.C1();
          _nw167.__ctor((_this).f17, ((_this).f22).plus((_this).f22));
          _795_v102 = _nw167;
          let _796_v103;
          _796_v103 = _module.D3.create_DC3(!(true));
          let _797_v104;
          let _nw168 = new _module.C2();
          _nw168.__ctor(_796_v103, (_795_v102.f25) || (!((_this).f17)), _795_v102.f25);
          _797_v104 = _nw168;
        }
        let _798_v105;
        _798_v105 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f23)[_module.__default.safeIndex(_module.__default.fm0((_this).f17, globalState), new BigNumber(((_this).f23).length))],new BigNumber(536));
        let _799_v106;
        _799_v106 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,(_this).f22);
        let _800_v107;
        _800_v107 = _dafny.MultiSet.fromElements(new BigNumber(34), p0, new BigNumber((_799_v106).length), new BigNumber(((_this).fm3(globalState)).length));
        _798_v105 = (_798_v105).Merge(_dafny.Map.Empty.slice().updateUnsafe(((_this).f23)[_module.__default.safeIndex(new BigNumber(-351), new BigNumber(((_this).f23).length))],new BigNumber((_800_v107).cardinality())));
      } else {
        let _801_v108;
        _801_v108 = _module.D9.create_DC16(new BigNumber(457), false);
        let _802_v109;
        let _nw169 = Array((new BigNumber(20)).toNumber());
        _nw169[(_dafny.ZERO).toNumber()] = (_this).f17;
        _nw169[(_dafny.ONE).toNumber()] = _dafny.Seq.contains(_683_v4, (_this).f22);
        _nw169[(new BigNumber(2)).toNumber()] = (_this).f17;
        _nw169[(new BigNumber(3)).toNumber()] = (_680_v1)[_module.__default.safeIndex((_this).f22, new BigNumber((_680_v1).length))];
        _nw169[(new BigNumber(4)).toNumber()] = p1;
        _nw169[(new BigNumber(5)).toNumber()] = p1;
        _nw169[(new BigNumber(6)).toNumber()] = _this.f16;
        _nw169[(new BigNumber(7)).toNumber()] = (_this).f17;
        _nw169[(new BigNumber(8)).toNumber()] = _this.f16;
        _nw169[(new BigNumber(9)).toNumber()] = (_680_v1)[_module.__default.safeIndex((_this).f22, new BigNumber((_680_v1).length))];
        _nw169[(new BigNumber(10)).toNumber()] = _dafny.Seq.IsPrefixOf((_this).f23, (_this).f23);
        _nw169[(new BigNumber(11)).toNumber()] = _this.f16;
        _nw169[(new BigNumber(12)).toNumber()] = !((p0).isLessThan((_this).f22));
        _nw169[(new BigNumber(13)).toNumber()] = true;
        _nw169[(new BigNumber(14)).toNumber()] = p1;
        _nw169[(new BigNumber(15)).toNumber()] = _this.f16;
        _nw169[(new BigNumber(16)).toNumber()] = (_801_v108).dtor_cf24;
        _nw169[(new BigNumber(17)).toNumber()] = true;
        _nw169[(new BigNumber(18)).toNumber()] = p1;
        _nw169[(new BigNumber(19)).toNumber()] = ((_this.f16) ? (p1) : ((_this).f17));
        _802_v109 = _nw169;
        let _803_v110;
        _803_v110 = _dafny.Seq.of(_680_v1);
        let _index126 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_802_v109).length));
        (_802_v109)[_index126] = _dafny.Seq.IsPrefixOf(_dafny.Seq.of(_680_v1, _680_v1), _803_v110);
        let _index127 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_802_v109).length));
        (_802_v109)[_index127] = !(_this.f16);
        let _804_v111;
        _804_v111 = _dafny.Set.fromElements((_this).f17);
        let _805_v112;
        _805_v112 = _dafny.Seq.of(_767_v76);
        let _806_v113;
        let _nw170 = Array((new BigNumber(3)).toNumber());
        _nw170[(_dafny.ZERO).toNumber()] = (_this).f22;
        _nw170[(_dafny.ONE).toNumber()] = p0;
        _nw170[(new BigNumber(2)).toNumber()] = p0;
        _806_v113 = _nw170;
        let _807_v114;
        _807_v114 = _module.D9.create_DC14(_806_v113);
        let _808_v115;
        _808_v115 = _dafny.Seq.of(_807_v114, _807_v114, _807_v114);
        let _809_v116;
        _809_v116 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,(_this).f17);
        let _810_v117;
        _810_v117 = _dafny.Set.fromElements(_809_v116);
        let _811_v118;
        let _nw171 = Array((new BigNumber(23)).toNumber());
        _nw171[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.update((_this).f23, _module.__default.safeIndex(new BigNumber((_804_v111).length), new BigNumber(((_this).f23).length)), _677_v0)).length);
        _nw171[(_dafny.ONE).toNumber()] = (_this).f22;
        _nw171[(new BigNumber(2)).toNumber()] = p0;
        _nw171[(new BigNumber(3)).toNumber()] = new BigNumber(((_this).f23).length);
        _nw171[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(((_this).f22).multipliedBy((_this).f22));
        _nw171[(new BigNumber(5)).toNumber()] = new BigNumber(((_dafny.MultiSet.fromElements(!((_802_v109)[_module.__default.safeIndex(new BigNumber(605), new BigNumber((_802_v109).length))]), p1)).Intersect((_805_v112)[_module.__default.safeIndex(new BigNumber(450), new BigNumber((_805_v112).length))])).cardinality());
        _nw171[(new BigNumber(6)).toNumber()] = (p0).plus((_this).f22);
        _nw171[(new BigNumber(7)).toNumber()] = (_this).f22;
        _nw171[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt((_this).f22, (_this).f22);
        _nw171[(new BigNumber(9)).toNumber()] = p0;
        _nw171[(new BigNumber(10)).toNumber()] = new BigNumber((_808_v115).length);
        _nw171[(new BigNumber(11)).toNumber()] = (_this).f22;
        _nw171[(new BigNumber(12)).toNumber()] = (_this).f22;
        _nw171[(new BigNumber(13)).toNumber()] = p0;
        _nw171[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus((_this).f22);
        _nw171[(new BigNumber(15)).toNumber()] = (new BigNumber(-897)).minus(p0);
        _nw171[(new BigNumber(16)).toNumber()] = p0;
        _nw171[(new BigNumber(17)).toNumber()] = new BigNumber(((_this).f23).length);
        _nw171[(new BigNumber(18)).toNumber()] = (_this).f22;
        _nw171[(new BigNumber(19)).toNumber()] = new BigNumber((_810_v117).length);
        _nw171[(new BigNumber(20)).toNumber()] = p0;
        _nw171[(new BigNumber(21)).toNumber()] = (_this).f22;
        _nw171[(new BigNumber(22)).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_683_v4, _dafny.Seq.update(_683_v4, _module.__default.safeIndex(p0, new BigNumber((_683_v4).length)), new BigNumber((_767_v76).cardinality()))))).cardinality());
        _811_v118 = _nw171;
        let _nw172 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
        _811_v118 = _nw172;
        let _index128 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_806_v113).length));
        (_806_v113)[_index128] = p0;
        let _index129 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_806_v113).length));
        (_806_v113)[_index129] = (_this).f22;
        if (_module.__default.fm31((_802_v109)[_module.__default.safeIndex(new BigNumber(605), new BigNumber((_802_v109).length))], (_767_v76).equals(_dafny.MultiSet.FromArray(_680_v1)), globalState)) {
          let _index130 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_802_v109).length));
          (_802_v109)[_index130] = !((_this).f22).isEqualTo(new BigNumber((_683_v4).length));
          let _812_v119;
          let _nw173 = new _module.C1();
          _nw173.__ctor(((_806_v113)[_module.__default.safeIndex(new BigNumber(761), new BigNumber((_806_v113).length))]).isLessThan(new BigNumber(((_this).f23).length)), ((_this).f22).minus(new BigNumber((_dafny.Seq.UnicodeFromString("slvmstdm")).length)));
          _812_v119 = _nw173;
          (globalState).f7 = p0;
          let _813_v120;
          let _nw174 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
          _813_v120 = _nw174;
          let _814_v121;
          _814_v121 = _dafny.Set.fromElements(new BigNumber((_804_v111).length));
          let _815_v122;
          _815_v122 = _dafny.Map.Empty.slice().updateUnsafe(_814_v121,(_this).f17);
          let _index131 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_813_v120).length));
          (_813_v120)[_index131] = _815_v122;
          let _816_v123;
          _816_v123 = _dafny.Seq.of(_804_v111);
          let _index132 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_813_v120).length));
          let _index133 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_802_v109).length));
          let _rhs132 = _module.__default.fm43(p1, globalState);
          let _rhs133 = _module.__default.fm51(_module.__default.fm0(_this.f16, globalState), new BigNumber(631), _dafny.Seq.Concat(_680_v1, _dafny.Seq.of(_this.f16)), _816_v123, globalState);
          let _rhs134 = (p1) === (_this.f16);
          let _lhs90 = _this;
          let _lhs91 = _813_v120;
          let _lhs92 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_813_v120).length));
          let _lhs93 = _802_v109;
          let _lhs94 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_802_v109).length));
          _lhs90.f16 = _rhs132;
          _lhs91[_lhs92] = _rhs133;
          _lhs93[_lhs94] = _rhs134;
          let _817_v124;
          _817_v124 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_dafny.Seq.update((_803_v110)[_module.__default.safeIndex(p0, new BigNumber((_803_v110).length))], _module.__default.safeIndex(p0, new BigNumber(((_803_v110)[_module.__default.safeIndex(p0, new BigNumber((_803_v110).length))]).length)), (_this).f17), _680_v1), _680_v1);
          let _818_v125;
          _818_v125 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("axlauykx"), (_this).f23, (_this).f23);
          let _index134 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_802_v109).length));
          let _rhs135 = _683_v4;
          let _rhs136 = _this.f16;
          let _rhs137 = (((_817_v124).contains(_dafny.Seq.Concat(_680_v1, _dafny.Seq.update(_680_v1, _module.__default.safeIndex(new BigNumber((_818_v125).length), new BigNumber((_680_v1).length)), _812_v119.f25)))) ? ((_817_v124).get(_dafny.Seq.Concat(_680_v1, _dafny.Seq.update(_680_v1, _module.__default.safeIndex(new BigNumber((_818_v125).length), new BigNumber((_680_v1).length)), _812_v119.f25)))) : (p0));
          let _rhs138 = _module.__default.fm39(globalState);
          let _lhs95 = _802_v109;
          let _lhs96 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_802_v109).length));
          let _lhs97 = globalState;
          _683_v4 = _rhs135;
          _lhs95[_lhs96] = _rhs136;
          _lhs97.f7 = _rhs137;
          _677_v0 = _rhs138;
        } else {
          let _819_v126;
          _819_v126 = _dafny.Seq.UnicodeFromString("srahniy");
          _819_v126 = _dafny.Seq.Concat(((true) ? (_819_v126) : (_819_v126)), _dafny.Seq.Concat(_819_v126, (_this).f23));
          _680_v1 = _680_v1;
          let _820_v127;
          _820_v127 = _dafny.Set.fromElements(_dafny.Seq.Concat((_this).f23, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-495)), ((_821_v0) => function (_822_i4) {
            return _821_v0;
          })(_677_v0))), (_this).f23, (_this).f23, _dafny.Seq.Concat((_this).f23, _819_v126));
          let _823_v128;
          let _init27 = function (_824_i5) {
            return _dafny.Seq.UnicodeFromString("btxaxkf");
          };
          let _nw175 = Array((new BigNumber(2)).toNumber());
          for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw175.length); _i0_27++) {
            _nw175[_i0_27] = _init27(new BigNumber(_i0_27));
          }
          _823_v128 = _nw175;
          let _index135 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_823_v128).length));
          (_823_v128)[_index135] = _819_v126;
          let _825_v129;
          _825_v129 = _dafny.MultiSet.fromElements(_677_v0, _677_v0);
          let _index136 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_802_v109).length));
          let _index137 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_823_v128).length));
          let _index138 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_802_v109).length));
          let _rhs139 = _820_v127;
          let _rhs140 = ((!(!(p1)) || ((_802_v109)[_module.__default.safeIndex(new BigNumber(605), new BigNumber((_802_v109).length))])) ? (p1) : ((_825_v129).IsProperSubsetOf(_dafny.MultiSet.fromElements(new _dafny.CodePoint('w'.codePointAt(0)), _677_v0))));
          let _rhs141 = (_this).f23;
          let _rhs142 = (_this).f17;
          let _lhs98 = _802_v109;
          let _lhs99 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_802_v109).length));
          let _lhs100 = _823_v128;
          let _lhs101 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_823_v128).length));
          let _lhs102 = _802_v109;
          let _lhs103 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_802_v109).length));
          _820_v127 = _rhs139;
          _lhs98[_lhs99] = _rhs140;
          _lhs100[_lhs101] = _rhs141;
          _lhs102[_lhs103] = _rhs142;
          let _index139 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_802_v109).length));
          (_802_v109)[_index139] = (((p0).isLessThan((((_682_v3).contains((_683_v4)[_module.__default.safeIndex((_this).f22, new BigNumber((_683_v4).length))])) ? ((_682_v3).get((_683_v4)[_module.__default.safeIndex((_this).f22, new BigNumber((_683_v4).length))])) : (new BigNumber(459))))) ? ((p1) && ((_this).f17)) : ((_this).f17));
          (_this).f16 = (new BigNumber(846)).isLessThan(p0);
        }
        (globalState).f2 = (p0).minus(_module.__default.fm0(p1, globalState));
      }
      r0 = (_dafny.ZERO).minus((_dafny.ZERO).minus((new BigNumber(192)).minus(new BigNumber((_dafny.Seq.Concat((_this).f23, _dafny.Seq.Create(_module.__default.abs(new BigNumber(863)), ((_826_v0) => function (_827_i6) {
        return _826_v0;
      })(_677_v0)))).length))));
      return r0;
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _828_v0;
      _828_v0 = _dafny.Map.Empty.slice().updateUnsafe(p3,_this.f16);
      let _829_v1;
      _829_v1 = _module.D7.create_DC11(_828_v0);
      let _830_v2;
      _830_v2 = _dafny.Seq.of((_this).f17, p2);
      let _831_v3;
      _831_v3 = _module.D3.create_DC4(!((_this).f17), new BigNumber(((_this).f23).length), new BigNumber(((_829_v1).dtor_cf15).length), (_830_v2)[_module.__default.safeIndex(new BigNumber((_830_v2).length), new BigNumber((_830_v2).length))], (_this).f17);
      let _source10 = _module.D3.create_DC5(_831_v3);
      if (_source10.is_DC4) {
        let _832___mcc_h0 = (_source10).cf4;
        let _833___mcc_h1 = (_source10).cf5;
        let _834___mcc_h2 = (_source10).cf6;
        let _835___mcc_h3 = (_source10).cf7;
        let _836___mcc_h4 = (_source10).cf8;
        let _837_cf8 = _836___mcc_h4;
        let _838_cf7 = _835___mcc_h3;
        let _839_cf6 = _834___mcc_h2;
        let _840_cf5 = _833___mcc_h1;
        let _841_cf4 = _832___mcc_h0;
        _838_cf7 = _838_cf7;
        let _842_v4;
        _842_v4 = _dafny.Map.Empty.slice().updateUnsafe(false,_839_cf6);
        (globalState).f7 = new BigNumber((_dafny.Seq.Concat(_module.__default.fm52(_838_cf7, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_837_cf8,p1)).length), new BigNumber((_842_v4).length), (_this).f22, globalState), _830_v2)).length);
        let _843_v5;
        let _nw176 = Array((new BigNumber(9)).toNumber());
        _nw176[(_dafny.ZERO).toNumber()] = p3;
        _nw176[(_dafny.ONE).toNumber()] = _840_cf5;
        _nw176[(new BigNumber(2)).toNumber()] = new BigNumber(621);
        _nw176[(new BigNumber(3)).toNumber()] = p1;
        _nw176[(new BigNumber(4)).toNumber()] = p1;
        _nw176[(new BigNumber(5)).toNumber()] = _839_cf6;
        _nw176[(new BigNumber(6)).toNumber()] = (_this).f22;
        _nw176[(new BigNumber(7)).toNumber()] = (_this).f22;
        _nw176[(new BigNumber(8)).toNumber()] = (new BigNumber((_dafny.Seq.UnicodeFromString("pfxolsjs")).length)).minus(_840_cf5);
        _843_v5 = _nw176;
        let _index140 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_843_v5).length));
        (_843_v5)[_index140] = p1;
        let _844_v6;
        _844_v6 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(((_this).f23).length)));
        let _index141 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_843_v5).length));
        (_843_v5)[_index141] = (_844_v6)[_module.__default.safeIndex(p1, new BigNumber((_844_v6).length))];
        let _845_v7;
        _845_v7 = _dafny.Map.Empty.slice().updateUnsafe(true,_844_v6);
        _845_v7 = (_845_v7).update((_830_v2)[_module.__default.safeIndex(_840_cf5, new BigNumber((_830_v2).length))], _844_v6);
      } else if (_source10.is_DC3) {
        let _846___mcc_h5 = (_source10).cf3;
        let _847_cf3 = _846___mcc_h5;
        let _848_v8;
        let _nw177 = Array((new BigNumber(9)).toNumber());
        _nw177[(_dafny.ZERO).toNumber()] = _847_cf3;
        _nw177[(_dafny.ONE).toNumber()] = (_this).f17;
        _nw177[(new BigNumber(2)).toNumber()] = (_this).f17;
        _nw177[(new BigNumber(3)).toNumber()] = true;
        _nw177[(new BigNumber(4)).toNumber()] = (p1).isEqualTo(new BigNumber(935));
        _nw177[(new BigNumber(5)).toNumber()] = (_this).f17;
        _nw177[(new BigNumber(6)).toNumber()] = p2;
        _nw177[(new BigNumber(7)).toNumber()] = p2;
        _nw177[(new BigNumber(8)).toNumber()] = false;
        _848_v8 = _nw177;
        let _index142 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_848_v8).length));
        (_848_v8)[_index142] = _module.__default.fm43((_this).f17, globalState);
        let _index143 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_848_v8).length));
        (_848_v8)[_index143] = !(p1).isEqualTo(new BigNumber(((_this).f23).length));
        let _849_v9;
        let _nw178 = new _module.C2();
        _nw178.__ctor(_module.__default.fm53((_this).f22, globalState), p2, !((_module.__default.fm40((_this).f17, globalState)).contains(!(p2))));
        _849_v9 = _nw178;
        let _850_v10;
        _850_v10 = new _dafny.CodePoint('w'.codePointAt(0));
        _850_v10 = _850_v10;
        let _851_v11;
        _851_v11 = _dafny.Seq.of(new BigNumber((_module.__default.fm42(_847_cf3, new BigNumber(905), (_this).f22, _847_cf3, globalState)).length));
        let _index144 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_848_v8).length));
        (_848_v8)[_index144] = _dafny.Seq.contains(_851_v11, _module.__default.safeModuloInt(new BigNumber((_830_v2).length), (_this).f22));
      } else {
        let _852___mcc_h6 = (_source10).cf9;
        let _853_cf9 = _852___mcc_h6;
        let _854_v12;
        let _init28 = ((_855_p1) => function (_856_i0) {
          return (new BigNumber(-949)).isLessThan(_855_p1);
        })(p1);
        let _nw179 = Array((new BigNumber(12)).toNumber());
        for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw179.length); _i0_28++) {
          _nw179[_i0_28] = _init28(new BigNumber(_i0_28));
        }
        _854_v12 = _nw179;
        let _index145 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_854_v12).length));
        (_854_v12)[_index145] = _this.f16;
        let _index146 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_854_v12).length));
        (_854_v12)[_index146] = true;
        let _857_v13;
        _857_v13 = _dafny.Map.Empty.slice().updateUnsafe(_830_v2,_module.__default.safeModuloInt(p3, p3));
        _857_v13 = _857_v13;
        r0 = _module.__default.safeModuloInt((new BigNumber(-582)).plus(p3), p3);
        let _index147 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_854_v12).length));
        (_854_v12)[_index147] = (_854_v12)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_854_v12).length))];
      }
      (globalState).f2 = (_dafny.ZERO).minus(new BigNumber((function (_source11) {
        let _858___mcc_h7 = _source11;
        let _859_cf0 = _858___mcc_h7;
        return _dafny.Seq.Concat((_this).f23, (_this).f23);
      }(_830_v2)).length));
      let _860_v14;
      let _init29 = function (_861_i1) {
        return _module.__default.safeModuloInt(_861_i1, new BigNumber((_dafny.Set.fromElements((_this).f22)).length));
      };
      let _nw180 = Array((new BigNumber(24)).toNumber());
      for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw180.length); _i0_29++) {
        _nw180[_i0_29] = _init29(new BigNumber(_i0_29));
      }
      _860_v14 = _nw180;
      _860_v14 = _860_v14;
      let _862_v15;
      let _nw181 = new _module.C1();
      _nw181.__ctor(!(_this.f16) || (!(p2)), (_this).f22);
      _862_v15 = _nw181;
      let _863_v16;
      _863_v16 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_this).f22);
      _863_v16 = (_863_v16).update(new BigNumber(606), p3);
      r0 = (_dafny.ZERO).minus(p3);
      r0 = p3;
      return r0;
    }
    m11(p0, p1, p2, p3, globalState) {
      let _this = this;
      (_this).f16 = !((_this).f17) || (_module.__default.fm31(_this.f16, _this.f16, globalState));
      let _864_v0;
      let _nw182 = Array((new BigNumber(29)).toNumber()).fill(false);
      _864_v0 = _nw182;
      let _index148 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length));
      (_864_v0)[_index148] = (_this).f17;
      let _index149 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length));
      (_864_v0)[_index149] = _module.__default.fm31(_this.f16, true, globalState);
      if (!((_864_v0)[_module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length))]) || (_this.f16)) {
        (globalState).f2 = _module.__default.safeModuloInt((_this).f22, (p0).plus((_this).f22));
        let _index150 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length));
        (_864_v0)[_index150] = p1;
        let _865_v1;
        let _nw183 = new _module.C1();
        _nw183.__ctor((_864_v0)[_module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length))], (p0).plus(p0));
        _865_v1 = _nw183;
        _865_v1 = _865_v1;
        let _866_v2;
        _866_v2 = _dafny.Map.Empty.slice().updateUnsafe(false,p1);
        let _867_v3;
        _867_v3 = _dafny.Map.Empty.slice().updateUnsafe(_866_v2,new BigNumber(844));
        (globalState).f7 = ((_dafny.ZERO).minus((((_867_v3).contains(_866_v2)) ? ((_867_v3).get(_866_v2)) : ((_dafny.ZERO).minus(new BigNumber(((_this).f23).length)))))).multipliedBy(_module.__default.fm0((_864_v0)[_module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length))], globalState));
        let _868_v4;
        _868_v4 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f17);
        let _869_v5;
        _869_v5 = _dafny.Seq.of(_868_v4);
        let _870_v6;
        _870_v6 = _module.D7.create_DC11(((_869_v5)[_module.__default.safeIndex((_this).f22, new BigNumber((_869_v5).length))]).Merge(_868_v4));
        _870_v6 = _870_v6;
      } else {
        let _rhs143 = _this.f16;
        let _rhs144 = (_dafny.ZERO).minus(p0);
        let _lhs104 = _this;
        let _lhs105 = globalState;
        _lhs104.f16 = _rhs143;
        _lhs105.f7 = _rhs144;
        let _index151 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length));
        (_864_v0)[_index151] = (_864_v0)[_module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length))];
        let _871_v7;
        _871_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,_dafny.Map.Empty.slice().updateUnsafe(p0,true));
        let _872_v8;
        _872_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,(_864_v0)[_module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length))]);
        let _pat_let_tv7 = _872_v8;
        let _source12 = function (_pat_let9_0) {
          return function (_873_dt__update__tmp_h0) {
            return function (_pat_let10_0) {
              return function (_874_dt__update_hcf15_h0) {
                return _module.D7.create_DC11(_874_dt__update_hcf15_h0);
              }(_pat_let10_0);
            }(_pat_let_tv7);
          }(_pat_let9_0);
        }(_module.D7.create_DC11((((_871_v7).contains(p1)) ? ((_871_v7).get(p1)) : (_872_v8))));
        if (_source12.is_DC12) {
          (globalState).f7 = (_this).f22;
          let _875_v9;
          _875_v9 = new _dafny.CodePoint('h'.codePointAt(0));
          (_this).f16 = (new BigNumber((((p1) ? ((_this).f23) : (_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(719)), function (_876_i0) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          }), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(719)), function (_877_i0) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          })).length)), _875_v9)))).length)).isLessThanOrEqualTo((_this).f22);
          let _index152 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length));
          (_864_v0)[_index152] = (_864_v0)[_module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length))];
          let _878_v10;
          let _init30 = ((_879_p1) => function (_880_i1) {
            return _dafny.Set.fromElements(_879_p1, _this.f16);
          })(p1);
          let _nw184 = Array((new BigNumber(24)).toNumber());
          for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw184.length); _i0_30++) {
            _nw184[_i0_30] = _init30(new BigNumber(_i0_30));
          }
          _878_v10 = _nw184;
          let _881_v11;
          let _nw185 = new _module.C0();
          _nw185.__ctor(_878_v10, (_this).f22);
          _881_v11 = _nw185;
        } else {
          let _882___mcc_h0 = (_source12).cf15;
          let _883_cf15 = _882___mcc_h0;
          let _884_v12;
          _884_v12 = _dafny.Seq.of((_this).f17, (_864_v0)[_module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length))], (_this).f17, (((_this).f17) ? (true) : (false)), (p1) || (_this.f16));
          let _885_v13;
          _885_v13 = _dafny.Seq.of(new BigNumber(344), (_this).f22);
          _884_v12 = _module.__default.fm52((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(734)), ((_886_p0) => function (_887_i2) {
            return ((_this).f23)[_module.__default.safeIndex(_886_p0, new BigNumber(((_this).f23).length))];
          })(p0))).length)).isLessThan(p0), (_this).f22, (((_this).f17) ? (new BigNumber((_885_v13).length)) : ((_this).f22)), _module.__default.fm0(_module.__default.fm31((_864_v0)[_module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length))], (_864_v0)[_module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length))], globalState), globalState), globalState);
          let _888_v14;
          let _init31 = ((_889_v12) => function (_890_i3) {
            return _889_v12;
          })(_884_v12);
          let _nw186 = Array((new BigNumber(15)).toNumber());
          for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw186.length); _i0_31++) {
            _nw186[_i0_31] = _init31(new BigNumber(_i0_31));
          }
          _888_v14 = _nw186;
          let _index153 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_888_v14).length));
          (_888_v14)[_index153] = _884_v12;
          let _index154 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_888_v14).length));
          (_888_v14)[_index154] = _dafny.Seq.Concat(_884_v12, _884_v12);
          (_this).f16 = (_this).f17;
          let _891_v15;
          _891_v15 = _dafny.Set.fromElements((_864_v0)[_module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length))], false);
          let _892_v16;
          _892_v16 = _module.D15.create_DC29(_891_v15);
          (globalState).f7 = new BigNumber(((_891_v15).Difference(((_892_v16).dtor_cf43).Difference(_891_v15))).length);
        }
        (globalState).f7 = ((_this).f22).minus((_this).f22);
        let _893_v17;
        _893_v17 = _dafny.Set.fromElements((_this).f17);
        let _894_v18;
        _894_v18 = _module.D15.create_DC29(_893_v17);
        let _895_v19;
        _895_v19 = _module.D15.create_DC31(_894_v18);
        _895_v19 = _module.D15.create_DC31(_894_v18);
      }
      let _896_v20;
      _896_v20 = _dafny.Map.Empty.slice().updateUnsafe(_module.D13.create_DC26(),_module.__default.safeDivisionInt((_dafny.ZERO).minus((_this).f22), p0));
      _896_v20 = (_896_v20).Merge(_896_v20);
      let _index155 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length));
      (_864_v0)[_index155] = (p0).isLessThan((_this).f22);
      let _pat_let_tv8 = p1;
      let _pat_let_tv9 = _864_v0;
      let _pat_let_tv10 = _864_v0;
      let _index156 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length));
      let _index157 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length));
      let _index158 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length));
      let _rhs145 = function (_source13) {
        if (_source13.is_DC30) {
          let _897___mcc_h1 = (_source13).cf44;
          let _898_cf44 = _897___mcc_h1;
          return _pat_let_tv8;
        } else if (_source13.is_DC29) {
          let _899___mcc_h2 = (_source13).cf43;
          let _900_cf43 = _899___mcc_h2;
          return !((_pat_let_tv10)[_module.__default.safeIndex(new BigNumber(852), new BigNumber((_pat_let_tv9).length))]);
        } else {
          let _901___mcc_h3 = (_source13).cf45;
          let _902_cf45 = _901___mcc_h3;
          return _this.f16;
        }
      }(_module.D15.create_DC29(_module.__default.fm44((_this).f17, (_864_v0)[_module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length))], globalState)));
      let _rhs146 = (_dafny.ZERO).minus((_dafny.ZERO).minus((((_864_v0)[_module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length))]) ? ((_dafny.ZERO).minus((p0).minus(p0))) : ((p0).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-710)), function (_903_i4) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      })).length))))));
      let _rhs147 = (_864_v0)[_module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length))];
      let _rhs148 = _module.__default.fm0(p1, globalState);
      let _rhs149 = true;
      let _lhs106 = _864_v0;
      let _lhs107 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length));
      let _lhs108 = globalState;
      let _lhs109 = _864_v0;
      let _lhs110 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length));
      let _lhs111 = globalState;
      let _lhs112 = _864_v0;
      let _lhs113 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_864_v0).length));
      _lhs106[_lhs107] = _rhs145;
      _lhs108.f2 = _rhs146;
      _lhs109[_lhs110] = _rhs147;
      _lhs111.f2 = _rhs148;
      _lhs112[_lhs113] = _rhs149;
      return;
    }
    get f22() {
      let _this = this;
      return _this._f22;
    };
    get f23() {
      let _this = this;
      return _this._f23;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f16 = false;
      this._f17 = false;
      this._f21 = false;
    }
    _parentTraits() {
      return [_module.T2, _module.T1];
    }
    get f16() {
      let _this = this;
      return _this._f16;
    };
    set f16(value) {
      let _this = this;
      _this._f16 = value;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    __ctor(f21, f16, f17) {
      let _this = this;
      (_this)._f21 = f21;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      return;
    }
    fm12(p0, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(((((_this).f17) ? (_dafny.Seq.UnicodeFromString("bsrmhd")) : (_dafny.Seq.UnicodeFromString("bv")))).length), new BigNumber(((_dafny.Set.fromElements(_this.f16, false)).Difference(_dafny.Set.fromElements(!(_this.f16)))).length))))));
    };
    fm13(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of((_this).f21, (_this).f17), _dafny.Seq.of(_this.f16)), _dafny.Seq.Concat(_dafny.Seq.of((_this).f21, (_this).f21), _dafny.Seq.of(true, !((_this).f17))));
    };
    fm3(globalState) {
      let _this = this;
      return (function () {
        let _coll42 = new _dafny.Map();
        for (const _compr_42 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("teahcor")).length))).length),_dafny.Seq.of(new BigNumber(-549)))).Keys.Elements) {
          let _904_v0 = _compr_42;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("teahcor")).length))).length),_dafny.Seq.of(new BigNumber(-549)))).contains(_904_v0)) {
            _coll42.push([(_904_v0).minus(new BigNumber(912)),new BigNumber(-8)]);
          }
        }
        return _coll42;
      }()).Merge((function () {
        let _coll43 = new _dafny.Map();
        for (const _compr_43 of (_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-808)))).Elements) {
          let _905_v1 = _compr_43;
          if (_dafny.Seq.contains(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-808))), _905_v1)) {
            _coll43.push([_module.__default.safeModuloInt(_905_v1, new BigNumber((_dafny.Seq.of(new BigNumber(922), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(309)), function (_906_i0) {
              return new _dafny.CodePoint('y'.codePointAt(0));
            })).length)), new BigNumber(380))).length)),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of((_this).f21, _this.f16, _this.f16, (_this).f17, true)).length))).cardinality())]);
          }
        }
        return _coll43;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements((_this).f21, (_this).f17, false)).cardinality()),new BigNumber((_dafny.Set.fromElements((_this).f17, _this.f16)).length))).length),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(951))).cardinality()))));
    };
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.Seq.UnicodeFromString("");
      let r3 = false;
      let _907_v0;
      _907_v0 = _dafny.Seq.of(p2);
      let _908_v1;
      _908_v1 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_907_v0).length)),p2);
      r1 = _module.__default.fm27(p2, p2, (p2).plus(new BigNumber((_908_v1).length)), _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(p1, globalState),p2), globalState);
      if (true) {
        r3 = (p2).isLessThan((_dafny.ZERO).minus(new BigNumber((_907_v0).length)));
        let _909_v2;
        _909_v2 = _dafny.Seq.of(!(p0), (_this).f21, (_this).f17, _this.f16);
        let _910_v3;
        _910_v3 = _909_v2;
        let _911_v4;
        _911_v4 = _module.D7.create_DC12();
        let _912_v5;
        _912_v5 = _dafny.Map.Empty.slice().updateUnsafe(_910_v3,_911_v4);
        if (!(_912_v5).contains(_910_v3)) {
          let _913_v6;
          _913_v6 = _dafny.Seq.UnicodeFromString("ojvcuvnba");
          r2 = _913_v6;
          let _914_v7;
          let _nw187 = Array((new BigNumber(7)).toNumber());
          _nw187[(_dafny.ZERO).toNumber()] = _909_v2;
          _nw187[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_909_v2, _909_v2);
          _nw187[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_909_v2, _909_v2);
          _nw187[(new BigNumber(3)).toNumber()] = _909_v2;
          _nw187[(new BigNumber(4)).toNumber()] = _909_v2;
          _nw187[(new BigNumber(5)).toNumber()] = _909_v2;
          _nw187[(new BigNumber(6)).toNumber()] = (((_this).f21) ? (_909_v2) : (_909_v2));
          _914_v7 = _nw187;
          _914_v7 = _914_v7;
          let _915_v8;
          _915_v8 = _dafny.Seq.of(_913_v6, _dafny.Seq.UnicodeFromString("dnpswr"));
          let _916_v9;
          _916_v9 = _dafny.Seq.of((_915_v8)[_module.__default.safeIndex(p2, new BigNumber((_915_v8).length))]);
          let _917_v10;
          _917_v10 = _dafny.Set.fromElements(p2);
          let _918_v11;
          _918_v11 = _dafny.MultiSet.fromElements(_this.f16, (_this).f17);
          let _919_v12;
          let _nw188 = Array((new BigNumber(13)).toNumber());
          _nw188[(_dafny.ZERO).toNumber()] = _this.f16;
          _nw188[(_dafny.ONE).toNumber()] = _dafny.areEqual((_916_v9)[_module.__default.safeIndex(new BigNumber(-154), new BigNumber((_916_v9).length))], _913_v6);
          _nw188[(new BigNumber(2)).toNumber()] = (_this).f21;
          _nw188[(new BigNumber(3)).toNumber()] = (_this).f17;
          _nw188[(new BigNumber(4)).toNumber()] = (_this).f21;
          _nw188[(new BigNumber(5)).toNumber()] = (p2).isLessThanOrEqualTo(p2);
          _nw188[(new BigNumber(6)).toNumber()] = p1;
          _nw188[(new BigNumber(7)).toNumber()] = ((!(true)) ? (p0) : ((_this).f17));
          _nw188[(new BigNumber(8)).toNumber()] = p0;
          _nw188[(new BigNumber(9)).toNumber()] = _module.__default.fm27(new BigNumber(254), (_this).fm12(new BigNumber((_917_v10).length), globalState), p2, _dafny.Map.Empty.slice().updateUnsafe(p2,p2), globalState);
          _nw188[(new BigNumber(10)).toNumber()] = p1;
          _nw188[(new BigNumber(11)).toNumber()] = !(_this.f16);
          _nw188[(new BigNumber(12)).toNumber()] = (_909_v2)[_module.__default.safeIndex(new BigNumber((_918_v11).cardinality()), new BigNumber((_909_v2).length))];
          _919_v12 = _nw188;
          let _index159 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_919_v12).length));
          (_919_v12)[_index159] = p1;
          let _920_v13;
          let _nw189 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _920_v13 = _nw189;
          let _921_v14;
          _921_v14 = _dafny.Set.fromElements(p1);
          let _922_v15;
          _922_v15 = _module.D3.create_DC4(false, new BigNumber(42), p2, (_this).f21, true);
          let _923_v16;
          _923_v16 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,p1);
          let _index160 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_919_v12).length));
          let _rhs150 = _919_v12;
          let _rhs151 = ((_921_v14).Union(_921_v14)).IsProperSubsetOf(_dafny.Set.fromElements(false, !(p0)));
          let _rhs152 = _module.__default.safeDivisionInt(p2, new BigNumber(256));
          let _rhs153 = _module.__default.fm28(false, _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(p2))).length)), _922_v15, ((_923_v16).update(!(true), !((_this).f21))).Merge(_923_v16), globalState);
          let _rhs154 = _920_v13;
          let _lhs114 = _919_v12;
          let _lhs115 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_919_v12).length));
          let _lhs116 = globalState;
          _919_v12 = _rhs150;
          _lhs114[_lhs115] = _rhs151;
          _lhs116.f2 = _rhs152;
          _913_v6 = _rhs153;
          _920_v13 = _rhs154;
          (globalState).f7 = p2;
          let _924_v17;
          _924_v17 = _module.D3.create_DC3(_dafny.areEqual(_dafny.Seq.UnicodeFromString("vgncwsm"), _913_v6));
          let _rhs155 = ((!(p0) || (p0)) ? (_module.__default.fm29(globalState)) : (_924_v17));
          let _rhs156 = _dafny.Seq.UnicodeFromString("ebjbv");
          _924_v17 = _rhs155;
          _913_v6 = _rhs156;
        } else {
          let _925_v18;
          _925_v18 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f16);
          let _926_v19;
          _926_v19 = _dafny.Set.fromElements(p2);
          let _927_v20;
          _927_v20 = _dafny.MultiSet.fromElements(new BigNumber((_module.__default.fm28((_this).f21, _908_v1, _module.D3.create_DC4(p1, p2, p2, (_this).f17, (_this).f17), _925_v18, globalState)).length), new BigNumber((_926_v19).length), p2);
          let _928_v21;
          let _nw190 = Array((new BigNumber(4)).toNumber()).fill(false);
          _928_v21 = _nw190;
          let _929_v22;
          _929_v22 = _dafny.MultiSet.fromElements(_928_v21, _928_v21);
          let _930_v23;
          _930_v23 = _dafny.Map.Empty.slice().updateUnsafe(p2,_927_v20);
          let _931_v24;
          let _nw191 = Array((new BigNumber(23)).toNumber());
          _nw191[(_dafny.ZERO).toNumber()] = (_dafny.MultiSet.fromElements(p2, p2)).update((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-855)), ((_932_p3, _933_p2) => function (_934_i0) {
            return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_932_p3,_dafny.Seq.of(_933_p2, _933_p2, _933_p2, new BigNumber(673)))).length);
          })(p3, p2))).length)), _module.__default.abs(p2));
          _nw191[(_dafny.ONE).toNumber()] = _927_v20;
          _nw191[(new BigNumber(2)).toNumber()] = (_dafny.MultiSet.fromElements(p2)).Union(_dafny.MultiSet.fromElements(new BigNumber((_907_v0).length), new BigNumber((_dafny.Seq.UnicodeFromString("bq")).length), p2));
          _nw191[(new BigNumber(3)).toNumber()] = _927_v20;
          _nw191[(new BigNumber(4)).toNumber()] = _927_v20;
          _nw191[(new BigNumber(5)).toNumber()] = (_927_v20).Difference((_927_v20).update(_module.__default.fm0(p0, globalState), _module.__default.abs(new BigNumber((_909_v2).length))));
          _nw191[(new BigNumber(6)).toNumber()] = (_927_v20).Union(_927_v20);
          _nw191[(new BigNumber(7)).toNumber()] = _927_v20;
          _nw191[(new BigNumber(8)).toNumber()] = _927_v20;
          _nw191[(new BigNumber(9)).toNumber()] = _927_v20;
          _nw191[(new BigNumber(10)).toNumber()] = _927_v20;
          _nw191[(new BigNumber(11)).toNumber()] = _dafny.MultiSet.fromElements(p2, new BigNumber((_929_v22).cardinality()), p2);
          _nw191[(new BigNumber(12)).toNumber()] = _927_v20;
          _nw191[(new BigNumber(13)).toNumber()] = _927_v20;
          _nw191[(new BigNumber(14)).toNumber()] = (_927_v20).Union(_927_v20);
          _nw191[(new BigNumber(15)).toNumber()] = (((_930_v23).contains(p2)) ? ((_930_v23).get(p2)) : (_927_v20));
          _nw191[(new BigNumber(16)).toNumber()] = _dafny.MultiSet.fromElements(p2);
          _nw191[(new BigNumber(17)).toNumber()] = _module.__default.fm30(_this.f16, globalState);
          _nw191[(new BigNumber(18)).toNumber()] = _927_v20;
          _nw191[(new BigNumber(19)).toNumber()] = _dafny.MultiSet.FromArray(_907_v0);
          _nw191[(new BigNumber(20)).toNumber()] = _dafny.MultiSet.fromElements(p2, p2);
          _nw191[(new BigNumber(21)).toNumber()] = _927_v20;
          _nw191[(new BigNumber(22)).toNumber()] = _927_v20;
          _931_v24 = _nw191;
          let _index161 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_931_v24).length));
          (_931_v24)[_index161] = _927_v20;
          let _index162 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_931_v24).length));
          let _rhs157 = (_dafny.MultiSet.FromArray(_907_v0)).Union(_927_v20);
          let _rhs158 = p2;
          let _lhs117 = _931_v24;
          let _lhs118 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_931_v24).length));
          let _lhs119 = globalState;
          _lhs117[_lhs118] = _rhs157;
          _lhs119.f7 = _rhs158;
          let _935_v25;
          _935_v25 = _dafny.Seq.UnicodeFromString("fk");
          let _936_v26;
          let _nw192 = new _module.C3();
          _nw192.__ctor(p2, _935_v25, false, !(true) || ((_this).f17));
          _936_v26 = _nw192;
          let _index163 = _module.__default.safeIndex(new BigNumber(16), new BigNumber((_928_v21).length));
          (_928_v21)[_index163] = (_this).f21;
          let _index164 = _module.__default.safeIndex(new BigNumber(16), new BigNumber((_928_v21).length));
          (_928_v21)[_index164] = p1;
          let _937_v27;
          _937_v27 = _dafny.Map.Empty.slice().updateUnsafe(!((_928_v21)[_module.__default.safeIndex(new BigNumber(16), new BigNumber((_928_v21).length))]),p2);
          let _938_v28;
          _938_v28 = _dafny.Map.Empty.slice().updateUnsafe(_937_v27,new BigNumber(((_module.__default.fm32(p0, true, p1, p0, globalState)).Merge(_937_v27)).length));
          let _939_v30;
          _939_v30 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_this.f16,p2));
          _938_v28 = (_938_v28).Merge((_938_v28).Merge(function () {
            let _coll44 = new _dafny.Map();
            for (const _compr_44 of (_939_v30).Elements) {
              let _940_v29 = _compr_44;
              if ((_939_v30).contains(_940_v29)) {
                _coll44.push([_940_v29,(_936_v26).f22]);
              }
            }
            return _coll44;
          }()));
          let _index165 = _module.__default.safeIndex(new BigNumber(16), new BigNumber((_928_v21).length));
          (_928_v21)[_index165] = _this.f16;
        }
        (globalState).f7 = p2;
        (globalState).f2 = p2;
        (globalState).f2 = _module.__default.safeDivisionInt(p2, p2);
      } else {
        let _941_v31;
        _941_v31 = _dafny.Map.Empty.slice().updateUnsafe(((p0) ? ((_this).f17) : (_module.__default.fm27(new BigNumber(137), p2, p2, _908_v1, globalState))),_module.__default.fm43(_this.f16, globalState));
        let _942_v33;
        _942_v33 = _dafny.Map.Empty.slice().updateUnsafe(p2,_908_v1);
        let _943_v34;
        _943_v34 = _module.D11.create_DC19(_908_v1);
        let _944_v35;
        _944_v35 = _dafny.Seq.of(function () {
          let _coll45 = new _dafny.Map();
          for (const _compr_45 of _dafny.IntegerRange(new BigNumber(650), new BigNumber(-586))) {
            let _945_v32 = _compr_45;
            if (((new BigNumber(650)).isLessThanOrEqualTo(_945_v32)) && ((_945_v32).isLessThan(new BigNumber(-586)))) {
              _coll45.push([(_945_v32).minus(new BigNumber(-436)),p2]);
            }
          }
          return _coll45;
        }(), (((_942_v33).contains(p2)) ? ((_942_v33).get(p2)) : ((_943_v34).dtor_cf27)));
        _941_v31 = (_941_v31).update(!(_dafny.Seq.contains(_944_v35, _908_v1)), _this.f16);
        let _946_v36;
        _946_v36 = _module.D3.create_DC3((_this).f17);
        let _947_v37;
        let _nw193 = new _module.C2();
        _nw193.__ctor(_946_v36, true, true);
        _947_v37 = _nw193;
        (globalState).f2 = _module.__default.fm0(p1, globalState);
        let _948_v38;
        _948_v38 = _dafny.Seq.UnicodeFromString("gx");
        r2 = _dafny.Seq.Concat(_948_v38, _948_v38);
        let _949_v39;
        _949_v39 = _dafny.Set.fromElements((_this).f17, p0, (_this).f17);
        let _950_v40;
        _950_v40 = _module.D13.create_DC26();
        let _951_v41;
        _951_v41 = _dafny.Set.fromElements(_950_v40);
        let _952_v42;
        _952_v42 = _module.D3.create_DC4((_this).f17, ((_module.__default.fm43((_this).f17, globalState)) ? (new BigNumber((_949_v39).length)) : (new BigNumber((_951_v41).length))), p2, (_this).f17, p1);
        let _953_v43;
        _953_v43 = _dafny.MultiSet.fromElements(!((((_941_v31).contains((_this).f17)) ? ((_941_v31).get((_this).f17)) : (_this.f16))));
        let _954_v44;
        _954_v44 = _dafny.MultiSet.fromElements(p2);
        let _rhs159 = _907_v0;
        let _rhs160 = ((p1) ? (false) : (_this.f16));
        let _rhs161 = (((new BigNumber((_953_v43).cardinality())).isLessThan(p2)) ? (_module.D3.create_DC4(p0, p2, p2, (_this).f17, p1)) : (_module.D3.create_DC4(!(_module.__default.fm43(_this.f16, globalState)), new BigNumber((_954_v44).cardinality()), (_dafny.ZERO).minus(p2), p0, p1)));
        let _rhs162 = p2;
        let _rhs163 = !(!(false)) || (_this.f16);
        let _lhs120 = _this;
        let _lhs121 = globalState;
        let _lhs122 = _this;
        _907_v0 = _rhs159;
        _lhs120.f16 = _rhs160;
        _952_v42 = _rhs161;
        _lhs121.f2 = _rhs162;
        _lhs122.f16 = _rhs163;
      }
      let _955_v45;
      _955_v45 = _dafny.Seq.of(true);
      let _956_v46;
      _956_v46 = new _dafny.CodePoint('o'.codePointAt(0));
      let _rhs164 = (((_dafny.MultiSet.fromElements(p2)).contains(p2)) ? ((_955_v45)) : (_955_v45));
      let _rhs165 = _956_v46;
      _955_v45 = _rhs164;
      _956_v46 = _rhs165;
      if ((new BigNumber((_dafny.MultiSet.fromElements((_this).f17)).cardinality())).isEqualTo(p2)) {
        r3 = _this.f16;
        (globalState).f7 = p2;
        let _957_v47;
        let _init32 = ((_958_p1, _959_p2, _960_p0) => function (_961_i1) {
          return _module.D9.create_DC16(new BigNumber((_dafny.MultiSet.fromElements(_958_p1)).cardinality()), (_module.D9.create_DC16(_959_p2, _960_p0)).dtor_cf24);
        })(p1, p2, p0);
        let _nw194 = Array((new BigNumber(13)).toNumber());
        for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw194.length); _i0_32++) {
          _nw194[_i0_32] = _init32(new BigNumber(_i0_32));
        }
        _957_v47 = _nw194;
        _957_v47 = ((p1) ? (_957_v47) : (_957_v47));
        _956_v46 = p3;
        r0 = p2;
      } else {
        if ((_this).f21) {
          r1 = ((true) ? ((_this).f21) : ((p2).isEqualTo((_dafny.ZERO).minus(p2))));
          let _962_v48;
          _962_v48 = _dafny.Map.Empty.slice().updateUnsafe(_907_v0,p2);
          let _963_v49;
          _963_v49 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f21);
          let _964_v50;
          _964_v50 = _dafny.MultiSet.fromElements(p2);
          let _965_v51;
          let _nw195 = Array((new BigNumber(4)).toNumber());
          _nw195[(_dafny.ZERO).toNumber()] = (_this).fm12(p2, globalState);
          _nw195[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_962_v48).length),p2)).length)).multipliedBy(new BigNumber((_963_v49).length)));
          _nw195[(new BigNumber(2)).toNumber()] = p2;
          _nw195[(new BigNumber(3)).toNumber()] = (p2).minus(new BigNumber((_964_v50).cardinality()));
          _965_v51 = _nw195;
          let _index166 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_965_v51).length));
          (_965_v51)[_index166] = p2;
          let _index167 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_965_v51).length));
          (_965_v51)[_index167] = p2;
          let _966_v53;
          _966_v53 = _dafny.MultiSet.fromElements((_this).f17);
          let _967_v54;
          _967_v54 = _dafny.Seq.of((_963_v49).Merge(function () {
            let _coll46 = new _dafny.Map();
            for (const _compr_46 of _dafny.IntegerRange(new BigNumber(533), new BigNumber(487))) {
              let _968_v52 = _compr_46;
              if (((new BigNumber(533)).isLessThanOrEqualTo(_968_v52)) && ((_968_v52).isLessThan(new BigNumber(487)))) {
                _coll46.push([_module.__default.safeModuloInt(_968_v52, new BigNumber((_908_v1).length)),p1]);
              }
            }
            return _coll46;
          }()), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_966_v53).cardinality()),_this.f16));
          _967_v54 = _967_v54;
          r3 = !(((_this).f21) || (!(_this.f16)));
          let _969_v55;
          let _nw196 = Array((new BigNumber(15)).toNumber()).fill([]);
          _969_v55 = _nw196;
          _969_v55 = _969_v55;
        } else {
          (globalState).f2 = p2;
          let _970_v56;
          let _nw197 = new _module.C1();
          _nw197.__ctor(_this.f16, p2);
          _970_v56 = _nw197;
          let _971_v57;
          _971_v57 = _dafny.MultiSet.fromElements(_970_v56);
          _971_v57 = _971_v57;
          (_this).f16 = true;
          let _972_v58;
          _972_v58 = _dafny.Seq.UnicodeFromString("gdfcanlo");
          let _973_v59;
          let _nw198 = Array((new BigNumber(8)).toNumber());
          _nw198[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_972_v58).length), p2);
          _nw198[(_dafny.ONE).toNumber()] = p2;
          _nw198[(new BigNumber(2)).toNumber()] = _module.__default.safeModuloInt(p2, p2);
          _nw198[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus((p2).minus(p2));
          _nw198[(new BigNumber(4)).toNumber()] = new BigNumber((_972_v58).length);
          _nw198[(new BigNumber(5)).toNumber()] = p2;
          _nw198[(new BigNumber(6)).toNumber()] = (p2).multipliedBy(p2);
          _nw198[(new BigNumber(7)).toNumber()] = new BigNumber((_972_v58).length);
          _973_v59 = _nw198;
          _973_v59 = _973_v59;
          let _974_v60;
          _974_v60 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
          let _975_v61;
          _975_v61 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_974_v60).length)),_this.f16);
          let _rhs166 = _970_v56.f25;
          let _rhs167 = ((_dafny.Seq.contains(_972_v58, p3)) ? ((_this).f21) : (false));
          let _rhs168 = !(_975_v61).equals(_975_v61);
          let _rhs169 = (_dafny.ZERO).minus(p2);
          let _lhs123 = _this;
          let _lhs124 = _970_v56;
          let _lhs125 = globalState;
          r3 = _rhs166;
          _lhs123.f16 = _rhs167;
          _lhs124.f25 = _rhs168;
          _lhs125.f2 = _rhs169;
        }
        let _976_v62;
        let _nw199 = new _module.C1();
        _nw199.__ctor(_this.f16, p2);
        _976_v62 = _nw199;
        let _977_v63;
        let _nw200 = new _module.C1();
        _nw200.__ctor(_976_v62.f25, p2);
        _977_v63 = _nw200;
        let _978_v64;
        _978_v64 = _module.D3.create_DC3(_976_v62.f25);
        let _979_v65;
        let _nw201 = new _module.C2();
        _nw201.__ctor(_978_v64, p1, false);
        _979_v65 = _nw201;
        let _rhs170 = _976_v62;
        let _rhs171 = _this.f16;
        let _rhs172 = _979_v65;
        let _lhs126 = _976_v62;
        _977_v63 = _rhs170;
        _lhs126.f25 = _rhs171;
        _979_v65 = _rhs172;
        let _980_v66;
        let _nw202 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _980_v66 = _nw202;
        let _981_v67;
        _981_v67 = _dafny.Map.Empty.slice().updateUnsafe(p0,_980_v66);
        let _982_v68;
        _982_v68 = _dafny.Set.fromElements(_980_v66, _980_v66, (((_981_v67).contains(_module.__default.fm43(_this.f16, globalState))) ? ((_981_v67).get(_module.__default.fm43(_this.f16, globalState))) : (_980_v66)));
        _982_v68 = _982_v68;
        (globalState).f7 = (p2).minus(_module.__default.safeDivisionInt(new BigNumber(598), p2));
      }
      (_this).f16 = _module.__default.fm43((_this).f21, globalState);
      _955_v45 = _955_v45;
      r0 = p2;
      let _983_v69;
      _983_v69 = _dafny.MultiSet.fromElements(_module.D15.create_DC30(_956_v46));
      let _984_v71;
      _984_v71 = _dafny.Set.fromElements(_983_v69, _983_v69, _983_v69, _983_v69, _dafny.MultiSet.fromElements(_module.D15.create_DC30(_956_v46)));
      r1 = !(!(function () {
        let _coll47 = new _dafny.Map();
        for (const _compr_47 of (_984_v71).Elements) {
          let _985_v70 = _compr_47;
          if ((_984_v71).contains(_985_v70)) {
            _coll47.push([_985_v70,new BigNumber(374)]);
          }
        }
        return _coll47;
      }()).contains(_983_v69)) || ((_this).f21);
      let _986_v72;
      _986_v72 = _dafny.Seq.UnicodeFromString("wo");
      r2 = _dafny.Seq.Concat(_986_v72, _986_v72);
      r3 = true;
      return [r0, r1, r2, r3];
    }
    m2(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _987_v0;
      let _nw203 = Array((new BigNumber(6)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _987_v0 = _nw203;
      let _988_v1;
      _988_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_987_v0);
      let _989_v2;
      _989_v2 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
      let _990_v3;
      _990_v3 = _dafny.Map.Empty.slice().updateUnsafe(_988_v1,(((_989_v2).contains(p0)) ? ((_989_v2).get(p0)) : ((_this).f21)));
      let _991_v4;
      _991_v4 = _dafny.Seq.of(_this.f16, (((_990_v3).contains(_988_v1)) ? ((_990_v3).get(_988_v1)) : (_this.f16)));
      let _992_v5;
      _992_v5 = _991_v4;
      let _source14 = ((false) ? (_992_v5) : (_992_v5));
      let _993___mcc_h0 = _source14;
      let _994_cf0 = _993___mcc_h0;
      if ((_this).f21) {
        let _995_v6;
        let _nw204 = new _module.C1();
        _nw204.__ctor(_this.f16, _module.__default.fm0((_this).f21, globalState));
        _995_v6 = _nw204;
        let _996_v7;
        _996_v7 = _module.D14.create_DC28(p0, (_this).f17, _995_v6);
        let _rhs173 = _module.__default.fm0(p1, globalState);
        let _rhs174 = p1;
        let _rhs175 = (_991_v4)[_module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_991_v4).length))];
        let _rhs176 = (_996_v7).dtor_cf41;
        let _lhs127 = globalState;
        let _lhs128 = _this;
        let _lhs129 = _this;
        let _lhs130 = _this;
        _lhs127.f2 = _rhs173;
        _lhs128.f16 = _rhs174;
        _lhs129.f16 = _rhs175;
        _lhs130.f16 = _rhs176;
        _991_v4 = _994_cf0;
        (_this).f16 = true;
        let _997_v8;
        _997_v8 = new _dafny.CodePoint('f'.codePointAt(0));
        _997_v8 = _997_v8;
        let _998_v9;
        let _999_v10;
        let _1000_v11;
        let _out20;
        let _out21;
        let _out22;
        let _outcollector7 = (_this).m10((_this).f17, _module.__default.safeModuloInt(_995_v6.f14, new BigNumber((_dafny.Seq.of(_dafny.Seq.of(_995_v6.f14))).length)), (p0).plus(_995_v6.f14), p1, globalState);
        _out20 = _outcollector7[0];
        _out21 = _outcollector7[1];
        _out22 = _outcollector7[2];
        _998_v9 = _out20;
        _999_v10 = _out21;
        _1000_v11 = _out22;
      } else {
        let _1001_v12;
        let _init33 = function (_1002_i0) {
          return _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("e"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(257)), function (_1003_i1) {
            return new _dafny.CodePoint('x'.codePointAt(0));
          }));
        };
        let _nw205 = Array((new BigNumber(18)).toNumber());
        for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw205.length); _i0_33++) {
          _nw205[_i0_33] = _init33(new BigNumber(_i0_33));
        }
        _1001_v12 = _nw205;
        let _index168 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_1001_v12).length));
        (_1001_v12)[_index168] = (_this).f17;
        let _index169 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_1001_v12).length));
        let _rhs177 = _1001_v12;
        let _rhs178 = _this.f16;
        let _lhs131 = _1001_v12;
        let _lhs132 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_1001_v12).length));
        _1001_v12 = _rhs177;
        _lhs131[_lhs132] = _rhs178;
        let _1004_v13;
        let _nw206 = new _module.C1();
        _nw206.__ctor(false, p0);
        _1004_v13 = _nw206;
        let _1005_v14;
        let _nw207 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1005_v14 = _nw207;
        let _1006_v15;
        _1006_v15 = _dafny.Seq.UnicodeFromString("uiwdx");
        let _index170 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_1005_v14).length));
        (_1005_v14)[_index170] = _1006_v15;
        let _1007_v16;
        _1007_v16 = _module.D12.create_DC22(_dafny.Seq.Create(_module.__default.abs(new BigNumber(385)), function (_1008_i2) {
  return new _dafny.CodePoint('a'.codePointAt(0));
}));
        let _index171 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_1005_v14).length));
        (_1005_v14)[_index171] = (_1007_v16).dtor_cf29;
        (globalState).f2 = (p0).plus(new BigNumber(520));
        let _index172 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_1001_v12).length));
        let _rhs179 = (((_989_v2).contains(((!(p1)) ? (p0) : (p0)))) ? ((_989_v2).get(((!(p1)) ? (p0) : (p0)))) : ((_this).f21));
        let _rhs180 = _1004_v13.f25;
        let _lhs133 = _this;
        let _lhs134 = _1001_v12;
        let _lhs135 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_1001_v12).length));
        _lhs133.f16 = _rhs179;
        _lhs134[_lhs135] = _rhs180;
      }
      let _1009_v17;
      _1009_v17 = _dafny.Seq.UnicodeFromString("ecyugm");
      let _1010_v18;
      let _nw208 = new _module.C1();
      _nw208.__ctor(_this.f16, p0);
      _1010_v18 = _nw208;
      let _1011_v19;
      _1011_v19 = _module.D14.create_DC28(new BigNumber((_1009_v17).length), (_this).f17, _1010_v18);
      let _source15 = function (_pat_let11_0) {
        return function (_1012_dt__update__tmp_h0) {
          return function (_pat_let12_0) {
            return function (_1013_dt__update_hcf41_h0) {
              return _module.D14.create_DC28((_1012_dt__update__tmp_h0).dtor_cf40, _1013_dt__update_hcf41_h0, (_1012_dt__update__tmp_h0).dtor_cf42);
            }(_pat_let12_0);
          }((_this).f17);
        }(_pat_let11_0);
      }(_1011_v19);
      if (_source15.is_DC28) {
        let _1014___mcc_h1 = (_source15).cf40;
        let _1015___mcc_h2 = (_source15).cf41;
        let _1016___mcc_h3 = (_source15).cf42;
        let _1017_cf42 = _1016___mcc_h3;
        let _1018_cf41 = _1015___mcc_h2;
        let _1019_cf40 = _1014___mcc_h1;
        let _1020_v20;
        _1020_v20 = _dafny.Set.fromElements(_this.f16);
        let _1021_v21;
        _1021_v21 = _module.D15.create_DC29(_1020_v20);
        let _1022_v22;
        _1022_v22 = _dafny.Seq.of(new BigNumber(242));
        let _1023_v23;
        _1023_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1010_v18.f14,_1017_cf42.f14);
        let _pat_let_tv11 = _1017_cf42;
        let _pat_let_tv12 = _1022_v22;
        let _pat_let_tv13 = _1023_v23;
        let _pat_let_tv14 = globalState;
        let _1024_v24;
        let _nw209 = Array((new BigNumber(10)).toNumber());
        _nw209[(_dafny.ZERO).toNumber()] = _1021_v21;
        _nw209[(_dafny.ONE).toNumber()] = _1021_v21;
        _nw209[(new BigNumber(2)).toNumber()] = _1021_v21;
        _nw209[(new BigNumber(3)).toNumber()] = _module.D15.create_DC29(_dafny.Set.fromElements(_1018_cf41));
        _nw209[(new BigNumber(4)).toNumber()] = _1021_v21;
        _nw209[(new BigNumber(5)).toNumber()] = _module.D15.create_DC29(_1020_v20);
        _nw209[(new BigNumber(6)).toNumber()] = _module.D15.create_DC29(_1020_v20);
        _nw209[(new BigNumber(7)).toNumber()] = _module.D15.create_DC29(_1020_v20);
        _nw209[(new BigNumber(8)).toNumber()] = _1021_v21;
        _nw209[(new BigNumber(9)).toNumber()] = function (_pat_let13_0) {
          return function (_1025_dt__update__tmp_h1) {
            return function (_pat_let14_0) {
              return function (_1026_dt__update_hcf43_h0) {
                return _module.D15.create_DC29(_1026_dt__update_hcf43_h0);
              }(_pat_let14_0);
            }(_dafny.Set.fromElements(_module.__default.fm27(_pat_let_tv11.f14, new BigNumber(-518), new BigNumber((_pat_let_tv12).length), _pat_let_tv13, _pat_let_tv14)));
          }(_pat_let13_0);
        }(_1021_v21);
        _1024_v24 = _nw209;
        _1024_v24 = _1024_v24;
        let _1027_v25;
        let _nw210 = Array((new BigNumber(25)).toNumber()).fill([]);
        _1027_v25 = _nw210;
        _1027_v25 = _1027_v25;
        (globalState).f7 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_1010_v18.f14));
        let _1028_v26;
        let _nw211 = new _module.C3();
        _nw211.__ctor(p0, _1009_v17, p1, _1018_cf41);
        _1028_v26 = _nw211;
      } else {
        let _1029___mcc_h4 = (_source15).cf39;
        let _1030_cf39 = _1029___mcc_h4;
        let _1031_v27;
        _1031_v27 = _module.D9.create_DC16(p0, (_this).f17);
        let _1032_v28;
        _1032_v28 = _dafny.MultiSet.fromElements((_1031_v27).dtor_cf24, _this.f16, (_this).f21, (_this).f21, (_this).f21);
        let _1033_v29;
        _1033_v29 = _dafny.Set.fromElements(_1010_v18.f14);
        let _1034_v30;
        _1034_v30 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1032_v28).cardinality()),_1033_v29);
        let _1035_v31;
        _1035_v31 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        _1034_v30 = (_1034_v30).update(_module.__default.fm0(!(_this.f16), globalState), function () {
          let _coll48 = new _dafny.Set();
          for (const _compr_48 of (_1035_v31).Keys.Elements) {
            let _1036_v32 = _compr_48;
            if ((_1035_v31).contains(_1036_v32)) {
              _coll48.add(_module.__default.safeDivisionInt(_1036_v32, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_module.D4.create_DC6(_dafny.Map.Empty.slice().updateUnsafe(true,!(true))), _module.D4.create_DC6(_dafny.Map.Empty.slice().updateUnsafe(false,true)))).length)))).length)));
            }
          }
          return _coll48;
        }());
        (_1010_v18).f14 = _1010_v18.f14;
        let _1037_v33;
        let _init34 = ((_1038_v17) => function (_1039_i3) {
          return (_1039_i3).multipliedBy(new BigNumber((_1038_v17).length));
        })(_1009_v17);
        let _nw212 = Array((new BigNumber(20)).toNumber());
        for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw212.length); _i0_34++) {
          _nw212[_i0_34] = _init34(new BigNumber(_i0_34));
        }
        _1037_v33 = _nw212;
        let _index173 = _module.__default.safeIndex(new BigNumber(988), new BigNumber((_1037_v33).length));
        (_1037_v33)[_index173] = _1010_v18.f14;
        let _1040_v34;
        _1040_v34 = _module.D3.create_DC4((_994_cf0)[_module.__default.safeIndex(_1010_v18.f14, new BigNumber((_994_cf0).length))], new BigNumber(-260), _1010_v18.f14, true, _this.f16);
        let _index174 = _module.__default.safeIndex(new BigNumber(988), new BigNumber((_1037_v33).length));
        (_1037_v33)[_index174] = (_1040_v34).dtor_cf6;
        let _1041_v35;
        _1041_v35 = _dafny.Seq.of(_1040_v34, _1040_v34);
        let _1042_v36;
        _1042_v36 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,_1041_v35);
        _1042_v36 = (_1042_v36).update(_this.f16, _1041_v35);
      }
      let _1043_v37;
      let _nw213 = Array((_dafny.ONE).toNumber());
      _nw213[(_dafny.ZERO).toNumber()] = _1010_v18.f14;
      _1043_v37 = _nw213;
      let _index175 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_1043_v37).length));
      (_1043_v37)[_index175] = p0;
      let _1044_v38;
      _1044_v38 = new _dafny.CodePoint('a'.codePointAt(0));
      let _index176 = _module.__default.safeIndex(new BigNumber(589), new BigNumber((_987_v0).length));
      (_987_v0)[_index176] = _1044_v38;
      let _1045_v39;
      _1045_v39 = _dafny.Seq.of(p0);
      let _1046_v40;
      let _nw214 = Array((new BigNumber(9)).toNumber()).fill(_module.D12.Default());
      _1046_v40 = _nw214;
      let _1047_v41;
      _1047_v41 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-487)), ((_1048_v18) => function (_1049_i4) {
        return _1048_v18.f14;
      })(_1010_v18))).length),_1046_v40);
      let _1050_v42;
      let _nw215 = Array((new BigNumber(11)).toNumber());
      _nw215[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_1045_v39)).cardinality())),_1046_v40);
      _nw215[(_dafny.ONE).toNumber()] = _1047_v41;
      _nw215[(new BigNumber(2)).toNumber()] = _1047_v41;
      _nw215[(new BigNumber(3)).toNumber()] = (_1047_v41).Merge(_1047_v41);
      _nw215[(new BigNumber(4)).toNumber()] = (_1047_v41).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,_1046_v40));
      _nw215[(new BigNumber(5)).toNumber()] = _1047_v41;
      _nw215[(new BigNumber(6)).toNumber()] = (_1047_v41).Merge(_1047_v41);
      _nw215[(new BigNumber(7)).toNumber()] = _1047_v41;
      _nw215[(new BigNumber(8)).toNumber()] = _1047_v41;
      _nw215[(new BigNumber(9)).toNumber()] = _1047_v41;
      _nw215[(new BigNumber(10)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,_1046_v40);
      _1050_v42 = _nw215;
      let _index177 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_1050_v42).length));
      (_1050_v42)[_index177] = _1047_v41;
      let _index178 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_1043_v37).length));
      let _index179 = _module.__default.safeIndex(new BigNumber(589), new BigNumber((_987_v0).length));
      let _index180 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_1050_v42).length));
      let _rhs181 = _1010_v18.f14;
      let _rhs182 = _1044_v38;
      let _rhs183 = _1047_v41;
      let _lhs136 = _1043_v37;
      let _lhs137 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_1043_v37).length));
      let _lhs138 = _987_v0;
      let _lhs139 = _module.__default.safeIndex(new BigNumber(589), new BigNumber((_987_v0).length));
      let _lhs140 = _1050_v42;
      let _lhs141 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_1050_v42).length));
      _lhs136[_lhs137] = _rhs181;
      _lhs138[_lhs139] = _rhs182;
      _lhs140[_lhs141] = _rhs183;
      if (_dafny.Seq.contains(_dafny.Seq.update(_dafny.Seq.of(p0), _module.__default.safeIndex(new BigNumber(21), new BigNumber((_dafny.Seq.of(p0)).length)), new BigNumber((_1009_v17).length)), (_1043_v37)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_1043_v37).length))])) {
        (globalState).f2 = new BigNumber((_dafny.Seq.of(_this.f16, p1, !((_this).f17), _module.__default.fm43((_991_v4)[_module.__default.safeIndex((_1043_v37)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_1043_v37).length))], new BigNumber((_991_v4).length))], globalState))).length);
        let _1051_v43;
        let _nw216 = Array((new BigNumber(21)).toNumber());
        _nw216[(_dafny.ZERO).toNumber()] = (_994_cf0)[_module.__default.safeIndex(_1010_v18.f14, new BigNumber((_994_cf0).length))];
        _nw216[(_dafny.ONE).toNumber()] = (_this).f21;
        _nw216[(new BigNumber(2)).toNumber()] = true;
        _nw216[(new BigNumber(3)).toNumber()] = _this.f16;
        _nw216[(new BigNumber(4)).toNumber()] = (_this).f21;
        _nw216[(new BigNumber(5)).toNumber()] = (_module.D3.create_DC4(p1, new BigNumber(257), _1010_v18.f14, (_this).f17, false)).dtor_cf4;
        _nw216[(new BigNumber(6)).toNumber()] = (_994_cf0)[_module.__default.safeIndex(_1010_v18.f14, new BigNumber((_994_cf0).length))];
        _nw216[(new BigNumber(7)).toNumber()] = p1;
        _nw216[(new BigNumber(8)).toNumber()] = !((_this).f17);
        _nw216[(new BigNumber(9)).toNumber()] = (_this).f17;
        _nw216[(new BigNumber(10)).toNumber()] = (((_989_v2).contains(_1010_v18.f14)) ? ((_989_v2).get(_1010_v18.f14)) : ((_this).f21));
        _nw216[(new BigNumber(11)).toNumber()] = (_this).f21;
        _nw216[(new BigNumber(12)).toNumber()] = (_this).f21;
        _nw216[(new BigNumber(13)).toNumber()] = (_this).f21;
        _nw216[(new BigNumber(14)).toNumber()] = !((_991_v4)[_module.__default.safeIndex(new BigNumber((_1009_v17).length), new BigNumber((_991_v4).length))]);
        _nw216[(new BigNumber(15)).toNumber()] = _this.f16;
        _nw216[(new BigNumber(16)).toNumber()] = !((_this).f17);
        _nw216[(new BigNumber(17)).toNumber()] = false;
        _nw216[(new BigNumber(18)).toNumber()] = true;
        _nw216[(new BigNumber(19)).toNumber()] = (_this).f21;
        _nw216[(new BigNumber(20)).toNumber()] = (_this).f21;
        _1051_v43 = _nw216;
        let _1052_v44;
        _1052_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1009_v17,_1051_v43);
        _1052_v44 = ((_1052_v44).Merge(_1052_v44)).Merge(_1052_v44);
        let _1053_v45;
        _1053_v45 = _module.D4.create_DC7(_992_v5, _1010_v18.f14);
        let _1054_v46;
        _1054_v46 = _dafny.Map.Empty.slice().updateUnsafe(p1,true);
        let _1055_v47;
        _1055_v47 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1054_v46).length),(_1045_v39)[_module.__default.safeIndex(new BigNumber((_989_v2).length), new BigNumber((_1045_v39).length))]);
        let _1056_v48;
        _1056_v48 = _module.D7.create_DC12();
        let _1057_v49;
        _1057_v49 = _dafny.Seq.of(_1056_v48, _1056_v48, _1056_v48);
        let _1058_v50;
        _1058_v50 = _dafny.MultiSet.fromElements(_1010_v18.f14, (_1043_v37)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_1043_v37).length))], (_1043_v37)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_1043_v37).length))]);
        let _1059_v51;
        _1059_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1057_v49,_1058_v50);
        let _1060_v52;
        _1060_v52 = _module.D3.create_DC4(!(!((_this).f17)), _1010_v18.f14, new BigNumber(((((_1059_v51).contains(_dafny.Seq.update(_1057_v49, _module.__default.safeIndex(p0, new BigNumber((_1057_v49).length)), _1056_v48))) ? ((_1059_v51).get(_dafny.Seq.update(_1057_v49, _module.__default.safeIndex(p0, new BigNumber((_1057_v49).length)), _1056_v48))) : (_1058_v50))).cardinality()), _this.f16, _module.__default.fm27((_1043_v37)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_1043_v37).length))], new BigNumber(36), p0, _1055_v47, globalState));
        _1009_v17 = _module.__default.fm28(!(_1010_v18.f14).isEqualTo((_1053_v45).dtor_cf12), _1055_v47, _1060_v52, _module.__default.fm40(!(p1), globalState), globalState);
        let _index181 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_1051_v43).length));
        (_1051_v43)[_index181] = _this.f16;
        let _index182 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_1051_v43).length));
        (_1051_v43)[_index182] = p1;
        let _1061_v53;
        _1061_v53 = _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!((_this).f17),(_1051_v43)[_module.__default.safeIndex(new BigNumber(372), new BigNumber((_1051_v43).length))])).length), new BigNumber(684), new BigNumber((_989_v2).length), _1010_v18.f14, (_1045_v39)[_module.__default.safeIndex(p0, new BigNumber((_1045_v39).length))]);
        let _1062_v54;
        _1062_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1061_v53,false);
        let _1063_v55;
        _1063_v55 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements((_1060_v52).dtor_cf6, (_1043_v37)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_1043_v37).length))], _1010_v18.f14),(_this).f17));
        _1062_v54 = ((_1063_v55)[_module.__default.safeIndex(new BigNumber((_994_cf0).length), new BigNumber((_1063_v55).length))]).update(_module.__default.fm49(new BigNumber(-344), new BigNumber(-264), (_1051_v43)[_module.__default.safeIndex(new BigNumber(372), new BigNumber((_1051_v43).length))], p0, globalState), _this.f16);
      } else {
        (_this).f16 = !((_this).fm12(new BigNumber((_dafny.Seq.UnicodeFromString("hhqnbdg")).length), globalState)).isEqualTo(_1010_v18.f14);
        _991_v4 = _991_v4;
        _1044_v38 = new _dafny.CodePoint('r'.codePointAt(0));
        (globalState).f7 = p0;
        let _1064_v56;
        let _nw217 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Set.Empty);
        _1064_v56 = _nw217;
        let _1065_v57;
        let _nw218 = new _module.C0();
        _nw218.__ctor(_1064_v56, _1010_v18.f14);
        _1065_v57 = _nw218;
      }
      let _1066_v58;
      _1066_v58 = _dafny.Set.fromElements(p0);
      let _1067_v59;
      _1067_v59 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_this.f16, globalState),_989_v2);
      let _1068_v60;
      _1068_v60 = _dafny.Seq.UnicodeFromString("llljm");
      let _1069_v61;
      _1069_v61 = _dafny.MultiSet.fromElements(_this.f16, (_this).f21);
      let _1070_v62;
      _1070_v62 = _module.D9.create_DC16(p0, (_this).f17);
      let _1071_v63;
      _1071_v63 = _dafny.Map.Empty.slice().updateUnsafe((_1070_v62).dtor_cf23,p0);
      let _1072_v64;
      let _nw219 = Array((new BigNumber(20)).toNumber());
      _nw219[(_dafny.ZERO).toNumber()] = !(!(true));
      _nw219[(_dafny.ONE).toNumber()] = !(_1066_v58).equals(_module.__default.fm49(p0, p0, _this.f16, p0, globalState));
      _nw219[(new BigNumber(2)).toNumber()] = (_this).f17;
      _nw219[(new BigNumber(3)).toNumber()] = p1;
      _nw219[(new BigNumber(4)).toNumber()] = _this.f16;
      _nw219[(new BigNumber(5)).toNumber()] = _this.f16;
      _nw219[(new BigNumber(6)).toNumber()] = !((_this).f17);
      _nw219[(new BigNumber(7)).toNumber()] = p1;
      _nw219[(new BigNumber(8)).toNumber()] = p1;
      _nw219[(new BigNumber(9)).toNumber()] = (_this).f17;
      _nw219[(new BigNumber(10)).toNumber()] = _module.__default.fm27(new BigNumber((_1067_v59).length), p0, new BigNumber((_dafny.Seq.update(_1068_v60, _module.__default.safeIndex(new BigNumber((_1069_v61).cardinality()), new BigNumber((_1068_v60).length)), new _dafny.CodePoint('o'.codePointAt(0)))).length), _1071_v63, globalState);
      _nw219[(new BigNumber(11)).toNumber()] = p1;
      _nw219[(new BigNumber(12)).toNumber()] = _module.__default.fm27(p0, (_this).fm12(p0, globalState), p0, _1071_v63, globalState);
      _nw219[(new BigNumber(13)).toNumber()] = (_this).f21;
      _nw219[(new BigNumber(14)).toNumber()] = (p0).isEqualTo(p0);
      _nw219[(new BigNumber(15)).toNumber()] = _this.f16;
      _nw219[(new BigNumber(16)).toNumber()] = _this.f16;
      _nw219[(new BigNumber(17)).toNumber()] = p1;
      _nw219[(new BigNumber(18)).toNumber()] = p1;
      _nw219[(new BigNumber(19)).toNumber()] = (((_this).f21) ? ((_this).f17) : ((_this).f21));
      _1072_v64 = _nw219;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1072_v64).length))) {
        let _1073_i5 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1073_i5)) && ((_1073_i5).isLessThan(new BigNumber((_1072_v64).length))))) {
          (_1072_v64)[(_1073_i5)] = (p0).isEqualTo(p0);
        }
      }
      let _1074_v65;
      _1074_v65 = _dafny.Map.Empty.slice().updateUnsafe(_1071_v63,(p0).multipliedBy((_1070_v62).dtor_cf23));
      _1074_v65 = (_1074_v65).update(_1071_v63, p0);
      let _rhs184 = !(p1);
      let _rhs185 = p1;
      let _lhs142 = _this;
      let _lhs143 = _this;
      _lhs142.f16 = _rhs184;
      _lhs143.f16 = _rhs185;
      if (!(!(false))) {
        let _index183 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_1072_v64).length));
        (_1072_v64)[_index183] = (_this).f17;
        let _index184 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_1072_v64).length));
        (_1072_v64)[_index184] = _this.f16;
        (_this).f16 = _this.f16;
        _1068_v60 = _1068_v60;
        let _1075_v66;
        _1075_v66 = _dafny.Set.fromElements((_this).f21);
        let _1076_v67;
        _1076_v67 = _module.D12.create_DC23((_this).f21, p0, new BigNumber((_1075_v66).length));
        let _source16 = _1076_v67;
        if (_source16.is_DC23) {
          let _1077___mcc_h5 = (_source16).cf30;
          let _1078___mcc_h6 = (_source16).cf31;
          let _1079___mcc_h7 = (_source16).cf32;
          let _1080_cf32 = _1079___mcc_h7;
          let _1081_cf31 = _1078___mcc_h6;
          let _1082_cf30 = _1077___mcc_h5;
          _1082_cf30 = (_this).f17;
          let _index185 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_1072_v64).length));
          (_1072_v64)[_index185] = _this.f16;
          let _1083_v68;
          let _init35 = ((_1084_cf30, _1085_p1) => function (_1086_i6) {
            return (_dafny.Map.Empty.slice().updateUnsafe(_1084_cf30,(_this).f17)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1084_cf30,_1085_p1));
          })(_1082_cf30, p1);
          let _nw220 = Array((new BigNumber(28)).toNumber());
          for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw220.length); _i0_35++) {
            _nw220[_i0_35] = _init35(new BigNumber(_i0_35));
          }
          _1083_v68 = _nw220;
          let _1087_v69;
          _1087_v69 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f17),_this.f16);
          let _index186 = _module.__default.safeIndex(new BigNumber(402), new BigNumber((_1083_v68).length));
          (_1083_v68)[_index186] = (_1087_v69).Merge((_1087_v69).update((_1072_v64)[_module.__default.safeIndex(new BigNumber(722), new BigNumber((_1072_v64).length))], (_this).f21));
          let _1088_v70;
          _1088_v70 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_1072_v64)[_module.__default.safeIndex(new BigNumber(722), new BigNumber((_1072_v64).length))],(_1072_v64)[_module.__default.safeIndex(new BigNumber(722), new BigNumber((_1072_v64).length))]), _1087_v69);
          let _index187 = _module.__default.safeIndex(new BigNumber(402), new BigNumber((_1083_v68).length));
          (_1083_v68)[_index187] = (_1088_v70)[_module.__default.safeIndex(new BigNumber(((_1075_v66).Union(_1075_v66)).length), new BigNumber((_1088_v70).length))];
          let _1089_v71;
          _1089_v71 = new _dafny.CodePoint('e'.codePointAt(0));
          _1089_v71 = _1089_v71;
        } else if (_source16.is_DC24) {
          let _1090___mcc_h8 = (_source16).cf33;
          let _1091___mcc_h9 = (_source16).cf34;
          let _1092___mcc_h10 = (_source16).cf35;
          let _1093___mcc_h11 = (_source16).cf36;
          let _1094___mcc_h12 = (_source16).cf37;
          let _1095_cf37 = _1094___mcc_h12;
          let _1096_cf36 = _1093___mcc_h11;
          let _1097_cf35 = _1092___mcc_h10;
          let _1098_cf34 = _1091___mcc_h9;
          let _1099_cf33 = _1090___mcc_h8;
          let _1100_v72;
          let _init36 = ((_1101_v66) => function (_1102_i7) {
            return _1101_v66;
          })(_1075_v66);
          let _nw221 = Array((new BigNumber(6)).toNumber());
          for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw221.length); _i0_36++) {
            _nw221[_i0_36] = _init36(new BigNumber(_i0_36));
          }
          _1100_v72 = _nw221;
          let _1103_v73;
          let _nw222 = new _module.C0();
          _nw222.__ctor(_1100_v72, p0);
          _1103_v73 = _nw222;
          (globalState).f7 = _1103_v73.f27;
          let _1104_v74;
          _1104_v74 = _dafny.Map.Empty.slice().updateUnsafe(_1099_cf33,p0);
          let _1105_v75;
          _1105_v75 = _dafny.MultiSet.fromElements(_1104_v74, _1104_v74, _1104_v74, _1104_v74, _1104_v74);
          let _1106_v76;
          _1106_v76 = _dafny.Seq.of(_module.__default.fm54(p0, (_this).f21, _1103_v73.f27, !(true), globalState));
          let _1107_v77;
          _1107_v77 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(p1,p0), _1104_v74, _1104_v74);
          _1105_v75 = (((false) ? (_1105_v75) : (_1105_v75))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.Concat((_1106_v76)[_module.__default.safeIndex(_1103_v73.f27, new BigNumber((_1106_v76).length))], _dafny.Seq.update(_1107_v77, _module.__default.safeIndex(p0, new BigNumber((_1107_v77).length)), _1104_v74))));
          let _1108_v78;
          let _nw223 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
          _1108_v78 = _nw223;
          let _index188 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_1108_v78).length));
          (_1108_v78)[_index188] = _1103_v73.f27;
          let _index189 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_1108_v78).length));
          let _rhs186 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(331)), function (_1109_i8) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          });
          let _rhs187 = _1103_v73.f27;
          let _rhs188 = _this.f16;
          let _lhs144 = _1108_v78;
          let _lhs145 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_1108_v78).length));
          _1068_v60 = _rhs186;
          _lhs144[_lhs145] = _rhs187;
          _1096_cf36 = _rhs188;
        } else {
          let _1110___mcc_h13 = (_source16).cf29;
          let _1111_cf29 = _1110___mcc_h13;
          (globalState).f2 = p0;
          _1068_v60 = _1068_v60;
          let _1112_v79;
          _1112_v79 = _dafny.Seq.of(p0, p0);
          let _1113_v80;
          let _nw224 = new _module.C1();
          _nw224.__ctor(_this.f16, p0);
          _1113_v80 = _nw224;
          let _1114_v81;
          _1114_v81 = _module.D3.create_DC4((_module.D14.create_DC28(p0, (_this).f17, _1113_v80)).dtor_cf41, _1113_v80.f14, new BigNumber((_module.__default.fm42(false, new BigNumber((_1066_v58).length), _1113_v80.f14, p1, globalState)).length), (_1072_v64)[_module.__default.safeIndex(new BigNumber(722), new BigNumber((_1072_v64).length))], (_this).f21);
          let _1115_v82;
          _1115_v82 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,_this.f16);
          (globalState).f7 = _module.__default.safeDivisionInt((_1112_v79)[_module.__default.safeIndex(new BigNumber((_module.__default.fm28(p1, _1071_v63, _1114_v81, _1115_v82, globalState)).length), new BigNumber((_1112_v79).length))], (_module.__default.fm0(_this.f16, globalState)).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,_1113_v80.f14)).length)));
          let _1116_v83;
          _1116_v83 = new _dafny.CodePoint('p'.codePointAt(0));
          let _1117_v84;
          let _nw225 = Array((new BigNumber(9)).toNumber());
          _nw225[(_dafny.ZERO).toNumber()] = _1075_v66;
          _nw225[(_dafny.ONE).toNumber()] = _dafny.Set.fromElements((_this).f17);
          _nw225[(new BigNumber(2)).toNumber()] = _dafny.Set.fromElements(p1);
          _nw225[(new BigNumber(3)).toNumber()] = _1075_v66;
          _nw225[(new BigNumber(4)).toNumber()] = _1075_v66;
          _nw225[(new BigNumber(5)).toNumber()] = _1075_v66;
          _nw225[(new BigNumber(6)).toNumber()] = _1075_v66;
          _nw225[(new BigNumber(7)).toNumber()] = _dafny.Set.fromElements((_1072_v64)[_module.__default.safeIndex(new BigNumber(722), new BigNumber((_1072_v64).length))]);
          _nw225[(new BigNumber(8)).toNumber()] = _1075_v66;
          _1117_v84 = _nw225;
          let _1118_v85;
          let _nw226 = new _module.C0();
          _nw226.__ctor(_1117_v84, new BigNumber((_1112_v79).length));
          _1118_v85 = _nw226;
          let _1119_v86;
          _1119_v86 = _dafny.Map.Empty.slice().updateUnsafe(_1118_v85,_1116_v83);
          let _1120_v87;
          _1120_v87 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1116_v83);
          let _index190 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_1072_v64).length));
          let _rhs189 = (((_this).f21) ? ((((_1119_v86).contains(_1118_v85)) ? ((_1119_v86).get(_1118_v85)) : (_1116_v83))) : ((((_1120_v87).contains(new BigNumber(312))) ? ((_1120_v87).get(new BigNumber(312))) : (_1116_v83))));
          let _rhs190 = _1068_v60;
          let _rhs191 = (_this).f17;
          let _rhs192 = !(!((_this).f17)) || ((_this).f21);
          let _lhs146 = _this;
          let _lhs147 = _1072_v64;
          let _lhs148 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_1072_v64).length));
          _1116_v83 = _rhs189;
          _1068_v60 = _rhs190;
          _lhs146.f16 = _rhs191;
          _lhs147[_lhs148] = _rhs192;
        }
        (globalState).f2 = _module.__default.fm0((_1072_v64)[_module.__default.safeIndex(new BigNumber(722), new BigNumber((_1072_v64).length))], globalState);
      } else {
        (_this).f16 = _this.f16;
        let _index191 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_1072_v64).length));
        (_1072_v64)[_index191] = !(p1);
        let _index192 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_1072_v64).length));
        let _rhs193 = !(!(true) || (!(p1)));
        let _rhs194 = p0;
        let _lhs149 = _1072_v64;
        let _lhs150 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_1072_v64).length));
        _lhs149[_lhs150] = _rhs193;
        r0 = _rhs194;
        let _index193 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_1072_v64).length));
        (_1072_v64)[_index193] = _this.f16;
        let _1121_v88;
        _1121_v88 = new _dafny.CodePoint('q'.codePointAt(0));
        let _1122_v89;
        let _nw227 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Set.Empty);
        _1122_v89 = _nw227;
        let _1123_v90;
        let _nw228 = new _module.C0();
        _nw228.__ctor(_1122_v89, p0);
        _1123_v90 = _nw228;
        let _1124_v91;
        let _nw229 = new _module.C1();
        _nw229.__ctor((_1072_v64)[_module.__default.safeIndex(new BigNumber(973), new BigNumber((_1072_v64).length))], new BigNumber((_dafny.Seq.of(_1123_v90, _1123_v90)).length));
        _1124_v91 = _nw229;
        let _1125_v92;
        _1125_v92 = _dafny.Seq.of(_1124_v91, _1124_v91, _1124_v91, _1124_v91);
        let _1126_v93;
        let _nw230 = new _module.C3();
        _nw230.__ctor(_1124_v91.f14, _dafny.Seq.Create(_module.__default.abs(new BigNumber(354)), ((_1127_v88) => function (_1128_i9) {
          return _1127_v88;
        })(_1121_v88)), (_this).f21, p1);
        _1126_v93 = _nw230;
        let _1129_v94;
        _1129_v94 = _dafny.Map.Empty.slice().updateUnsafe(_1126_v93,_1123_v90.f27);
        let _1130_v95;
        _1130_v95 = _dafny.MultiSet.fromElements(_1124_v91.f14, new BigNumber(846));
        let _1131_v96;
        let _nw231 = Array((new BigNumber(19)).toNumber());
        _nw231[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("te")).length);
        _nw231[(_dafny.ONE).toNumber()] = p0;
        _nw231[(new BigNumber(2)).toNumber()] = p0;
        _nw231[(new BigNumber(3)).toNumber()] = new BigNumber((_1071_v63).length);
        _nw231[(new BigNumber(4)).toNumber()] = p0;
        _nw231[(new BigNumber(5)).toNumber()] = p0;
        _nw231[(new BigNumber(6)).toNumber()] = p0;
        _nw231[(new BigNumber(7)).toNumber()] = new BigNumber((_1066_v58).length);
        _nw231[(new BigNumber(8)).toNumber()] = p0;
        _nw231[(new BigNumber(9)).toNumber()] = p0;
        _nw231[(new BigNumber(10)).toNumber()] = new BigNumber((_1125_v92).length);
        _nw231[(new BigNumber(11)).toNumber()] = _1124_v91.f14;
        _nw231[(new BigNumber(12)).toNumber()] = new BigNumber((_991_v4).length);
        _nw231[(new BigNumber(13)).toNumber()] = new BigNumber((_1129_v94).length);
        _nw231[(new BigNumber(14)).toNumber()] = new BigNumber(((_1126_v93).f23).length);
        _nw231[(new BigNumber(15)).toNumber()] = _1124_v91.f14;
        _nw231[(new BigNumber(16)).toNumber()] = new BigNumber((_1130_v95).cardinality());
        _nw231[(new BigNumber(17)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("ld")).length);
        _nw231[(new BigNumber(18)).toNumber()] = _1123_v90.f27;
        _1131_v96 = _nw231;
        let _1132_v97;
        let _init37 = function (_1133_i10) {
          return false;
        };
        let _nw232 = Array((new BigNumber(27)).toNumber());
        for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw232.length); _i0_37++) {
          _nw232[_i0_37] = _init37(new BigNumber(_i0_37));
        }
        _1132_v97 = _nw232;
        let _1134_v98;
        _1134_v98 = _module.D9.create_DC15(p1, _this.f16, _1066_v58, _1131_v96, _1132_v97);
        let _pat_let_tv15 = _1132_v97;
        let _pat_let_tv16 = p1;
        let _1135_v99;
        let _nw233 = Array((new BigNumber(28)).toNumber());
        _nw233[(_dafny.ZERO).toNumber()] = _1134_v98;
        _nw233[(_dafny.ONE).toNumber()] = _module.D9.create_DC15(_this.f16, (_this).f17, _1066_v58, _1131_v96, _1132_v97);
        _nw233[(new BigNumber(2)).toNumber()] = _1134_v98;
        _nw233[(new BigNumber(3)).toNumber()] = _1134_v98;
        _nw233[(new BigNumber(4)).toNumber()] = function (_pat_let15_0) {
          return function (_1138_dt__update__tmp_h3) {
            return function (_pat_let18_0) {
              return function (_1139_dt__update_hcf18_h0) {
                return _module.D9.create_DC15(_1139_dt__update_hcf18_h0, (_1138_dt__update__tmp_h3).dtor_cf19, (_1138_dt__update__tmp_h3).dtor_cf20, (_1138_dt__update__tmp_h3).dtor_cf21, (_1138_dt__update__tmp_h3).dtor_cf22);
              }(_pat_let18_0);
            }(_pat_let_tv16);
          }(_pat_let15_0);
        }(function (_pat_let16_0) {
          return function (_1136_dt__update__tmp_h2) {
            return function (_pat_let17_0) {
              return function (_1137_dt__update_hcf22_h0) {
                return _module.D9.create_DC15((_1136_dt__update__tmp_h2).dtor_cf18, (_1136_dt__update__tmp_h2).dtor_cf19, (_1136_dt__update__tmp_h2).dtor_cf20, (_1136_dt__update__tmp_h2).dtor_cf21, _1137_dt__update_hcf22_h0);
              }(_pat_let17_0);
            }(_pat_let_tv15);
          }(_pat_let16_0);
        }(_1134_v98));
        _nw233[(new BigNumber(5)).toNumber()] = _1134_v98;
        _nw233[(new BigNumber(6)).toNumber()] = _1134_v98;
        _nw233[(new BigNumber(7)).toNumber()] = _1134_v98;
        _nw233[(new BigNumber(8)).toNumber()] = _module.D9.create_DC15((_this).f21, p1, _1066_v58, _1131_v96, _1132_v97);
        _nw233[(new BigNumber(9)).toNumber()] = _1134_v98;
        _nw233[(new BigNumber(10)).toNumber()] = _1134_v98;
        _nw233[(new BigNumber(11)).toNumber()] = _1134_v98;
        _nw233[(new BigNumber(12)).toNumber()] = _1134_v98;
        _nw233[(new BigNumber(13)).toNumber()] = _1134_v98;
        _nw233[(new BigNumber(14)).toNumber()] = _1134_v98;
        _nw233[(new BigNumber(15)).toNumber()] = _1134_v98;
        _nw233[(new BigNumber(16)).toNumber()] = function (_pat_let19_0) {
          return function (_1140_dt__update__tmp_h4) {
            return function (_pat_let20_0) {
              return function (_1141_dt__update_hcf19_h0) {
                return _module.D9.create_DC15((_1140_dt__update__tmp_h4).dtor_cf18, _1141_dt__update_hcf19_h0, (_1140_dt__update__tmp_h4).dtor_cf20, (_1140_dt__update__tmp_h4).dtor_cf21, (_1140_dt__update__tmp_h4).dtor_cf22);
              }(_pat_let20_0);
            }(true);
          }(_pat_let19_0);
        }(_module.D9.create_DC15(p1, (_this).f17, _1066_v58, _1131_v96, _1132_v97));
        _nw233[(new BigNumber(17)).toNumber()] = _1134_v98;
        _nw233[(new BigNumber(18)).toNumber()] = _1134_v98;
        _nw233[(new BigNumber(19)).toNumber()] = _1134_v98;
        _nw233[(new BigNumber(20)).toNumber()] = _1134_v98;
        _nw233[(new BigNumber(21)).toNumber()] = _1134_v98;
        _nw233[(new BigNumber(22)).toNumber()] = _module.D9.create_DC15((_1072_v64)[_module.__default.safeIndex(new BigNumber(973), new BigNumber((_1072_v64).length))], (_1072_v64)[_module.__default.safeIndex(new BigNumber(973), new BigNumber((_1072_v64).length))], _1066_v58, _1131_v96, _1132_v97);
        _nw233[(new BigNumber(23)).toNumber()] = _1134_v98;
        _nw233[(new BigNumber(24)).toNumber()] = _1134_v98;
        _nw233[(new BigNumber(25)).toNumber()] = _module.D9.create_DC15((_this).f17, (_1072_v64)[_module.__default.safeIndex(new BigNumber(973), new BigNumber((_1072_v64).length))], _dafny.Set.fromElements((_1126_v93).f22), _1131_v96, _1132_v97);
        _nw233[(new BigNumber(26)).toNumber()] = _1134_v98;
        _nw233[(new BigNumber(27)).toNumber()] = _module.D9.create_DC15(p1, (_this).f17, _1066_v58, _1131_v96, _1132_v97);
        _1135_v99 = _nw233;
        let _1142_v100;
        _1142_v100 = _dafny.Seq.of(_1135_v99, _1135_v99, _1135_v99);
        let _1143_v101;
        _1143_v101 = _module.D3.create_DC3((_1072_v64)[_module.__default.safeIndex(new BigNumber(973), new BigNumber((_1072_v64).length))]);
        let _1144_v102;
        _1144_v102 = _dafny.Seq.of(_1143_v101);
        let _1145_v103;
        _1145_v103 = _module.D12.create_DC24(_this.f16, _dafny.Seq.update(_991_v4, _module.__default.safeIndex(p0, new BigNumber((_991_v4).length)), p1), _1142_v100, (_1143_v101).dtor_cf3, _1144_v102);
        let _1146_v104;
        _1146_v104 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm43(!((_this).f17), globalState),p1);
        let _rhs195 = _1121_v88;
        let _rhs196 = ((_1145_v103).dtor_cf33) || (_this.f16);
        let _rhs197 = _module.__default.fm0((((_1146_v104).contains((_this).f17)) ? ((_1146_v104).get((_this).f17)) : ((_this).f17)), globalState);
        let _lhs151 = _this;
        let _lhs152 = globalState;
        _1121_v88 = _rhs195;
        _lhs151.f16 = _rhs196;
        _lhs152.f2 = _rhs197;
        let _1147_v105;
        let _nw234 = new _module.C2();
        _nw234.__ctor(_1143_v101, (_991_v4)[_module.__default.safeIndex((_1126_v93).f22, new BigNumber((_991_v4).length))], p1);
        _1147_v105 = _nw234;
      }
      let _1148_v106;
      _1148_v106 = _module.D3.create_DC3((_this).f21);
      let _source17 = _1148_v106;
      if (_source17.is_DC4) {
        let _1149___mcc_h14 = (_source17).cf4;
        let _1150___mcc_h15 = (_source17).cf5;
        let _1151___mcc_h16 = (_source17).cf6;
        let _1152___mcc_h17 = (_source17).cf7;
        let _1153___mcc_h18 = (_source17).cf8;
        let _1154_cf8 = _1153___mcc_h18;
        let _1155_cf7 = _1152___mcc_h17;
        let _1156_cf6 = _1151___mcc_h16;
        let _1157_cf5 = _1150___mcc_h15;
        let _1158_cf4 = _1149___mcc_h14;
        let _1159_v107;
        _1159_v107 = _module.D16.create_DC32(_1069_v61);
        let _1160_v108;
        let _1161_v109;
        let _1162_v110;
        let _out23;
        let _out24;
        let _out25;
        let _outcollector8 = (_this).m10((_this).f17, p0, ((_this.f16) ? (new BigNumber(((_1159_v107).dtor_cf46).cardinality())) : (new BigNumber(92))), _1158_cf4, globalState);
        _out23 = _outcollector8[0];
        _out24 = _outcollector8[1];
        _out25 = _outcollector8[2];
        _1160_v108 = _out23;
        _1161_v109 = _out24;
        _1162_v110 = _out25;
        (_this).f16 = !(!_dafny.Seq.contains(_991_v4, _1160_v108));
        if (_dafny.Seq.IsPrefixOf(_module.__default.fm42(_this.f16, new BigNumber(537), (_this).fm12(new BigNumber((_dafny.Seq.UnicodeFromString("mtf")).length), globalState), true, globalState), _1068_v60)) {
          let _1163_v111;
          let _init38 = ((_1164_v4) => function (_1165_i11) {
            return _1164_v4;
          })(_991_v4);
          let _nw235 = Array((new BigNumber(21)).toNumber());
          for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw235.length); _i0_38++) {
            _nw235[_i0_38] = _init38(new BigNumber(_i0_38));
          }
          _1163_v111 = _nw235;
          _1163_v111 = _1163_v111;
          let _1166_v112;
          _1166_v112 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1068_v60);
          _1157_cf5 = (_module.__default.safeModuloInt(p0, p0)).minus(new BigNumber((_1166_v112).length));
          let _1167_v113;
          _1167_v113 = _dafny.Map.Empty.slice().updateUnsafe(_1154_cf8,p1);
          _1167_v113 = (_1167_v113).update(((_1162_v110) ? (_1154_cf8) : (_this.f16)), p1);
          let _1168_v114;
          _1168_v114 = _dafny.MultiSet.fromElements(_1156_cf6, _1156_cf6, new BigNumber(((_1069_v61).Union(_1069_v61)).cardinality()));
          let _1169_v115;
          _1169_v115 = _dafny.Map.Empty.slice().updateUnsafe(_1168_v114,p0);
          let _rhs198 = _1155_cf7;
          let _rhs199 = _1168_v114;
          let _rhs200 = (_dafny.ZERO).minus(new BigNumber(((_module.__default.fm55(true, p0, globalState)).dtor_cf46).cardinality()));
          let _rhs201 = new BigNumber(340);
          let _rhs202 = ((_1168_v114).Difference(_dafny.MultiSet.fromElements(new BigNumber((_1068_v60).length), new BigNumber((_1169_v115).length)))).Intersect(((_1158_cf4) ? (_1168_v114) : (_dafny.MultiSet.fromElements(_1157_cf5))));
          let _lhs153 = globalState;
          _1162_v110 = _rhs198;
          _1168_v114 = _rhs199;
          _lhs153.f2 = _rhs200;
          _1157_cf5 = _rhs201;
          _1168_v114 = _rhs202;
          let _1170_v116;
          let _1171_v117;
          let _1172_v118;
          let _out26;
          let _out27;
          let _out28;
          let _outcollector9 = (_this).m10(_1158_cf4, _1157_cf5, _1157_cf5, _1158_cf4, globalState);
          _out26 = _outcollector9[0];
          _out27 = _outcollector9[1];
          _out28 = _outcollector9[2];
          _1170_v116 = _out26;
          _1171_v117 = _out27;
          _1172_v118 = _out28;
        } else {
          let _index194 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_1072_v64).length));
          (_1072_v64)[_index194] = _this.f16;
          let _index195 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_1072_v64).length));
          (_1072_v64)[_index195] = (((_989_v2).contains(_module.__default.fm0(_1155_cf7, globalState))) ? ((_989_v2).get(_module.__default.fm0(_1155_cf7, globalState))) : (!(_module.__default.fm31(_1155_cf7, _1158_cf4, globalState))));
          _1160_v108 = p1;
          let _1173_v119;
          _1173_v119 = _dafny.Set.fromElements(_this.f16);
          let _1174_v120;
          _1174_v120 = _dafny.Map.Empty.slice().updateUnsafe(_1158_cf4,p1);
          (globalState).f7 = (((_1069_v61).contains((new BigNumber((_1071_v63).length)).isLessThanOrEqualTo(new BigNumber((_1173_v119).length)))) ? ((_1069_v61).get((new BigNumber((_1071_v63).length)).isLessThanOrEqualTo(new BigNumber((_1173_v119).length)))) : (new BigNumber((_1174_v120).length)));
          let _1175_v121;
          _1175_v121 = new _dafny.CodePoint('s'.codePointAt(0));
          let _index196 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_987_v0).length));
          (_987_v0)[_index196] = _1175_v121;
          let _index197 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_987_v0).length));
          (_987_v0)[_index197] = _1175_v121;
          let _1176_v122;
          let _nw236 = new _module.C1();
          _nw236.__ctor(!(p1) || (_1162_v110), _1156_cf6);
          _1176_v122 = _nw236;
        }
        let _1177_v123;
        let _nw237 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
        _1177_v123 = _nw237;
        let _index198 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_1177_v123).length));
        (_1177_v123)[_index198] = _module.__default.safeModuloInt((_dafny.ZERO).minus(p0), new BigNumber((_1068_v60).length));
        let _index199 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_1177_v123).length));
        (_1177_v123)[_index199] = p0;
      } else if (_source17.is_DC3) {
        let _1178___mcc_h19 = (_source17).cf3;
        let _1179_cf3 = _1178___mcc_h19;
        if ((_this).f17) {
          let _1180_v124;
          _1180_v124 = _dafny.MultiSet.fromElements(p0);
          let _1181_v125;
          _1181_v125 = _dafny.Map.Empty.slice().updateUnsafe(true,p0);
          let _1182_v126;
          _1182_v126 = _dafny.MultiSet.fromElements((((_1180_v124).contains(p0)) ? ((_1180_v124).get(p0)) : (p0)), new BigNumber((_1181_v125).length));
          let _1183_v127;
          let _nw238 = new _module.C3();
          _nw238.__ctor(_module.__default.safeDivisionInt(p0, p0), _1068_v60, p1, (_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_1069_v61).cardinality())))).IsSubsetOf(_1182_v126));
          _1183_v127 = _nw238;
          _991_v4 = _991_v4;
          let _1184_v128;
          let _nw239 = Array((new BigNumber(4)).toNumber());
          _nw239[(_dafny.ZERO).toNumber()] = (_1183_v127).f22;
          _nw239[(_dafny.ONE).toNumber()] = (_1183_v127).f22;
          _nw239[(new BigNumber(2)).toNumber()] = new BigNumber(((_1183_v127).f23).length);
          _nw239[(new BigNumber(3)).toNumber()] = p0;
          _1184_v128 = _nw239;
          let _1185_v129;
          _1185_v129 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_1184_v128);
          let _1186_v130;
          _1186_v130 = _dafny.Seq.of(p0);
          let _1187_v131;
          let _nw240 = Array((new BigNumber(15)).toNumber());
          _nw240[(_dafny.ZERO).toNumber()] = ((p1) ? (new BigNumber(-291)) : (new BigNumber((_1185_v129).length)));
          _nw240[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("jsmajdos")).length);
          _nw240[(new BigNumber(2)).toNumber()] = (_1183_v127).f22;
          _nw240[(new BigNumber(3)).toNumber()] = ((_dafny.ZERO).minus(p0)).minus(new BigNumber(((_this).fm13((_1183_v127).f22, (_this).f21, false, globalState)).length));
          _nw240[(new BigNumber(4)).toNumber()] = (_1183_v127).f22;
          _nw240[(new BigNumber(5)).toNumber()] = (_1183_v127).f22;
          _nw240[(new BigNumber(6)).toNumber()] = (_1183_v127).f22;
          _nw240[(new BigNumber(7)).toNumber()] = (new BigNumber(865)).plus(new BigNumber((_dafny.Seq.UnicodeFromString("mksbkds")).length));
          _nw240[(new BigNumber(8)).toNumber()] = _module.__default.safeModuloInt((_1183_v127).f22, (_1183_v127).f22);
          _nw240[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("sjpp")).length);
          _nw240[(new BigNumber(10)).toNumber()] = (_1186_v130)[_module.__default.safeIndex(p0, new BigNumber((_1186_v130).length))];
          _nw240[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(p0));
          _nw240[(new BigNumber(12)).toNumber()] = (_1183_v127).f22;
          _nw240[(new BigNumber(13)).toNumber()] = new BigNumber(((_1183_v127).f23).length);
          _nw240[(new BigNumber(14)).toNumber()] = p0;
          _1187_v131 = _nw240;
          let _index200 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_1184_v128).length));
          (_1184_v128)[_index200] = p0;
          let _index201 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_1184_v128).length));
          (_1184_v128)[_index201] = p0;
          let _1188_v132;
          let _nw241 = Array((_dafny.ONE).toNumber());
          _nw241[(_dafny.ZERO).toNumber()] = _1072_v64;
          _1188_v132 = _nw241;
          let _1189_v133;
          let _init39 = ((_1190_v60) => function (_1191_i12) {
            return _1190_v60;
          })(_1068_v60);
          let _nw242 = Array((new BigNumber(26)).toNumber());
          for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw242.length); _i0_39++) {
            _nw242[_i0_39] = _init39(new BigNumber(_i0_39));
          }
          _1189_v133 = _nw242;
          let _index202 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_1189_v133).length));
          (_1189_v133)[_index202] = _dafny.Seq.Concat(_1068_v60, _dafny.Seq.UnicodeFromString("catdtpru"));
          let _index203 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_1184_v128).length));
          let _index204 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_1184_v128).length));
          let _index205 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_1189_v133).length));
          let _rhs203 = p0;
          let _rhs204 = _1188_v132;
          let _rhs205 = (_1183_v127).f22;
          let _rhs206 = _dafny.Seq.Concat(_1068_v60, (_1183_v127).f23);
          let _lhs154 = _1184_v128;
          let _lhs155 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_1184_v128).length));
          let _lhs156 = _1184_v128;
          let _lhs157 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_1184_v128).length));
          let _lhs158 = _1189_v133;
          let _lhs159 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_1189_v133).length));
          _lhs154[_lhs155] = _rhs203;
          _1188_v132 = _rhs204;
          _lhs156[_lhs157] = _rhs205;
          _lhs158[_lhs159] = _rhs206;
          let _index206 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_1072_v64).length));
          (_1072_v64)[_index206] = _1179_cf3;
          let _index207 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_1184_v128).length));
          let _index208 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_1072_v64).length));
          let _index209 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_1184_v128).length));
          let _rhs207 = _this.f16;
          let _rhs208 = (_1183_v127).f22;
          let _rhs209 = _module.__default.fm27((_1184_v128)[_module.__default.safeIndex(new BigNumber(423), new BigNumber((_1184_v128).length))], (p0).plus((_dafny.ZERO).minus((_dafny.ZERO).minus(p0))), (_dafny.ZERO).minus(_module.__default.fm0((_this).f17, globalState)), _1071_v63, globalState);
          let _rhs210 = p0;
          let _rhs211 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(883)), function (_1192_i13) {
            return new _dafny.CodePoint('c'.codePointAt(0));
          });
          let _lhs160 = _1184_v128;
          let _lhs161 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_1184_v128).length));
          let _lhs162 = _1072_v64;
          let _lhs163 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_1072_v64).length));
          let _lhs164 = _1184_v128;
          let _lhs165 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_1184_v128).length));
          _1179_cf3 = _rhs207;
          _lhs160[_lhs161] = _rhs208;
          _lhs162[_lhs163] = _rhs209;
          _lhs164[_lhs165] = _rhs210;
          _1068_v60 = _rhs211;
        } else {
          let _1193_v134;
          let _init40 = ((_1194_p0) => function (_1195_i14) {
            return (_1195_i14).minus(_1194_p0);
          })(p0);
          let _nw243 = Array((new BigNumber(21)).toNumber());
          for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw243.length); _i0_40++) {
            _nw243[_i0_40] = _init40(new BigNumber(_i0_40));
          }
          _1193_v134 = _nw243;
          let _1196_v135;
          _1196_v135 = _dafny.Set.fromElements(_1072_v64);
          let _1197_v136;
          let _nw244 = new _module.C1();
          _nw244.__ctor((_this).f17, new BigNumber((_1196_v135).length));
          _1197_v136 = _nw244;
          let _1198_v137;
          _1198_v137 = _module.D14.create_DC28(p0, (_this).f21, _1197_v136);
          let _1199_v138;
          _1199_v138 = _dafny.Map.Empty.slice().updateUnsafe(_1193_v134,_1198_v137);
          _1199_v138 = (_1199_v138).update(_1193_v134, _1198_v137);
          let _rhs212 = (_1069_v61).update(_1179_cf3, _module.__default.abs(_module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_991_v4)[_module.__default.safeIndex(_1197_v136.f14, new BigNumber((_991_v4).length))],(_this).f17)).length), _1197_v136.f14)));
          let _rhs213 = !(!(_1179_cf3));
          let _rhs214 = _1197_v136.f14;
          let _lhs166 = _this;
          let _lhs167 = globalState;
          _1069_v61 = _rhs212;
          _lhs166.f16 = _rhs213;
          _lhs167.f2 = _rhs214;
          (_this).f16 = false;
          (globalState).f2 = (_module.__default.safeDivisionInt(_1197_v136.f14, p0)).multipliedBy((_dafny.ZERO).minus(_module.__default.fm0(_this.f16, globalState)));
          _1072_v64 = _1072_v64;
        }
        let _1200_v139;
        _1200_v139 = _dafny.Map.Empty.slice().updateUnsafe(_1179_cf3,_module.__default.safeDivisionInt(p0, new BigNumber((_dafny.Seq.UnicodeFromString("c")).length)));
        _1200_v139 = (_1200_v139).update(true, new BigNumber((_dafny.Seq.UnicodeFromString("n")).length));
        let _1201_v140;
        let _nw245 = new _module.C1();
        _nw245.__ctor(p1, p0);
        _1201_v140 = _nw245;
        let _1202_v141;
        _1202_v141 = _module.D16.create_DC33(_1179_cf3, (_this).f17, new BigNumber((_1071_v63).length), (_this).f17);
        let _1203_v142;
        _1203_v142 = _module.D16.create_DC34((_this).f17, _1201_v140, (_1202_v141).dtor_cf49);
        let _1204_v143;
        _1204_v143 = _dafny.MultiSet.fromElements(_1203_v142);
        let _1205_v144;
        _1205_v144 = _dafny.Map.Empty.slice().updateUnsafe(!(p1),_1204_v143);
        let _1206_v145;
        _1206_v145 = _dafny.MultiSet.fromElements(p0, p0, new BigNumber(913));
        _1205_v144 = (_1205_v144).update((_1206_v145).IsSubsetOf(_1206_v145), _1204_v143);
        let _1207_v146;
        _1207_v146 = _dafny.Map.Empty.slice().updateUnsafe(_1179_cf3,true);
        _1207_v146 = (_1207_v146).update((_this).f21, (_this).f21);
      } else {
        let _1208___mcc_h20 = (_source17).cf9;
        let _1209_cf9 = _1208___mcc_h20;
        let _1210_v147;
        let _init41 = ((_1211_p0) => function (_1212_i15) {
          return (_1212_i15).plus(_1211_p0);
        })(p0);
        let _nw246 = Array((new BigNumber(25)).toNumber());
        for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw246.length); _i0_41++) {
          _nw246[_i0_41] = _init41(new BigNumber(_i0_41));
        }
        _1210_v147 = _nw246;
        let _1213_v148;
        _1213_v148 = _dafny.Seq.of(_1210_v147, _1210_v147, _1210_v147);
        (_this).f16 = _dafny.Seq.IsPrefixOf(_1213_v148, _1213_v148);
        let _1214_v149;
        _1214_v149 = new _dafny.CodePoint('j'.codePointAt(0));
        _1214_v149 = _1214_v149;
        (globalState).f7 = p0;
        let _1215_v150;
        _1215_v150 = _dafny.Seq.of(p0);
        let _1216_v151;
        _1216_v151 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,new BigNumber((_dafny.Set.fromElements(_module.__default.fm0((_this).f21, globalState), p0, p0)).length));
        let _1217_v152;
        _1217_v152 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_1215_v150, _dafny.Seq.of(new BigNumber((_1216_v151).length), new BigNumber((_1071_v63).length)))).length),new _dafny.CodePoint('x'.codePointAt(0)));
        (_this).f16 = !((p0).plus(new BigNumber((_1217_v152).length))).isEqualTo(p0);
      }
      r0 = new BigNumber(536);
      return r0;
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _1218_v0;
      _1218_v0 = _dafny.Set.fromElements(p3, p1, p3, p3);
      let _1219_v1;
      _1219_v1 = _dafny.Seq.UnicodeFromString("vgfxhsafg");
      let _1220_v2;
      _1220_v2 = _dafny.Seq.of(new BigNumber((_1219_v1).length), p3);
      let _1221_v3;
      _1221_v3 = _module.D3.create_DC4((_this).f21, new BigNumber((_1219_v1).length), p1, (_this).f17, !(_module.__default.fm43((_this).f17, globalState)));
      let _1222_v4;
      let _nw247 = Array((new BigNumber(13)).toNumber());
      _nw247[(_dafny.ZERO).toNumber()] = p1;
      _nw247[(_dafny.ONE).toNumber()] = p3;
      _nw247[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(p3);
      _nw247[(new BigNumber(3)).toNumber()] = new BigNumber((_1218_v0).length);
      _nw247[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus((_this).fm12(p1, globalState));
      _nw247[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(_1220_v2)).cardinality());
      _nw247[(new BigNumber(6)).toNumber()] = new BigNumber((_1219_v1).length);
      _nw247[(new BigNumber(7)).toNumber()] = p1;
      _nw247[(new BigNumber(8)).toNumber()] = p3;
      _nw247[(new BigNumber(9)).toNumber()] = p1;
      _nw247[(new BigNumber(10)).toNumber()] = p3;
      _nw247[(new BigNumber(11)).toNumber()] = (_1221_v3).dtor_cf5;
      _nw247[(new BigNumber(12)).toNumber()] = new BigNumber(84);
      _1222_v4 = _nw247;
      let _1223_v5;
      _1223_v5 = _dafny.Map.Empty.slice().updateUnsafe(_1222_v4,(p3).isLessThan(new BigNumber((_module.__default.fm49(p3, p3, (_this).f21, p1, globalState)).length)));
      _1223_v5 = (_1223_v5).Merge(_1223_v5);
      let _1224_v6;
      _1224_v6 = _dafny.Set.fromElements(p2, p2, (_this).f17, p2, (_this).f17);
      let _1225_v7;
      _1225_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,new BigNumber(400));
      let _1226_v8;
      _1226_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1224_v6,(_1225_v7).Merge(_1225_v7));
      _1226_v8 = (_1226_v8).update(_1224_v6, _1225_v7);
      let _1227_v9;
      _1227_v9 = _dafny.Map.Empty.slice().updateUnsafe(p3,p1);
      (_this).f16 = ((new BigNumber((_1219_v1).length)).plus(new BigNumber((_module.__default.fm38((_this).f21, new BigNumber((_1227_v9).length), p3, globalState)).length))).isEqualTo(new BigNumber(352));
      let _1228_v10;
      _1228_v10 = _module.D3.create_DC3(true);
      if ((_1228_v10).dtor_cf3) {
        (_this).f16 = (_this).f21;
        let _1229_v11;
        _1229_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1219_v1,_this.f16);
        _1229_v11 = (_1229_v11).update(_1219_v1, false);
        (_this).f16 = ((_this.f16) ? ((_this).f17) : (!(false)));
        let _1230_v12;
        _1230_v12 = _dafny.Seq.of(_this.f16);
        (_this).f16 = (_1230_v12)[_module.__default.safeIndex(p3, new BigNumber((_1230_v12).length))];
        let _1231_v13;
        _1231_v13 = _dafny.MultiSet.fromElements(p1);
        let _1232_v14;
        _1232_v14 = _dafny.Map.Empty.slice().updateUnsafe(p3,_1231_v13);
        _1232_v14 = (_1232_v14).update(p3, _1231_v13);
      } else {
        (_this).f16 = (_this).f17;
        let _1233_v15;
        _1233_v15 = _dafny.Map.Empty.slice().updateUnsafe(true,p0);
        _1233_v15 = (_1233_v15).update(_this.f16, p0);
        let _1234_v16;
        _1234_v16 = _dafny.Seq.of(_module.__default.fm27(p3, p1, p1, (_1227_v9).update(p1, p1), globalState));
        r0 = (p3).minus(new BigNumber((_dafny.Seq.Concat(_1234_v16, _1234_v16)).length));
        let _1235_v17;
        let _nw248 = new _module.C1();
        _nw248.__ctor((_this).f21, p1);
        _1235_v17 = _nw248;
        let _1236_v18;
        _1236_v18 = _module.D14.create_DC28(p1, (_this).f17, _1235_v17);
        let _index210 = _module.__default.safeIndex(new BigNumber(322), new BigNumber((_1222_v4).length));
        (_1222_v4)[_index210] = (_dafny.ZERO).minus((_1236_v18).dtor_cf40);
        let _index211 = _module.__default.safeIndex(new BigNumber(322), new BigNumber((_1222_v4).length));
        (_1222_v4)[_index211] = (_dafny.ZERO).minus(new BigNumber(-460));
        let _1237_v19;
        _1237_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1225_v7,p2);
        let _1238_v20;
        _1238_v20 = _dafny.Seq.of(_1219_v1, _1219_v1);
        let _1239_v21;
        let _nw249 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
        _1239_v21 = _nw249;
        let _1240_v22;
        _1240_v22 = _dafny.Seq.of(_1239_v21, _1239_v21);
        let _1241_v23;
        let _nw250 = Array((new BigNumber(18)).toNumber());
        _nw250[(_dafny.ZERO).toNumber()] = (_1234_v16)[_module.__default.safeIndex(new BigNumber((_1234_v16).length), new BigNumber((_1234_v16).length))];
        _nw250[(_dafny.ONE).toNumber()] = _this.f16;
        _nw250[(new BigNumber(2)).toNumber()] = !((_this).f21);
        _nw250[(new BigNumber(3)).toNumber()] = _this.f16;
        _nw250[(new BigNumber(4)).toNumber()] = (_this).f21;
        _nw250[(new BigNumber(5)).toNumber()] = (_this).f21;
        _nw250[(new BigNumber(6)).toNumber()] = p2;
        _nw250[(new BigNumber(7)).toNumber()] = p2;
        _nw250[(new BigNumber(8)).toNumber()] = (_this).f17;
        _nw250[(new BigNumber(9)).toNumber()] = ((true) ? ((_this).f21) : ((((_1237_v19).contains(_1225_v7)) ? ((_1237_v19).get(_1225_v7)) : (false))));
        _nw250[(new BigNumber(10)).toNumber()] = _this.f16;
        _nw250[(new BigNumber(11)).toNumber()] = _this.f16;
        _nw250[(new BigNumber(12)).toNumber()] = p2;
        _nw250[(new BigNumber(13)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.update((_1238_v20)[_module.__default.safeIndex(_1235_v17.f14, new BigNumber((_1238_v20).length))], _module.__default.safeIndex(new BigNumber((_1234_v16).length), new BigNumber(((_1238_v20)[_module.__default.safeIndex(_1235_v17.f14, new BigNumber((_1238_v20).length))]).length)), p0), _1219_v1);
        _nw250[(new BigNumber(14)).toNumber()] = _dafny.areEqual(_1240_v22, _1240_v22);
        _nw250[(new BigNumber(15)).toNumber()] = _this.f16;
        _nw250[(new BigNumber(16)).toNumber()] = ((_this).f17) === ((_1234_v16)[_module.__default.safeIndex(p3, new BigNumber((_1234_v16).length))]);
        _nw250[(new BigNumber(17)).toNumber()] = (_this).f17;
        _1241_v23 = _nw250;
        let _nw251 = Array((new BigNumber(16)).toNumber()).fill(false);
        _1241_v23 = _nw251;
      }
      let _1242_v24;
      let _nw252 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Set.Empty);
      _1242_v24 = _nw252;
      let _1243_v25;
      _1243_v25 = _dafny.Seq.of(false, _this.f16);
      let _1244_v26;
      let _nw253 = new _module.C0();
      _nw253.__ctor(_1242_v24, new BigNumber((_1243_v25).length));
      _1244_v26 = _nw253;
      if (!(!((_this).f17))) {
        (_this).f16 = _this.f16;
        let _1245_v27;
        _1245_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1222_v4,(_dafny.ZERO).minus(p1));
        let _1246_v28;
        _1246_v28 = _dafny.MultiSet.fromElements(_1219_v1);
        let _1247_v29;
        _1247_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1246_v28,_1245_v27);
        let _1248_v30;
        _1248_v30 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1245_v27);
        let _1249_v31;
        let _nw254 = Array((new BigNumber(8)).toNumber());
        _nw254[(_dafny.ZERO).toNumber()] = _1245_v27;
        _nw254[(_dafny.ONE).toNumber()] = _1245_v27;
        _nw254[(new BigNumber(2)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1222_v4,p3)).Merge(_1245_v27);
        _nw254[(new BigNumber(3)).toNumber()] = _1245_v27;
        _nw254[(new BigNumber(4)).toNumber()] = (_1245_v27).Merge(_1245_v27);
        _nw254[(new BigNumber(5)).toNumber()] = (((_1247_v29).contains(_dafny.MultiSet.fromElements(_1219_v1))) ? ((_1247_v29).get(_dafny.MultiSet.fromElements(_1219_v1))) : ((((_1248_v30).contains(p0)) ? ((_1248_v30).get(p0)) : (_1245_v27))));
        _nw254[(new BigNumber(6)).toNumber()] = _1245_v27;
        _nw254[(new BigNumber(7)).toNumber()] = _1245_v27;
        _1249_v31 = _nw254;
        let _index212 = _module.__default.safeIndex(new BigNumber(502), new BigNumber((_1249_v31).length));
        (_1249_v31)[_index212] = _1245_v27;
        let _index213 = _module.__default.safeIndex(new BigNumber(502), new BigNumber((_1249_v31).length));
        (_1249_v31)[_index213] = _1245_v27;
        let _1250_v32;
        let _nw255 = Array((new BigNumber(21)).toNumber()).fill(false);
        _1250_v32 = _nw255;
        _1250_v32 = (((_1218_v0).IsDisjointFrom(_module.__default.fm49(new BigNumber((_1218_v0).length), p1, !(!((_this).f17)), _1244_v26.f27, globalState))) ? (_1250_v32) : (_1250_v32));
        (globalState).f7 = new BigNumber(387);
        _1250_v32 = _1250_v32;
      } else {
        (_this).f16 = (_1244_v26.f27).isLessThanOrEqualTo(p1);
        let _1251_v33;
        _1251_v33 = _dafny.Seq.of(_1243_v25);
        let _1252_v34;
        let _nw256 = new _module.C3();
        _nw256.__ctor(new BigNumber((_dafny.MultiSet.FromArray((_1251_v33)[_module.__default.safeIndex(p3, new BigNumber((_1251_v33).length))])).cardinality()), _1219_v1, (_this).f21, (_this).f21);
        _1252_v34 = _nw256;
        let _1253_v35;
        _1253_v35 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(430)), ((_1254_p0) => function (_1255_i0) {
          return _1254_p0;
        })(p0)),_1252_v34);
        let _1256_v36;
        let _nw257 = Array((new BigNumber(21)).toNumber());
        _nw257[(_dafny.ZERO).toNumber()] = _1253_v35;
        _nw257[(_dafny.ONE).toNumber()] = _1253_v35;
        _nw257[(new BigNumber(2)).toNumber()] = (_1253_v35).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("m"),_1252_v34));
        _nw257[(new BigNumber(3)).toNumber()] = _1253_v35;
        _nw257[(new BigNumber(4)).toNumber()] = _1253_v35;
        _nw257[(new BigNumber(5)).toNumber()] = (_1253_v35).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1219_v1,_1252_v34));
        _nw257[(new BigNumber(6)).toNumber()] = (_1253_v35).update(_1219_v1, _1252_v34);
        _nw257[(new BigNumber(7)).toNumber()] = (_1253_v35).Merge(_1253_v35);
        _nw257[(new BigNumber(8)).toNumber()] = (_1253_v35).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1219_v1,_1252_v34));
        _nw257[(new BigNumber(9)).toNumber()] = _1253_v35;
        _nw257[(new BigNumber(10)).toNumber()] = _1253_v35;
        _nw257[(new BigNumber(11)).toNumber()] = _1253_v35;
        _nw257[(new BigNumber(12)).toNumber()] = _1253_v35;
        _nw257[(new BigNumber(13)).toNumber()] = (_1253_v35).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1219_v1,_1252_v34));
        _nw257[(new BigNumber(14)).toNumber()] = _1253_v35;
        _nw257[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1219_v1,_1252_v34);
        _nw257[(new BigNumber(16)).toNumber()] = _1253_v35;
        _nw257[(new BigNumber(17)).toNumber()] = _1253_v35;
        _nw257[(new BigNumber(18)).toNumber()] = ((p2) ? (_1253_v35) : (_1253_v35));
        _nw257[(new BigNumber(19)).toNumber()] = _1253_v35;
        _nw257[(new BigNumber(20)).toNumber()] = _1253_v35;
        _1256_v36 = _nw257;
        let _index214 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_1256_v36).length));
        (_1256_v36)[_index214] = _1253_v35;
        let _index215 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_1256_v36).length));
        (_1256_v36)[_index215] = _1253_v35;
        let _1257_v37;
        _1257_v37 = _module.D15.create_DC30(p0);
        let _1258_v38;
        _1258_v38 = _module.D15.create_DC31(_1257_v37);
        _1258_v38 = _1258_v38;
        (globalState).f2 = _1244_v26.f27;
        let _1259_v39;
        _1259_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_1219_v1);
        _1219_v1 = (((_1259_v39).contains(p2)) ? ((_1259_v39).get(p2)) : (_1219_v1));
      }
      let _1260_v40;
      _1260_v40 = _dafny.MultiSet.fromElements(((_this).f17) === ((_this).f17), !((_this).f21));
      let _1261_v41;
      _1261_v41 = _module.D15.create_DC29(_1224_v6);
      r0 = (((_1260_v40).contains((new BigNumber((_1224_v6).length)).isLessThan(new BigNumber(864)))) ? ((_1260_v40).get((new BigNumber((_1224_v6).length)).isLessThan(new BigNumber(864)))) : (((_module.__default.fm43((_this).f21, globalState)) ? (new BigNumber(((_1261_v41).dtor_cf43).length)) : (p3))));
      return r0;
    }
    m10(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _module.D6.Default();
      let r2 = false;
      let _1262_v0;
      _1262_v0 = _dafny.Seq.UnicodeFromString("fbxik");
      let _1263_v1;
      _1263_v1 = _dafny.Seq.of(p2, new BigNumber((_1262_v0).length), p2, p1, p1);
      (globalState).f7 = ((_1263_v1)[_module.__default.safeIndex(p2, new BigNumber((_1263_v1).length))]).multipliedBy(p1);
      let _1264_i0;
      _1264_i0 = _dafny.ZERO;
      L5: {
        while ((_module.__default.fm0(p3, globalState)).isLessThanOrEqualTo(p1)) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1264_i0)) {
              break L5;
            }
            _1264_i0 = (_1264_i0).plus(_dafny.ONE);
            let _1265_v2;
            _1265_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,p0);
            let _1266_v3;
            _1266_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_1265_v2);
            let _1267_v4;
            _1267_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1262_v0,(_dafny.ZERO).minus(new BigNumber((_1266_v3).length)));
            let _1268_v5;
            _1268_v5 = _module.D13.create_DC25(_1267_v4);
            let _pat_let_tv17 = _1267_v4;
            let _source18 = (((p2).isLessThanOrEqualTo(p2)) ? (_1268_v5) : (function (_pat_let21_0) {
              return function (_1269_dt__update__tmp_h0) {
                return function (_pat_let22_0) {
                  return function (_1270_dt__update_hcf38_h0) {
                    return _module.D13.create_DC25(_1270_dt__update_hcf38_h0);
                  }(_pat_let22_0);
                }(_pat_let_tv17);
              }(_pat_let21_0);
            }(_1268_v5)));
            if (_source18.is_DC26) {
              let _1271_v6;
              _1271_v6 = _module.D3.create_DC3((_this).f21);
              let _1272_v7;
              let _nw258 = new _module.C2();
              _nw258.__ctor(_1271_v6, p0, p0);
              _1272_v7 = _nw258;
              _1263_v1 = _1263_v1;
              (_this).f16 = (!(false)) === ((_this).f17);
              let _1273_v8;
              let _nw259 = Array((new BigNumber(9)).toNumber()).fill(false);
              _1273_v8 = _nw259;
              let _1274_v9;
              _1274_v9 = _dafny.Set.fromElements(new BigNumber(19));
              let _1275_v10;
              _1275_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1273_v8,(_1274_v9).Intersect(_dafny.Set.fromElements((_module.D9.create_DC16(new BigNumber((_1265_v2).length), !((_this).f21))).dtor_cf23)));
              let _rhs215 = _dafny.Map.Empty.slice().updateUnsafe(_1273_v8,_1274_v9);
              let _rhs216 = (_this).f21;
              let _rhs217 = _dafny.Seq.UnicodeFromString("jlvv");
              let _rhs218 = p1;
              let _rhs219 = true;
              let _lhs168 = globalState;
              _1275_v10 = _rhs215;
              r2 = _rhs216;
              _1262_v0 = _rhs217;
              _lhs168.f7 = _rhs218;
              r0 = _rhs219;
            } else {
              let _1276___mcc_h0 = (_source18).cf38;
              let _1277_cf38 = _1276___mcc_h0;
              (globalState).f7 = (p1).multipliedBy((((_this).f17) ? (p1) : (p1)));
              let _1278_v11;
              _1278_v11 = new _dafny.CodePoint('q'.codePointAt(0));
              let _1279_v12;
              let _nw260 = Array((new BigNumber(22)).toNumber());
              _nw260[(_dafny.ZERO).toNumber()] = _1278_v11;
              _nw260[(_dafny.ONE).toNumber()] = (_1262_v0)[_module.__default.safeIndex((_dafny.ZERO).minus(p2), new BigNumber((_1262_v0).length))];
              _nw260[(new BigNumber(2)).toNumber()] = _1278_v11;
              _nw260[(new BigNumber(3)).toNumber()] = _1278_v11;
              _nw260[(new BigNumber(4)).toNumber()] = _1278_v11;
              _nw260[(new BigNumber(5)).toNumber()] = _1278_v11;
              _nw260[(new BigNumber(6)).toNumber()] = _1278_v11;
              _nw260[(new BigNumber(7)).toNumber()] = _1278_v11;
              _nw260[(new BigNumber(8)).toNumber()] = _1278_v11;
              _nw260[(new BigNumber(9)).toNumber()] = _1278_v11;
              _nw260[(new BigNumber(10)).toNumber()] = _1278_v11;
              _nw260[(new BigNumber(11)).toNumber()] = _1278_v11;
              _nw260[(new BigNumber(12)).toNumber()] = _1278_v11;
              _nw260[(new BigNumber(13)).toNumber()] = _1278_v11;
              _nw260[(new BigNumber(14)).toNumber()] = _1278_v11;
              _nw260[(new BigNumber(15)).toNumber()] = _1278_v11;
              _nw260[(new BigNumber(16)).toNumber()] = _1278_v11;
              _nw260[(new BigNumber(17)).toNumber()] = _1278_v11;
              _nw260[(new BigNumber(18)).toNumber()] = _1278_v11;
              _nw260[(new BigNumber(19)).toNumber()] = _1278_v11;
              _nw260[(new BigNumber(20)).toNumber()] = _1278_v11;
              _nw260[(new BigNumber(21)).toNumber()] = _1278_v11;
              _1279_v12 = _nw260;
              _1279_v12 = _1279_v12;
              (globalState).f7 = p1;
              let _1280_v13;
              _1280_v13 = _dafny.Seq.of(p3);
              (globalState).f2 = (p1).multipliedBy(_module.__default.safeDivisionInt(new BigNumber((_1280_v13).length), (_dafny.ZERO).minus(p1)));
            }
            let _1281_v14;
            _1281_v14 = new _dafny.CodePoint('p'.codePointAt(0));
            let _1282_v15;
            _1282_v15 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_1281_v14),(_this).f17);
            let _1283_v16;
            _1283_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1282_v15,_dafny.Map.Empty.slice().updateUnsafe(p3,false));
            let _1284_v17;
            _1284_v17 = _dafny.MultiSet.fromElements(_1281_v14, new _dafny.CodePoint('w'.codePointAt(0)), _1281_v14);
            _1283_v16 = (_1283_v16).update((_1282_v15).update(_1284_v17, _this.f16), _1265_v2);
            (globalState).f2 = p2;
            (globalState).f7 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(((p3) ? (p1) : (p1)),p1)).length);
          }
        }
      }
      let _1285_v18;
      _1285_v18 = _dafny.Set.fromElements(p2, new BigNumber((_dafny.Seq.UnicodeFromString("kgey")).length));
      let _1286_v21;
      _1286_v21 = _dafny.Seq.of(p0);
      let _1287_v22;
      let _nw261 = Array((new BigNumber(14)).toNumber());
      _nw261[(_dafny.ZERO).toNumber()] = p1;
      _nw261[(_dafny.ONE).toNumber()] = new BigNumber((_1286_v21).length);
      _nw261[(new BigNumber(2)).toNumber()] = p1;
      _nw261[(new BigNumber(3)).toNumber()] = p1;
      _nw261[(new BigNumber(4)).toNumber()] = new BigNumber((_1286_v21).length);
      _nw261[(new BigNumber(5)).toNumber()] = new BigNumber(-92);
      _nw261[(new BigNumber(6)).toNumber()] = new BigNumber(12);
      _nw261[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(p2);
      _nw261[(new BigNumber(8)).toNumber()] = new BigNumber(170);
      _nw261[(new BigNumber(9)).toNumber()] = new BigNumber(903);
      _nw261[(new BigNumber(10)).toNumber()] = p1;
      _nw261[(new BigNumber(11)).toNumber()] = new BigNumber((_1263_v1).length);
      _nw261[(new BigNumber(12)).toNumber()] = p2;
      _nw261[(new BigNumber(13)).toNumber()] = p2;
      _1287_v22 = _nw261;
      let _1288_v23;
      let _nw262 = Array((new BigNumber(27)).toNumber()).fill(false);
      _1288_v23 = _nw262;
      let _1289_v24;
      _1289_v24 = _module.D9.create_DC15(false, (_this).f21, function () {
  let _coll49 = new _dafny.Set();
  for (const _compr_49 of _dafny.IntegerRange(new BigNumber(413), new BigNumber(-815))) {
    let _1290_v20 = _compr_49;
    if (((new BigNumber(413)).isLessThanOrEqualTo(_1290_v20)) && ((_1290_v20).isLessThan(new BigNumber(-815)))) {
      _coll49.add((_1290_v20).minus(p1));
    }
  }
  return _coll49;
}(), _1287_v22, _1288_v23);
      let _1291_v25;
      _1291_v25 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1288_v23);
      let _1292_v26;
      _1292_v26 = _module.D9.create_DC15(false, _this.f16, (_1289_v24).dtor_cf20, _1287_v22, (((_1291_v25).contains(p1)) ? ((_1291_v25).get(p1)) : (_1288_v23)));
      let _1293_v28;
      let _nw263 = Array((new BigNumber(10)).toNumber());
      _nw263[(_dafny.ZERO).toNumber()] = (_1285_v18).Intersect(function () {
        let _coll50 = new _dafny.Set();
        for (const _compr_50 of (_dafny.Seq.of(p2)).Elements) {
          let _1294_v19 = _compr_50;
          if (_dafny.Seq.contains(_dafny.Seq.of(p2), _1294_v19)) {
            _coll50.add((_1294_v19).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("aracab")).length)));
          }
        }
        return _coll50;
      }());
      _nw263[(_dafny.ONE).toNumber()] = ((_1289_v24).dtor_cf20).Union(function () {
        let _coll51 = new _dafny.Set();
        for (const _compr_51 of _dafny.IntegerRange(new BigNumber(474), new BigNumber(449))) {
          let _1295_v27 = _compr_51;
          if (((new BigNumber(474)).isLessThanOrEqualTo(_1295_v27)) && ((_1295_v27).isLessThan(new BigNumber(449)))) {
            _coll51.add((_1295_v27).plus(p2));
          }
        }
        return _coll51;
      }());
      _nw263[(new BigNumber(2)).toNumber()] = _module.__default.fm49(new BigNumber(767), p2, _this.f16, (_dafny.ZERO).minus(p1), globalState);
      _nw263[(new BigNumber(3)).toNumber()] = _1285_v18;
      _nw263[(new BigNumber(4)).toNumber()] = (_1285_v18).Difference(_1285_v18);
      _nw263[(new BigNumber(5)).toNumber()] = _1285_v18;
      _nw263[(new BigNumber(6)).toNumber()] = (_1285_v18).Union(_1285_v18);
      _nw263[(new BigNumber(7)).toNumber()] = _1285_v18;
      _nw263[(new BigNumber(8)).toNumber()] = (((_this).f21) ? (_1285_v18) : (_1285_v18));
      _nw263[(new BigNumber(9)).toNumber()] = _1285_v18;
      _1293_v28 = _nw263;
      for (const _guard_loop_6 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1293_v28).length))) {
        let _1296_i1 = _guard_loop_6;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1296_i1)) && ((_1296_i1).isLessThan(new BigNumber((_1293_v28).length))))) {
          (_1293_v28)[(_1296_i1)] = (_1285_v18).Difference(_1285_v18);
        }
      }
      let _1297_v29;
      _1297_v29 = _dafny.MultiSet.fromElements((_this).f21);
      let _1298_v30;
      _1298_v30 = _module.D16.create_DC32((_1297_v29).Intersect(_1297_v29));
      let _source19 = _1298_v30;
      if (_source19.is_DC33) {
        let _1299___mcc_h1 = (_source19).cf47;
        let _1300___mcc_h2 = (_source19).cf48;
        let _1301___mcc_h3 = (_source19).cf49;
        let _1302___mcc_h4 = (_source19).cf50;
        let _1303_cf50 = _1302___mcc_h4;
        let _1304_cf49 = _1301___mcc_h3;
        let _1305_cf48 = _1300___mcc_h2;
        let _1306_cf47 = _1299___mcc_h1;
        (globalState).f7 = p2;
        let _1307_v31;
        _1307_v31 = new _dafny.CodePoint('i'.codePointAt(0));
        let _1308_v32;
        _1308_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1303_cf50,_1287_v22);
        let _rhs220 = _1307_v31;
        let _rhs221 = _1304_cf49;
        let _rhs222 = (_1308_v32).update(p3, _1287_v22);
        let _lhs169 = globalState;
        _1307_v31 = _rhs220;
        _lhs169.f7 = _rhs221;
        _1308_v32 = _rhs222;
        let _1309_v33;
        let _nw264 = Array((new BigNumber(9)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1309_v33 = _nw264;
        let _index216 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_1309_v33).length));
        (_1309_v33)[_index216] = ((true) ? (_1307_v31) : (_1307_v31));
        let _index217 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_1309_v33).length));
        (_1309_v33)[_index217] = _1307_v31;
        (globalState).f7 = (_1304_cf49).multipliedBy((((_this).f17) ? (new BigNumber((_module.__default.fm52(_1306_cf47, new BigNumber((_1285_v18).length), _1304_cf49, p2, globalState)).length)) : (p1)));
      } else if (_source19.is_DC34) {
        let _1310___mcc_h5 = (_source19).cf51;
        let _1311___mcc_h6 = (_source19).cf52;
        let _1312___mcc_h7 = (_source19).cf53;
        let _1313_cf53 = _1312___mcc_h7;
        let _1314_cf52 = _1311___mcc_h6;
        let _1315_cf51 = _1310___mcc_h5;
        let _index218 = _module.__default.safeIndex(new BigNumber(258), new BigNumber((_1288_v23).length));
        (_1288_v23)[_index218] = _1315_cf51;
        let _1316_v34;
        let _init42 = function (_1317_i2) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        };
        let _nw265 = Array((new BigNumber(2)).toNumber());
        for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw265.length); _i0_42++) {
          _nw265[_i0_42] = _init42(new BigNumber(_i0_42));
        }
        _1316_v34 = _nw265;
        let _1318_v35;
        _1318_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1315_cf51,_1316_v34);
        let _1319_v36;
        _1319_v36 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,(((_1318_v35).contains((_this).f21)) ? ((_1318_v35).get((_this).f21)) : (_1316_v34)));
        let _1320_v37;
        _1320_v37 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,p1);
        let _index219 = _module.__default.safeIndex(new BigNumber(258), new BigNumber((_1288_v23).length));
        let _rhs223 = ((new BigNumber((_1318_v35).length)).minus(p1)).minus(new BigNumber((_1286_v21).length));
        let _rhs224 = _module.__default.fm43(!(_1320_v37).contains(p3), globalState);
        let _lhs170 = globalState;
        let _lhs171 = _1288_v23;
        let _lhs172 = _module.__default.safeIndex(new BigNumber(258), new BigNumber((_1288_v23).length));
        _lhs170.f7 = _rhs223;
        _lhs171[_lhs172] = _rhs224;
        let _1321_v38;
        _1321_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1314_cf52.f14,new BigNumber((_1262_v0).length));
        (_1314_cf52).f14 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of((_this).f21, true)).length), (((_1321_v38).contains(_1313_cf53)) ? ((_1321_v38).get(_1313_cf53)) : (new BigNumber((_1263_v1).length))));
        let _1322_v39;
        let _nw266 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Set.Empty);
        _1322_v39 = _nw266;
        let _1323_v40;
        let _nw267 = new _module.C0();
        _nw267.__ctor(_1322_v39, p2);
        _1323_v40 = _nw267;
        (_this).f16 = !((_1297_v29).IsProperSubsetOf((_1298_v30).dtor_cf46));
      } else if (_source19.is_DC32) {
        let _1324___mcc_h8 = (_source19).cf46;
        let _1325_cf46 = _1324___mcc_h8;
        if ((_this).f17) {
          r0 = p0;
          let _index220 = _module.__default.safeIndex(new BigNumber(604), new BigNumber((_1287_v22).length));
          (_1287_v22)[_index220] = p1;
          let _index221 = _module.__default.safeIndex(new BigNumber(604), new BigNumber((_1287_v22).length));
          (_1287_v22)[_index221] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(p1, p2));
          let _1326_v41;
          let _nw268 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
          _1326_v41 = _nw268;
          let _1327_v42;
          _1327_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1326_v41,p0);
          let _1328_v43;
          _1328_v43 = _dafny.MultiSet.fromElements(_1326_v41, _1326_v41);
          let _index222 = _module.__default.safeIndex(new BigNumber(604), new BigNumber((_1287_v22).length));
          let _rhs225 = _1292_v26;
          let _rhs226 = (_this).f17;
          let _rhs227 = (_1328_v43).IsProperSubsetOf(_1328_v43);
          let _rhs228 = p2;
          let _rhs229 = _1327_v42;
          let _lhs173 = _1287_v22;
          let _lhs174 = _module.__default.safeIndex(new BigNumber(604), new BigNumber((_1287_v22).length));
          _1289_v24 = _rhs225;
          r0 = _rhs226;
          r0 = _rhs227;
          _lhs173[_lhs174] = _rhs228;
          _1327_v42 = _rhs229;
          (globalState).f7 = (_dafny.ZERO).minus(((_1287_v22)[_module.__default.safeIndex(new BigNumber(604), new BigNumber((_1287_v22).length))]).multipliedBy(p2));
          let _1329_v45;
          _1329_v45 = _dafny.Seq.of(function () {
            let _coll52 = new _dafny.Set();
            for (const _compr_52 of _dafny.IntegerRange(new BigNumber(370), new BigNumber(-776))) {
              let _1330_v44 = _compr_52;
              if (((new BigNumber(370)).isLessThanOrEqualTo(_1330_v44)) && ((_1330_v44).isLessThan(new BigNumber(-776)))) {
                _coll52.add((_1330_v44).plus(p2));
              }
            }
            return _coll52;
          }(), (_1285_v18).Intersect(_dafny.Set.fromElements(p2, p2, new BigNumber((_1262_v0).length))));
          _1285_v18 = (_1329_v45)[_module.__default.safeIndex((_1287_v22)[_module.__default.safeIndex(new BigNumber(604), new BigNumber((_1287_v22).length))], new BigNumber((_1329_v45).length))];
        } else {
          let _1331_v46;
          _1331_v46 = _module.D12.create_DC23(p0, p2, new BigNumber(182));
          let _rhs230 = ((_1331_v46).dtor_cf32).isEqualTo((_1263_v1)[_module.__default.safeIndex(p1, new BigNumber((_1263_v1).length))]);
          let _rhs231 = (_1286_v21)[_module.__default.safeIndex(p2, new BigNumber((_1286_v21).length))];
          r0 = _rhs230;
          r0 = _rhs231;
          r2 = p3;
          let _1332_v47;
          _1332_v47 = _module.D15.create_DC30(new _dafny.CodePoint('m'.codePointAt(0)));
          let _1333_v48;
          _1333_v48 = new _dafny.CodePoint('b'.codePointAt(0));
          _1332_v47 = function (_pat_let23_0) {
            return function (_1334_dt__update__tmp_h1) {
              return function (_pat_let24_0) {
                return function (_1335_dt__update_hcf44_h0) {
                  return _module.D15.create_DC30(_1335_dt__update_hcf44_h0);
                }(_pat_let24_0);
              }(new _dafny.CodePoint('m'.codePointAt(0)));
            }(_pat_let23_0);
          }(_module.D15.create_DC30(_1333_v48));
          let _1336_v49;
          _1336_v49 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1291_v25).length),(_1286_v21)[_module.__default.safeIndex(p2, new BigNumber((_1286_v21).length))]);
          _1336_v49 = (_1336_v49).update(p2, !((_this).f21));
          let _1337_v50;
          _1337_v50 = _dafny.Map.Empty.slice().updateUnsafe(p2,p1);
          let _1338_v52;
          _1338_v52 = _dafny.Map.Empty.slice().updateUnsafe((_1337_v50).Merge(function () {
            let _coll53 = new _dafny.Map();
            for (const _compr_53 of _dafny.IntegerRange(new BigNumber(549), new BigNumber(-902))) {
              let _1339_v51 = _compr_53;
              if (((new BigNumber(549)).isLessThanOrEqualTo(_1339_v51)) && ((_1339_v51).isLessThan(new BigNumber(-902)))) {
                _coll53.push([(_1339_v51).plus(p1),p2]);
              }
            }
            return _coll53;
          }()),_module.__default.fm31(!(true), !(p3), globalState));
          _1338_v52 = _1338_v52;
        }
        (globalState).f7 = p1;
        let _1340_v53;
        _1340_v53 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("pwnthagln")),_1263_v1);
        _1263_v1 = (((_1340_v53).contains(_dafny.Seq.of(_1262_v0))) ? ((_1340_v53).get(_dafny.Seq.of(_1262_v0))) : (_module.__default.fm35(p1, true, new BigNumber((_1262_v0).length), new _dafny.CodePoint('m'.codePointAt(0)), globalState)));
        let _1341_v54;
        _1341_v54 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(p2, p2),p2);
        let _1342_v55;
        _1342_v55 = _dafny.MultiSet.fromElements(p2);
        _1341_v54 = ((_dafny.Map.Empty.slice().updateUnsafe(p1,p2)).update(p1, p2)).update(p2, (((_1342_v55).contains(p1)) ? ((_1342_v55).get(p1)) : (p1)));
      } else {
        let _1343___mcc_h9 = (_source19).cf54;
        let _1344_cf54 = _1343___mcc_h9;
        let _index223 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_1287_v22).length));
        (_1287_v22)[_index223] = p1;
        let _index224 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_1287_v22).length));
        (_1287_v22)[_index224] = new BigNumber(952);
        (_this).f16 = !((_this).f21);
        _1262_v0 = (_module.D12.create_DC22(_dafny.Seq.UnicodeFromString("tncsoxvl"))).dtor_cf29;
        let _source20 = _module.__default.fm56(p1, (p2).multipliedBy(p1), globalState);
        if (_source20.is_DC28) {
          let _1345___mcc_h10 = (_source20).cf40;
          let _1346___mcc_h11 = (_source20).cf41;
          let _1347___mcc_h12 = (_source20).cf42;
          let _1348_cf42 = _1347___mcc_h12;
          let _1349_cf41 = _1346___mcc_h11;
          let _1350_cf40 = _1345___mcc_h10;
          r2 = !(p2).isEqualTo(_1348_cf42.f14);
          let _1351_v56;
          _1351_v56 = _dafny.MultiSet.fromElements(_1288_v23, _1288_v23);
          r2 = !((_1351_v56).Union(_1351_v56)).equals(_dafny.MultiSet.fromElements(_1288_v23));
          r0 = _this.f16;
          let _1352_v57;
          _1352_v57 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1349_cf41);
          let _1353_v58;
          _1353_v58 = _dafny.Set.fromElements(_1352_v57);
          let _1354_v60;
          _1354_v60 = new _dafny.CodePoint('b'.codePointAt(0));
          let _1355_v61;
          _1355_v61 = _module.D6.create_DC9(_1354_v60);
          let _rhs232 = (function () {
            let _coll54 = new _dafny.Set();
            for (const _compr_54 of (_dafny.Map.Empty.slice().updateUnsafe(_1352_v57,_1348_cf42.f14)).Keys.Elements) {
              let _1356_v59 = _compr_54;
              if ((_dafny.Map.Empty.slice().updateUnsafe(_1352_v57,_1348_cf42.f14)).contains(_1356_v59)) {
                _coll54.add(_1356_v59);
              }
            }
            return _coll54;
          }()).Union(_module.__default.fm57(globalState));
          let _rhs233 = p3;
          let _rhs234 = (_this).f21;
          let _rhs235 = ((_module.__default.fm27(new BigNumber((_1262_v0).length), p1, _1348_cf42.f14, _dafny.Map.Empty.slice().updateUnsafe(_1348_cf42.f14,_1348_cf42.f14), globalState)) ? (_1355_v61) : (_1355_v61));
          let _rhs236 = _module.__default.fm35(new BigNumber(92), (_1349_cf41) === (p3), _1348_cf42.f14, ((_module.__default.fm43(_this.f16, globalState)) ? (_1354_v60) : ((_1355_v61).dtor_cf14)), globalState);
          let _lhs175 = _this;
          _1353_v58 = _rhs232;
          r0 = _rhs233;
          _lhs175.f16 = _rhs234;
          r1 = _rhs235;
          _1263_v1 = _rhs236;
        } else {
          let _1357___mcc_h13 = (_source20).cf39;
          let _1358_cf39 = _1357___mcc_h13;
          let _1359_v62;
          let _nw269 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
          _1359_v62 = _nw269;
          _1359_v62 = _1359_v62;
          let _index225 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_1288_v23).length));
          (_1288_v23)[_index225] = _this.f16;
          let _index226 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_1288_v23).length));
          (_1288_v23)[_index226] = !(!(!(!(_1297_v29).equals((_1297_v29).Union(_1297_v29)))));
          let _nw270 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
          _1359_v62 = _nw270;
          let _index227 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_1288_v23).length));
          (_1288_v23)[_index227] = (_1286_v21)[_module.__default.safeIndex(p2, new BigNumber((_1286_v21).length))];
        }
      }
      let _hi9 = p1;
      for (let _1360_i3 = (_dafny.ZERO).minus((p1).multipliedBy(p1)); _1360_i3.isLessThan(_hi9); _1360_i3 = _1360_i3.plus(_dafny.ONE)) {
        _1286_v21 = _1286_v21;
        let _1361_v64;
        _1361_v64 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,p2);
        let _1362_v65;
        _1362_v65 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_1361_v64).length)),_1263_v1);
        let _1363_v66;
        _1363_v66 = _dafny.Map.Empty.slice().updateUnsafe((function () {
          let _coll55 = new _dafny.Map();
          for (const _compr_55 of (_dafny.Set.fromElements(p1, _1360_i3)).Elements) {
            let _1364_v63 = _compr_55;
            if ((_dafny.Set.fromElements(p1, _1360_i3)).contains(_1364_v63)) {
              _coll55.push([(_1364_v63).plus(new BigNumber((_dafny.Seq.UnicodeFromString("puh")).length)),_1263_v1]);
            }
          }
          return _coll55;
        }()).Merge(_1362_v65),_1285_v18);
        _1363_v66 = (_1363_v66).update(_1362_v65, _1285_v18);
        let _1365_v67;
        let _nw271 = new _module.C3();
        _nw271.__ctor(_1360_i3, _1262_v0, (_1360_i3).isEqualTo(_1360_i3), !((_this).f21));
        _1365_v67 = _nw271;
        (globalState).f2 = (_dafny.ZERO).minus((new BigNumber(((_1365_v67).f23).length)).plus(_1360_i3));
      }
      let _1366_v68;
      _1366_v68 = _dafny.MultiSet.fromElements(p2);
      if (_module.__default.fm31((_1366_v68).IsDisjointFrom(_1366_v68), _dafny.Seq.IsPrefixOf(_1262_v0, _1262_v0), globalState)) {
        (globalState).f7 = _module.__default.safeDivisionInt((p1).plus(p2), p2);
        let _1367_v69;
        _1367_v69 = _dafny.Set.fromElements(_this.f16);
        (globalState).f7 = _module.__default.safeModuloInt((p2).minus(p1), (new BigNumber((_1367_v69).length)).minus((_dafny.ZERO).minus(new BigNumber((_1262_v0).length))));
        let _1368_v70;
        _1368_v70 = _dafny.Seq.of(_1286_v21);
        let _1369_v71;
        let _nw272 = Array((new BigNumber(17)).toNumber());
        _nw272[(_dafny.ZERO).toNumber()] = _1286_v21;
        _nw272[(_dafny.ONE).toNumber()] = _1286_v21;
        _nw272[(new BigNumber(2)).toNumber()] = _module.__default.fm52((_this).f21, new BigNumber(-962), p1, new BigNumber(129), globalState);
        _nw272[(new BigNumber(3)).toNumber()] = _1286_v21;
        _nw272[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_1286_v21, _1286_v21);
        _nw272[(new BigNumber(5)).toNumber()] = _module.__default.fm52((_this).f21, new BigNumber((_1263_v1).length), p1, p2, globalState);
        _nw272[(new BigNumber(6)).toNumber()] = _1286_v21;
        _nw272[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_1286_v21, (_1368_v70)[_module.__default.safeIndex(p1, new BigNumber((_1368_v70).length))]);
        _nw272[(new BigNumber(8)).toNumber()] = _1286_v21;
        _nw272[(new BigNumber(9)).toNumber()] = _1286_v21;
        _nw272[(new BigNumber(10)).toNumber()] = _1286_v21;
        _nw272[(new BigNumber(11)).toNumber()] = _1286_v21;
        _nw272[(new BigNumber(12)).toNumber()] = _1286_v21;
        _nw272[(new BigNumber(13)).toNumber()] = _1286_v21;
        _nw272[(new BigNumber(14)).toNumber()] = _1286_v21;
        _nw272[(new BigNumber(15)).toNumber()] = _dafny.Seq.of(true);
        _nw272[(new BigNumber(16)).toNumber()] = _1286_v21;
        _1369_v71 = _nw272;
        _1369_v71 = _1369_v71;
        let _1370_v72;
        _1370_v72 = new _dafny.CodePoint('g'.codePointAt(0));
        _1262_v0 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.update(_1262_v0, _module.__default.safeIndex(p1, new BigNumber((_1262_v0).length)), _1370_v72), _1262_v0), _1262_v0);
        let _index228 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_1288_v23).length));
        (_1288_v23)[_index228] = (_this).f17;
        let _index229 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_1288_v23).length));
        (_1288_v23)[_index229] = false;
      } else {
        let _1371_v74;
        _1371_v74 = _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((_1297_v29).cardinality())), function () {
          let _coll56 = new _dafny.Set();
          for (const _compr_56 of (_1263_v1).Elements) {
            let _1372_v73 = _compr_56;
            if (_dafny.Seq.contains(_1263_v1, _1372_v73)) {
              _coll56.add(_module.__default.safeDivisionInt(_1372_v73, new BigNumber(772)));
            }
          }
          return _coll56;
        }(), _1285_v18);
        let _1373_v75;
        _1373_v75 = _1371_v74;
        let _1374_v76;
        _1374_v76 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f17) || (!((_this).f21)),(_1373_v75));
        _1374_v76 = (_1374_v76).update((new BigNumber(735)).isLessThan(p1), _1371_v74);
        (globalState).f2 = (new BigNumber(334)).plus(p1);
        r2 = !(true);
        let _1375_v77;
        let _nw273 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Set.Empty);
        _1375_v77 = _nw273;
        let _1376_v78;
        _1376_v78 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1375_v77);
        let _1377_v79;
        let _nw274 = new _module.C0();
        _nw274.__ctor(((p0) ? (_1375_v77) : ((((_1376_v78).contains(p2)) ? ((_1376_v78).get(p2)) : (_1375_v77)))), p2);
        _1377_v79 = _nw274;
        let _1378_v80;
        _1378_v80 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1377_v79.f27);
        let _1379_v81;
        _1379_v81 = (_1378_v80).Merge(_1378_v80);
        _1379_v81 = _1379_v81;
      }
      r0 = _this.f16;
      let _1380_v82;
      _1380_v82 = _dafny.Map.Empty.slice().updateUnsafe(p2,true);
      let _1381_v83;
      _1381_v83 = _module.D7.create_DC11(_1380_v82);
      r1 = function (_source21) {
        if (_source21.is_DC12) {
          return _module.D6.create_DC9(new _dafny.CodePoint('n'.codePointAt(0)));
        } else {
          let _1382___mcc_h14 = (_source21).cf15;
          let _1383_cf15 = _1382___mcc_h14;
          return _module.D6.create_DC9(new _dafny.CodePoint('i'.codePointAt(0)));
        }
      }(_1381_v83);
      r2 = (_this).f21;
      return [r0, r1, r2];
    }
    get f21() {
      let _this = this;
      return _this._f21;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f16 = false;
      this._f17 = false;
    }
    _parentTraits() {
      return [_module.T2, _module.T1];
    }
    get f16() {
      let _this = this;
      return _this._f16;
    };
    set f16(value) {
      let _this = this;
      _this._f16 = value;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    __ctor(f16, f17) {
      let _this = this;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      return;
    }
    fm12(p0, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Set.fromElements((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-22),(_this).f17)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('s'.codePointAt(0)))).length),false)))).length);
    };
    fm13(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Seq.of(!((_this).f17));
    };
    fm3(globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(586),new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f17, (_this).f17))).cardinality()))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-351),new BigNumber(589)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_this).f17)).length),_this.f16)).length))).length),new BigNumber(355)));
    };
    fm23(p0, globalState) {
      let _this = this;
      return false;
    };
    fm24(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.MultiSet.fromElements(((_module.D7.create_DC11(function () {
  let _coll57 = new _dafny.Map();
  for (const _compr_57 of _dafny.IntegerRange(new BigNumber(226), new BigNumber(398))) {
    let _1384_v0 = _compr_57;
    if (((new BigNumber(226)).isLessThanOrEqualTo(_1384_v0)) && ((_1384_v0).isLessThan(new BigNumber(398)))) {
      _coll57.push([_module.__default.safeDivisionInt(_1384_v0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(40)), function (_1385_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      })).length)),(_this).f17]);
    }
  }
  return _coll57;
}())).dtor_cf15).Merge(function () {
        let _coll58 = new _dafny.Map();
        for (const _compr_58 of _dafny.IntegerRange(new BigNumber(-579), new BigNumber(47))) {
          let _1386_v1 = _compr_58;
          if (((new BigNumber(-579)).isLessThanOrEqualTo(_1386_v1)) && ((_1386_v1).isLessThan(new BigNumber(47)))) {
            _coll58.push([_module.__default.safeDivisionInt(_1386_v1, new BigNumber(190)),(_this).f17]);
          }
        }
        return _coll58;
      }()), ((_this.f16) ? (function () {
        let _coll59 = new _dafny.Map();
        for (const _compr_59 of (function () {
          let _coll60 = new _dafny.Set();
          for (const _compr_60 of _dafny.IntegerRange(new BigNumber(109), new BigNumber(285))) {
            let _1387_v3 = _compr_60;
            if (((new BigNumber(109)).isLessThanOrEqualTo(_1387_v3)) && ((_1387_v3).isLessThan(new BigNumber(285)))) {
              _coll60.add((_1387_v3).multipliedBy(new BigNumber((_dafny.Seq.of((_this).f17)).length)));
            }
          }
          return _coll60;
        }()).Elements) {
          let _1388_v2 = _compr_59;
          if ((function () {
            let _coll61 = new _dafny.Set();
            for (const _compr_61 of _dafny.IntegerRange(new BigNumber(109), new BigNumber(285))) {
              let _1389_v3 = _compr_61;
              if (((new BigNumber(109)).isLessThanOrEqualTo(_1389_v3)) && ((_1389_v3).isLessThan(new BigNumber(285)))) {
                _coll61.add((_1389_v3).multipliedBy(new BigNumber((_dafny.Seq.of((_this).f17)).length)));
              }
            }
            return _coll61;
          }()).contains(_1388_v2)) {
            _coll59.push([_module.__default.safeDivisionInt(_1388_v2, new BigNumber(595)),!((_this).f17)]);
          }
        }
        return _coll59;
      }()) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(400),true))), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(392),_this.f16), (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements((_this).f17, true)).cardinality()),(_this).f17)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements((_this).f17)).length),(_this).f17)), _dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_this.f16));
    };
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.Seq.UnicodeFromString("");
      let r3 = false;
      (_this).f16 = !(p1);
      let _1390_v0;
      _1390_v0 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f17) === (false),(new BigNumber(-772)).plus(p2));
      _1390_v0 = (_1390_v0).update(_this.f16, p2);
      let _1391_v1;
      _1391_v1 = _dafny.Seq.UnicodeFromString("aaw");
      (globalState).f2 = (new BigNumber((_1391_v1).length)).plus((_dafny.ZERO).minus((_this).fm12(p2, globalState)));
      let _1392_v2;
      _1392_v2 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("qsjaaalnq"));
      let _1393_v3;
      _1393_v3 = _dafny.Map.Empty.slice().updateUnsafe((p2).plus(p2),(_1392_v2).Union(_1392_v2));
      let _1394_v4;
      _1394_v4 = _dafny.Map.Empty.slice().updateUnsafe(p2,p1);
      _1393_v3 = (_1393_v3).update(new BigNumber((((p0) ? (_dafny.Map.Empty.slice().updateUnsafe(p2,p0)) : (_1394_v4))).length), ((_1392_v2).update(_1391_v1, _module.__default.abs(p2))).Intersect(_1392_v2));
      let _1395_v5;
      let _nw275 = Array((new BigNumber(18)).toNumber()).fill(_module.D3.Default());
      _1395_v5 = _nw275;
      for (const _guard_loop_7 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1395_v5).length))) {
        let _1396_i0 = _guard_loop_7;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1396_i0)) && ((_1396_i0).isLessThan(new BigNumber((_1395_v5).length))))) {
          (_1395_v5)[(_1396_i0)] = _module.D3.create_DC3(_this.f16);
        }
      }
      let _1397_v6;
      _1397_v6 = _dafny.Seq.of((_this).fm23(_this.f16, globalState));
      let _1398_v7;
      _1398_v7 = _dafny.Map.Empty.slice().updateUnsafe(false,_1397_v6);
      _1398_v7 = _1398_v7;
      r0 = (p2).multipliedBy(p2);
      r1 = (_this).f17;
      r2 = _dafny.Seq.UnicodeFromString("suk");
      r3 = _this.f16;
      return [r0, r1, r2, r3];
    }
    m2(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _1399_v0;
      let _nw276 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
      _1399_v0 = _nw276;
      let _index230 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length));
      (_1399_v0)[_index230] = p0;
      let _index231 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length));
      let _rhs237 = p0;
      let _rhs238 = ((_this).fm23(p1, globalState)) || ((_this).f17);
      let _lhs176 = _1399_v0;
      let _lhs177 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length));
      let _lhs178 = _this;
      _lhs176[_lhs177] = _rhs237;
      _lhs178.f16 = _rhs238;
      let _1400_v1;
      _1400_v1 = _dafny.Seq.UnicodeFromString("ma");
      _1400_v1 = _1400_v1;
      let _1401_i0;
      _1401_i0 = _dafny.ZERO;
      L6: {
        while (_this.f16) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1401_i0)) {
              break L6;
            }
            _1401_i0 = (_1401_i0).plus(_dafny.ONE);
            (globalState).f7 = p0;
            let _1402_v2;
            _1402_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,_this.f16);
            let _1403_v3;
            _1403_v3 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,(((_1402_v2).contains(p1)) ? ((_1402_v2).get(p1)) : (_this.f16)));
            let _1404_v4;
            _1404_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1402_v2).length),!(!(false)));
            let _1405_v5;
            _1405_v5 = _module.D7.create_DC11(_1404_v4);
            let _source22 = _1405_v5;
            if (_source22.is_DC12) {
              (_this).f16 = (_this).fm23(p1, globalState);
              let _1406_v6;
              _1406_v6 = new _dafny.CodePoint('i'.codePointAt(0));
              let _1407_v7;
              _1407_v7 = _dafny.Set.fromElements(_1406_v6, _1406_v6, _1406_v6, _1406_v6, new _dafny.CodePoint('h'.codePointAt(0)));
              (globalState).f2 = new BigNumber((((_1407_v7).Intersect(function () {
                let _coll62 = new _dafny.Set();
                for (const _compr_62 of (_dafny.MultiSet.FromArray(_1400_v1)).Elements) {
                  let _1408_v8 = _compr_62;
                  if ((_dafny.MultiSet.FromArray(_1400_v1)).contains(_1408_v8)) {
                    _coll62.add(_1408_v8);
                  }
                }
                return _coll62;
              }())).Intersect(_1407_v7)).length);
              let _1409_v9;
              let _nw277 = Array((new BigNumber(2)).toNumber()).fill(false);
              _1409_v9 = _nw277;
              let _1410_v10;
              _1410_v10 = _1409_v9;
              let _1411_v11;
              _1411_v11 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_1410_v10);
              let _1412_v12;
              _1412_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,_1411_v11);
              let _1413_v13;
              _1413_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1409_v9,(((_1412_v12).contains((_this).f17)) ? ((_1412_v12).get((_this).f17)) : (_1411_v11)));
              let _rhs239 = (_this.f16) === (p1);
              let _rhs240 = _1413_v13;
              let _rhs241 = p1;
              let _lhs179 = _this;
              let _lhs180 = _this;
              _lhs179.f16 = _rhs239;
              _1413_v13 = _rhs240;
              _lhs180.f16 = _rhs241;
              let _1414_v14;
              _1414_v14 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,_1409_v9);
              let _1415_v15;
              _1415_v15 = _dafny.Map.Empty.slice().updateUnsafe((((_1404_v4).contains(new BigNumber((_1414_v14).length))) ? ((_1404_v4).get(new BigNumber((_1414_v14).length))) : (_this.f16)),(p0).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1399_v0,p0)).length)));
              let _1416_v16;
              _1416_v16 = _1415_v15;
              _1415_v15 = ((_1416_v16)).Merge(_1415_v15);
            } else {
              let _1417___mcc_h0 = (_source22).cf15;
              let _1418_cf15 = _1417___mcc_h0;
              let _1419_v17;
              _1419_v17 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
              let _1420_v18;
              _1420_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1400_v1,new BigNumber(((_1419_v17).Merge(_module.__default.fm25((_1399_v0)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length))], globalState))).length));
              let _1421_v19;
              _1421_v19 = _dafny.Seq.of((_this).f17);
              let _1422_v20;
              _1422_v20 = _dafny.MultiSet.fromElements(_1421_v19);
              let _1423_v21;
              _1423_v21 = new _dafny.CodePoint('v'.codePointAt(0));
              let _1424_v22;
              _1424_v22 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1423_v21);
              let _1425_v23;
              _1425_v23 = _dafny.Map.Empty.slice().updateUnsafe((((_1422_v20).contains(_1421_v19)) ? ((_1422_v20).get(_1421_v19)) : (p0)),_module.__default.fm26((_1399_v0)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length))], (_1399_v0)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length))], p0, new BigNumber((_1424_v22).length), globalState));
              _1420_v18 = (_1420_v18).update(_1400_v1, ((_1399_v0)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length))]).plus(new BigNumber((_1425_v23).length)));
              let _1426_v24;
              _1426_v24 = _dafny.Map.Empty.slice().updateUnsafe((_1399_v0)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length))],_1400_v1);
              _1426_v24 = (_1426_v24).update((_1399_v0)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length))], _1400_v1);
              let _1427_v25;
              let _nw278 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Set.Empty);
              _1427_v25 = _nw278;
              let _1428_v26;
              let _nw279 = new _module.C0();
              _nw279.__ctor(_1427_v25, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(350)), ((_1429_v21) => function (_1430_i1) {
                return _1429_v21;
              })(_1423_v21))).length));
              _1428_v26 = _nw279;
              (_this).f16 = _this.f16;
            }
            (_this).f16 = (_this).f17;
            let _1431_v27;
            _1431_v27 = _dafny.Seq.of(p0, (_1399_v0)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length))], new BigNumber(((_this).fm13((_1399_v0)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length))], p1, (_this).f17, globalState)).length), new BigNumber((_dafny.Seq.UnicodeFromString("jx")).length));
            let _1432_v28;
            _1432_v28 = _dafny.Seq.of(_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(774)), ((_1433_p0) => function (_1434_i2) {
              return _1433_p0;
            })(p0)), _1431_v27), _this.f16);
            _1432_v28 = _1432_v28;
          }
        }
      }
      if ((_this).f17) {
        let _1435_v29;
        _1435_v29 = _dafny.Seq.of(p1);
        let _1436_v30;
        _1436_v30 = _dafny.Map.Empty.slice().updateUnsafe((_1399_v0)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length))],new BigNumber((_1435_v29).length));
        (_this).f16 = !(_1436_v30).contains(new BigNumber(-434));
        let _1437_v31;
        _1437_v31 = new _dafny.CodePoint('y'.codePointAt(0));
        let _index232 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length));
        let _rhs242 = (_1399_v0)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length))];
        let _rhs243 = _1437_v31;
        let _rhs244 = ((_1399_v0)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length))]).plus((_1399_v0)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length))]);
        let _lhs181 = globalState;
        let _lhs182 = _1399_v0;
        let _lhs183 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length));
        _lhs181.f7 = _rhs242;
        _1437_v31 = _rhs243;
        _lhs182[_lhs183] = _rhs244;
        let _1438_v32;
        let _nw280 = Array((new BigNumber(14)).toNumber());
        _nw280[(_dafny.ZERO).toNumber()] = _1400_v1;
        _nw280[(_dafny.ONE).toNumber()] = _1400_v1;
        _nw280[(new BigNumber(2)).toNumber()] = _1400_v1;
        _nw280[(new BigNumber(3)).toNumber()] = _1400_v1;
        _nw280[(new BigNumber(4)).toNumber()] = _1400_v1;
        _nw280[(new BigNumber(5)).toNumber()] = _1400_v1;
        _nw280[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_1400_v1, _1400_v1);
        _nw280[(new BigNumber(7)).toNumber()] = _dafny.Seq.UnicodeFromString("sfbtphsjv");
        _nw280[(new BigNumber(8)).toNumber()] = _1400_v1;
        _nw280[(new BigNumber(9)).toNumber()] = _1400_v1;
        _nw280[(new BigNumber(10)).toNumber()] = _1400_v1;
        _nw280[(new BigNumber(11)).toNumber()] = _1400_v1;
        _nw280[(new BigNumber(12)).toNumber()] = _dafny.Seq.UnicodeFromString("ocouv");
        _nw280[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_1400_v1, _1400_v1);
        _1438_v32 = _nw280;
        let _index233 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_1438_v32).length));
        (_1438_v32)[_index233] = _dafny.Seq.UnicodeFromString("xaxrbysaq");
        let _index234 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_1438_v32).length));
        (_1438_v32)[_index234] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-244)), ((_1439_v31) => function (_1440_i3) {
          return _1439_v31;
        })(_1437_v31)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(308)), function (_1441_i4) {
          return new _dafny.CodePoint('p'.codePointAt(0));
        }));
        let _1442_v33;
        let _nw281 = new _module.C1();
        _nw281.__ctor(p1, _module.__default.safeModuloInt(p0, new BigNumber((_1436_v30).length)));
        _1442_v33 = _nw281;
        let _1443_v34;
        _1443_v34 = _dafny.Set.fromElements(_1437_v31, _1437_v31);
        let _1444_v35;
        let _nw282 = new _module.C4();
        _nw282.__ctor(p1, (_this).f17, (_1443_v34).IsSubsetOf(_1443_v34));
        _1444_v35 = _nw282;
      } else {
        let _1445_v36;
        _1445_v36 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f17);
        let _1446_v37;
        _1446_v37 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_1400_v1).length));
        (_this).f16 = (_module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_1445_v36).length)), (_dafny.ZERO).minus(p0))).isLessThanOrEqualTo((((_1446_v37).contains(p1)) ? ((_1446_v37).get(p1)) : ((_1399_v0)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length))])));
        let _1447_v38;
        _1447_v38 = _dafny.Seq.of(_this.f16);
        let _1448_v39;
        _1448_v39 = _1447_v38;
        let _1449_v40;
        _1449_v40 = _module.D4.create_DC7(_1448_v39, new BigNumber(357));
        let _1450_v41;
        _1450_v41 = _dafny.Seq.of(p0);
        let _1451_v42;
        _1451_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1449_v40,(_1450_v41)[_module.__default.safeIndex((_1399_v0)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length))], new BigNumber((_1450_v41).length))]);
        (globalState).f2 = (_dafny.ZERO).minus((((_1451_v42).contains(_1449_v40)) ? ((_1451_v42).get(_1449_v40)) : ((_1399_v0)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length))])));
        r0 = _module.__default.fm0(p1, globalState);
        (_this).f16 = _module.__default.fm43(((_1399_v0)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length))]).isLessThan(new BigNumber((_1450_v41).length)), globalState);
        let _1452_v43;
        _1452_v43 = _module.D3.create_DC3(_this.f16);
        let _1453_v44;
        let _nw283 = new _module.C2();
        _nw283.__ctor(_1452_v43, p1, (_this).f17);
        _1453_v44 = _nw283;
      }
      let _1454_v45;
      _1454_v45 = _dafny.Seq.of(p0);
      let _1455_v46;
      _1455_v46 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1454_v45).length),p1);
      let _1456_v48;
      _1456_v48 = _dafny.MultiSet.fromElements((_1399_v0)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length))]);
      let _1457_v49;
      _1457_v49 = _module.D7.create_DC11(function () {
  let _coll63 = new _dafny.Map();
  for (const _compr_63 of (_1456_v48).Elements) {
    let _1458_v47 = _compr_63;
    if ((_1456_v48).contains(_1458_v47)) {
      _coll63.push([(_1458_v47).plus((_1399_v0)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length))]),_this.f16]);
    }
  }
  return _coll63;
}());
      let _pat_let_tv18 = _1399_v0;
      let _pat_let_tv19 = _1399_v0;
      let _pat_let_tv20 = _1399_v0;
      let _pat_let_tv21 = _1399_v0;
      let _index235 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length));
      (_1399_v0)[_index235] = function (_source23) {
        if (_source23.is_DC12) {
          return (_pat_let_tv19)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_pat_let_tv18).length))];
        } else {
          let _1459___mcc_h1 = (_source23).cf15;
          let _1460_cf15 = _1459___mcc_h1;
          return (_pat_let_tv21)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_pat_let_tv20).length))];
        }
      }((((_this).fm23((_this).f17, globalState)) ? (_module.D7.create_DC11(_1455_v46)) : (_1457_v49)));
      let _1461_v50;
      let _nw284 = new _module.C1();
      _nw284.__ctor((_this).f17, p0);
      _1461_v50 = _nw284;
      let _1462_v51;
      _1462_v51 = _module.D14.create_DC28((_1399_v0)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length))], p1, _1461_v50);
      let _1463_v52;
      _1463_v52 = _module.D3.create_DC4(p1, p0, (_1462_v51).dtor_cf40, _dafny.Seq.IsPrefixOf(_1454_v45, _dafny.Seq.update(_1454_v45, _module.__default.safeIndex(p0, new BigNumber((_1454_v45).length)), _1461_v50.f14)), (_this.f16) || (true));
      let _source24 = _1463_v52;
      if (_source24.is_DC4) {
        let _1464___mcc_h2 = (_source24).cf4;
        let _1465___mcc_h3 = (_source24).cf5;
        let _1466___mcc_h4 = (_source24).cf6;
        let _1467___mcc_h5 = (_source24).cf7;
        let _1468___mcc_h6 = (_source24).cf8;
        let _1469_cf8 = _1468___mcc_h6;
        let _1470_cf7 = _1467___mcc_h5;
        let _1471_cf6 = _1466___mcc_h4;
        let _1472_cf5 = _1465___mcc_h3;
        let _1473_cf4 = _1464___mcc_h2;
        let _index236 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length));
        (_1399_v0)[_index236] = _1471_cf6;
        let _1474_v54;
        _1474_v54 = _dafny.Seq.of(function () {
          let _coll64 = new _dafny.Map();
          for (const _compr_64 of _dafny.IntegerRange(new BigNumber(699), new BigNumber(566))) {
            let _1475_v53 = _compr_64;
            if (((new BigNumber(699)).isLessThanOrEqualTo(_1475_v53)) && ((_1475_v53).isLessThan(new BigNumber(566)))) {
              _coll64.push([(_1475_v53).multipliedBy(new BigNumber((_1400_v1).length)),true]);
            }
          }
          return _coll64;
        }());
        let _1476_v55;
        _1476_v55 = _dafny.Map.Empty.slice().updateUnsafe((_1474_v54)[_module.__default.safeIndex(p0, new BigNumber((_1474_v54).length))],_1472_cf5);
        let _1477_v56;
        let _nw285 = Array((new BigNumber(7)).toNumber());
        _nw285[(_dafny.ZERO).toNumber()] = (_1476_v55).update(_1455_v46, (_1399_v0)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length))]);
        _nw285[(_dafny.ONE).toNumber()] = _1476_v55;
        _nw285[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_1457_v49).dtor_cf15,(_1399_v0)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length))]);
        _nw285[(new BigNumber(3)).toNumber()] = (_1476_v55).Merge(_1476_v55);
        _nw285[(new BigNumber(4)).toNumber()] = _1476_v55;
        _nw285[(new BigNumber(5)).toNumber()] = _1476_v55;
        _nw285[(new BigNumber(6)).toNumber()] = _1476_v55;
        _1477_v56 = _nw285;
        let _index237 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_1477_v56).length));
        (_1477_v56)[_index237] = _1476_v55;
        let _1478_v57;
        let _nw286 = Array((new BigNumber(26)).toNumber()).fill(false);
        _1478_v57 = _nw286;
        let _index238 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_1478_v57).length));
        (_1478_v57)[_index238] = (_this).f17;
        let _index239 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_1477_v56).length));
        let _index240 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_1478_v57).length));
        let _rhs245 = ((_1476_v55).Merge(function () {
          let _coll65 = new _dafny.Map();
          for (const _compr_65 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_1461_v50.f14,true),new BigNumber(163))).Keys.Elements) {
            let _1479_v58 = _compr_65;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_1461_v50.f14,true),new BigNumber(163))).contains(_1479_v58)) {
              _coll65.push([_1479_v58,(_dafny.ZERO).minus(_1461_v50.f14)]);
            }
          }
          return _coll65;
        }())).Merge(_1476_v55);
        let _rhs246 = _this.f16;
        let _lhs184 = _1477_v56;
        let _lhs185 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_1477_v56).length));
        let _lhs186 = _1478_v57;
        let _lhs187 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_1478_v57).length));
        _lhs184[_lhs185] = _rhs245;
        _lhs186[_lhs187] = _rhs246;
        let _1480_v59;
        let _init43 = ((_1481_v1, _1482_v46) => function (_1483_i5) {
          return _dafny.Seq.update(_1481_v1, _module.__default.safeIndex(new BigNumber((_1482_v46).length), new BigNumber((_1481_v1).length)), new _dafny.CodePoint('h'.codePointAt(0)));
        })(_1400_v1, _1455_v46);
        let _nw287 = Array((new BigNumber(2)).toNumber());
        for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw287.length); _i0_43++) {
          _nw287[_i0_43] = _init43(new BigNumber(_i0_43));
        }
        _1480_v59 = _nw287;
        let _1484_v60;
        _1484_v60 = _1454_v45;
        let _rhs247 = _1461_v50.f14;
        let _rhs248 = false;
        let _rhs249 = ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("pqlcg")).length))).Difference(_1456_v48)).IsDisjointFrom(_dafny.MultiSet.FromArray((_1484_v60)));
        let _rhs250 = _1480_v59;
        let _lhs188 = globalState;
        _lhs188.f7 = _rhs247;
        _1473_cf4 = _rhs248;
        _1469_cf8 = _rhs249;
        _1480_v59 = _rhs250;
        let _1485_v61;
        _1485_v61 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,_this.f16);
        let _1486_v62;
        _1486_v62 = new _dafny.CodePoint('h'.codePointAt(0));
        let _index241 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_1478_v57).length));
        let _index242 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_1478_v57).length));
        let _rhs251 = _dafny.Seq.Concat(_1400_v1, _dafny.Seq.Concat(_1400_v1, _1400_v1));
        let _rhs252 = (((_1485_v61).contains((!(_1470_cf7)) || (_this.f16))) ? ((_1485_v61).get((!(_1470_cf7)) || (_this.f16))) : ((_1478_v57)[_module.__default.safeIndex(new BigNumber(568), new BigNumber((_1478_v57).length))]));
        let _rhs253 = _dafny.Seq.IsProperPrefixOf(_1400_v1, _dafny.Seq.update(_1400_v1, _module.__default.safeIndex((_1454_v45)[_module.__default.safeIndex((_1399_v0)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length))], new BigNumber((_1454_v45).length))], new BigNumber((_1400_v1).length)), _1486_v62));
        let _lhs189 = _1478_v57;
        let _lhs190 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_1478_v57).length));
        let _lhs191 = _1478_v57;
        let _lhs192 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_1478_v57).length));
        _1400_v1 = _rhs251;
        _lhs189[_lhs190] = _rhs252;
        _lhs191[_lhs192] = _rhs253;
      } else if (_source24.is_DC3) {
        let _1487___mcc_h7 = (_source24).cf3;
        let _1488_cf3 = _1487___mcc_h7;
        let _1489_v63;
        let _nw288 = Array((new BigNumber(8)).toNumber());
        _nw288[(_dafny.ZERO).toNumber()] = _1488_cf3;
        _nw288[(_dafny.ONE).toNumber()] = (_this).f17;
        _nw288[(new BigNumber(2)).toNumber()] = _this.f16;
        _nw288[(new BigNumber(3)).toNumber()] = p1;
        _nw288[(new BigNumber(4)).toNumber()] = !(!(_this.f16));
        _nw288[(new BigNumber(5)).toNumber()] = false;
        _nw288[(new BigNumber(6)).toNumber()] = p1;
        _nw288[(new BigNumber(7)).toNumber()] = _this.f16;
        _1489_v63 = _nw288;
        let _index243 = _module.__default.safeIndex(new BigNumber(340), new BigNumber((_1489_v63).length));
        (_1489_v63)[_index243] = _1488_cf3;
        let _1490_v64;
        _1490_v64 = _dafny.Seq.of(p1, false, false);
        let _index244 = _module.__default.safeIndex(new BigNumber(340), new BigNumber((_1489_v63).length));
        (_1489_v63)[_index244] = _module.__default.fm31((_this).f17, (_1490_v64)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("xus")).length), new BigNumber((_1490_v64).length))], globalState);
        let _index245 = _module.__default.safeIndex(new BigNumber(340), new BigNumber((_1489_v63).length));
        (_1489_v63)[_index245] = (_this).f17;
        let _index246 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length));
        (_1399_v0)[_index246] = new BigNumber((_1454_v45).length);
        let _1491_v65;
        _1491_v65 = _module.D9.create_DC16(p0, (_1489_v63)[_module.__default.safeIndex(new BigNumber(340), new BigNumber((_1489_v63).length))]);
        let _1492_v66;
        _1492_v66 = _module.D9.create_DC17(_1491_v65);
        let _1493_v67;
        _1493_v67 = _dafny.Map.Empty.slice().updateUnsafe((_1489_v63)[_module.__default.safeIndex(new BigNumber(340), new BigNumber((_1489_v63).length))],_1488_cf3);
        let _1494_v68;
        _1494_v68 = _dafny.Map.Empty.slice().updateUnsafe(!(_1488_cf3),_1493_v67);
        let _rhs254 = _module.D9.create_DC17(_module.D9.create_DC16(new BigNumber((_1454_v45).length), (_this).f17));
        let _rhs255 = ((_1494_v68).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f16,_1493_v67))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(_1488_cf3),_1493_v67));
        _1492_v66 = _rhs254;
        _1494_v68 = _rhs255;
      } else {
        let _1495___mcc_h8 = (_source24).cf9;
        let _1496_cf9 = _1495___mcc_h8;
        let _1497_v69;
        let _init44 = ((_1498_v1) => function (_1499_i6) {
          return _dafny.Seq.IsProperPrefixOf(_1498_v1, _1498_v1);
        })(_1400_v1);
        let _nw289 = Array((new BigNumber(21)).toNumber());
        for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw289.length); _i0_44++) {
          _nw289[_i0_44] = _init44(new BigNumber(_i0_44));
        }
        _1497_v69 = _nw289;
        let _index247 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_1497_v69).length));
        (_1497_v69)[_index247] = (_this).f17;
        let _index248 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_1497_v69).length));
        (_1497_v69)[_index248] = (_this).f17;
        _1399_v0 = _1399_v0;
        let _index249 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length));
        (_1399_v0)[_index249] = (_1399_v0)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length))];
        if (true) {
          let _1500_v70;
          _1500_v70 = _dafny.MultiSet.fromElements(p1, p1, (_this).f17);
          let _1501_v71;
          _1501_v71 = _dafny.Seq.of(_this.f16);
          (_1461_v50).f14 = new BigNumber(((_1500_v70).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of((_this).f17, (_1497_v69)[_module.__default.safeIndex(new BigNumber(263), new BigNumber((_1497_v69).length))], !((_this).f17), (_this).f17), _1501_v71)))).cardinality());
          let _1502_v72;
          let _nw290 = Array((new BigNumber(8)).toNumber()).fill([]);
          _1502_v72 = _nw290;
          let _1503_v73;
          let _nw291 = Array((new BigNumber(2)).toNumber()).fill([]);
          _1503_v73 = _nw291;
          let _index250 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_1502_v72).length));
          (_1502_v72)[_index250] = _1503_v73;
          let _index251 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_1502_v72).length));
          (_1502_v72)[_index251] = _1503_v73;
          let _1504_v74;
          _1504_v74 = _dafny.Seq.of(_1399_v0, (((_this).f17) ? (_1399_v0) : (_1399_v0)), _1399_v0, _1399_v0, _1399_v0);
          let _index252 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length));
          (_1399_v0)[_index252] = new BigNumber((_1504_v74).length);
          let _1505_v75;
          let _nw292 = new _module.C4();
          _nw292.__ctor(false, _this.f16, true);
          _1505_v75 = _nw292;
          let _1506_v76;
          _1506_v76 = new _dafny.CodePoint('w'.codePointAt(0));
          let _1507_v77;
          _1507_v77 = _module.D6.create_DC9(_1506_v76);
          let _1508_v78;
          let _nw293 = Array((new BigNumber(5)).toNumber());
          _nw293[(_dafny.ZERO).toNumber()] = _1507_v77;
          _nw293[(_dafny.ONE).toNumber()] = _1507_v77;
          _nw293[(new BigNumber(2)).toNumber()] = _1507_v77;
          _nw293[(new BigNumber(3)).toNumber()] = _1507_v77;
          _nw293[(new BigNumber(4)).toNumber()] = _1507_v77;
          _1508_v78 = _nw293;
          _1508_v78 = _1508_v78;
        } else {
          let _1509_v79;
          _1509_v79 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1400_v1).length),_1400_v1);
          _1509_v79 = (_1509_v79).update(_1461_v50.f14, _1400_v1);
          let _1510_v80;
          _1510_v80 = _dafny.Seq.of((_this).f17);
          let _1511_v81;
          _1511_v81 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(603)), ((_1512_v1) => function (_1513_i7) {
            return new BigNumber((_1512_v1).length);
          })(_1400_v1))).length));
          let _1514_v82;
          _1514_v82 = _module.D9.create_DC15((_1497_v69)[_module.__default.safeIndex(new BigNumber(263), new BigNumber((_1497_v69).length))], (_this).f17, _1511_v81, _1399_v0, _1497_v69);
          let _1515_v83;
          let _nw294 = Array((new BigNumber(5)).toNumber());
          _nw294[(_dafny.ZERO).toNumber()] = _1514_v82;
          _nw294[(_dafny.ONE).toNumber()] = _1514_v82;
          _nw294[(new BigNumber(2)).toNumber()] = _1514_v82;
          _nw294[(new BigNumber(3)).toNumber()] = _1514_v82;
          _nw294[(new BigNumber(4)).toNumber()] = _1514_v82;
          _1515_v83 = _nw294;
          let _1516_v84;
          _1516_v84 = _dafny.Seq.of(_module.D3.create_DC3(_this.f16));
          let _1517_v85;
          _1517_v85 = _module.D12.create_DC24(_this.f16, _1510_v80, _dafny.Seq.of(_1515_v83), (_this).f17, _1516_v84);
          let _1518_v86;
          _1518_v86 = _dafny.Seq.of((_1517_v85).dtor_cf36, (_1497_v69)[_module.__default.safeIndex(new BigNumber(263), new BigNumber((_1497_v69).length))], (_this).f17);
          (_this).f16 = ((new BigNumber(476)).multipliedBy(new BigNumber((_1510_v80).length))).isLessThan(p0);
          let _index253 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1399_v0).length));
          (_1399_v0)[_index253] = new BigNumber(140);
          _1454_v45 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1454_v45, _1454_v45), _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of((_module.D16.create_DC34(true, _1461_v50, _1461_v50.f14)).dtor_cf53, _1461_v50.f14), _1454_v45), _module.__default.safeIndex(new BigNumber(87), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_module.D16.create_DC34(true, _1461_v50, _1461_v50.f14)).dtor_cf53, _1461_v50.f14), _1454_v45)).length)), _1461_v50.f14));
          _1497_v69 = _1497_v69;
        }
      }
      r0 = _module.__default.safeDivisionInt(p0, _1461_v50.f14);
      return r0;
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _1519_v0;
      _1519_v0 = _dafny.Seq.UnicodeFromString("cdkvfv");
      let _1520_v1;
      _1520_v1 = _dafny.Map.Empty.slice().updateUnsafe(_1519_v0,!(p1).isEqualTo(p1));
      if ((((_1520_v1).contains(_1519_v0)) ? ((_1520_v1).get(_1519_v0)) : ((_this).f17))) {
        let _1521_v2;
        _1521_v2 = _dafny.MultiSet.fromElements(p0);
        let _1522_v3;
        _1522_v3 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_1519_v0).length));
        let _1523_v4;
        _1523_v4 = _dafny.Seq.of(_module.__default.fm31(_this.f16, _module.__default.fm27(new BigNumber((_dafny.Seq.of((_this).f17)).length), new BigNumber((_1521_v2).cardinality()), p3, _1522_v3, globalState), globalState), _this.f16);
        _1523_v4 = _dafny.Seq.update(_1523_v4, _module.__default.safeIndex(p1, new BigNumber((_1523_v4).length)), !(false));
        let _1524_v5;
        _1524_v5 = _dafny.Map.Empty.slice().updateUnsafe(false,p2);
        let _1525_v6;
        _1525_v6 = _module.D4.create_DC6(_1524_v5);
        let _1526_v7;
        _1526_v7 = _dafny.Map.Empty.slice().updateUnsafe(p3,_this.f16);
        let _1527_v8;
        let _nw295 = Array((new BigNumber(25)).toNumber());
        _nw295[(_dafny.ZERO).toNumber()] = p1;
        _nw295[(_dafny.ONE).toNumber()] = _module.__default.fm0((_this).f17, globalState);
        _nw295[(new BigNumber(2)).toNumber()] = p1;
        _nw295[(new BigNumber(3)).toNumber()] = p1;
        _nw295[(new BigNumber(4)).toNumber()] = p1;
        _nw295[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(p1);
        _nw295[(new BigNumber(6)).toNumber()] = new BigNumber(558);
        _nw295[(new BigNumber(7)).toNumber()] = p3;
        _nw295[(new BigNumber(8)).toNumber()] = p3;
        _nw295[(new BigNumber(9)).toNumber()] = new BigNumber(841);
        _nw295[(new BigNumber(10)).toNumber()] = p3;
        _nw295[(new BigNumber(11)).toNumber()] = p1;
        _nw295[(new BigNumber(12)).toNumber()] = p1;
        _nw295[(new BigNumber(13)).toNumber()] = new BigNumber(762);
        _nw295[(new BigNumber(14)).toNumber()] = p3;
        _nw295[(new BigNumber(15)).toNumber()] = p1;
        _nw295[(new BigNumber(16)).toNumber()] = p3;
        _nw295[(new BigNumber(17)).toNumber()] = new BigNumber(((_1525_v6).dtor_cf10).length);
        _nw295[(new BigNumber(18)).toNumber()] = p3;
        _nw295[(new BigNumber(19)).toNumber()] = p1;
        _nw295[(new BigNumber(20)).toNumber()] = (_this).fm12(new BigNumber(865), globalState);
        _nw295[(new BigNumber(21)).toNumber()] = _module.__default.fm0((_this).f17, globalState);
        _nw295[(new BigNumber(22)).toNumber()] = new BigNumber(399);
        _nw295[(new BigNumber(23)).toNumber()] = new BigNumber(((_1526_v7).update(p1, !((_this).f17))).length);
        _nw295[(new BigNumber(24)).toNumber()] = p1;
        _1527_v8 = _nw295;
        let _1528_v9;
        _1528_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1519_v0,_1527_v8);
        _1528_v9 = (_1528_v9).update(_dafny.Seq.Concat(_1519_v0, _1519_v0), _1527_v8);
        let _1529_v10;
        let _nw296 = new _module.C1();
        _nw296.__ctor(_this.f16, p3);
        _1529_v10 = _nw296;
        let _1530_v11;
        _1530_v11 = _module.D14.create_DC28(p1, (_this).f17, _1529_v10);
        let _1531_v12;
        _1531_v12 = _dafny.Seq.of(p3, new BigNumber((_1519_v0).length), _1529_v10.f14, (_dafny.ZERO).minus(p3));
        let _1532_v13;
        let _nw297 = Array((new BigNumber(29)).toNumber()).fill(false);
        _1532_v13 = _nw297;
        let _1533_v14;
        _1533_v14 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm27(p1, (_1530_v11).dtor_cf40, new BigNumber((_1531_v12).length), _1522_v3, globalState),_1532_v13);
        let _1534_v15;
        _1534_v15 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_1523_v4).length)), p3);
        _1533_v14 = (_1533_v14).update((_1534_v15).IsProperSubsetOf(_1534_v15), _1532_v13);
        let _index254 = _module.__default.safeIndex(new BigNumber(10), new BigNumber((_1532_v13).length));
        (_1532_v13)[_index254] = (new BigNumber(732)).isLessThan(new BigNumber(912));
        let _index255 = _module.__default.safeIndex(new BigNumber(10), new BigNumber((_1532_v13).length));
        (_1532_v13)[_index255] = !(_this.f16) || (false);
        (globalState).f2 = _module.__default.fm0((_this).f17, globalState);
      } else {
        let _1535_v16;
        _1535_v16 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,p2);
        let _1536_v17;
        _1536_v17 = _dafny.Seq.of((_this).f17);
        let _1537_v18;
        let _nw298 = new _module.C3();
        _nw298.__ctor((p1).multipliedBy(new BigNumber((_1519_v0).length)), _1519_v0, (((_1535_v16).contains((_1536_v17)[_module.__default.safeIndex(new BigNumber((_1535_v16).length), new BigNumber((_1536_v17).length))])) ? ((_1535_v16).get((_1536_v17)[_module.__default.safeIndex(new BigNumber((_1535_v16).length), new BigNumber((_1536_v17).length))])) : (p2)), true);
        _1537_v18 = _nw298;
        let _1538_v19;
        let _nw299 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
        _1538_v19 = _nw299;
        let _index256 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1538_v19).length));
        (_1538_v19)[_index256] = p1;
        let _index257 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1538_v19).length));
        let _rhs256 = p3;
        let _rhs257 = _1519_v0;
        let _rhs258 = new BigNumber(563);
        let _lhs193 = globalState;
        let _lhs194 = _1538_v19;
        let _lhs195 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1538_v19).length));
        _lhs193.f2 = _rhs256;
        _1519_v0 = _rhs257;
        _lhs194[_lhs195] = _rhs258;
        if ((_this).f17) {
          let _1539_v20;
          let _init45 = function (_1540_i0) {
            return _this.f16;
          };
          let _nw300 = Array((new BigNumber(12)).toNumber());
          for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw300.length); _i0_45++) {
            _nw300[_i0_45] = _init45(new BigNumber(_i0_45));
          }
          _1539_v20 = _nw300;
          let _index258 = _module.__default.safeIndex(new BigNumber(370), new BigNumber((_1539_v20).length));
          (_1539_v20)[_index258] = true;
          let _1541_v21;
          _1541_v21 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_1537_v18).f22,(_this).f17));
          let _index259 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1538_v19).length));
          let _index260 = _module.__default.safeIndex(new BigNumber(370), new BigNumber((_1539_v20).length));
          let _rhs259 = ((p2) ? (_dafny.Seq.UnicodeFromString("wikghvh")) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(792)), ((_1542_p0) => function (_1543_i1) {
            return _1542_p0;
          })(p0))));
          let _rhs260 = _module.__default.safeDivisionInt(new BigNumber((_1541_v21).length), (_dafny.ZERO).minus(p3));
          let _rhs261 = new BigNumber((_1519_v0).length);
          let _rhs262 = (_this).f17;
          let _rhs263 = (_1537_v18).f22;
          let _lhs196 = globalState;
          let _lhs197 = _1538_v19;
          let _lhs198 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1538_v19).length));
          let _lhs199 = _1539_v20;
          let _lhs200 = _module.__default.safeIndex(new BigNumber(370), new BigNumber((_1539_v20).length));
          _1519_v0 = _rhs259;
          _lhs196.f7 = _rhs260;
          _lhs197[_lhs198] = _rhs261;
          _lhs199[_lhs200] = _rhs262;
          r0 = _rhs263;
          let _index261 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1538_v19).length));
          (_1538_v19)[_index261] = (_1538_v19)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_1538_v19).length))];
          let _index262 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1538_v19).length));
          (_1538_v19)[_index262] = (((_this).f17) ? ((_1537_v18).f22) : ((_1537_v18).f22));
          let _1544_v22;
          _1544_v22 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-561)), ((_1545_p0) => function (_1546_i2) {
            return _1545_p0;
          })(p0))).length),_dafny.Seq.Concat(_1519_v0, _1519_v0));
          _1544_v22 = (_1544_v22).update(new BigNumber(683), (_1537_v18).f23);
          (globalState).f7 = p3;
        } else {
          let _index263 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1538_v19).length));
          let _index264 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1538_v19).length));
          let _index265 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1538_v19).length));
          let _rhs264 = _module.__default.safeDivisionInt((_1537_v18).f22, (_1537_v18).f22);
          let _rhs265 = (_1537_v18).f22;
          let _rhs266 = (_1537_v18).f22;
          let _rhs267 = p1;
          let _lhs201 = _1538_v19;
          let _lhs202 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1538_v19).length));
          let _lhs203 = _1538_v19;
          let _lhs204 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1538_v19).length));
          let _lhs205 = _1538_v19;
          let _lhs206 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1538_v19).length));
          let _lhs207 = globalState;
          _lhs201[_lhs202] = _rhs264;
          _lhs203[_lhs204] = _rhs265;
          _lhs205[_lhs206] = _rhs266;
          _lhs207.f2 = _rhs267;
          let _1547_v23;
          let _nw301 = new _module.C4();
          _nw301.__ctor((_this).f17, _this.f16, p2);
          _1547_v23 = _nw301;
          let _1548_v24;
          _1548_v24 = _dafny.MultiSet.fromElements(((true) ? ((_1547_v23).f21) : (p2)), p2, _this.f16);
          let _rhs268 = ((!(_this.f16) || ((_this).f17)) ? (_1548_v24) : (((p2) ? (_1548_v24) : (_1548_v24))));
          let _rhs269 = p3;
          _1548_v24 = _rhs268;
          r0 = _rhs269;
          let _1549_v25;
          _1549_v25 = _dafny.Set.fromElements(new BigNumber(-158));
          let _1550_v26;
          let _nw302 = Array((new BigNumber(21)).toNumber()).fill(false);
          _1550_v26 = _nw302;
          let _1551_v27;
          _1551_v27 = _module.D9.create_DC15((_this).f17, (_this).f17, _1549_v25, _1538_v19, _1550_v26);
          let _1552_v28;
          let _nw303 = new _module.C1();
          _nw303.__ctor((_1551_v27).dtor_cf18, p1);
          _1552_v28 = _nw303;
          let _1553_v29;
          _1553_v29 = _module.D14.create_DC28((_1538_v19)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_1538_v19).length))], p2, _1552_v28);
          _1549_v25 = _dafny.Set.fromElements((_1537_v18).f22, (_1553_v29).dtor_cf40, new BigNumber(((_1537_v18).f23).length));
          let _1554_v30;
          _1554_v30 = _module.D3.create_DC3(p2);
          let _1555_v31;
          let _nw304 = new _module.C2();
          _nw304.__ctor(_1554_v30, (_1547_v23).f21, !(!((_this).f17)));
          _1555_v31 = _nw304;
        }
        let _index266 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1538_v19).length));
        (_1538_v19)[_index266] = (_1537_v18).f22;
        let _1556_v32;
        _1556_v32 = _dafny.Set.fromElements((_1538_v19)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_1538_v19).length))], (_1537_v18).f22);
        let _1557_v33;
        _1557_v33 = _dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber((_1556_v32).length));
        let _1558_v36;
        _1558_v36 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f16),function () {
          let _coll66 = new _dafny.Map();
          for (const _compr_66 of _dafny.IntegerRange(new BigNumber(8), new BigNumber(230))) {
            let _1559_v35 = _compr_66;
            if (((new BigNumber(8)).isLessThanOrEqualTo(_1559_v35)) && ((_1559_v35).isLessThan(new BigNumber(230)))) {
              _coll66.push([(_1559_v35).multipliedBy((_1537_v18).f22),new BigNumber(333)]);
            }
          }
          return _coll66;
        }());
        let _1560_v37;
        _1560_v37 = _dafny.Seq.of(p1);
        let _1561_v38;
        _1561_v38 = _module.D11.create_DC19(_1557_v33);
        let _1562_v39;
        let _nw305 = Array((new BigNumber(28)).toNumber());
        _nw305[(_dafny.ZERO).toNumber()] = (_1557_v33).Merge(_1557_v33);
        _nw305[(_dafny.ONE).toNumber()] = (function () {
          let _coll67 = new _dafny.Map();
          for (const _compr_67 of _dafny.IntegerRange(new BigNumber(736), new BigNumber(62))) {
            let _1563_v34 = _compr_67;
            if (((new BigNumber(736)).isLessThanOrEqualTo(_1563_v34)) && ((_1563_v34).isLessThan(new BigNumber(62)))) {
              _coll67.push([(_1563_v34).plus((_1538_v19)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_1538_v19).length))]),new BigNumber(((_1537_v18).f23).length)]);
            }
          }
          return _coll67;
        }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(77),p3));
        _nw305[(new BigNumber(2)).toNumber()] = (_1557_v33).Merge((((_1558_v36).contains(_this.f16)) ? ((_1558_v36).get(_this.f16)) : (_1557_v33)));
        _nw305[(new BigNumber(3)).toNumber()] = _1557_v33;
        _nw305[(new BigNumber(4)).toNumber()] = (_1557_v33).Merge(_1557_v33);
        _nw305[(new BigNumber(5)).toNumber()] = _1557_v33;
        _nw305[(new BigNumber(6)).toNumber()] = _1557_v33;
        _nw305[(new BigNumber(7)).toNumber()] = _1557_v33;
        _nw305[(new BigNumber(8)).toNumber()] = _1557_v33;
        _nw305[(new BigNumber(9)).toNumber()] = _1557_v33;
        _nw305[(new BigNumber(10)).toNumber()] = _1557_v33;
        _nw305[(new BigNumber(11)).toNumber()] = (_1557_v33).Merge(_dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber(871)));
        _nw305[(new BigNumber(12)).toNumber()] = _1557_v33;
        _nw305[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_1538_v19)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_1538_v19).length))],(_dafny.ZERO).minus((_1537_v18).f22));
        _nw305[(new BigNumber(14)).toNumber()] = (_1557_v33).Merge(_1557_v33);
        _nw305[(new BigNumber(15)).toNumber()] = _1557_v33;
        _nw305[(new BigNumber(16)).toNumber()] = (_1557_v33).Merge(_1557_v33);
        _nw305[(new BigNumber(17)).toNumber()] = (_1557_v33).Merge(_1557_v33);
        _nw305[(new BigNumber(18)).toNumber()] = _1557_v33;
        _nw305[(new BigNumber(19)).toNumber()] = _1557_v33;
        _nw305[(new BigNumber(20)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(p1,(_1538_v19)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_1538_v19).length))])).update(new BigNumber((_1560_v37).length), _module.__default.fm0(_this.f16, globalState));
        _nw305[(new BigNumber(21)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p1),new BigNumber(((_1537_v18).f23).length));
        _nw305[(new BigNumber(22)).toNumber()] = _1557_v33;
        _nw305[(new BigNumber(23)).toNumber()] = (_1557_v33).Merge((_1561_v38).dtor_cf27);
        _nw305[(new BigNumber(24)).toNumber()] = ((_1557_v33).update(_module.__default.fm0(false, globalState), (_1538_v19)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_1538_v19).length))])).Merge(_1557_v33);
        _nw305[(new BigNumber(25)).toNumber()] = (_1557_v33).update((_dafny.ZERO).minus(new BigNumber((_1557_v33).length)), new BigNumber(-433));
        _nw305[(new BigNumber(26)).toNumber()] = _1557_v33;
        _nw305[(new BigNumber(27)).toNumber()] = (_1557_v33).Merge(_1557_v33);
        _1562_v39 = _nw305;
        let _index267 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_1562_v39).length));
        (_1562_v39)[_index267] = ((_1557_v33).update(p1, (_1538_v19)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_1538_v19).length))])).Merge(_1557_v33);
        let _1564_v40;
        _1564_v40 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(744),false);
        let _1565_v41;
        _1565_v41 = _dafny.Seq.of(_1519_v0);
        let _1566_v42;
        _1566_v42 = _dafny.Seq.of((_1565_v41)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_this.f16, true, (_this).f17, (_this).f17)).length), new BigNumber((_1565_v41).length))]);
        let _index268 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_1562_v39).length));
        let _rhs270 = _this.f16;
        let _rhs271 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1564_v40).length),p1);
        let _rhs272 = ((_this.f16) ? (p2) : (!((_this).f17)));
        let _rhs273 = _dafny.areEqual(_1566_v42, _1566_v42);
        let _rhs274 = ((_1537_v18).f22).minus(p1);
        let _lhs208 = _this;
        let _lhs209 = _1562_v39;
        let _lhs210 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_1562_v39).length));
        let _lhs211 = _this;
        let _lhs212 = _this;
        let _lhs213 = globalState;
        _lhs208.f16 = _rhs270;
        _lhs209[_lhs210] = _rhs271;
        _lhs211.f16 = _rhs272;
        _lhs212.f16 = _rhs273;
        _lhs213.f2 = _rhs274;
      }
      let _1567_v43;
      _1567_v43 = _dafny.Seq.of(_this.f16);
      let _1568_v44;
      _1568_v44 = _dafny.Set.fromElements(false);
      let _1569_v45;
      let _nw306 = Array((new BigNumber(18)).toNumber());
      _nw306[(_dafny.ZERO).toNumber()] = (_this).fm23(false, globalState);
      _nw306[(_dafny.ONE).toNumber()] = p2;
      _nw306[(new BigNumber(2)).toNumber()] = !(p2);
      _nw306[(new BigNumber(3)).toNumber()] = false;
      _nw306[(new BigNumber(4)).toNumber()] = p2;
      _nw306[(new BigNumber(5)).toNumber()] = !(!((p1).isLessThan(p1)));
      _nw306[(new BigNumber(6)).toNumber()] = p2;
      _nw306[(new BigNumber(7)).toNumber()] = false;
      _nw306[(new BigNumber(8)).toNumber()] = p2;
      _nw306[(new BigNumber(9)).toNumber()] = !(p2);
      _nw306[(new BigNumber(10)).toNumber()] = _this.f16;
      _nw306[(new BigNumber(11)).toNumber()] = (_1567_v43)[_module.__default.safeIndex(p1, new BigNumber((_1567_v43).length))];
      _nw306[(new BigNumber(12)).toNumber()] = (_1567_v43)[_module.__default.safeIndex(p1, new BigNumber((_1567_v43).length))];
      _nw306[(new BigNumber(13)).toNumber()] = (((_this).f17) ? ((_this).f17) : (p2));
      _nw306[(new BigNumber(14)).toNumber()] = (false) && (false);
      _nw306[(new BigNumber(15)).toNumber()] = (_1568_v44).IsDisjointFrom(_1568_v44);
      _nw306[(new BigNumber(16)).toNumber()] = p2;
      _nw306[(new BigNumber(17)).toNumber()] = (_this).f17;
      _1569_v45 = _nw306;
      for (const _guard_loop_8 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1569_v45).length))) {
        let _1570_i3 = _guard_loop_8;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1570_i3)) && ((_1570_i3).isLessThan(new BigNumber((_1569_v45).length))))) {
          (_1569_v45)[(_1570_i3)] = true;
        }
      }
      let _1571_v46;
      _1571_v46 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(741));
      let _1572_v47;
      _1572_v47 = _module.D3.create_DC4(_this.f16, p1, p1, !(_this.f16), (_this).f17);
      let _1573_v48;
      _1573_v48 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).fm23(_this.f16, globalState));
      let _1574_v49;
      _1574_v49 = _dafny.Seq.of(_1573_v48);
      let _1575_v50;
      _1575_v50 = _dafny.Set.fromElements(p1, p3, new BigNumber((_module.__default.fm28((_this).f17, _1571_v46, _1572_v47, (_1574_v49)[_module.__default.safeIndex(p1, new BigNumber((_1574_v49).length))], globalState)).length));
      let _1576_v51;
      _1576_v51 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1575_v50).length),(_this).f17);
      let _1577_v53;
      _1577_v53 = _module.D7.create_DC11((_1576_v51).Merge(function () {
  let _coll68 = new _dafny.Map();
  for (const _compr_68 of _dafny.IntegerRange(new BigNumber(-134), new BigNumber(203))) {
    let _1578_v52 = _compr_68;
    if (((new BigNumber(-134)).isLessThanOrEqualTo(_1578_v52)) && ((_1578_v52).isLessThan(new BigNumber(203)))) {
      _coll68.push([(_1578_v52).plus(new BigNumber(-376)),(_this).f17]);
    }
  }
  return _coll68;
}()));
      let _source25 = _1577_v53;
      if (_source25.is_DC12) {
        (globalState).f2 = p3;
        let _1579_v54;
        _1579_v54 = new _dafny.CodePoint('d'.codePointAt(0));
        let _1580_v55;
        _1580_v55 = _dafny.Seq.of(_module.__default.fm42((_this).f17, new BigNumber((_1575_v50).length), p1, (_this).f17, globalState));
        let _1581_v56;
        let _nw307 = Array((new BigNumber(15)).toNumber()).fill([]);
        _1581_v56 = _nw307;
        let _1582_v57;
        let _init46 = function (_1583_i4) {
          return (_1583_i4).multipliedBy((new BigNumber((_dafny.MultiSet.fromElements((_this).f17, (_this).f17)).cardinality())));
        };
        let _nw308 = Array((new BigNumber(28)).toNumber());
        for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw308.length); _i0_46++) {
          _nw308[_i0_46] = _init46(new BigNumber(_i0_46));
        }
        _1582_v57 = _nw308;
        let _index269 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_1581_v56).length));
        (_1581_v56)[_index269] = _1582_v57;
        let _1584_v58;
        let _init47 = function (_1585_i5) {
          return _dafny.Seq.of(_this.f16);
        };
        let _nw309 = Array((new BigNumber(8)).toNumber());
        for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw309.length); _i0_47++) {
          _nw309[_i0_47] = _init47(new BigNumber(_i0_47));
        }
        _1584_v58 = _nw309;
        let _1586_v59;
        _1586_v59 = _dafny.Map.Empty.slice().updateUnsafe(_1584_v58,p3);
        let _index270 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_1581_v56).length));
        let _rhs275 = p2;
        let _rhs276 = new _dafny.CodePoint('y'.codePointAt(0));
        let _rhs277 = _1580_v55;
        let _rhs278 = _1582_v57;
        let _rhs279 = ((_1586_v59).update(_1584_v58, p1)).Merge((_1586_v59).Merge(_1586_v59));
        let _lhs214 = _this;
        let _lhs215 = _1581_v56;
        let _lhs216 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_1581_v56).length));
        _lhs214.f16 = _rhs275;
        _1579_v54 = _rhs276;
        _1580_v55 = _rhs277;
        _lhs215[_lhs216] = _rhs278;
        _1586_v59 = _rhs279;
        r0 = p3;
        (_this).f16 = (_this).f17;
      } else {
        let _1587___mcc_h0 = (_source25).cf15;
        let _1588_cf15 = _1587___mcc_h0;
        (_this).f16 = !((_this).f17) || ((_this).f17);
        (globalState).f2 = p3;
        let _1589_v60;
        let _nw310 = Array((new BigNumber(6)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1589_v60 = _nw310;
        let _index271 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_1589_v60).length));
        (_1589_v60)[_index271] = p0;
        let _1590_v61;
        let _init48 = ((_1591_p1) => function (_1592_i6) {
          return (_1592_i6).minus(_1591_p1);
        })(p1);
        let _nw311 = Array((new BigNumber(29)).toNumber());
        for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw311.length); _i0_48++) {
          _nw311[_i0_48] = _init48(new BigNumber(_i0_48));
        }
        _1590_v61 = _nw311;
        let _1593_v62;
        _1593_v62 = _module.D9.create_DC15((_this).f17, p2, _1575_v50, _1590_v61, _1569_v45);
        let _1594_v63;
        let _nw312 = Array((new BigNumber(22)).toNumber());
        _nw312[(_dafny.ZERO).toNumber()] = _module.D9.create_DC15(p2, p2, _dafny.Set.fromElements(p3), _1590_v61, _1569_v45);
        _nw312[(_dafny.ONE).toNumber()] = _1593_v62;
        _nw312[(new BigNumber(2)).toNumber()] = _1593_v62;
        _nw312[(new BigNumber(3)).toNumber()] = _1593_v62;
        _nw312[(new BigNumber(4)).toNumber()] = _1593_v62;
        _nw312[(new BigNumber(5)).toNumber()] = _module.D9.create_DC15(p2, _this.f16, _1575_v50, _1590_v61, _1569_v45);
        _nw312[(new BigNumber(6)).toNumber()] = _1593_v62;
        _nw312[(new BigNumber(7)).toNumber()] = _1593_v62;
        _nw312[(new BigNumber(8)).toNumber()] = _1593_v62;
        _nw312[(new BigNumber(9)).toNumber()] = _1593_v62;
        _nw312[(new BigNumber(10)).toNumber()] = _module.D9.create_DC15(p2, false, _1575_v50, _1590_v61, _1569_v45);
        _nw312[(new BigNumber(11)).toNumber()] = _1593_v62;
        _nw312[(new BigNumber(12)).toNumber()] = _1593_v62;
        _nw312[(new BigNumber(13)).toNumber()] = _1593_v62;
        _nw312[(new BigNumber(14)).toNumber()] = _1593_v62;
        _nw312[(new BigNumber(15)).toNumber()] = _1593_v62;
        _nw312[(new BigNumber(16)).toNumber()] = _1593_v62;
        _nw312[(new BigNumber(17)).toNumber()] = _1593_v62;
        _nw312[(new BigNumber(18)).toNumber()] = _1593_v62;
        _nw312[(new BigNumber(19)).toNumber()] = _1593_v62;
        _nw312[(new BigNumber(20)).toNumber()] = _1593_v62;
        _nw312[(new BigNumber(21)).toNumber()] = _1593_v62;
        _1594_v63 = _nw312;
        let _1595_v64;
        _1595_v64 = _module.D18.create_DC37(_1594_v63);
        let _1596_v65;
        _1596_v65 = _dafny.Seq.of((_1595_v64).dtor_cf56);
        let _1597_v66;
        _1597_v66 = _module.D12.create_DC24(_dafny.Seq.IsProperPrefixOf(_1567_v43, _1567_v43), _1567_v43, _1596_v65, ((_this.f16) ? (p2) : (p2)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(445)), ((_1598_p2) => function (_1599_i7) {
  return _module.D3.create_DC3(_1598_p2);
})(p2)));
        let _index272 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_1589_v60).length));
        let _rhs280 = (_1519_v0)[_module.__default.safeIndex(p3, new BigNumber((_1519_v0).length))];
        let _rhs281 = _1597_v66;
        let _lhs217 = _1589_v60;
        let _lhs218 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_1589_v60).length));
        _lhs217[_lhs218] = _rhs280;
        _1597_v66 = _rhs281;
        let _1600_v67;
        _1600_v67 = _module.D3.create_DC3((_this).fm23(_this.f16, globalState));
        let _1601_v68;
        let _nw313 = new _module.C2();
        _nw313.__ctor(_1600_v67, _this.f16, (_module.__default.fm49(p3, p3, p2, p1, globalState)).IsProperSubsetOf(_1575_v50));
        _1601_v68 = _nw313;
      }
      let _hi10 = p3;
      for (let _1602_i8 = p3; _1602_i8.isLessThan(_hi10); _1602_i8 = _1602_i8.plus(_dafny.ONE)) {
        let _1603_v69;
        let _init49 = ((_1604_v43) => function (_1605_i9) {
          return _dafny.Seq.Concat(_1604_v43, _1604_v43);
        })(_1567_v43);
        let _nw314 = Array((new BigNumber(16)).toNumber());
        for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw314.length); _i0_49++) {
          _nw314[_i0_49] = _init49(new BigNumber(_i0_49));
        }
        _1603_v69 = _nw314;
        let _index273 = _module.__default.safeIndex(new BigNumber(835), new BigNumber((_1603_v69).length));
        (_1603_v69)[_index273] = _dafny.Seq.of(_this.f16, p2);
        let _1606_v70;
        let _nw315 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
        _1606_v70 = _nw315;
        let _1607_v71;
        _1607_v71 = _module.D9.create_DC15((_this).f17, p2, _1575_v50, _1606_v70, _1569_v45);
        let _pat_let_tv22 = _1575_v50;
        let _1608_v72;
        let _nw316 = Array((new BigNumber(23)).toNumber());
        _nw316[(_dafny.ZERO).toNumber()] = _1607_v71;
        _nw316[(_dafny.ONE).toNumber()] = _1607_v71;
        _nw316[(new BigNumber(2)).toNumber()] = _module.D9.create_DC15((_this).f17, (_this).f17, _1575_v50, _1606_v70, _1569_v45);
        _nw316[(new BigNumber(3)).toNumber()] = _1607_v71;
        _nw316[(new BigNumber(4)).toNumber()] = _1607_v71;
        _nw316[(new BigNumber(5)).toNumber()] = _1607_v71;
        _nw316[(new BigNumber(6)).toNumber()] = _1607_v71;
        _nw316[(new BigNumber(7)).toNumber()] = _1607_v71;
        _nw316[(new BigNumber(8)).toNumber()] = _1607_v71;
        _nw316[(new BigNumber(9)).toNumber()] = _1607_v71;
        _nw316[(new BigNumber(10)).toNumber()] = function (_pat_let25_0) {
          return function (_1609_dt__update__tmp_h0) {
            return function (_pat_let26_0) {
              return function (_1610_dt__update_hcf19_h0) {
                return function (_pat_let27_0) {
                  return function (_1611_dt__update_hcf20_h0) {
                    return _module.D9.create_DC15((_1609_dt__update__tmp_h0).dtor_cf18, _1610_dt__update_hcf19_h0, _1611_dt__update_hcf20_h0, (_1609_dt__update__tmp_h0).dtor_cf21, (_1609_dt__update__tmp_h0).dtor_cf22);
                  }(_pat_let27_0);
                }(_pat_let_tv22);
              }(_pat_let26_0);
            }((_this).f17);
          }(_pat_let25_0);
        }(_1607_v71);
        _nw316[(new BigNumber(11)).toNumber()] = _1607_v71;
        _nw316[(new BigNumber(12)).toNumber()] = _1607_v71;
        _nw316[(new BigNumber(13)).toNumber()] = _1607_v71;
        _nw316[(new BigNumber(14)).toNumber()] = _module.D9.create_DC15(p2, (_this).f17, _1575_v50, _1606_v70, _1569_v45);
        _nw316[(new BigNumber(15)).toNumber()] = _1607_v71;
        _nw316[(new BigNumber(16)).toNumber()] = _1607_v71;
        _nw316[(new BigNumber(17)).toNumber()] = _1607_v71;
        _nw316[(new BigNumber(18)).toNumber()] = _1607_v71;
        _nw316[(new BigNumber(19)).toNumber()] = _1607_v71;
        _nw316[(new BigNumber(20)).toNumber()] = _module.D9.create_DC15(p2, (_this).f17, _1575_v50, _1606_v70, _1569_v45);
        _nw316[(new BigNumber(21)).toNumber()] = _1607_v71;
        _nw316[(new BigNumber(22)).toNumber()] = _1607_v71;
        _1608_v72 = _nw316;
        let _1612_v73;
        _1612_v73 = _dafny.Seq.of(_1608_v72);
        let _1613_v74;
        _1613_v74 = _module.D3.create_DC3(true);
        let _1614_v75;
        _1614_v75 = _dafny.Seq.of(_1613_v74);
        let _1615_v76;
        _1615_v76 = _module.D12.create_DC24(false, _1567_v43, _1612_v73, p2, _1614_v75);
        let _index274 = _module.__default.safeIndex(new BigNumber(835), new BigNumber((_1603_v69).length));
        (_1603_v69)[_index274] = ((((_this).f17) ? (_1615_v76) : (_module.D12.create_DC24(_this.f16, _1567_v43, _1612_v73, (_this).f17, _1614_v75)))).dtor_cf34;
        let _1616_v77;
        let _nw317 = new _module.C1();
        _nw317.__ctor(_this.f16, _1602_i8);
        _1616_v77 = _nw317;
        (globalState).f2 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(p2), _dafny.Seq.of(!(true)))).length);
        let _1617_v78;
        _1617_v78 = _dafny.Seq.of(p1);
        let _1618_v79;
        _1618_v79 = (_1603_v69)[_module.__default.safeIndex(new BigNumber(835), new BigNumber((_1603_v69).length))];
        let _1619_v80;
        _1619_v80 = _module.D4.create_DC7(_1618_v79, p3);
        let _1620_v81;
        _1620_v81 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_1617_v78).length)),_1619_v80);
        let _1621_v82;
        _1621_v82 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber(((_module.__default.fm58(false, p2, new BigNumber(498), globalState)).Merge(_1620_v81)).length));
        _1621_v82 = (_1621_v82).update((_this).f17, p1);
      }
      let _1622_v83;
      let _nw318 = new _module.C1();
      _nw318.__ctor((p3).isLessThan(p1), p1);
      _1622_v83 = _nw318;
      let _hi11 = p3;
      for (let _1623_i10 = p1; _1623_i10.isLessThan(_hi11); _1623_i10 = _1623_i10.plus(_dafny.ONE)) {
        let _1624_v84;
        _1624_v84 = _dafny.MultiSet.fromElements(p0);
        let _1625_v85;
        let _init50 = function (_1626_i11) {
          return _module.__default.safeModuloInt(_1626_i11, new BigNumber(-761));
        };
        let _nw319 = Array((new BigNumber(10)).toNumber());
        for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw319.length); _i0_50++) {
          _nw319[_i0_50] = _init50(new BigNumber(_i0_50));
        }
        _1625_v85 = _nw319;
        let _1627_v86;
        _1627_v86 = _module.D9.create_DC15(_module.__default.fm27(new BigNumber((_1624_v84).cardinality()), _1623_i10, p3, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-678),new BigNumber(9)), globalState), p2, _1575_v50, _1625_v85, _1569_v45);
        let _source26 = _1627_v86;
        if (_source26.is_DC15) {
          let _1628___mcc_h1 = (_source26).cf18;
          let _1629___mcc_h2 = (_source26).cf19;
          let _1630___mcc_h3 = (_source26).cf20;
          let _1631___mcc_h4 = (_source26).cf21;
          let _1632___mcc_h5 = (_source26).cf22;
          let _1633_cf22 = _1632___mcc_h5;
          let _1634_cf21 = _1631___mcc_h4;
          let _1635_cf20 = _1630___mcc_h3;
          let _1636_cf19 = _1629___mcc_h2;
          let _1637_cf18 = _1628___mcc_h1;
          let _1638_v87;
          _1638_v87 = _module.D3.create_DC4(_1622_v83.f25, (_dafny.ZERO).minus(_1623_i10), _1623_i10, !(_module.__default.fm31(_1622_v83.f25, true, globalState)), _1637_cf18);
          let _1639_v88;
          _1639_v88 = _module.D3.create_DC5(_1638_v87);
          let _1640_v89;
          _1640_v89 = _module.D3.create_DC5(_1639_v88);
          let _1641_v90;
          _1641_v90 = _dafny.Map.Empty.slice().updateUnsafe(_1640_v89,!(_1637_cf18));
          let _1642_v91;
          _1642_v91 = _module.D3.create_DC4(p2, p1, new BigNumber(965), _this.f16, (((_1641_v90).contains(_1640_v89)) ? ((_1641_v90).get(_1640_v89)) : (_1636_cf19)));
          let _1643_v92;
          _1643_v92 = _module.D3.create_DC5(_1639_v88);
          _1640_v89 = _1640_v89;
          let _1644_v93;
          _1644_v93 = _1567_v43;
          let _1645_v94;
          _1645_v94 = _module.D4.create_DC7(_1644_v93, _module.__default.safeDivisionInt(p1, p1));
          _1645_v94 = _1645_v94;
          (globalState).f7 = (new BigNumber(-594)).minus((((_1571_v46).contains(p3)) ? ((_1571_v46).get(p3)) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_1622_v83.f25)).length))));
          (globalState).f7 = (_this).fm12(new BigNumber((_1568_v44).length), globalState);
        } else if (_source26.is_DC16) {
          let _1646___mcc_h6 = (_source26).cf23;
          let _1647___mcc_h7 = (_source26).cf24;
          let _1648_cf24 = _1647___mcc_h7;
          let _1649_cf23 = _1646___mcc_h6;
          let _1650_v95;
          let _nw320 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1650_v95 = _nw320;
          _1650_v95 = _1650_v95;
          let _index275 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_1625_v85).length));
          (_1625_v85)[_index275] = new BigNumber(974);
          let _1651_v96;
          _1651_v96 = _dafny.Seq.of(new BigNumber((_1519_v0).length));
          let _1652_v97;
          let _nw321 = new _module.C1();
          _nw321.__ctor(_this.f16, new BigNumber((_1651_v96).length));
          _1652_v97 = _nw321;
          let _index276 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_1625_v85).length));
          (_1625_v85)[_index276] = ((p2) ? (_module.__default.safeDivisionInt(p1, new BigNumber((_1519_v0).length))) : (_module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm31(p2, p2, globalState),_1652_v97)).length), p1)));
          let _1653_v98;
          _1653_v98 = new _dafny.CodePoint('p'.codePointAt(0));
          _1653_v98 = p0;
          (globalState).f2 = (p1).plus(new BigNumber((_1576_v51).length));
        } else if (_source26.is_DC14) {
          let _1654___mcc_h8 = (_source26).cf17;
          let _1655_cf17 = _1654___mcc_h8;
          let _1656_v99;
          let _init51 = ((_1657_p3) => function (_1658_i12) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(425)), ((_1659_p3) => function (_1660_i13) {
              return _1659_p3;
            })(_1657_p3));
          })(p3);
          let _nw322 = Array((new BigNumber(23)).toNumber());
          for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw322.length); _i0_51++) {
            _nw322[_i0_51] = _init51(new BigNumber(_i0_51));
          }
          _1656_v99 = _nw322;
          let _1661_v100;
          _1661_v100 = _dafny.Map.Empty.slice().updateUnsafe(_1656_v99,(p1).plus(p1));
          let _1662_v101;
          _1662_v101 = _dafny.MultiSet.fromElements(_1519_v0);
          _1661_v100 = (_1661_v100).update(_1656_v99, (((_1662_v101).contains(_1519_v0)) ? ((_1662_v101).get(_1519_v0)) : (new BigNumber((_1576_v51).length))));
          let _1663_v102;
          _1663_v102 = _module.D9.create_DC16(p3, p2);
          r0 = _module.__default.safeDivisionInt(p1, (_dafny.ZERO).minus((function (_pat_let28_0) {
            return function (_1664_dt__update__tmp_h1) {
              return function (_pat_let29_0) {
                return function (_1665_dt__update_hcf24_h0) {
                  return _module.D9.create_DC16((_1664_dt__update__tmp_h1).dtor_cf23, _1665_dt__update_hcf24_h0);
                }(_pat_let29_0);
              }((_this).f17);
            }(_pat_let28_0);
          }(_1663_v102)).dtor_cf23));
          let _rhs282 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(524)), function (_1666_i14) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          });
          let _rhs283 = (_1623_i10).minus((p1).plus(p3));
          let _lhs219 = globalState;
          _1519_v0 = _rhs282;
          _lhs219.f7 = _rhs283;
          (_1622_v83).f25 = false;
        } else {
          let _1667___mcc_h9 = (_source26).cf25;
          let _1668_cf25 = _1667___mcc_h9;
          let _rhs284 = p1;
          let _rhs285 = (_module.__default.fm40(_1622_v83.f25, globalState)).Merge((_1573_v48).Merge(_1573_v48));
          let _rhs286 = ((!(!(p2))) ? (_1569_v45) : (_1569_v45));
          let _lhs220 = globalState;
          _lhs220.f7 = _rhs284;
          _1573_v48 = _rhs285;
          _1569_v45 = _rhs286;
          let _1669_v103;
          _1669_v103 = _dafny.MultiSet.fromElements(_1623_i10);
          let _1670_v104;
          _1670_v104 = _module.D19.create_DC39(_1669_v103);
          let _1671_v105;
          _1671_v105 = _module.D19.create_DC39((_1670_v104).dtor_cf60);
          (globalState).f7 = new BigNumber((((_1669_v103).Union((_1671_v105).dtor_cf60)).Union((_1669_v103).Union(_1669_v103))).cardinality());
          let _index277 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_1625_v85).length));
          (_1625_v85)[_index277] = p1;
          let _1672_v106;
          let _nw323 = Array((new BigNumber(12)).toNumber()).fill(_module.D9.Default());
          _1672_v106 = _nw323;
          let _1673_v107;
          _1673_v107 = _dafny.Seq.of(_1672_v106, _1672_v106, _1672_v106);
          let _1674_v108;
          _1674_v108 = _module.D3.create_DC3(_1622_v83.f25);
          let _1675_v109;
          _1675_v109 = _dafny.Seq.of(_1674_v108, _1674_v108, _module.D3.create_DC3(!(p2)), _module.D3.create_DC3(true), _1674_v108);
          let _1676_v110;
          _1676_v110 = _module.D12.create_DC24((_this).f17, _1567_v43, _1673_v107, _1622_v83.f25, _1675_v109);
          let _1677_v111;
          _1677_v111 = _dafny.MultiSet.fromElements(_1676_v110, _1676_v110, _1676_v110, _1676_v110, _1676_v110);
          let _1678_v112;
          _1678_v112 = _module.D20.create_DC41((_dafny.MultiSet.fromElements(_1676_v110, _1676_v110)).update(_1676_v110, _module.__default.abs(new BigNumber((_1519_v0).length))));
          let _1679_v113;
          _1679_v113 = _dafny.MultiSet.fromElements(_1622_v83.f25, _1622_v83.f25);
          let _1680_v114;
          _1680_v114 = _dafny.Map.Empty.slice().updateUnsafe(_1679_v113,_this.f16);
          let _index278 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_1625_v85).length));
          let _rhs287 = !(!((_1567_v43)[_module.__default.safeIndex(p1, new BigNumber((_1567_v43).length))]));
          let _rhs288 = ((_1677_v111).Difference((_1678_v112).dtor_cf64)).IsProperSubsetOf(_dafny.MultiSet.fromElements(_1676_v110, _module.D12.create_DC24(p2, _1567_v43, _1673_v107, _1622_v83.f25, _1675_v109)));
          let _rhs289 = _module.__default.safeModuloInt(p3, new BigNumber(((_1680_v114).Merge(_1680_v114)).length));
          let _lhs221 = _1622_v83;
          let _lhs222 = _1622_v83;
          let _lhs223 = _1625_v85;
          let _lhs224 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_1625_v85).length));
          _lhs221.f25 = _rhs287;
          _lhs222.f25 = _rhs288;
          _lhs223[_lhs224] = _rhs289;
          let _1681_v115;
          _1681_v115 = _module.D4.create_DC6(_1573_v48);
          let _1682_v116;
          let _nw324 = new _module.C1();
          _nw324.__ctor(p2, (_1625_v85)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_1625_v85).length))]);
          _1682_v116 = _nw324;
          let _1683_v117;
          _1683_v117 = _module.D16.create_DC34((_this).f17, _1682_v116, p1);
          let _1684_v118;
          _1684_v118 = _module.D16.create_DC34(p2, (_1683_v117).dtor_cf52, _1682_v116.f14);
          let _rhs290 = p1;
          let _rhs291 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(new BigNumber((_1573_v48).length), new BigNumber(((_1681_v115).dtor_cf10).length)), (_1683_v117).dtor_cf53);
          let _rhs292 = p2;
          let _rhs293 = !(_1622_v83.f25);
          let _rhs294 = _1682_v116.f14;
          let _lhs225 = globalState;
          let _lhs226 = _this;
          let _lhs227 = _1622_v83;
          let _lhs228 = globalState;
          _lhs225.f7 = _rhs290;
          r0 = _rhs291;
          _lhs226.f16 = _rhs292;
          _lhs227.f25 = _rhs293;
          _lhs228.f2 = _rhs294;
        }
        let _1685_v119;
        _1685_v119 = _dafny.Map.Empty.slice().updateUnsafe(_1624_v84,_1569_v45);
        let _1686_v120;
        let _nw325 = Array((new BigNumber(4)).toNumber());
        _nw325[(_dafny.ZERO).toNumber()] = _1569_v45;
        _nw325[(_dafny.ONE).toNumber()] = _1569_v45;
        _nw325[(new BigNumber(2)).toNumber()] = _1569_v45;
        _nw325[(new BigNumber(3)).toNumber()] = (((_1685_v119).contains(_1624_v84)) ? ((_1685_v119).get(_1624_v84)) : (_1569_v45));
        _1686_v120 = _nw325;
        _1686_v120 = _1686_v120;
        let _1687_v121;
        let _nw326 = new _module.C4();
        _nw326.__ctor(_1622_v83.f25, _this.f16, _1622_v83.f25);
        _1687_v121 = _nw326;
        let _index279 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_1569_v45).length));
        (_1569_v45)[_index279] = _this.f16;
        let _index280 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_1569_v45).length));
        (_1569_v45)[_index280] = _1622_v83.f25;
      }
      r0 = new BigNumber((_dafny.Seq.Concat(_1567_v43, _dafny.Seq.of((_this).f17))).length);
      return r0;
    }
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f16 = false;
      this._f14 = _dafny.ZERO;
      this._f17 = false;
      this._f19 = _dafny.ZERO;
      this._f20 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0, _module.T2];
    }
    get f16() {
      let _this = this;
      return _this._f16;
    };
    set f16(value) {
      let _this = this;
      _this._f16 = value;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    set f14(value) {
      let _this = this;
      _this._f14 = value;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    __ctor(f19, f20, f16, f17, f14) {
      let _this = this;
      (_this)._f19 = f19;
      (_this)._f20 = f20;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f14 = f14;
      return;
    }
    fm3(globalState) {
      let _this = this;
      return function () {
        let _coll69 = new _dafny.Map();
        for (const _compr_69 of _dafny.IntegerRange(new BigNumber(425), new BigNumber(-562))) {
          let _1688_v0 = _compr_69;
          if (((new BigNumber(425)).isLessThanOrEqualTo(_1688_v0)) && ((_1688_v0).isLessThan(new BigNumber(-562)))) {
            _coll69.push([(_1688_v0).minus((_this).f19),(_this).f19]);
          }
        }
        return _coll69;
      }();
    };
    fm12(p0, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f19, (_this).f19));
    };
    fm13(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.of(_this.f16, (_this).f20), _dafny.Seq.Concat(_dafny.Seq.of(_this.f16), _dafny.Seq.of((_this).f20, (_this).f20, _this.f16)));
    };
    fm18(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _source27 = _module.D4.create_DC6(_dafny.Map.Empty.slice().updateUnsafe(_this.f16,(_this).f20));
      if (_source27.is_DC7) {
        let _1689___mcc_h0 = (_source27).cf11;
        let _1690___mcc_h1 = (_source27).cf12;
        let _1691_cf12 = _1690___mcc_h1;
        let _1692_cf11 = _1689___mcc_h0;
        return (_1691_cf12).plus(_this.f14);
      } else {
        let _1693___mcc_h2 = (_source27).cf10;
        let _1694_cf10 = _1693___mcc_h2;
        return (_this).f19;
      }
    };
    fm19(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Seq.of((_dafny.ZERO).minus(_this.f14), _this.f14);
    };
    m2(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _hi12 = ((_this).f19).plus((_this).f19);
      for (let _1695_i0 = (_this).f19; _1695_i0.isLessThan(_hi12); _1695_i0 = _1695_i0.plus(_dafny.ONE)) {
        let _1696_v0;
        _1696_v0 = _dafny.Seq.UnicodeFromString("usc");
        _1696_v0 = _dafny.Seq.Concat(_1696_v0, ((_this.f16) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(425)), function (_1697_i1) {
          return new _dafny.CodePoint('a'.codePointAt(0));
        })) : (_1696_v0)));
        let _1698_v1;
        let _nw327 = Array((new BigNumber(26)).toNumber()).fill([]);
        _1698_v1 = _nw327;
        _1698_v1 = _1698_v1;
        (globalState).f2 = new BigNumber(495);
        let _1699_v2;
        let _nw328 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
        _1699_v2 = _nw328;
        let _index281 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_1699_v2).length));
        (_1699_v2)[_index281] = p0;
        let _1700_v3;
        _1700_v3 = _module.D3.create_DC3(p1);
        let _1701_v4;
        _1701_v4 = _dafny.Set.fromElements((_1700_v3).dtor_cf3, !((_this).f17));
        let _1702_v5;
        _1702_v5 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_1701_v4).length)), new BigNumber((_1701_v4).length));
        let _1703_v6;
        _1703_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1702_v5).length),_this.f14);
        let _1704_v7;
        _1704_v7 = _dafny.MultiSet.fromElements(_1703_v6, _1703_v6, (_1703_v6).Merge(_1703_v6));
        let _1705_v8;
        _1705_v8 = new _dafny.CodePoint('m'.codePointAt(0));
        let _1706_v9;
        _1706_v9 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true),_1705_v8);
        let _1707_v10;
        _1707_v10 = _dafny.Seq.of((_this).f20);
        let _1708_v11;
        _1708_v11 = _dafny.Seq.of(_1707_v10, _dafny.Seq.of(p1));
        let _index282 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_1699_v2).length));
        let _rhs295 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((new BigNumber((_1701_v4).length)).plus(p0)), _module.__default.safeModuloInt(p0, new BigNumber((_1706_v9).length)));
        let _rhs296 = !((_module.D3.create_DC4(_this.f16, (_this).f19, p0, (_this).f20, false)).dtor_cf8);
        let _rhs297 = _module.__default.fm20(_1695_i0, new BigNumber(((_1708_v11)[_module.__default.safeIndex(_this.f14, new BigNumber((_1708_v11).length))]).length), globalState);
        let _rhs298 = false;
        let _lhs229 = _1699_v2;
        let _lhs230 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_1699_v2).length));
        let _lhs231 = _this;
        let _lhs232 = _this;
        _lhs229[_lhs230] = _rhs295;
        _lhs231.f16 = _rhs296;
        _1704_v7 = _rhs297;
        _lhs232.f16 = _rhs298;
      }
      let _1709_i2;
      _1709_i2 = _dafny.ZERO;
      L7: {
        while (((p0).isLessThanOrEqualTo(p0)) || (p1)) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1709_i2)) {
              break L7;
            }
            _1709_i2 = (_1709_i2).plus(_dafny.ONE);
            let _1710_v12;
            let _nw329 = Array((new BigNumber(13)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
            _1710_v12 = _nw329;
            let _1711_v13;
            _1711_v13 = _dafny.Map.Empty.slice().updateUnsafe(false,_module.__default.fm21(globalState));
            let _1712_v14;
            _1712_v14 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_1711_v13);
            let _1713_v15;
            _1713_v15 = _dafny.Seq.of((_this).f17, p1);
            let _rhs299 = _1710_v12;
            let _rhs300 = (_this).fm18(true, new BigNumber(50), _1712_v14, (_this).f19, globalState);
            let _rhs301 = _dafny.Seq.contains(_1713_v15, _module.__default.fm21(globalState));
            let _lhs233 = globalState;
            let _lhs234 = _this;
            _1710_v12 = _rhs299;
            _lhs233.f2 = _rhs300;
            _lhs234.f16 = _rhs301;
            (_this).f16 = !((_this).f20) || ((_this).f20);
            let _1714_v16;
            _1714_v16 = new _dafny.CodePoint('v'.codePointAt(0));
            let _1715_v17;
            _1715_v17 = _module.D6.create_DC9(_1714_v16);
            _1714_v16 = (_1715_v17).dtor_cf14;
            (globalState).f7 = p0;
          }
        }
      }
      let _1716_v18;
      let _init52 = function (_1717_i3) {
        return (_this).f20;
      };
      let _nw330 = Array((new BigNumber(24)).toNumber());
      for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw330.length); _i0_52++) {
        _nw330[_i0_52] = _init52(new BigNumber(_i0_52));
      }
      _1716_v18 = _nw330;
      let _1718_v19;
      _1718_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1716_v18,((_this).f19).minus(_this.f14));
      let _1719_v20;
      _1719_v20 = _dafny.Seq.UnicodeFromString("yc");
      let _1720_v21;
      _1720_v21 = _dafny.MultiSet.fromElements((_this).f19, p0, _this.f14);
      let _1721_v22;
      _1721_v22 = _dafny.Seq.of(new BigNumber((_1719_v20).length), _this.f14, _this.f14, new BigNumber((_1720_v21).cardinality()));
      let _1722_v23;
      _1722_v23 = _dafny.Seq.of(_1716_v18);
      let _1723_v24;
      _1723_v24 = _dafny.Seq.of((_1722_v23)[_module.__default.safeIndex(_this.f14, new BigNumber((_1722_v23).length))]);
      let _1724_v25;
      _1724_v25 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm22(!((_this).f20), globalState),_1719_v20);
      let _rhs302 = _1716_v18;
      let _rhs303 = _dafny.Map.Empty.slice().updateUnsafe((_1723_v24)[_module.__default.safeIndex(_this.f14, new BigNumber((_1723_v24).length))],(_dafny.ZERO).minus(p0));
      let _rhs304 = _this.f14;
      let _rhs305 = _dafny.Seq.Concat(_1721_v22, _dafny.Seq.Concat(_1721_v22, _dafny.Seq.Create(_module.__default.abs(new BigNumber(776)), function (_1725_i4) {
        return (_this).f19;
      })));
      let _rhs306 = _module.__default.safeModuloInt((_this).fm12((_1721_v22)[_module.__default.safeIndex(new BigNumber((_1721_v22).length), new BigNumber((_1721_v22).length))], globalState), new BigNumber((_1724_v25).length));
      let _lhs235 = _this;
      let _lhs236 = globalState;
      _1716_v18 = _rhs302;
      _1718_v19 = _rhs303;
      _lhs235.f14 = _rhs304;
      _1721_v22 = _rhs305;
      _lhs236.f7 = _rhs306;
      let _1726_v26;
      let _init53 = function (_1727_i5) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      };
      let _nw331 = Array((new BigNumber(2)).toNumber());
      for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw331.length); _i0_53++) {
        _nw331[_i0_53] = _init53(new BigNumber(_i0_53));
      }
      _1726_v26 = _nw331;
      let _1728_v27;
      _1728_v27 = new _dafny.CodePoint('k'.codePointAt(0));
      let _index283 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_1726_v26).length));
      (_1726_v26)[_index283] = _1728_v27;
      let _index284 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_1726_v26).length));
      (_1726_v26)[_index284] = _1728_v27;
      let _source28 = _module.D6.create_DC9(new _dafny.CodePoint('x'.codePointAt(0)));
      if (_source28.is_DC10) {
        let _1729_v28;
        _1729_v28 = _module.D6.create_DC10();
        let _1730_v29;
        _1730_v29 = _dafny.Set.fromElements(new BigNumber((_1719_v20).length));
        let _1731_v30;
        _1731_v30 = _dafny.Set.fromElements(_1730_v29);
        let _1732_v31;
        let _nw332 = Array((new BigNumber(2)).toNumber());
        _nw332[(_dafny.ZERO).toNumber()] = new BigNumber((_1731_v30).length);
        _nw332[(_dafny.ONE).toNumber()] = p0;
        _1732_v31 = _nw332;
        let _1733_v32;
        _1733_v32 = _dafny.Seq.of(_1732_v31, _1732_v31, _1732_v31);
        let _1734_v33;
        _1734_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1729_v28,(_1733_v32)[_module.__default.safeIndex(_this.f14, new BigNumber((_1733_v32).length))]);
        _1734_v33 = ((true) ? (_1734_v33) : (_1734_v33));
        let _1735_v34;
        let _nw333 = new _module.C4();
        _nw333.__ctor((_this).f17, _this.f16, p1);
        _1735_v34 = _nw333;
        let _1736_v35;
        _1736_v35 = _dafny.Seq.of(_1735_v34, _1735_v34);
        _1736_v35 = _dafny.Seq.Concat(_1736_v35, _1736_v35);
        let _1737_v36;
        _1737_v36 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_1721_v22),_dafny.Seq.Concat(_1719_v20, _dafny.Seq.UnicodeFromString("eiktmgg")));
        let _1738_v37;
        _1738_v37 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
        let _1739_v38;
        _1739_v38 = _dafny.Map.Empty.slice().updateUnsafe((_1726_v26)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_1726_v26).length))],p1);
        let _rhs307 = ((_1737_v36).Merge(_1737_v36)).update(_1720_v21, (((((_1738_v37).contains(new BigNumber((_1739_v38).length))) ? ((_1738_v37).get(new BigNumber((_1739_v38).length))) : (p1))) ? (_1719_v20) : (_1719_v20)));
        let _rhs308 = p0;
        let _rhs309 = new BigNumber((_dafny.Seq.Concat(_1719_v20, _1719_v20)).length);
        let _lhs237 = globalState;
        let _lhs238 = globalState;
        _1737_v36 = _rhs307;
        _lhs237.f2 = _rhs308;
        _lhs238.f7 = _rhs309;
        let _1740_v39;
        let _1741_v40;
        let _1742_v41;
        let _1743_v42;
        let _out29;
        let _out30;
        let _out31;
        let _out32;
        let _outcollector10 = (_1735_v34).m6((_this).f17, !((_1735_v34).f17), p0, (_1726_v26)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_1726_v26).length))], globalState);
        _out29 = _outcollector10[0];
        _out30 = _outcollector10[1];
        _out31 = _outcollector10[2];
        _out32 = _outcollector10[3];
        _1740_v39 = _out29;
        _1741_v40 = _out30;
        _1742_v41 = _out31;
        _1743_v42 = _out32;
      } else {
        let _1744___mcc_h0 = (_source28).cf14;
        let _1745_cf14 = _1744___mcc_h0;
        let _1746_v43;
        _1746_v43 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,p1);
        _1719_v20 = _dafny.Seq.Concat(_1719_v20, _dafny.Seq.Concat(_module.__default.fm42(p1, (_this).f19, new BigNumber((_1746_v43).length), (_this).f17, globalState), _1719_v20));
        (globalState).f7 = p0;
        let _1747_v44;
        let _init54 = function (_1748_i6) {
          return (_1748_i6).minus(_this.f14);
        };
        let _nw334 = Array((_dafny.ONE).toNumber());
        for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw334.length); _i0_54++) {
          _nw334[_i0_54] = _init54(new BigNumber(_i0_54));
        }
        _1747_v44 = _nw334;
        _1747_v44 = _1747_v44;
        let _1749_v45;
        _1749_v45 = _dafny.Seq.of((_this).f20);
        let _1750_v46;
        _1750_v46 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,(_this).f20);
        _1749_v45 = _module.__default.fm52((((_1750_v46).contains(_this.f14)) ? ((_1750_v46).get(_this.f14)) : (p1)), (_dafny.ZERO).minus((p0).minus(p0)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,(_this).f20)).length), (_dafny.ZERO).minus(_module.__default.fm0((_this).f17, globalState)), globalState);
      }
      let _1751_v47;
      let _init55 = ((_1752_p0) => function (_1753_i7) {
        return _module.__default.safeModuloInt(_1753_i7, _1752_p0);
      })(p0);
      let _nw335 = Array((new BigNumber(16)).toNumber());
      for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw335.length); _i0_55++) {
        _nw335[_i0_55] = _init55(new BigNumber(_i0_55));
      }
      _1751_v47 = _nw335;
      _1751_v47 = _1751_v47;
      r0 = _this.f14;
      return r0;
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _1754_v0;
      _1754_v0 = _dafny.Seq.of((_this).f17);
      let _1755_v1;
      _1755_v1 = _1754_v0;
      let _source29 = _1755_v1;
      let _1756___mcc_h0 = _source29;
      let _1757_cf0 = _1756___mcc_h0;
      let _1758_v2;
      let _init56 = function (_1759_i0) {
        return (_1759_i0).multipliedBy((_this).f19);
      };
      let _nw336 = Array((new BigNumber(21)).toNumber());
      for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw336.length); _i0_56++) {
        _nw336[_i0_56] = _init56(new BigNumber(_i0_56));
      }
      _1758_v2 = _nw336;
      let _1760_v3;
      _1760_v3 = _dafny.MultiSet.fromElements(_1758_v2, _1758_v2);
      let _1761_v4;
      let _nw337 = new _module.C4();
      _nw337.__ctor(!((((_1760_v3).contains(_1758_v2)) ? ((_1760_v3).get(_1758_v2)) : (_this.f14))).isEqualTo((_this).f19), ((_this.f16) ? (false) : (p2)), (_this).f17);
      _1761_v4 = _nw337;
      let _1762_v5;
      let _nw338 = new _module.C1();
      _nw338.__ctor(_this.f16, p3);
      _1762_v5 = _nw338;
      let _1763_v6;
      let _nw339 = Array((new BigNumber(29)).toNumber()).fill([]);
      _1763_v6 = _nw339;
      let _1764_v7;
      let _nw340 = Array((new BigNumber(26)).toNumber()).fill(_module.D15.Default());
      _1764_v7 = _nw340;
      let _1765_v8;
      _1765_v8 = _dafny.Seq.UnicodeFromString("gumglvsn");
      let _1766_v9;
      _1766_v9 = _dafny.Seq.of(p1, ((_this).f19).plus(p3), p3);
      let _1767_v10;
      _1767_v10 = _dafny.Seq.of(_1764_v7, _1764_v7, _1764_v7);
      let _1768_v11;
      _1768_v11 = _dafny.Seq.of(_1764_v7, _1764_v7, _1764_v7, (_1767_v10)[_module.__default.safeIndex(p1, new BigNumber((_1767_v10).length))]);
      let _rhs310 = new BigNumber(-202);
      let _rhs311 = (_dafny.Seq.IsPrefixOf(_1765_v8, _1765_v8)) || (!((_this).f20));
      let _rhs312 = (_1766_v9)[_module.__default.safeIndex((_this).f19, new BigNumber((_1766_v9).length))];
      let _rhs313 = _1763_v6;
      let _rhs314 = (_1768_v11)[_module.__default.safeIndex(p3, new BigNumber((_1768_v11).length))];
      let _lhs239 = globalState;
      let _lhs240 = _1762_v5;
      let _lhs241 = globalState;
      _lhs239.f2 = _rhs310;
      _lhs240.f25 = _rhs311;
      _lhs241.f7 = _rhs312;
      _1763_v6 = _rhs313;
      _1764_v7 = _rhs314;
      (_1762_v5).f25 = _this.f16;
      let _1769_i1;
      _1769_i1 = _dafny.ZERO;
      L8: {
        while ((_this).f17) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1769_i1)) {
              break L8;
            }
            _1769_i1 = (_1769_i1).plus(_dafny.ONE);
            (_this).f14 = p3;
            let _1770_v12;
            _1770_v12 = _module.D9.create_DC16(_this.f14, _this.f16);
            let _1771_v13;
            _1771_v13 = _dafny.Seq.of(_this.f14, (_1770_v12).dtor_cf23, p3, (_dafny.ZERO).minus(_this.f14), _module.__default.fm0(p2, globalState));
            _1771_v13 = _1771_v13;
            let _1772_v14;
            _1772_v14 = _module.D4.create_DC7(_1755_v1, new BigNumber((_dafny.Seq.UnicodeFromString("gpuicrr")).length));
            let _pat_let_tv23 = p3;
            let _pat_let_tv24 = _1755_v1;
            let _pat_let_tv25 = _1755_v1;
            let _1773_v15;
            let _nw341 = Array((new BigNumber(21)).toNumber());
            _nw341[(_dafny.ZERO).toNumber()] = _1772_v14;
            _nw341[(_dafny.ONE).toNumber()] = function (_pat_let30_0) {
              return function (_1774_dt__update__tmp_h0) {
                return function (_pat_let31_0) {
                  return function (_1775_dt__update_hcf12_h0) {
                    return _module.D4.create_DC7((_1774_dt__update__tmp_h0).dtor_cf11, _1775_dt__update_hcf12_h0);
                  }(_pat_let31_0);
                }(_pat_let_tv23);
              }(_pat_let30_0);
            }(_module.D4.create_DC7(_dafny.Seq.of(_module.__default.fm31(p2, (_this).f17, globalState)), p1));
            _nw341[(new BigNumber(2)).toNumber()] = _module.D4.create_DC7(_1755_v1, (_this).f19);
            _nw341[(new BigNumber(3)).toNumber()] = _1772_v14;
            _nw341[(new BigNumber(4)).toNumber()] = _1772_v14;
            _nw341[(new BigNumber(5)).toNumber()] = _module.__default.fm59(globalState);
            _nw341[(new BigNumber(6)).toNumber()] = _1772_v14;
            _nw341[(new BigNumber(7)).toNumber()] = _module.D4.create_DC7(_1755_v1, p3);
            _nw341[(new BigNumber(8)).toNumber()] = _module.D4.create_DC7(_1755_v1, p1);
            _nw341[(new BigNumber(9)).toNumber()] = _1772_v14;
            _nw341[(new BigNumber(10)).toNumber()] = _1772_v14;
            _nw341[(new BigNumber(11)).toNumber()] = _module.D4.create_DC7(_1755_v1, (_this).f19);
            _nw341[(new BigNumber(12)).toNumber()] = _1772_v14;
            _nw341[(new BigNumber(13)).toNumber()] = _module.D4.create_DC7(_1755_v1, (_this).f19);
            _nw341[(new BigNumber(14)).toNumber()] = function (_pat_let32_0) {
              return function (_1776_dt__update__tmp_h1) {
                return function (_pat_let33_0) {
                  return function (_1777_dt__update_hcf11_h0) {
                    return _module.D4.create_DC7(_1777_dt__update_hcf11_h0, (_1776_dt__update__tmp_h1).dtor_cf12);
                  }(_pat_let33_0);
                }(_pat_let_tv24);
              }(_pat_let32_0);
            }(_module.__default.fm59(globalState));
            _nw341[(new BigNumber(15)).toNumber()] = _1772_v14;
            _nw341[(new BigNumber(16)).toNumber()] = _1772_v14;
            _nw341[(new BigNumber(17)).toNumber()] = _1772_v14;
            _nw341[(new BigNumber(18)).toNumber()] = function (_pat_let34_0) {
              return function (_1778_dt__update__tmp_h2) {
                return function (_pat_let35_0) {
                  return function (_1779_dt__update_hcf11_h1) {
                    return _module.D4.create_DC7(_1779_dt__update_hcf11_h1, (_1778_dt__update__tmp_h2).dtor_cf12);
                  }(_pat_let35_0);
                }(_pat_let_tv25);
              }(_pat_let34_0);
            }(_1772_v14);
            _nw341[(new BigNumber(19)).toNumber()] = _module.__default.fm59(globalState);
            _nw341[(new BigNumber(20)).toNumber()] = _1772_v14;
            _1773_v15 = _nw341;
            let _index285 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_1773_v15).length));
            (_1773_v15)[_index285] = _module.D4.create_DC7(_1755_v1, p1);
            let _index286 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_1773_v15).length));
            (_1773_v15)[_index286] = _1772_v14;
            let _1780_v16;
            _1780_v16 = _dafny.MultiSet.fromElements((_this).f19);
            _1780_v16 = ((_1780_v16).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.update(_1771_v13, _module.__default.safeIndex((_this).f19, new BigNumber((_1771_v13).length)), new BigNumber(-990))))).update(p1, _module.__default.abs(p3));
          }
        }
      }
      let _1781_v17;
      let _init57 = function (_1782_i2) {
        return (_this).f20;
      };
      let _nw342 = Array((_dafny.ONE).toNumber());
      for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw342.length); _i0_57++) {
        _nw342[_i0_57] = _init57(new BigNumber(_i0_57));
      }
      _1781_v17 = _nw342;
      let _1783_v18;
      _1783_v18 = _dafny.Seq.of(p0, new _dafny.CodePoint('b'.codePointAt(0)));
      let _index287 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1781_v17).length));
      (_1781_v17)[_index287] = _module.__default.fm27(p3, new BigNumber(599), new BigNumber((_dafny.MultiSet.FromArray(_1783_v18)).cardinality()), _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(p2, globalState),new BigNumber((_1754_v0).length)), globalState);
      let _1784_v19;
      _1784_v19 = _module.D12.create_DC23((_this).f17, (_this).fm12(p3, globalState), p3);
      let _1785_v20;
      _1785_v20 = _dafny.Map.Empty.slice().updateUnsafe(p3,_1784_v19);
      let _index288 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1781_v17).length));
      let _rhs315 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("sm"), _1783_v18);
      let _rhs316 = (_this).f20;
      let _rhs317 = (new BigNumber(706)).multipliedBy(((_this.f16) ? (new BigNumber(387)) : (new BigNumber((_1785_v20).length))));
      let _rhs318 = _module.__default.safeModuloInt(_this.f14, ((_this).f19).multipliedBy(p1));
      let _rhs319 = (((_1754_v0)[_module.__default.safeIndex(p3, new BigNumber((_1754_v0).length))]) ? (!(_this.f16) || ((_this).f20)) : ((_this).f17));
      let _lhs242 = _this;
      let _lhs243 = _this;
      let _lhs244 = globalState;
      let _lhs245 = _this;
      let _lhs246 = _1781_v17;
      let _lhs247 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1781_v17).length));
      _lhs242.f16 = _rhs315;
      _lhs243.f16 = _rhs316;
      _lhs244.f7 = _rhs317;
      _lhs245.f14 = _rhs318;
      _lhs246[_lhs247] = _rhs319;
      let _1786_v21;
      _1786_v21 = _dafny.Map.Empty.slice().updateUnsafe(p3,p2);
      (_this).f16 = ((_this).f17) === ((((_1786_v21).contains((_this).f19)) ? ((_1786_v21).get((_this).f19)) : (p2)));
      let _index289 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1781_v17).length));
      (_1781_v17)[_index289] = !((_this).f20);
      let _1787_v22;
      _1787_v22 = _dafny.MultiSet.fromElements(p1);
      _1787_v22 = _1787_v22;
      r0 = (p3).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f14)));
      return r0;
    }
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.Seq.UnicodeFromString("");
      let r3 = false;
      (globalState).f7 = _this.f14;
      let _1788_v0;
      _1788_v0 = _dafny.Seq.UnicodeFromString("i");
      let _hi13 = _module.__default.safeModuloInt((_this).f19, new BigNumber((_dafny.Seq.of(new BigNumber((_1788_v0).length))).length));
      for (let _1789_i0 = p2; _1789_i0.isLessThan(_hi13); _1789_i0 = _1789_i0.plus(_dafny.ONE)) {
        let _1790_v1;
        _1790_v1 = _module.D11.create_DC20();
        _1790_v1 = _module.__default.fm60((_this).f17, (new BigNumber(543)).multipliedBy(new BigNumber(728)), p3, globalState);
        let _1791_v2;
        let _nw343 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Map.Empty);
        _1791_v2 = _nw343;
        let _1792_v3;
        _1792_v3 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,new BigNumber((_1788_v0).length));
        let _index290 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1791_v2).length));
        (_1791_v2)[_index290] = _1792_v3;
        let _index291 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1791_v2).length));
        (_1791_v2)[_index291] = (_1792_v3).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,_1789_i0));
        let _1793_v4;
        let _nw344 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
        _1793_v4 = _nw344;
        let _index292 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_1793_v4).length));
        (_1793_v4)[_index292] = _1789_i0;
        let _index293 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_1793_v4).length));
        (_1793_v4)[_index293] = new BigNumber(-545);
        _1788_v0 = _dafny.Seq.Concat(_1788_v0, _1788_v0);
      }
      let _1794_v5;
      let _nw345 = Array((new BigNumber(24)).toNumber()).fill(false);
      _1794_v5 = _nw345;
      for (const _guard_loop_9 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1794_v5).length))) {
        let _1795_i1 = _guard_loop_9;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1795_i1)) && ((_1795_i1).isLessThan(new BigNumber((_1794_v5).length))))) {
          (_1794_v5)[(_1795_i1)] = !(((!((_this).f17)) ? (_this.f16) : (p0)));
        }
      }
      let _hi14 = (p2).multipliedBy(p2);
      for (let _1796_i2 = (_this).f19; _1796_i2.isLessThan(_hi14); _1796_i2 = _1796_i2.plus(_dafny.ONE)) {
        let _1797_v6;
        _1797_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1788_v0,_1796_i2);
        let _1798_v7;
        _1798_v7 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
        let _1799_v8;
        _1799_v8 = _dafny.MultiSet.fromElements((_this).f17, (_this).f17);
        _1797_v6 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.update(_1788_v0, _module.__default.safeIndex(new BigNumber((_1798_v7).length), new BigNumber((_1788_v0).length)), p3), _module.__default.safeIndex((((_1799_v8).contains(p1)) ? ((_1799_v8).get(p1)) : ((_this).f19)), new BigNumber((_dafny.Seq.update(_1788_v0, _module.__default.safeIndex(new BigNumber((_1798_v7).length), new BigNumber((_1788_v0).length)), p3)).length)), p3), _dafny.Seq.UnicodeFromString("hr")),_module.__default.safeModuloInt((_this).f19, _1796_i2));
        let _1800_v9;
        _1800_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1788_v0,!(!((_this).f20)));
        let _1801_v10;
        _1801_v10 = _dafny.Seq.of(false);
        let _1802_v11;
        _1802_v11 = _dafny.Set.fromElements(_1801_v10, _1801_v10);
        let _1803_v12;
        let _nw346 = Array((new BigNumber(7)).toNumber());
        _nw346[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1788_v0,_this.f16);
        _nw346[(_dafny.ONE).toNumber()] = _1800_v9;
        _nw346[(new BigNumber(2)).toNumber()] = _1800_v9;
        _nw346[(new BigNumber(3)).toNumber()] = _1800_v9;
        _nw346[(new BigNumber(4)).toNumber()] = _module.__default.fm61(_1802_v11, new BigNumber((_1799_v8).cardinality()), globalState);
        _nw346[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("aqpgv"),p1);
        _nw346[(new BigNumber(6)).toNumber()] = _1800_v9;
        _1803_v12 = _nw346;
        let _index294 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_1803_v12).length));
        (_1803_v12)[_index294] = _1800_v9;
        let _1804_v13;
        _1804_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,(_this).f17);
        let _1805_v14;
        _1805_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1804_v13).length),p2);
        let _1806_v15;
        _1806_v15 = _dafny.Map.Empty.slice().updateUnsafe((((_1805_v14).contains(_1796_i2)) ? ((_1805_v14).get(_1796_i2)) : ((_this).f19)),p0);
        let _1807_v16;
        let _nw347 = new _module.C4();
        _nw347.__ctor((((_1806_v15).contains(_this.f14)) ? ((_1806_v15).get(_this.f14)) : ((_this).f17)), _this.f16, ((p1) ? (_this.f16) : (p1)));
        _1807_v16 = _nw347;
        let _1808_v17;
        _1808_v17 = new _dafny.CodePoint('e'.codePointAt(0));
        let _1809_v19;
        _1809_v19 = _dafny.Seq.of(_1788_v0);
        let _index295 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_1803_v12).length));
        let _rhs320 = ((function () {
          let _coll70 = new _dafny.Map();
          for (const _compr_70 of (_1809_v19).Elements) {
            let _1810_v18 = _compr_70;
            if (_dafny.Seq.contains(_1809_v19, _1810_v18)) {
              _coll70.push([_1810_v18,p1]);
            }
          }
          return _coll70;
        }()).Merge(_1800_v9)).Merge(_module.__default.fm61(_1802_v11, _this.f14, globalState));
        let _rhs321 = ((_1804_v13).Merge(_1804_v13)).Merge(_1804_v13);
        let _rhs322 = _1807_v16;
        let _rhs323 = _1808_v17;
        let _lhs248 = _1803_v12;
        let _lhs249 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_1803_v12).length));
        _lhs248[_lhs249] = _rhs320;
        _1804_v13 = _rhs321;
        _1807_v16 = _rhs322;
        _1808_v17 = _rhs323;
        let _1811_v20;
        _1811_v20 = _module.D3.create_DC3(p1);
        let _1812_v21;
        let _nw348 = new _module.C2();
        _nw348.__ctor(_1811_v20, (_this).f17, (_this).f17);
        _1812_v21 = _nw348;
        let _index296 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_1794_v5).length));
        (_1794_v5)[_index296] = !(p1) || ((_this).f20);
        let _index297 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_1794_v5).length));
        (_1794_v5)[_index297] = (_this).f20;
      }
      let _index298 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_1794_v5).length));
      (_1794_v5)[_index298] = !((_this).f19).isEqualTo((_dafny.ZERO).minus((_this).f19));
      let _1813_v22;
      _1813_v22 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f17);
      let _1814_v23;
      _1814_v23 = _module.D12.create_DC22(_1788_v0);
      let _pat_let_tv26 = p1;
      let _pat_let_tv27 = p1;
      let _index299 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_1794_v5).length));
      let _rhs324 = _this.f14;
      let _rhs325 = false;
      let _rhs326 = !(function (_source30) {
        if (_source30.is_DC23) {
          let _1815___mcc_h0 = (_source30).cf30;
          let _1816___mcc_h1 = (_source30).cf31;
          let _1817___mcc_h2 = (_source30).cf32;
          let _1818_cf32 = _1817___mcc_h2;
          let _1819_cf31 = _1816___mcc_h1;
          let _1820_cf30 = _1815___mcc_h0;
          return _pat_let_tv26;
        } else if (_source30.is_DC24) {
          let _1821___mcc_h3 = (_source30).cf33;
          let _1822___mcc_h4 = (_source30).cf34;
          let _1823___mcc_h5 = (_source30).cf35;
          let _1824___mcc_h6 = (_source30).cf36;
          let _1825___mcc_h7 = (_source30).cf37;
          let _1826_cf37 = _1825___mcc_h7;
          let _1827_cf36 = _1824___mcc_h6;
          let _1828_cf35 = _1823___mcc_h5;
          let _1829_cf34 = _1822___mcc_h4;
          let _1830_cf33 = _1821___mcc_h3;
          return _1830_cf33;
        } else {
          let _1831___mcc_h8 = (_source30).cf29;
          let _1832_cf29 = _1831___mcc_h8;
          return _pat_let_tv27;
        }
      }(_1814_v23));
      let _rhs327 = (_this).f20;
      let _rhs328 = (_1813_v22).Merge(_1813_v22);
      let _lhs250 = globalState;
      let _lhs251 = _this;
      let _lhs252 = _1794_v5;
      let _lhs253 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_1794_v5).length));
      _lhs250.f7 = _rhs324;
      _lhs251.f16 = _rhs325;
      r1 = _rhs326;
      _lhs252[_lhs253] = _rhs327;
      _1813_v22 = _rhs328;
      let _1833_v24;
      let _init58 = ((_1834_p1, _1835_p0) => function (_1836_i3) {
        return _dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(!(!(_1834_p1)), _1835_p0, _1834_p1)).length));
      })(p1, p0);
      let _nw349 = Array((new BigNumber(27)).toNumber());
      for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw349.length); _i0_58++) {
        _nw349[_i0_58] = _init58(new BigNumber(_i0_58));
      }
      _1833_v24 = _nw349;
      let _1837_v25;
      _1837_v25 = _dafny.MultiSet.fromElements((_this).f19);
      let _1838_v26;
      _1838_v26 = _dafny.Set.fromElements(false);
      let _1839_v27;
      _1839_v27 = _dafny.Seq.of(_this.f14, (_this).f19);
      let _1840_v28;
      _1840_v28 = _dafny.Seq.of(new BigNumber((_1839_v27).length), (_this).f19);
      let _index300 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_1833_v24).length));
      (_1833_v24)[_index300] = (_1837_v25).update(new BigNumber((_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_1838_v26).length)), _module.__default.fm38(p1, p2, new BigNumber(816), globalState), _1839_v27, _1840_v28)).length), _module.__default.abs((_this).f19));
      let _1841_v29;
      _1841_v29 = _dafny.MultiSet.fromElements(_this.f16);
      let _1842_v30;
      _1842_v30 = _dafny.Map.Empty.slice().updateUnsafe((_1841_v29).IsSubsetOf(_module.__default.fm1(globalState)),(_1837_v25).update(new BigNumber(322), _module.__default.abs(new BigNumber((_1788_v0).length))));
      let _index301 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_1833_v24).length));
      (_1833_v24)[_index301] = (((_1842_v30).contains((_1794_v5)[_module.__default.safeIndex(new BigNumber(443), new BigNumber((_1794_v5).length))])) ? ((_1842_v30).get((_1794_v5)[_module.__default.safeIndex(new BigNumber(443), new BigNumber((_1794_v5).length))])) : (_dafny.MultiSet.FromArray(_1839_v27)));
      r0 = ((_this).fm18((_this).f20, _module.__default.fm0((_1794_v5)[_module.__default.safeIndex(new BigNumber(443), new BigNumber((_1794_v5).length))], globalState), function () {
        let _coll71 = new _dafny.Map();
        for (const _compr_71 of _dafny.IntegerRange(new BigNumber(-924), new BigNumber(707))) {
          let _1843_v31 = _compr_71;
          if (((new BigNumber(-924)).isLessThanOrEqualTo(_1843_v31)) && ((_1843_v31).isLessThan(new BigNumber(707)))) {
            _coll71.push([(_1843_v31).plus(p2),_1813_v22]);
          }
        }
        return _coll71;
      }(), (_dafny.ZERO).minus((_this).f19), globalState)).plus(new BigNumber(704));
      let _1844_v32;
      _1844_v32 = _dafny.Set.fromElements(p2, p2, (_dafny.ZERO).minus(_this.f14), _this.f14);
      r1 = !((_1844_v32).IsProperSubsetOf(_1844_v32));
      r2 = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("dljlsf"), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.UnicodeFromString("dljlsf")).length)), p3);
      r3 = (_this).f17;
      return [r0, r1, r2, r3];
    }
    get f19() {
      let _this = this;
      return _this._f19;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f16 = false;
      this._f14 = _dafny.ZERO;
      this._f17 = false;
    }
    _parentTraits() {
      return [_module.T2, _module.T0, _module.T1];
    }
    get f16() {
      let _this = this;
      return _this._f16;
    };
    set f16(value) {
      let _this = this;
      _this._f16 = value;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    set f14(value) {
      let _this = this;
      _this._f14 = value;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    __ctor(f16, f17, f14) {
      let _this = this;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f14 = f14;
      return;
    }
    fm12(p0, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus((_this.f14).plus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f14,(_this).f17)).length)), _dafny.Seq.of((_module.D3.create_DC4(true, _this.f14, _this.f14, (_this).f17, false)).dtor_cf5), _dafny.Seq.Create(_module.__default.abs(new BigNumber(499)), function (_1845_i0) {
        return _this.f14;
      }), (_dafny.Seq.of(_this.f14))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(504)), function (_1846_i1) {
        return _dafny.Seq.of(new BigNumber(91), _this.f14, _this.f14, _this.f14, _this.f14);
      }))).length)));
    };
    fm13(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of((_this).f17)), _dafny.Seq.Concat(_dafny.Seq.of(false, (_this).f17), _dafny.Seq.of(!(_this.f16))));
    };
    fm3(globalState) {
      let _this = this;
      if (_this.f16) {
        return _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_this.f14);
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_this.f14);
      }
    };
    fm16(p0, globalState) {
      let _this = this;
      return _this.f14;
    };
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.Seq.UnicodeFromString("");
      let r3 = false;
      let _1847_v0;
      _1847_v0 = _dafny.Seq.of(p1, (_this).f17);
      let _1848_v1;
      _1848_v1 = _dafny.Set.fromElements(p2, p2);
      let _1849_v2;
      _1849_v2 = _dafny.Map.Empty.slice().updateUnsafe(p2,p1);
      let _1850_v3;
      _1850_v3 = _dafny.Seq.of(new BigNumber(602), p2, p2);
      let _1851_v4;
      _1851_v4 = _dafny.Set.fromElements(p1, p0, _this.f16, _this.f16);
      let _1852_v5;
      let _nw350 = Array((new BigNumber(15)).toNumber());
      _nw350[(_dafny.ZERO).toNumber()] = (p2).isLessThanOrEqualTo(p2);
      _nw350[(_dafny.ONE).toNumber()] = (_this).f17;
      _nw350[(new BigNumber(2)).toNumber()] = p0;
      _nw350[(new BigNumber(3)).toNumber()] = !(!(!(p0)));
      _nw350[(new BigNumber(4)).toNumber()] = !(!_dafny.areEqual(_1847_v0, _1847_v0));
      _nw350[(new BigNumber(5)).toNumber()] = _this.f16;
      _nw350[(new BigNumber(6)).toNumber()] = (_this).f17;
      _nw350[(new BigNumber(7)).toNumber()] = p1;
      _nw350[(new BigNumber(8)).toNumber()] = (_1848_v1).IsDisjointFrom(_1848_v1);
      _nw350[(new BigNumber(9)).toNumber()] = !(!((((_this).f17) ? ((_this).f17) : ((((_1849_v2).contains(new BigNumber((_dafny.MultiSet.fromElements(_this.f16)).cardinality()))) ? ((_1849_v2).get(new BigNumber((_dafny.MultiSet.fromElements(_this.f16)).cardinality()))) : (_this.f16))))));
      _nw350[(new BigNumber(10)).toNumber()] = _module.__default.fm17(p1, p2, p2, new BigNumber((_1850_v3).length), globalState);
      _nw350[(new BigNumber(11)).toNumber()] = (_1851_v4).IsDisjointFrom(_1851_v4);
      _nw350[(new BigNumber(12)).toNumber()] = (_this).f17;
      _nw350[(new BigNumber(13)).toNumber()] = _this.f16;
      _nw350[(new BigNumber(14)).toNumber()] = _module.__default.fm17((_this).f17, _this.f14, _this.f14, new BigNumber((_dafny.MultiSet.FromArray(_1847_v0)).cardinality()), globalState);
      _1852_v5 = _nw350;
      for (const _guard_loop_10 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1852_v5).length))) {
        let _1853_i0 = _guard_loop_10;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1853_i0)) && ((_1853_i0).isLessThan(new BigNumber((_1852_v5).length))))) {
          (_1852_v5)[(_1853_i0)] = !_dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("gdxkk"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(520)), function (_1854_i1) {
            return new _dafny.CodePoint('e'.codePointAt(0));
          })), _dafny.Seq.UnicodeFromString("cj"));
        }
      }
      let _1855_i2;
      _1855_i2 = _dafny.ZERO;
      L9: {
        while (!(_module.__default.fm17((_this).f17, (p2).minus(_this.f14), _module.__default.safeModuloInt(p2, (_dafny.ZERO).minus(p2)), _this.f14, globalState))) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1855_i2)) {
              break L9;
            }
            _1855_i2 = (_1855_i2).plus(_dafny.ONE);
            let _1856_v6;
            _1856_v6 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
            _1856_v6 = (_1856_v6).update((_this).f17, p1);
            let _1857_v7;
            let _init59 = ((_1858_p3) => function (_1859_i3) {
              return _dafny.Seq.Create(_module.__default.abs(new BigNumber(501)), ((_1860_p3) => function (_1861_i4) {
                return _1860_p3;
              })(_1858_p3));
            })(p3);
            let _nw351 = Array((new BigNumber(27)).toNumber());
            for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw351.length); _i0_59++) {
              _nw351[_i0_59] = _init59(new BigNumber(_i0_59));
            }
            _1857_v7 = _nw351;
            _1857_v7 = _1857_v7;
            r3 = p0;
            let _1862_v8;
            let _nw352 = new _module.C4();
            _nw352.__ctor(p1, _this.f16, p0);
            _1862_v8 = _nw352;
          }
        }
      }
      let _1863_v9;
      _1863_v9 = _dafny.Seq.of(_1851_v4, (_1851_v4).Intersect(_1851_v4));
      _1863_v9 = _dafny.Seq.of((_1851_v4).Difference(_1851_v4));
      let _1864_v10;
      let _nw353 = new _module.C4();
      _nw353.__ctor(!_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("j"), p3), _this.f16, _this.f16);
      _1864_v10 = _nw353;
      let _1865_v11;
      let _nw354 = new _module.C5();
      _nw354.__ctor(_module.__default.fm43(p1, globalState), p1);
      _1865_v11 = _nw354;
      if (_this.f16) {
        if (((_module.__default.fm21(globalState)) ? ((_1864_v10).f21) : (!((_1864_v10).f21) || ((_1864_v10).f21)))) {
          let _1866_v12;
          _1866_v12 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
          _1866_v12 = (_1866_v12).Merge((_1866_v12).update(new BigNumber((_1866_v12).length), p2));
          let _1867_v13;
          let _nw355 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _1867_v13 = _nw355;
          let _nw356 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
          _1867_v13 = _nw356;
          let _1868_v14;
          _1868_v14 = _dafny.Set.fromElements(_1852_v5);
          r1 = (_1868_v14).IsProperSubsetOf(_1868_v14);
          _1850_v3 = _dafny.Seq.update(_dafny.Seq.Concat(_1850_v3, _dafny.Seq.Concat(_dafny.Seq.of(p2), _1850_v3)), _module.__default.safeIndex(new BigNumber(745), new BigNumber((_dafny.Seq.Concat(_1850_v3, _dafny.Seq.Concat(_dafny.Seq.of(p2), _1850_v3))).length)), p2);
          (globalState).f2 = _module.__default.safeModuloInt(p2, _this.f14);
        } else {
          let _1869_v15;
          _1869_v15 = new _dafny.CodePoint('h'.codePointAt(0));
          let _1870_v16;
          _1870_v16 = _dafny.Seq.UnicodeFromString("oiltx");
          let _1871_v17;
          _1871_v17 = _dafny.Seq.of(_1870_v16);
          let _rhs329 = p1;
          let _rhs330 = p3;
          let _rhs331 = (_1864_v10).f21;
          let _rhs332 = _1871_v17;
          let _rhs333 = new BigNumber((_1847_v0).length);
          let _lhs254 = globalState;
          r3 = _rhs329;
          _1869_v15 = _rhs330;
          r3 = _rhs331;
          _1871_v17 = _rhs332;
          _lhs254.f7 = _rhs333;
          r3 = (_1864_v10).f21;
          r3 = !(p1);
          let _index302 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_1852_v5).length));
          (_1852_v5)[_index302] = (_1864_v10).f21;
          let _1872_v18;
          _1872_v18 = _dafny.Set.fromElements(_1852_v5);
          let _index303 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_1852_v5).length));
          (_1852_v5)[_index303] = !((((_this).f17) ? (true) : (_dafny.areEqual(_1850_v3, _dafny.Seq.update(_dafny.Seq.of(p2, new BigNumber((_1872_v18).length)), _module.__default.safeIndex(_this.f14, new BigNumber((_dafny.Seq.of(p2, new BigNumber((_1872_v18).length))).length)), new BigNumber((_1849_v2).length))))));
          (globalState).f7 = (p2).plus((p2).plus(p2));
        }
        let _1873_v19;
        _1873_v19 = _dafny.Seq.UnicodeFromString("mdpuq");
        let _1874_v20;
        _1874_v20 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_this.f14);
        let _1875_v21;
        let _nw357 = Array((new BigNumber(6)).toNumber());
        _nw357[(_dafny.ZERO).toNumber()] = _1873_v19;
        _nw357[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(468)), ((_1876_v19, _1877_v3, _1878_p2) => function (_1879_i5) {
          return (_1876_v19)[_module.__default.safeIndex((_1877_v3)[_module.__default.safeIndex(_1878_p2, new BigNumber((_1877_v3).length))], new BigNumber((_1876_v19).length))];
        })(_1873_v19, _1850_v3, p2));
        _nw357[(new BigNumber(2)).toNumber()] = _module.__default.fm42(p1, new BigNumber((_1850_v3).length), new BigNumber(342), (_1864_v10).f21, globalState);
        _nw357[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ydpgyfg"), _1873_v19);
        _nw357[(new BigNumber(4)).toNumber()] = _1873_v19;
        _nw357[(new BigNumber(5)).toNumber()] = _module.__default.fm42(p0, new BigNumber((_1850_v3).length), (((_1874_v20).contains(p2)) ? ((_1874_v20).get(p2)) : (_this.f14)), (_1864_v10).f21, globalState);
        _1875_v21 = _nw357;
        let _index304 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_1875_v21).length));
        (_1875_v21)[_index304] = _1873_v19;
        let _index305 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_1875_v21).length));
        let _rhs334 = _dafny.Seq.Concat(_1873_v19, _1873_v19);
        let _rhs335 = (_1864_v10).f21;
        let _rhs336 = _dafny.Seq.UnicodeFromString("dgew");
        let _lhs255 = _1875_v21;
        let _lhs256 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_1875_v21).length));
        let _lhs257 = _this;
        _lhs255[_lhs256] = _rhs334;
        _lhs257.f16 = _rhs335;
        r2 = _rhs336;
        let _1880_v22;
        _1880_v22 = _1852_v5;
        let _1881_v23;
        _1881_v23 = _dafny.Map.Empty.slice().updateUnsafe((_1880_v22),p2);
        _1881_v23 = (_1881_v23).update(_1852_v5, (_1850_v3)[_module.__default.safeIndex(_this.f14, new BigNumber((_1850_v3).length))]);
        let _1882_v24;
        _1882_v24 = _dafny.Map.Empty.slice().updateUnsafe((_1864_v10).f21,(_this.f14).isLessThan((_1864_v10).fm12(_this.f14, globalState)));
        let _1883_v25;
        let _nw358 = Array((new BigNumber(9)).toNumber());
        _nw358[(_dafny.ZERO).toNumber()] = p2;
        _nw358[(_dafny.ONE).toNumber()] = p2;
        _nw358[(new BigNumber(2)).toNumber()] = new BigNumber((_1850_v3).length);
        _nw358[(new BigNumber(3)).toNumber()] = new BigNumber(27);
        _nw358[(new BigNumber(4)).toNumber()] = p2;
        _nw358[(new BigNumber(5)).toNumber()] = p2;
        _nw358[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(p2);
        _nw358[(new BigNumber(7)).toNumber()] = p2;
        _nw358[(new BigNumber(8)).toNumber()] = _this.f14;
        _1883_v25 = _nw358;
        let _1884_v26;
        _1884_v26 = _dafny.Set.fromElements(_1883_v25);
        let _index306 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_1852_v5).length));
        (_1852_v5)[_index306] = (_1884_v26).IsSubsetOf(_1884_v26);
        let _1885_v27;
        _1885_v27 = _dafny.Seq.of(_1882_v24, _1882_v24, _1882_v24, _module.__default.fm40(_this.f16, globalState), (_1882_v24).update(p0, p0));
        let _1886_v28;
        _1886_v28 = _module.D4.create_DC6((_1885_v27)[_module.__default.safeIndex(_this.f14, new BigNumber((_1885_v27).length))]);
        let _index307 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_1852_v5).length));
        let _rhs337 = (_this).f17;
        let _rhs338 = _this.f14;
        let _rhs339 = (_1886_v28).dtor_cf10;
        let _rhs340 = _module.__default.fm21(globalState);
        let _rhs341 = new BigNumber((_dafny.Seq.Concat((_1875_v21)[_module.__default.safeIndex(new BigNumber(879), new BigNumber((_1875_v21).length))], _dafny.Seq.UnicodeFromString("dbl"))).length);
        let _lhs258 = globalState;
        let _lhs259 = _1852_v5;
        let _lhs260 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_1852_v5).length));
        let _lhs261 = globalState;
        r1 = _rhs337;
        _lhs258.f7 = _rhs338;
        _1882_v24 = _rhs339;
        _lhs259[_lhs260] = _rhs340;
        _lhs261.f7 = _rhs341;
        let _index308 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_1852_v5).length));
        (_1852_v5)[_index308] = !(((_1865_v11).fm12((_dafny.ZERO).minus(p2), globalState)).isLessThanOrEqualTo(_this.f14));
      } else {
        let _1887_v29;
        _1887_v29 = _dafny.Seq.UnicodeFromString("bfbavpph");
        let _1888_v30;
        _1888_v30 = _dafny.Map.Empty.slice().updateUnsafe(false,p0);
        let _1889_v31;
        let _nw359 = new _module.C3();
        _nw359.__ctor(_this.f14, _dafny.Seq.update(_1887_v29, _module.__default.safeIndex(p2, new BigNumber((_1887_v29).length)), new _dafny.CodePoint('c'.codePointAt(0))), p1, (new BigNumber((_1888_v30).length)).isEqualTo(_this.f14));
        _1889_v31 = _nw359;
        let _index309 = _module.__default.safeIndex(new BigNumber(227), new BigNumber((_1852_v5).length));
        (_1852_v5)[_index309] = _dafny.Seq.IsPrefixOf(_dafny.Seq.of(p3, p3, p3, p3), (_1889_v31).f23);
        let _1890_v32;
        _1890_v32 = _dafny.MultiSet.fromElements(p2);
        let _index310 = _module.__default.safeIndex(new BigNumber(227), new BigNumber((_1852_v5).length));
        (_1852_v5)[_index310] = !((_1890_v32).IsDisjointFrom(((p1) ? (_1890_v32) : (_1890_v32))));
        let _1891_v33;
        let _init60 = ((_1892_v4) => function (_1893_i6) {
          return _1892_v4;
        })(_1851_v4);
        let _nw360 = Array((new BigNumber(19)).toNumber());
        for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw360.length); _i0_60++) {
          _nw360[_i0_60] = _init60(new BigNumber(_i0_60));
        }
        _1891_v33 = _nw360;
        let _1894_v34;
        let _nw361 = new _module.C0();
        _nw361.__ctor(_1891_v33, (_1889_v31).f22);
        _1894_v34 = _nw361;
        let _1895_v35;
        _1895_v35 = _dafny.Seq.of((_1849_v2).Merge(_1849_v2), _1849_v2, _1849_v2, _1849_v2, _1849_v2);
        let _rhs342 = (_this).f17;
        let _rhs343 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(284)), ((_1896_p2, _1897_v10) => function (_1898_i7) {
          return _dafny.Map.Empty.slice().updateUnsafe(_1896_p2,(_1897_v10).f21);
        })(p2, _1864_v10));
        r1 = _rhs342;
        _1895_v35 = _rhs343;
        (_1894_v34).f27 = _this.f14;
      }
      r0 = (_dafny.ZERO).minus(p2);
      let _1899_v36;
      _1899_v36 = _dafny.Seq.UnicodeFromString("nei");
      let _1900_v37;
      _1900_v37 = _dafny.Seq.of(_1899_v36);
      r1 = ((_this).f17) === (_dafny.areEqual(_dafny.Seq.of(_1899_v36, _dafny.Seq.update(_1899_v36, _module.__default.safeIndex(_this.f14, new BigNumber((_1899_v36).length)), p3), _1899_v36), _1900_v37));
      r2 = _dafny.Seq.Concat(_1899_v36, _1899_v36);
      let _1901_v38;
      _1901_v38 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.fromElements(p0, (_this).f17, true)).update(p1, _module.__default.abs(_this.f14)),!(!((_this).f17)));
      let _1902_v39;
      _1902_v39 = _dafny.MultiSet.fromElements(_this.f16, (_1864_v10).f21);
      r3 = (((_this).f17) ? ((((_1901_v38).contains(_1902_v39)) ? ((_1901_v38).get(_1902_v39)) : (!(_module.__default.fm31((_this).f17, p1, globalState))))) : (_dafny.areEqual(_1847_v0, _1847_v0)));
      return [r0, r1, r2, r3];
    }
    m2(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      if (p1) {
        let _1903_v0;
        let _nw362 = Array((new BigNumber(3)).toNumber()).fill(false);
        _1903_v0 = _nw362;
        let _1904_v1;
        _1904_v1 = _1903_v0;
        let _1905_v2;
        _1905_v2 = new _dafny.CodePoint('b'.codePointAt(0));
        let _1906_v3;
        _1906_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1905_v2,_1903_v0);
        let _1907_v4;
        let _nw363 = Array((new BigNumber(14)).toNumber());
        _nw363[(_dafny.ZERO).toNumber()] = _1903_v0;
        _nw363[(_dafny.ONE).toNumber()] = _1903_v0;
        _nw363[(new BigNumber(2)).toNumber()] = _1904_v1;
        _nw363[(new BigNumber(3)).toNumber()] = _1904_v1;
        _nw363[(new BigNumber(4)).toNumber()] = _1904_v1;
        _nw363[(new BigNumber(5)).toNumber()] = _1904_v1;
        _nw363[(new BigNumber(6)).toNumber()] = _1903_v0;
        _nw363[(new BigNumber(7)).toNumber()] = _1904_v1;
        _nw363[(new BigNumber(8)).toNumber()] = _1904_v1;
        _nw363[(new BigNumber(9)).toNumber()] = _1904_v1;
        _nw363[(new BigNumber(10)).toNumber()] = _1904_v1;
        _nw363[(new BigNumber(11)).toNumber()] = _1904_v1;
        _nw363[(new BigNumber(12)).toNumber()] = (((_1906_v3).contains(_1905_v2)) ? ((_1906_v3).get(_1905_v2)) : (_1903_v0));
        _nw363[(new BigNumber(13)).toNumber()] = _1904_v1;
        _1907_v4 = _nw363;
        let _1908_v5;
        let _nw364 = Array((new BigNumber(11)).toNumber());
        _nw364[(_dafny.ZERO).toNumber()] = _1907_v4;
        _nw364[(_dafny.ONE).toNumber()] = _1907_v4;
        _nw364[(new BigNumber(2)).toNumber()] = ((_module.__default.fm17(_this.f16, new BigNumber((_dafny.Set.fromElements(p0, p0, _this.f14)).length), p0, _this.f14, globalState)) ? (_1907_v4) : (_1907_v4));
        _nw364[(new BigNumber(3)).toNumber()] = _1907_v4;
        _nw364[(new BigNumber(4)).toNumber()] = _1907_v4;
        _nw364[(new BigNumber(5)).toNumber()] = _1907_v4;
        _nw364[(new BigNumber(6)).toNumber()] = _1907_v4;
        _nw364[(new BigNumber(7)).toNumber()] = ((true) ? (_1907_v4) : (_1907_v4));
        _nw364[(new BigNumber(8)).toNumber()] = _1907_v4;
        _nw364[(new BigNumber(9)).toNumber()] = _1907_v4;
        _nw364[(new BigNumber(10)).toNumber()] = _1907_v4;
        _1908_v5 = _nw364;
        let _index311 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_1908_v5).length));
        (_1908_v5)[_index311] = _1907_v4;
        let _1909_v6;
        let _nw365 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1909_v6 = _nw365;
        let _1910_v7;
        _1910_v7 = _dafny.Seq.UnicodeFromString("yefkcasyq");
        let _index312 = _module.__default.safeIndex(new BigNumber(415), new BigNumber((_1909_v6).length));
        (_1909_v6)[_index312] = _1910_v7;
        let _index313 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_1908_v5).length));
        let _index314 = _module.__default.safeIndex(new BigNumber(415), new BigNumber((_1909_v6).length));
        let _rhs344 = _1907_v4;
        let _rhs345 = _1910_v7;
        let _rhs346 = _this.f16;
        let _lhs262 = _1908_v5;
        let _lhs263 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_1908_v5).length));
        let _lhs264 = _1909_v6;
        let _lhs265 = _module.__default.safeIndex(new BigNumber(415), new BigNumber((_1909_v6).length));
        let _lhs266 = _this;
        _lhs262[_lhs263] = _rhs344;
        _lhs264[_lhs265] = _rhs345;
        _lhs266.f16 = _rhs346;
        let _1911_v8;
        _1911_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1903_v0,((_this.f16) ? (_1903_v0) : (_1903_v0)));
        _1911_v8 = (_1911_v8).update(_1903_v0, _1903_v0);
        let _1912_v9;
        let _nw366 = new _module.C6();
        _nw366.__ctor((_dafny.ZERO).minus(_this.f14), _this.f16, p1, (p0).isLessThan(_this.f14), p0);
        _1912_v9 = _nw366;
        _1912_v9 = _1912_v9;
        (globalState).f7 = p0;
        let _1913_v10;
        let _nw367 = Array((new BigNumber(17)).toNumber()).fill([]);
        _1913_v10 = _nw367;
        let _1914_v11;
        let _nw368 = new _module.C6();
        _nw368.__ctor((_1912_v9).f19, !(_this.f16), _this.f16, !((_1912_v9).f20), _this.f14);
        _1914_v11 = _nw368;
        let _1915_v12;
        _1915_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1914_v11,p0);
        let _1916_v13;
        _1916_v13 = _dafny.Map.Empty.slice().updateUnsafe(!(p1),_this.f14);
        let _1917_v14;
        _1917_v14 = _dafny.Set.fromElements((_1912_v9).f20, (_this).f17);
        let _1918_v15;
        _1918_v15 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_1914_v11.f16);
        let _1919_v16;
        _1919_v16 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).fm16(_1918_v15, globalState)),_1905_v2);
        let _1920_v17;
        let _nw369 = Array((new BigNumber(24)).toNumber());
        _nw369[(_dafny.ZERO).toNumber()] = p0;
        _nw369[(_dafny.ONE).toNumber()] = _this.f14;
        _nw369[(new BigNumber(2)).toNumber()] = (_1912_v9).f19;
        _nw369[(new BigNumber(3)).toNumber()] = _this.f14;
        _nw369[(new BigNumber(4)).toNumber()] = _this.f14;
        _nw369[(new BigNumber(5)).toNumber()] = _this.f14;
        _nw369[(new BigNumber(6)).toNumber()] = _this.f14;
        _nw369[(new BigNumber(7)).toNumber()] = new BigNumber(-166);
        _nw369[(new BigNumber(8)).toNumber()] = new BigNumber((_1915_v12).length);
        _nw369[(new BigNumber(9)).toNumber()] = p0;
        _nw369[(new BigNumber(10)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("pxtakl")).length);
        _nw369[(new BigNumber(11)).toNumber()] = (_1912_v9).f19;
        _nw369[(new BigNumber(12)).toNumber()] = new BigNumber((_1916_v13).length);
        _nw369[(new BigNumber(13)).toNumber()] = (_1912_v9).f19;
        _nw369[(new BigNumber(14)).toNumber()] = new BigNumber((_1917_v14).length);
        _nw369[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus(_this.f14);
        _nw369[(new BigNumber(16)).toNumber()] = _this.f14;
        _nw369[(new BigNumber(17)).toNumber()] = p0;
        _nw369[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus(_this.f14);
        _nw369[(new BigNumber(19)).toNumber()] = _this.f14;
        _nw369[(new BigNumber(20)).toNumber()] = _this.f14;
        _nw369[(new BigNumber(21)).toNumber()] = new BigNumber((_1919_v16).length);
        _nw369[(new BigNumber(22)).toNumber()] = p0;
        _nw369[(new BigNumber(23)).toNumber()] = new BigNumber(((_1909_v6)[_module.__default.safeIndex(new BigNumber(415), new BigNumber((_1909_v6).length))]).length);
        _1920_v17 = _nw369;
        let _index315 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_1913_v10).length));
        (_1913_v10)[_index315] = _1920_v17;
        let _1921_v18;
        _1921_v18 = _dafny.Seq.of((_1912_v9).f19);
        let _index316 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_1913_v10).length));
        let _rhs347 = (_module.D9.create_DC15(true, (_1912_v9).f20, function () {
  let _coll72 = new _dafny.Set();
  for (const _compr_72 of (_1921_v18).Elements) {
    let _1922_v19 = _compr_72;
    if (_dafny.Seq.contains(_1921_v18, _1922_v19)) {
      _coll72.add(_module.__default.safeModuloInt(_1922_v19, new BigNumber(82)));
    }
  }
  return _coll72;
}(), _1920_v17, _1903_v0)).dtor_cf21;
        let _rhs348 = (_1912_v9).f20;
        let _lhs267 = _1913_v10;
        let _lhs268 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_1913_v10).length));
        let _lhs269 = _1914_v11;
        _lhs267[_lhs268] = _rhs347;
        _lhs269.f16 = _rhs348;
      } else {
        let _1923_v20;
        _1923_v20 = _dafny.Seq.of((_dafny.ZERO).minus((_dafny.ZERO).minus(p0)), p0);
        (globalState).f2 = new BigNumber((_1923_v20).length);
        let _1924_v21;
        _1924_v21 = _dafny.MultiSet.fromElements((_this).f17);
        let _1925_v22;
        _1925_v22 = _dafny.Seq.of(false, (_this).f17);
        let _1926_v23;
        _1926_v23 = _module.D16.create_DC32(_1924_v21);
        let _1927_v24;
        _1927_v24 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _1928_v25;
        let _nw370 = Array((new BigNumber(23)).toNumber());
        _nw370[(_dafny.ZERO).toNumber()] = _1924_v21;
        _nw370[(_dafny.ONE).toNumber()] = (_1924_v21).Intersect(_dafny.MultiSet.FromArray(_1925_v22));
        _nw370[(new BigNumber(2)).toNumber()] = (_module.__default.fm1(globalState)).Union(_1924_v21);
        _nw370[(new BigNumber(3)).toNumber()] = (_dafny.MultiSet.FromArray(_1925_v22)).Difference(_1924_v21);
        _nw370[(new BigNumber(4)).toNumber()] = _1924_v21;
        _nw370[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.fromElements(_this.f16, true);
        _nw370[(new BigNumber(6)).toNumber()] = _1924_v21;
        _nw370[(new BigNumber(7)).toNumber()] = _1924_v21;
        _nw370[(new BigNumber(8)).toNumber()] = _1924_v21;
        _nw370[(new BigNumber(9)).toNumber()] = _module.__default.fm1(globalState);
        _nw370[(new BigNumber(10)).toNumber()] = _1924_v21;
        _nw370[(new BigNumber(11)).toNumber()] = (_1924_v21).Difference(_dafny.MultiSet.fromElements(false));
        _nw370[(new BigNumber(12)).toNumber()] = _1924_v21;
        _nw370[(new BigNumber(13)).toNumber()] = _1924_v21;
        _nw370[(new BigNumber(14)).toNumber()] = _1924_v21;
        _nw370[(new BigNumber(15)).toNumber()] = _1924_v21;
        _nw370[(new BigNumber(16)).toNumber()] = _1924_v21;
        _nw370[(new BigNumber(17)).toNumber()] = _1924_v21;
        _nw370[(new BigNumber(18)).toNumber()] = (_1926_v23).dtor_cf46;
        _nw370[(new BigNumber(19)).toNumber()] = (_1924_v21).Difference(_1924_v21);
        _nw370[(new BigNumber(20)).toNumber()] = (_1924_v21).update((((_1927_v24).contains(_this.f16)) ? ((_1927_v24).get(_this.f16)) : (_this.f16)), _module.__default.abs(_this.f14));
        _nw370[(new BigNumber(21)).toNumber()] = (_dafny.MultiSet.fromElements(_module.__default.fm21(globalState), p1, _this.f16)).update(_this.f16, _module.__default.abs((_dafny.ZERO).minus((_1923_v20)[_module.__default.safeIndex(_this.f14, new BigNumber((_1923_v20).length))])));
        _nw370[(new BigNumber(22)).toNumber()] = (_1926_v23).dtor_cf46;
        _1928_v25 = _nw370;
        let _index317 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_1928_v25).length));
        (_1928_v25)[_index317] = (_1924_v21).Union(_1924_v21);
        let _index318 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_1928_v25).length));
        (_1928_v25)[_index318] = _dafny.MultiSet.fromElements(!(false) || (_this.f16));
        let _1929_v26;
        _1929_v26 = _1923_v20;
        _1929_v26 = _1929_v26;
        let _1930_v27;
        _1930_v27 = _dafny.Seq.UnicodeFromString("vhiffetkh");
        (_this).f16 = !_dafny.areEqual(_1930_v27, _1930_v27);
        (_this).f16 = _this.f16;
      }
      let _1931_v28;
      let _nw371 = Array((new BigNumber(10)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _1931_v28 = _nw371;
      let _1932_v29;
      _1932_v29 = _dafny.Seq.of(_1931_v28);
      _1932_v29 = _1932_v29;
      let _1933_v30;
      _1933_v30 = _dafny.Set.fromElements(_this.f16, (_this).f17);
      let _hi15 = new BigNumber((_1933_v30).length);
      for (let _1934_i0 = _this.f14; _1934_i0.isLessThan(_hi15); _1934_i0 = _1934_i0.plus(_dafny.ONE)) {
        let _1935_v31;
        _1935_v31 = _dafny.Seq.of(p1);
        let _1936_v32;
        _1936_v32 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_dafny.ZERO).minus(_1934_i0));
        let _1937_v33;
        _1937_v33 = new _dafny.CodePoint('h'.codePointAt(0));
        let _1938_v34;
        let _nw372 = Array((new BigNumber(10)).toNumber());
        _nw372[(_dafny.ZERO).toNumber()] = p1;
        _nw372[(_dafny.ONE).toNumber()] = _dafny.Seq.contains(_1935_v31, _this.f16);
        _nw372[(new BigNumber(2)).toNumber()] = _module.__default.fm43(p1, globalState);
        _nw372[(new BigNumber(3)).toNumber()] = (((_this).f17) ? (_this.f16) : (p1));
        _nw372[(new BigNumber(4)).toNumber()] = _this.f16;
        _nw372[(new BigNumber(5)).toNumber()] = _this.f16;
        _nw372[(new BigNumber(6)).toNumber()] = !((_this).f17);
        _nw372[(new BigNumber(7)).toNumber()] = p1;
        _nw372[(new BigNumber(8)).toNumber()] = !(_1934_i0).isEqualTo(new BigNumber((_1936_v32).length));
        _nw372[(new BigNumber(9)).toNumber()] = !_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-425)), ((_1939_v33) => function (_1940_i1) {
          return _1939_v33;
        })(_1937_v33)), _1937_v33);
        _1938_v34 = _nw372;
        let _1941_v35;
        _1941_v35 = _dafny.Seq.UnicodeFromString("rlmtgv");
        let _1942_v36;
        _1942_v36 = _dafny.MultiSet.fromElements(_1941_v35, _1941_v35);
        let _index319 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_1938_v34).length));
        (_1938_v34)[_index319] = (_1942_v36).IsSubsetOf(_1942_v36);
        let _index320 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_1938_v34).length));
        (_1938_v34)[_index320] = ((_this).f17) === (_this.f16);
        _1937_v33 = _1937_v33;
        (_this).f16 = p1;
        let _1943_v37;
        let _nw373 = Array((new BigNumber(12)).toNumber()).fill(_dafny.MultiSet.Empty);
        _1943_v37 = _nw373;
        let _1944_v39;
        _1944_v39 = _dafny.Seq.of(new BigNumber(423), p0);
        let _1945_v40;
        _1945_v40 = _dafny.MultiSet.fromElements(function () {
          let _coll73 = new _dafny.Map();
          for (const _compr_73 of (_1944_v39).Elements) {
            let _1946_v38 = _compr_73;
            if (_dafny.Seq.contains(_1944_v39, _1946_v38)) {
              _coll73.push([_module.__default.safeDivisionInt(_1946_v38, new BigNumber((_1936_v32).length)),_module.D11.create_DC20()]);
            }
          }
          return _coll73;
        }());
        let _index321 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_1943_v37).length));
        (_1943_v37)[_index321] = (_1945_v40).Difference(_1945_v40);
        let _index322 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_1943_v37).length));
        (_1943_v37)[_index322] = ((_1945_v40).Difference(_1945_v40)).Difference(_1945_v40);
      }
      let _1947_v41;
      _1947_v41 = _dafny.Seq.UnicodeFromString("intkfrwlm");
      _1947_v41 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lmngcvt"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("iwta"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(277)), function (_1948_i2) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      })));
      let _hi16 = _this.f14;
      for (let _1949_i3 = _this.f14; _1949_i3.isLessThan(_hi16); _1949_i3 = _1949_i3.plus(_dafny.ONE)) {
        let _1950_v42;
        _1950_v42 = _dafny.Seq.of(new BigNumber(392));
        let _1951_v43;
        _1951_v43 = _dafny.MultiSet.fromElements(_dafny.Seq.of(p0, (_dafny.ZERO).minus(_this.f14)), _1950_v42);
        let _1952_v44;
        _1952_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1951_v43,(_1950_v42)[_module.__default.safeIndex(p0, new BigNumber((_1950_v42).length))]);
        _1952_v44 = (_1952_v44).update(_1951_v43, p0);
        (_this).f16 = (_module.__default.fm0((_this).f17, globalState)).isLessThan((_dafny.ZERO).minus(_1949_i3));
        (_this).f16 = true;
        (_this).f14 = _module.__default.safeDivisionInt(new BigNumber(677), _1949_i3);
      }
      _1947_v41 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(262)), function (_1953_i4) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      });
      r0 = _this.f14;
      return r0;
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _1954_v0;
      let _nw374 = Array((new BigNumber(24)).toNumber()).fill([]);
      _1954_v0 = _nw374;
      let _1955_v1;
      _1955_v1 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
      let _1956_v2;
      _1956_v2 = _dafny.Seq.UnicodeFromString("cnmueq");
      let _1957_v3;
      let _nw375 = Array((new BigNumber(17)).toNumber());
      _nw375[(_dafny.ZERO).toNumber()] = !(p2);
      _nw375[(_dafny.ONE).toNumber()] = true;
      _nw375[(new BigNumber(2)).toNumber()] = (_this).f17;
      _nw375[(new BigNumber(3)).toNumber()] = _module.__default.fm27(_this.f14, p1, _this.f14, (_1955_v1).update(new BigNumber(131), new BigNumber((_1956_v2).length)), globalState);
      _nw375[(new BigNumber(4)).toNumber()] = _module.__default.fm21(globalState);
      _nw375[(new BigNumber(5)).toNumber()] = true;
      _nw375[(new BigNumber(6)).toNumber()] = p2;
      _nw375[(new BigNumber(7)).toNumber()] = _this.f16;
      _nw375[(new BigNumber(8)).toNumber()] = p2;
      _nw375[(new BigNumber(9)).toNumber()] = (_this).f17;
      _nw375[(new BigNumber(10)).toNumber()] = (_this).f17;
      _nw375[(new BigNumber(11)).toNumber()] = _this.f16;
      _nw375[(new BigNumber(12)).toNumber()] = !((_this).f17);
      _nw375[(new BigNumber(13)).toNumber()] = _this.f16;
      _nw375[(new BigNumber(14)).toNumber()] = _this.f16;
      _nw375[(new BigNumber(15)).toNumber()] = false;
      _nw375[(new BigNumber(16)).toNumber()] = p2;
      _1957_v3 = _nw375;
      let _1958_v4;
      _1958_v4 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_1957_v3);
      let _index323 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_1954_v0).length));
      (_1954_v0)[_index323] = (((_1958_v4).contains((_this).f17)) ? ((_1958_v4).get((_this).f17)) : (_1957_v3));
      let _index324 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_1954_v0).length));
      (_1954_v0)[_index324] = _1957_v3;
      let _ingredients1 = [];
      for (const _guard_loop_11 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber(((_1954_v0)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_1954_v0).length))]).length))) {
        let _1959_i0 = _guard_loop_11;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1959_i0)) && ((_1959_i0).isLessThan(new BigNumber(((_1954_v0)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_1954_v0).length))]).length))))) {
          _ingredients1.push(_dafny.Tuple.of((_1954_v0)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_1954_v0).length))], (_1959_i0).toNumber(), _this.f16));
        }
      }
      for (const _tup1 of _ingredients1) {
        _tup1[0][_tup1[1]] = _tup1[2];
      }
      let _ingredients2 = [];
      for (const _guard_loop_12 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber(((_1954_v0)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_1954_v0).length))]).length))) {
        let _1960_i1 = _guard_loop_12;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1960_i1)) && ((_1960_i1).isLessThan(new BigNumber(((_1954_v0)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_1954_v0).length))]).length))))) {
          _ingredients2.push(_dafny.Tuple.of((_1954_v0)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_1954_v0).length))], (_1960_i1).toNumber(), true));
        }
      }
      for (const _tup2 of _ingredients2) {
        _tup2[0][_tup2[1]] = _tup2[2];
      }
      let _1961_v5;
      let _nw376 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
      _1961_v5 = _nw376;
      _1961_v5 = _1961_v5;
      (_this).f14 = (_dafny.ZERO).minus((new BigNumber((_1956_v2).length)).plus(p1));
      _1956_v2 = ((p2) ? (_dafny.Seq.UnicodeFromString("ukilfkar")) : (_1956_v2));
      let _1962_v6;
      _1962_v6 = _dafny.Seq.of(_this.f16);
      r0 = (new BigNumber((_1962_v6).length)).multipliedBy(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("eemycykf"), _module.__default.safeIndex(p3, new BigNumber((_dafny.Seq.UnicodeFromString("eemycykf")).length)), new _dafny.CodePoint('w'.codePointAt(0)))).length), _this.f14));
      return r0;
    }
    m9(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Set.Empty;
      let r2 = false;
      let r3 = _dafny.ZERO;
      let _1963_v0;
      let _nw377 = Array((new BigNumber(2)).toNumber()).fill([]);
      _1963_v0 = _nw377;
      let _1964_v1;
      let _nw378 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.of());
      _1964_v1 = _nw378;
      let _index325 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_1963_v0).length));
      (_1963_v0)[_index325] = _1964_v1;
      let _1965_v2;
      let _nw379 = new _module.C4();
      _nw379.__ctor(_this.f16, _this.f16, (_this).f17);
      _1965_v2 = _nw379;
      let _1966_v3;
      _1966_v3 = _dafny.Set.fromElements(_1965_v2, _1965_v2, _1965_v2, ((false) ? (_1965_v2) : (_1965_v2)));
      let _1967_v4;
      _1967_v4 = _dafny.Set.fromElements(_this.f14);
      let _index326 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_1963_v0).length));
      let _rhs349 = _1964_v1;
      let _rhs350 = !((_1967_v4).IsProperSubsetOf(_1967_v4));
      let _rhs351 = _1966_v3;
      let _lhs270 = _1963_v0;
      let _lhs271 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_1963_v0).length));
      _lhs270[_lhs271] = _rhs349;
      r0 = _rhs350;
      _1966_v3 = _rhs351;
      let _1968_v5;
      let _init61 = function (_1969_i1) {
        return (_1969_i1).plus(new BigNumber((_dafny.Seq.UnicodeFromString("tjeefai")).length));
      };
      let _nw380 = Array((new BigNumber(11)).toNumber());
      for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw380.length); _i0_61++) {
        _nw380[_i0_61] = _init61(new BigNumber(_i0_61));
      }
      _1968_v5 = _nw380;
      for (const _guard_loop_13 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1968_v5).length))) {
        let _1970_i0 = _guard_loop_13;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1970_i0)) && ((_1970_i0).isLessThan(new BigNumber((_1968_v5).length))))) {
          (_1968_v5)[(_1970_i0)] = (_1970_i0).minus(_module.__default.safeDivisionInt(new BigNumber(163), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!((_1965_v2).f21),_this.f14)).length)));
        }
      }
      let _1971_v6;
      _1971_v6 = _dafny.Seq.UnicodeFromString("r");
      _1971_v6 = _1971_v6;
      let _1972_v7;
      _1972_v7 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_1971_v6).length)), _this.f14, _this.f14, _this.f14, _this.f14);
      let _1973_v8;
      _1973_v8 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(false, globalState),_dafny.Seq.update(_1972_v7, _module.__default.safeIndex(new BigNumber(347), new BigNumber((_1972_v7).length)), _this.f14));
      if (_dafny.Seq.IsProperPrefixOf((((_1973_v8).contains(_module.__default.fm0((_this).f17, globalState))) ? ((_1973_v8).get(_module.__default.fm0((_this).f17, globalState))) : (_1972_v7)), (((_1973_v8).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f16,_dafny.Seq.of(_this.f14, _this.f14))).length))) ? ((_1973_v8).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f16,_dafny.Seq.of(_this.f14, _this.f14))).length))) : (_1972_v7)))) {
        let _1974_v9;
        _1974_v9 = new _dafny.CodePoint('j'.codePointAt(0));
        (globalState).f7 = (_dafny.ZERO).minus((new BigNumber((_dafny.Seq.update(_1971_v6, _module.__default.safeIndex(_this.f14, new BigNumber((_1971_v6).length)), _1974_v9)).length)).minus(_this.f14));
        let _1975_v10;
        let _init62 = ((_1976_v2) => function (_1977_i2) {
          return (_1976_v2).f21;
        })(_1965_v2);
        let _nw381 = Array((new BigNumber(4)).toNumber());
        for (let _i0_62 = 0; _i0_62 < new BigNumber(_nw381.length); _i0_62++) {
          _nw381[_i0_62] = _init62(new BigNumber(_i0_62));
        }
        _1975_v10 = _nw381;
        let _1978_v11;
        _1978_v11 = _dafny.MultiSet.fromElements(_1975_v10);
        let _1979_v12;
        let _nw382 = new _module.C6();
        _nw382.__ctor(_this.f14, (_1965_v2).f21, !((_1965_v2).f21), ((_1978_v11).update(_1975_v10, _module.__default.abs(new BigNumber((_1971_v6).length)))).IsSubsetOf((_1978_v11).update(_1975_v10, _module.__default.abs(_this.f14))), _this.f14);
        _1979_v12 = _nw382;
        let _1980_v13;
        _1980_v13 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_1975_v10);
        r2 = !((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_1980_v13).length)),_1968_v5)).contains((_1979_v12).f19));
        r3 = _this.f14;
        let _1981_v14;
        let _out33;
        _out33 = (_1979_v12).m3(_1974_v9, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(857)), ((_1982_v12) => function (_1983_i3) {
          return (_1982_v12).f19;
        })(_1979_v12))).length), (_this).f17, _this.f14, globalState);
        _1981_v14 = _out33;
      } else {
        let _1984_v15;
        _1984_v15 = new _dafny.CodePoint('m'.codePointAt(0));
        _1971_v6 = _dafny.Seq.update(_1971_v6, _module.__default.safeIndex(new BigNumber(455), new BigNumber((_1971_v6).length)), _1984_v15);
        let _1985_v16;
        _1985_v16 = _module.D9.create_DC16(_this.f14, _this.f16);
        let _1986_v17;
        _1986_v17 = _module.D9.create_DC17(_1985_v16);
        let _1987_v18;
        _1987_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1986_v17,_this.f16);
        let _rhs352 = _this.f14;
        let _rhs353 = !(_this.f14).isEqualTo(new BigNumber((_1987_v18).length));
        let _rhs354 = !((_this).f17);
        let _rhs355 = _this.f14;
        let _lhs272 = globalState;
        let _lhs273 = _this;
        let _lhs274 = globalState;
        _lhs272.f7 = _rhs352;
        r2 = _rhs353;
        _lhs273.f16 = _rhs354;
        _lhs274.f7 = _rhs355;
        (globalState).f2 = _this.f14;
        let _1988_v19;
        let _nw383 = Array((new BigNumber(7)).toNumber()).fill(false);
        _1988_v19 = _nw383;
        let _1989_v20;
        _1989_v20 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_1988_v19);
        _1989_v20 = (_1989_v20).Merge(_1989_v20);
        let _1990_v21;
        _1990_v21 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_this.f16);
        r2 = !((((_1990_v21).contains((_this).f17)) ? ((_1990_v21).get((_this).f17)) : ((_1965_v2).f21)));
      }
      let _1991_v22;
      _1991_v22 = _dafny.Seq.of((_1965_v2).f21, _this.f16, false, _this.f16, (_1965_v2).f21);
      let _hi17 = _this.f14;
      for (let _1992_i4 = new BigNumber((_dafny.Seq.update(_1991_v22, _module.__default.safeIndex(_this.f14, new BigNumber((_1991_v22).length)), (_this).f17)).length); _1992_i4.isLessThan(_hi17); _1992_i4 = _1992_i4.plus(_dafny.ONE)) {
        r2 = (_1965_v2).f21;
        let _1993_v23;
        _1993_v23 = _dafny.Map.Empty.slice().updateUnsafe(false,_this.f14);
        let _1994_v24;
        let _nw384 = new _module.C3();
        _nw384.__ctor(new BigNumber((_1971_v6).length), _dafny.Seq.UnicodeFromString("blldidk"), (_1965_v2).f21, (_this).f17);
        _1994_v24 = _nw384;
        let _1995_v25;
        let _nw385 = Array((new BigNumber(6)).toNumber());
        _nw385[(_dafny.ZERO).toNumber()] = _1994_v24;
        _nw385[(_dafny.ONE).toNumber()] = _1994_v24;
        _nw385[(new BigNumber(2)).toNumber()] = _1994_v24;
        _nw385[(new BigNumber(3)).toNumber()] = _1994_v24;
        _nw385[(new BigNumber(4)).toNumber()] = _1994_v24;
        _nw385[(new BigNumber(5)).toNumber()] = _1994_v24;
        _1995_v25 = _nw385;
        let _index327 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_1995_v25).length));
        (_1995_v25)[_index327] = _1994_v24;
        let _index328 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_1995_v25).length));
        let _rhs356 = _1968_v5;
        let _rhs357 = _1993_v23;
        let _rhs358 = _1971_v6;
        let _rhs359 = _1994_v24;
        let _lhs275 = _1995_v25;
        let _lhs276 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_1995_v25).length));
        _1968_v5 = _rhs356;
        _1993_v23 = _rhs357;
        _1971_v6 = _rhs358;
        _lhs275[_lhs276] = _rhs359;
        let _1996_v26;
        _1996_v26 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_1992_i4);
        let _1997_v28;
        _1997_v28 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_1965_v2).f21, globalState),(_this).f17);
        let _1998_v29;
        _1998_v29 = _module.D11.create_DC19((_1996_v26).Merge(function () {
  let _coll74 = new _dafny.Map();
  for (const _compr_74 of (_1997_v28).Keys.Elements) {
    let _1999_v27 = _compr_74;
    if ((_1997_v28).contains(_1999_v27)) {
      _coll74.push([_module.__default.safeModuloInt(_1999_v27, _1992_i4),new BigNumber(504)]);
    }
  }
  return _coll74;
}()));
        let _source31 = _1998_v29;
        if (_source31.is_DC20) {
          let _2000_v30;
          _2000_v30 = _module.D3.create_DC3(_this.f16);
          let _2001_v31;
          _2001_v31 = _dafny.Map.Empty.slice().updateUnsafe((_1965_v2).f21,(_1965_v2).f21);
          let _2002_v32;
          let _nw386 = new _module.C2();
          _nw386.__ctor(function (_pat_let36_0) {
            return function (_2003_dt__update__tmp_h0) {
              return function (_pat_let37_0) {
                return function (_2004_dt__update_hcf3_h0) {
                  return _module.D3.create_DC3(_2004_dt__update_hcf3_h0);
                }(_pat_let37_0);
              }(!(!((_this).f17)));
            }(_pat_let36_0);
          }(_2000_v30), (_1965_v2).f21, ((((_2001_v31).contains((_1965_v2).f21)) ? ((_2001_v31).get((_1965_v2).f21)) : ((_1965_v2).f21))) || ((_this).f17));
          _2002_v32 = _nw386;
          let _2005_v33;
          _2005_v33 = _dafny.Map.Empty.slice().updateUnsafe((_1965_v2).f21,_1972_v7);
          let _index329 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_1968_v5).length));
          (_1968_v5)[_index329] = _1992_i4;
          let _2006_v34;
          _2006_v34 = _dafny.MultiSet.fromElements(true);
          let _2007_v35;
          _2007_v35 = new _dafny.CodePoint('y'.codePointAt(0));
          let _2008_v36;
          _2008_v36 = _dafny.MultiSet.fromElements((_1994_v24).f22, _1992_i4);
          let _2009_v38;
          let _nw387 = Array((new BigNumber(22)).toNumber());
          _nw387[(_dafny.ZERO).toNumber()] = (_1965_v2).f21;
          _nw387[(_dafny.ONE).toNumber()] = _this.f16;
          _nw387[(new BigNumber(2)).toNumber()] = false;
          _nw387[(new BigNumber(3)).toNumber()] = (_this).f17;
          _nw387[(new BigNumber(4)).toNumber()] = (_1965_v2).f21;
          _nw387[(new BigNumber(5)).toNumber()] = _this.f16;
          _nw387[(new BigNumber(6)).toNumber()] = true;
          _nw387[(new BigNumber(7)).toNumber()] = (_1965_v2).f21;
          _nw387[(new BigNumber(8)).toNumber()] = _this.f16;
          _nw387[(new BigNumber(9)).toNumber()] = false;
          _nw387[(new BigNumber(10)).toNumber()] = (_1965_v2).f21;
          _nw387[(new BigNumber(11)).toNumber()] = _this.f16;
          _nw387[(new BigNumber(12)).toNumber()] = (_1965_v2).f21;
          _nw387[(new BigNumber(13)).toNumber()] = (_this).f17;
          _nw387[(new BigNumber(14)).toNumber()] = (_1965_v2).f21;
          _nw387[(new BigNumber(15)).toNumber()] = (_1965_v2).f21;
          _nw387[(new BigNumber(16)).toNumber()] = _this.f16;
          _nw387[(new BigNumber(17)).toNumber()] = _module.__default.fm43((_1965_v2).f21, globalState);
          _nw387[(new BigNumber(18)).toNumber()] = (_1965_v2).f21;
          _nw387[(new BigNumber(19)).toNumber()] = (_this).f17;
          _nw387[(new BigNumber(20)).toNumber()] = (_1965_v2).f21;
          _nw387[(new BigNumber(21)).toNumber()] = _this.f16;
          _2009_v38 = _nw387;
          let _2010_v39;
          _2010_v39 = _module.D9.create_DC15((_this).f17, (_1965_v2).f21, function () {
  let _coll75 = new _dafny.Set();
  for (const _compr_75 of (_1967_v4).Elements) {
    let _2011_v37 = _compr_75;
    if ((_1967_v4).contains(_2011_v37)) {
      _coll75.add(_module.__default.safeDivisionInt(_2011_v37, new BigNumber(891)));
    }
  }
  return _coll75;
}(), _1968_v5, _2009_v38);
          let _2012_v40;
          let _nw388 = Array((new BigNumber(6)).toNumber());
          _nw388[(_dafny.ZERO).toNumber()] = _2010_v39;
          _nw388[(_dafny.ONE).toNumber()] = _module.D9.create_DC15((_1991_v22)[_module.__default.safeIndex(new BigNumber((_1996_v26).length), new BigNumber((_1991_v22).length))], _module.__default.fm31((_1965_v2).f21, (_1965_v2).f21, globalState), _1967_v4, _1968_v5, _2009_v38);
          _nw388[(new BigNumber(2)).toNumber()] = _2010_v39;
          _nw388[(new BigNumber(3)).toNumber()] = _2010_v39;
          _nw388[(new BigNumber(4)).toNumber()] = _2010_v39;
          _nw388[(new BigNumber(5)).toNumber()] = _2010_v39;
          _2012_v40 = _nw388;
          let _2013_v41;
          _2013_v41 = _module.D12.create_DC24((_1965_v2).f21, _1991_v22, _dafny.Seq.of(_2012_v40), (_1965_v2).f21, _module.__default.fm62(globalState));
          let _2014_v42;
          let _nw389 = Array((new BigNumber(19)).toNumber());
          _nw389[(_dafny.ZERO).toNumber()] = (_1991_v22)[_module.__default.safeIndex(new BigNumber(165), new BigNumber((_1991_v22).length))];
          _nw389[(_dafny.ONE).toNumber()] = (_1965_v2).f21;
          _nw389[(new BigNumber(2)).toNumber()] = (_dafny.MultiSet.FromArray(_1991_v22)).IsDisjointFrom(_2006_v34);
          _nw389[(new BigNumber(3)).toNumber()] = (_this).f17;
          _nw389[(new BigNumber(4)).toNumber()] = (_1965_v2).f21;
          _nw389[(new BigNumber(5)).toNumber()] = _module.__default.fm31((_this).f17, _this.f16, globalState);
          _nw389[(new BigNumber(6)).toNumber()] = !_dafny.Seq.contains((_1994_v24).f23, _2007_v35);
          _nw389[(new BigNumber(7)).toNumber()] = (_1965_v2).f21;
          _nw389[(new BigNumber(8)).toNumber()] = (_1965_v2).f21;
          _nw389[(new BigNumber(9)).toNumber()] = _this.f16;
          _nw389[(new BigNumber(10)).toNumber()] = (_this).f17;
          _nw389[(new BigNumber(11)).toNumber()] = !(_2008_v36).contains((_dafny.ZERO).minus(_this.f14));
          _nw389[(new BigNumber(12)).toNumber()] = _this.f16;
          _nw389[(new BigNumber(13)).toNumber()] = (_2013_v41).dtor_cf33;
          _nw389[(new BigNumber(14)).toNumber()] = (_1965_v2).f21;
          _nw389[(new BigNumber(15)).toNumber()] = ((_1994_v24).f22).isLessThan((_1965_v2).fm12(_this.f14, globalState));
          _nw389[(new BigNumber(16)).toNumber()] = (_this).f17;
          _nw389[(new BigNumber(17)).toNumber()] = true;
          _nw389[(new BigNumber(18)).toNumber()] = (_1965_v2).f21;
          _2014_v42 = _nw389;
          let _index330 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_2014_v42).length));
          (_2014_v42)[_index330] = false;
          let _2015_v43;
          _2015_v43 = _dafny.Set.fromElements(_this.f16);
          let _index331 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_1968_v5).length));
          let _index332 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_2014_v42).length));
          let _rhs360 = _dafny.Map.Empty.slice().updateUnsafe((_1965_v2).f21,_dafny.Seq.Concat(_1972_v7, _1972_v7));
          let _rhs361 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("h")).length), new BigNumber((_1972_v7).length));
          let _rhs362 = (_1965_v2).f21;
          let _rhs363 = ((_2015_v43).Union(_2015_v43)).IsProperSubsetOf(_2015_v43);
          let _lhs277 = _1968_v5;
          let _lhs278 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_1968_v5).length));
          let _lhs279 = _2014_v42;
          let _lhs280 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_2014_v42).length));
          _2005_v33 = _rhs360;
          _lhs277[_lhs278] = _rhs361;
          _lhs279[_lhs280] = _rhs362;
          r2 = _rhs363;
          let _2016_v44;
          _2016_v44 = _dafny.Map.Empty.slice().updateUnsafe((_1994_v24).f23,!((_1965_v2).f21));
          let _2017_v45;
          _2017_v45 = _module.D21.create_DC43(_dafny.MultiSet.fromElements(_2007_v35, _2007_v35));
          let _2018_v46;
          _2018_v46 = _dafny.MultiSet.fromElements(_2007_v35);
          let _2019_v47;
          _2019_v47 = _dafny.Map.Empty.slice().updateUnsafe(_2007_v35,_2018_v46);
          let _2020_v48;
          _2020_v48 = _dafny.Seq.of((_2017_v45).dtor_cf68, (((_2019_v47).contains(new _dafny.CodePoint('m'.codePointAt(0)))) ? ((_2019_v47).get(new _dafny.CodePoint('m'.codePointAt(0)))) : (_dafny.MultiSet.FromArray(_1971_v6))));
          let _2021_v49;
          _2021_v49 = _module.D21.create_DC44((_1994_v24).f22, (_1965_v2).f21);
          let _index333 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_2014_v42).length));
          let _rhs364 = (_1994_v24).f22;
          let _rhs365 = !(_2016_v44).contains((_1994_v24).f23);
          let _rhs366 = ((_2020_v48)[_module.__default.safeIndex(_1992_i4, new BigNumber((_2020_v48).length))]).IsDisjointFrom(_2018_v46);
          let _rhs367 = (_module.__default.safeModuloInt((_2021_v49).dtor_cf69, (_1994_v24).f22)).isLessThan((_1994_v24).f22);
          let _lhs281 = globalState;
          let _lhs282 = _2014_v42;
          let _lhs283 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_2014_v42).length));
          _lhs281.f2 = _rhs364;
          r2 = _rhs365;
          r0 = _rhs366;
          _lhs282[_lhs283] = _rhs367;
          let _init63 = ((_2022_v24, _2023_v26) => function (_2024_i5) {
            return ((((_2023_v26).contains(_this.f14)) ? ((_2023_v26).get(_this.f14)) : (_this.f14))).isLessThanOrEqualTo((_2022_v24).f22);
          })(_1994_v24, _1996_v26);
          let _nw390 = Array((new BigNumber(13)).toNumber());
          for (let _i0_63 = 0; _i0_63 < new BigNumber(_nw390.length); _i0_63++) {
            _nw390[_i0_63] = _init63(new BigNumber(_i0_63));
          }
          _2009_v38 = _nw390;
        } else if (_source31.is_DC19) {
          let _2025___mcc_h0 = (_source31).cf27;
          let _2026_cf27 = _2025___mcc_h0;
          let _index334 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_1968_v5).length));
          (_1968_v5)[_index334] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(812)), ((_2027_v24) => function (_2028_i6) {
            return (_2027_v24).f22;
          })(_1994_v24))).length);
          let _2029_v50;
          _2029_v50 = _dafny.MultiSet.fromElements(_this.f14);
          let _index335 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_1968_v5).length));
          (_1968_v5)[_index335] = new BigNumber((_2029_v50).cardinality());
          (globalState).f7 = (_1994_v24).f22;
          let _2030_v51;
          let _out34;
          _out34 = (_1965_v2).m3(((_1994_v24).f23)[_module.__default.safeIndex(_1992_i4, new BigNumber(((_1994_v24).f23).length))], _this.f14, (_1965_v2).f21, _1992_i4, globalState);
          _2030_v51 = _out34;
          let _2031_v52;
          let _init64 = function (_2032_i7) {
            return (_this).f17;
          };
          let _nw391 = Array((new BigNumber(28)).toNumber());
          for (let _i0_64 = 0; _i0_64 < new BigNumber(_nw391.length); _i0_64++) {
            _nw391[_i0_64] = _init64(new BigNumber(_i0_64));
          }
          _2031_v52 = _nw391;
          let _2033_v53;
          _2033_v53 = _dafny.Set.fromElements(_2031_v52);
          _2026_cf27 = (_2026_cf27).update(new BigNumber(22), new BigNumber((_2033_v53).length));
        } else {
          let _2034___mcc_h1 = (_source31).cf28;
          let _2035_cf28 = _2034___mcc_h1;
          let _2036_v54;
          _2036_v54 = new _dafny.CodePoint('v'.codePointAt(0));
          let _2037_v55;
          _2037_v55 = _dafny.MultiSet.fromElements(_2036_v54, _2036_v54, _2036_v54);
          r2 = _module.__default.fm27((((_2037_v55).contains(_2036_v54)) ? ((_2037_v55).get(_2036_v54)) : ((((_2037_v55).contains(_2036_v54)) ? ((_2037_v55).get(_2036_v54)) : (_this.f14)))), (_1994_v24).f22, new BigNumber(676), _dafny.Map.Empty.slice().updateUnsafe((_1994_v24).f22,(_1994_v24).f22), globalState);
          let _2038_v56;
          _2038_v56 = _module.D3.create_DC3((_1965_v2).f21);
          let _2039_v57;
          let _nw392 = new _module.C2();
          _nw392.__ctor(_2038_v56, (_1965_v2).f21, (_1965_v2).f21);
          _2039_v57 = _nw392;
          let _2040_v58;
          _2040_v58 = _dafny.Map.Empty.slice().updateUnsafe((_1991_v22)[_module.__default.safeIndex((_1994_v24).f22, new BigNumber((_1991_v22).length))],_2039_v57);
          _2039_v57 = (((_1965_v2).f21) ? (_2039_v57) : ((((_1965_v2).f21) ? ((((_2040_v58).contains((_1965_v2).f21)) ? ((_2040_v58).get((_1965_v2).f21)) : (_2039_v57))) : (_2039_v57))));
          (_this).f16 = (_2038_v56).dtor_cf3;
          r0 = !((_this.f14).isLessThan((_1994_v24).f22));
        }
        let _2041_v59;
        let _nw393 = Array((new BigNumber(2)).toNumber());
        _nw393[(_dafny.ZERO).toNumber()] = (_1996_v26).Merge(_1996_v26);
        _nw393[(_dafny.ONE).toNumber()] = (((_this).f17) ? (_1996_v26) : (_1996_v26));
        _2041_v59 = _nw393;
        let _index336 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_2041_v59).length));
        (_2041_v59)[_index336] = _1996_v26;
        let _index337 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_2041_v59).length));
        (_2041_v59)[_index337] = (_1994_v24).fm3(globalState);
      }
      let _2042_v60;
      _2042_v60 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,(_this).f17);
      let _2043_v61;
      _2043_v61 = _dafny.MultiSet.fromElements(_this.f16);
      let _2044_v62;
      _2044_v62 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,new BigNumber((_1991_v22).length));
      let _2045_v63;
      _2045_v63 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_1972_v7).length));
      let _2046_v64;
      _2046_v64 = _module.D3.create_DC4((_this).f17, _this.f14, new BigNumber((_2045_v63).length), _this.f16, false);
      let _2047_v65;
      _2047_v65 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,_module.__default.fm28(_this.f16, _2044_v62, _2046_v64, _2042_v60, globalState));
      let _2048_v66;
      let _nw394 = new _module.C6();
      _nw394.__ctor(_this.f14, !(_2043_v61).contains((((_2042_v60).contains(_this.f16)) ? ((_2042_v60).get(_this.f16)) : (!((_1965_v2).f21)))), _this.f16, ((_2047_v65).update((_1965_v2).f21, (((_2047_v65).contains(true)) ? ((_2047_v65).get(true)) : (_1971_v6)))).equals(_2047_v65), new BigNumber(375));
      _2048_v66 = _nw394;
      r0 = (_1965_v2).f21;
      let _2049_v67;
      _2049_v67 = _1991_v22;
      let _pat_let_tv28 = _2048_v66;
      let _pat_let_tv29 = _2042_v60;
      let _pat_let_tv30 = _2048_v66;
      let _pat_let_tv31 = _2042_v60;
      let _pat_let_tv32 = _2048_v66;
      r1 = function (_source32) {
        let _2050___mcc_h2 = _source32;
        let _2051_cf0 = _2050___mcc_h2;
        return _dafny.Set.fromElements((_this).f17, (_pat_let_tv28).f20, !(!(_this.f16)), false, !((((_pat_let_tv31).contains((_pat_let_tv32).f20)) ? ((_pat_let_tv29).get((_pat_let_tv30).f20)) : ((_this).f17))));
      }(_2049_v67);
      r2 = (_2048_v66).f20;
      r3 = _this.f14;
      return [r0, r1, r2, r3];
    }
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
      this._f16 = false;
      this._f14 = _dafny.ZERO;
      this._f17 = false;
      this.f18 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T1, _module.T0, _module.T2];
    }
    get f16() {
      let _this = this;
      return _this._f16;
    };
    set f16(value) {
      let _this = this;
      _this._f16 = value;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    set f14(value) {
      let _this = this;
      _this._f14 = value;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    __ctor(f18, f16, f17, f14) {
      let _this = this;
      (_this).f18 = f18;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f14 = f14;
      return;
    }
    fm3(globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(88)), function (_2052_i0) {
        return _this.f14;
      })).length),_this.f14);
    };
    fm12(p0, globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus(_this.f14)).multipliedBy(((_this.f16) ? (new BigNumber((_dafny.Seq.UnicodeFromString("pbtospxwi")).length)) : (_this.f14)));
    };
    fm13(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Seq.of((_this).f17, _this.f16, (_this).f17);
    };
    fm14(globalState) {
      let _this = this;
      return _this.f16;
    };
    fm15(globalState) {
      let _this = this;
      return _dafny.MultiSet.fromElements(_this.f18, _this.f18, _this.f18, _dafny.Seq.Concat(_this.f18, _this.f18), _dafny.Seq.UnicodeFromString("kk"));
    };
    m2(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _2053_v0;
      _2053_v0 = _dafny.Seq.of(true, p1);
      let _2054_v1;
      let _nw395 = new _module.C4();
      _nw395.__ctor(_this.f16, _this.f16, (_2053_v0)[_module.__default.safeIndex(p0, new BigNumber((_2053_v0).length))]);
      _2054_v1 = _nw395;
      let _2055_v2;
      _2055_v2 = _dafny.Set.fromElements(_this.f14, p0, _this.f14, _this.f14, _module.__default.fm0((_this).f17, globalState));
      (_this).f16 = ((_module.__default.fm49(p0, p0, _this.f16, _this.f14, globalState)).Intersect(_2055_v2)).IsDisjointFrom(_dafny.Set.fromElements(p0));
      let _2056_v3;
      _2056_v3 = _dafny.Seq.of(_this.f14, _this.f14);
      let _2057_v4;
      _2057_v4 = _dafny.Seq.of(_2056_v3);
      let _2058_i0;
      _2058_i0 = _dafny.ZERO;
      L10: {
        while (!(!(_dafny.Seq.IsPrefixOf(_2056_v3, (_2057_v4)[_module.__default.safeIndex(new BigNumber((_this.f18).length), new BigNumber((_2057_v4).length))])))) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2058_i0)) {
              break L10;
            }
            _2058_i0 = (_2058_i0).plus(_dafny.ONE);
            (globalState).f7 = p0;
            let _2059_v5;
            _2059_v5 = _dafny.Set.fromElements((_2054_v1).f21, false, (_2054_v1).f21, _this.f16, _this.f16);
            let _2060_v6;
            let _nw396 = new _module.C1();
            _nw396.__ctor(_this.f16, p0);
            _2060_v6 = _nw396;
            let _2061_v7;
            _2061_v7 = _dafny.MultiSet.fromElements(_2056_v3);
            let _2062_v8;
            _2062_v8 = _module.D16.create_DC34(!(p1), _2060_v6, new BigNumber((_2061_v7).cardinality()));
            let _2063_v9;
            let _nw397 = Array((new BigNumber(28)).toNumber());
            _nw397[(_dafny.ZERO).toNumber()] = _this.f14;
            _nw397[(_dafny.ONE).toNumber()] = _this.f14;
            _nw397[(new BigNumber(2)).toNumber()] = p0;
            _nw397[(new BigNumber(3)).toNumber()] = (p0).multipliedBy(new BigNumber((_this.f18).length));
            _nw397[(new BigNumber(4)).toNumber()] = p0;
            _nw397[(new BigNumber(5)).toNumber()] = _this.f14;
            _nw397[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(p0);
            _nw397[(new BigNumber(7)).toNumber()] = new BigNumber((_2059_v5).length);
            _nw397[(new BigNumber(8)).toNumber()] = _module.__default.safeModuloInt(p0, p0);
            _nw397[(new BigNumber(9)).toNumber()] = p0;
            _nw397[(new BigNumber(10)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_this.f18, _this.f18)).length);
            _nw397[(new BigNumber(11)).toNumber()] = (_this.f14).minus(p0);
            _nw397[(new BigNumber(12)).toNumber()] = (p0).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(182)), function (_2064_i1) {
              return new _dafny.CodePoint('h'.codePointAt(0));
            })).length));
            _nw397[(new BigNumber(13)).toNumber()] = new BigNumber(652);
            _nw397[(new BigNumber(14)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of(_this.f16, p1)).length), (_dafny.ZERO).minus(new BigNumber((_2056_v3).length)));
            _nw397[(new BigNumber(15)).toNumber()] = _this.f14;
            _nw397[(new BigNumber(16)).toNumber()] = p0;
            _nw397[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_2056_v3).length));
            _nw397[(new BigNumber(18)).toNumber()] = _this.f14;
            _nw397[(new BigNumber(19)).toNumber()] = p0;
            _nw397[(new BigNumber(20)).toNumber()] = p0;
            _nw397[(new BigNumber(21)).toNumber()] = (p0).minus((_dafny.ZERO).minus(_this.f14));
            _nw397[(new BigNumber(22)).toNumber()] = new BigNumber(567);
            _nw397[(new BigNumber(23)).toNumber()] = (((_this).f17) ? (p0) : (new BigNumber(397)));
            _nw397[(new BigNumber(24)).toNumber()] = (_2062_v8).dtor_cf53;
            _nw397[(new BigNumber(25)).toNumber()] = _2060_v6.f14;
            _nw397[(new BigNumber(26)).toNumber()] = new BigNumber(672);
            _nw397[(new BigNumber(27)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_2053_v0, _2053_v0)).length);
            _2063_v9 = _nw397;
            let _index338 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_2063_v9).length));
            (_2063_v9)[_index338] = _2060_v6.f14;
            let _index339 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_2063_v9).length));
            let _rhs368 = (_dafny.ZERO).minus((_this.f14).minus(_2060_v6.f14));
            let _rhs369 = new BigNumber((_2053_v0).length);
            let _rhs370 = _2060_v6.f14;
            let _rhs371 = p1;
            let _lhs284 = _2063_v9;
            let _lhs285 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_2063_v9).length));
            let _lhs286 = globalState;
            let _lhs287 = _this;
            let _lhs288 = _this;
            _lhs284[_lhs285] = _rhs368;
            _lhs286.f2 = _rhs369;
            _lhs287.f14 = _rhs370;
            _lhs288.f16 = _rhs371;
            let _rhs372 = (_dafny.ZERO).minus((_2063_v9)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_2063_v9).length))]);
            let _rhs373 = _dafny.Seq.Concat(_2056_v3, _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(_module.__default.fm0((_2054_v1).f21, globalState)), _2056_v3), _module.__default.safeIndex((_2063_v9)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_2063_v9).length))], new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_module.__default.fm0((_2054_v1).f21, globalState)), _2056_v3)).length)), (_2063_v9)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_2063_v9).length))]));
            r0 = _rhs372;
            _2056_v3 = _rhs373;
            if (p1) {
              let _2065_v10;
              _2065_v10 = new _dafny.CodePoint('r'.codePointAt(0));
              let _2066_v11;
              _2066_v11 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_2065_v10, _2065_v10),(_2054_v1).f21);
              _2066_v11 = (_2066_v11).update(_dafny.Set.fromElements(new _dafny.CodePoint('w'.codePointAt(0)), _2065_v10, _2065_v10, _2065_v10, _2065_v10), _this.f16);
              (_this).f16 = (_this).f17;
              (globalState).f7 = new BigNumber((_dafny.Seq.UnicodeFromString("mkxdimuqr")).length);
              let _2067_v12;
              _2067_v12 = _dafny.Seq.of(_this.f18);
              let _2068_v13;
              _2068_v13 = _dafny.Seq.of((_2067_v12)[_module.__default.safeIndex(_this.f14, new BigNumber((_2067_v12).length))], _this.f18);
              let _2069_v14;
              let _nw398 = Array((new BigNumber(5)).toNumber());
              _nw398[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("dess"), _this.f18);
              _nw398[(_dafny.ONE).toNumber()] = _this.f18;
              _nw398[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jlwjpaav"), (_2068_v13)[_module.__default.safeIndex((_this).fm12(new BigNumber(-793), globalState), new BigNumber((_2068_v13).length))]);
              _nw398[(new BigNumber(3)).toNumber()] = _this.f18;
              _nw398[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_this.f18, _this.f18);
              _2069_v14 = _nw398;
              let _index340 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_2069_v14).length));
              (_2069_v14)[_index340] = _dafny.Seq.Concat(_this.f18, _this.f18);
              let _index341 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_2069_v14).length));
              (_2069_v14)[_index341] = _this.f18;
              let _2070_v15;
              _2070_v15 = _dafny.Map.Empty.slice().updateUnsafe((_2054_v1).f21,(_2069_v14)[_module.__default.safeIndex(new BigNumber(546), new BigNumber((_2069_v14).length))]);
              let _2071_v16;
              _2071_v16 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_2065_v10);
              let _2072_v17;
              _2072_v17 = _dafny.Map.Empty.slice().updateUnsafe((_2054_v1).f21,new BigNumber(((_2069_v14)[_module.__default.safeIndex(new BigNumber(546), new BigNumber((_2069_v14).length))]).length));
              let _rhs374 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_2063_v9)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_2063_v9).length))]));
              let _rhs375 = !(_module.__default.fm21(globalState)) || ((_2054_v1).f21);
              let _rhs376 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update((((_2070_v15).contains((_2054_v1).f21)) ? ((_2070_v15).get((_2054_v1).f21)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(352)), ((_2073_v10) => function (_2074_i2) {
                return _2073_v10;
              })(_2065_v10)))), _module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray(_2056_v3)).cardinality()), new BigNumber(((((_2070_v15).contains((_2054_v1).f21)) ? ((_2070_v15).get((_2054_v1).f21)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(352)), ((_2075_v10) => function (_2076_i2) {
                return _2075_v10;
              })(_2065_v10))))).length)), (((_2071_v16).contains(new BigNumber(556))) ? ((_2071_v16).get(new BigNumber(556))) : (_2065_v10))), _this.f18)).length));
              let _rhs377 = (((_2072_v17).equals(_2072_v17)) ? (new BigNumber(366)) : ((_2063_v9)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_2063_v9).length))]));
              let _lhs289 = globalState;
              let _lhs290 = _this;
              let _lhs291 = _2060_v6;
              let _lhs292 = _this;
              _lhs289.f7 = _rhs374;
              _lhs290.f16 = _rhs375;
              _lhs291.f14 = _rhs376;
              _lhs292.f14 = _rhs377;
            } else {
              let _2077_v18;
              let _nw399 = Array((new BigNumber(12)).toNumber()).fill(_module.D9.Default());
              _2077_v18 = _nw399;
              let _2078_v19;
              _2078_v19 = _dafny.Seq.of(_2077_v18);
              let _2079_v20;
              _2079_v20 = _module.D12.create_DC23(true, (_dafny.ZERO).minus((_2063_v9)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_2063_v9).length))]), p0);
              let _2080_v21;
              _2080_v21 = _dafny.MultiSet.fromElements((_2054_v1).f21, (_2079_v20).dtor_cf30);
              let _2081_v22;
              _2081_v22 = _dafny.Map.Empty.slice().updateUnsafe(_2060_v6.f14,_this.f14);
              let _2082_v23;
              _2082_v23 = _module.D3.create_DC3((_2054_v1).f21);
              let _2083_v24;
              _2083_v24 = _module.D12.create_DC24(_this.f16, _2053_v0, _2078_v19, (_2054_v1).f21, _dafny.Seq.of(_module.D3.create_DC3(_module.__default.fm27(new BigNumber((_2080_v21).cardinality()), (_2056_v3)[_module.__default.safeIndex(_this.f14, new BigNumber((_2056_v3).length))], _this.f14, _2081_v22, globalState)), _2082_v23));
              let _2084_v25;
              _2084_v25 = _dafny.Map.Empty.slice().updateUnsafe(_2083_v24,(_2054_v1).f21);
              let _2085_v26;
              _2085_v26 = _dafny.Map.Empty.slice().updateUnsafe(!(p1),_2084_v25);
              let _2086_v27;
              _2086_v27 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_2084_v25);
              let _2087_v28;
              let _nw400 = Array((new BigNumber(15)).toNumber());
              _nw400[(_dafny.ZERO).toNumber()] = _2084_v25;
              _nw400[(_dafny.ONE).toNumber()] = (_2084_v25).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2083_v24,_this.f16));
              _nw400[(new BigNumber(2)).toNumber()] = (_2084_v25).Merge(_2084_v25);
              _nw400[(new BigNumber(3)).toNumber()] = (_2084_v25).Merge(_2084_v25);
              _nw400[(new BigNumber(4)).toNumber()] = ((p1) ? ((((_2085_v26).contains((_2054_v1).f21)) ? ((_2085_v26).get((_2054_v1).f21)) : (_2084_v25))) : (_2084_v25));
              _nw400[(new BigNumber(5)).toNumber()] = _2084_v25;
              _nw400[(new BigNumber(6)).toNumber()] = (_2084_v25).Merge(_2084_v25);
              _nw400[(new BigNumber(7)).toNumber()] = (_2084_v25).Merge(_2084_v25);
              _nw400[(new BigNumber(8)).toNumber()] = _2084_v25;
              _nw400[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2083_v24,(_2053_v0)[_module.__default.safeIndex((_2063_v9)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_2063_v9).length))], new BigNumber((_2053_v0).length))]);
              _nw400[(new BigNumber(10)).toNumber()] = _2084_v25;
              _nw400[(new BigNumber(11)).toNumber()] = (((_2086_v27).contains(p0)) ? ((_2086_v27).get(p0)) : (_2084_v25));
              _nw400[(new BigNumber(12)).toNumber()] = _2084_v25;
              _nw400[(new BigNumber(13)).toNumber()] = _2084_v25;
              _nw400[(new BigNumber(14)).toNumber()] = _2084_v25;
              _2087_v28 = _nw400;
              let _nw401 = Array((new BigNumber(23)).toNumber());
              _nw401[(_dafny.ZERO).toNumber()] = _2084_v25;
              _nw401[(_dafny.ONE).toNumber()] = _2084_v25;
              _nw401[(new BigNumber(2)).toNumber()] = (_2084_v25).Merge(_2084_v25);
              _nw401[(new BigNumber(3)).toNumber()] = _2084_v25;
              _nw401[(new BigNumber(4)).toNumber()] = _2084_v25;
              _nw401[(new BigNumber(5)).toNumber()] = (_2084_v25).Merge(_2084_v25);
              _nw401[(new BigNumber(6)).toNumber()] = _2084_v25;
              _nw401[(new BigNumber(7)).toNumber()] = _2084_v25;
              _nw401[(new BigNumber(8)).toNumber()] = (_2084_v25).Merge(_2084_v25);
              _nw401[(new BigNumber(9)).toNumber()] = _2084_v25;
              _nw401[(new BigNumber(10)).toNumber()] = (_2084_v25).update(_2083_v24, (_2054_v1).f21);
              _nw401[(new BigNumber(11)).toNumber()] = ((_2084_v25).update(_2083_v24, (_2054_v1).f21)).update(_2083_v24, !((_this).f17));
              _nw401[(new BigNumber(12)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_2083_v24,(_2054_v1).f21)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2083_v24,_this.f16));
              _nw401[(new BigNumber(13)).toNumber()] = _2084_v25;
              _nw401[(new BigNumber(14)).toNumber()] = _2084_v25;
              _nw401[(new BigNumber(15)).toNumber()] = _2084_v25;
              _nw401[(new BigNumber(16)).toNumber()] = _2084_v25;
              _nw401[(new BigNumber(17)).toNumber()] = _2084_v25;
              _nw401[(new BigNumber(18)).toNumber()] = _2084_v25;
              _nw401[(new BigNumber(19)).toNumber()] = _2084_v25;
              _nw401[(new BigNumber(20)).toNumber()] = _2084_v25;
              _nw401[(new BigNumber(21)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2083_v24,(_this).f17);
              _nw401[(new BigNumber(22)).toNumber()] = _2084_v25;
              _2087_v28 = _nw401;
              (_this).f16 = (_2054_v1).f21;
              _2081_v22 = _2081_v22;
              let _2088_v29;
              let _nw402 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.of());
              _2088_v29 = _nw402;
              _2088_v29 = _2088_v29;
              (globalState).f2 = _module.__default.fm0((_2054_v1).f21, globalState);
            }
          }
        }
      }
      let _2089_v30;
      let _nw403 = Array((new BigNumber(25)).toNumber()).fill(_dafny.MultiSet.Empty);
      _2089_v30 = _nw403;
      let _index342 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_2089_v30).length));
      (_2089_v30)[_index342] = _module.__default.fm1(globalState);
      let _2090_v31;
      _2090_v31 = _dafny.MultiSet.fromElements(_this.f16);
      let _2091_v32;
      _2091_v32 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(p1),_2055_v2);
      let _2092_v33;
      _2092_v33 = _dafny.Set.fromElements((_this).f17, !((_2054_v1).f21), (_2054_v1).f21, (_2054_v1).f21);
      let _index343 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_2089_v30).length));
      let _rhs378 = (_2090_v31).update((new BigNumber(((_2091_v32).update(_2092_v33, _2055_v2)).length)).isLessThan(p0), _module.__default.abs(_this.f14));
      let _rhs379 = p0;
      let _lhs293 = _2089_v30;
      let _lhs294 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_2089_v30).length));
      _lhs293[_lhs294] = _rhs378;
      r0 = _rhs379;
      (globalState).f7 = _module.__default.safeModuloInt(new BigNumber(153), p0);
      let _2093_i3;
      _2093_i3 = _dafny.ZERO;
      L11: {
        while (false) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2093_i3)) {
              break L11;
            }
            _2093_i3 = (_2093_i3).plus(_dafny.ONE);
            (globalState).f2 = (_dafny.ZERO).minus((_dafny.ZERO).minus((p0).multipliedBy(p0)));
            (_this).f14 = (new BigNumber(82)).minus(new BigNumber((_this.f18).length));
            _2089_v30 = _2089_v30;
            let _2094_v34;
            _2094_v34 = new _dafny.CodePoint('w'.codePointAt(0));
            let _2095_v35;
            _2095_v35 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,_2094_v34);
            let _2096_v36;
            _2096_v36 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_2053_v0)).cardinality()),(_this).f17);
            let _2097_v37;
            _2097_v37 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm17((_2054_v1).f21, new BigNumber((_2095_v35).length), _this.f14, new BigNumber((_2096_v36).length), globalState),_2094_v34);
            let _2098_v38;
            _2098_v38 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2094_v34);
            let _rhs380 = !((_this).f17) || (!((_this).f17));
            let _rhs381 = (((_2097_v37).contains((_2098_v38).equals(_2098_v38))) ? ((_2097_v37).get((_2098_v38).equals(_2098_v38))) : (_2094_v34));
            let _rhs382 = _this.f14;
            let _lhs295 = _this;
            let _lhs296 = globalState;
            _lhs295.f16 = _rhs380;
            _2094_v34 = _rhs381;
            _lhs296.f7 = _rhs382;
          }
        }
      }
      r0 = p0;
      return r0;
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _2099_v0;
      _2099_v0 = _dafny.Seq.of(_this.f16, true);
      let _2100_v1;
      let _nw404 = Array((new BigNumber(26)).toNumber()).fill(_module.D9.Default());
      _2100_v1 = _nw404;
      let _2101_v2;
      _2101_v2 = _dafny.Seq.of(_2100_v1, _2100_v1);
      let _2102_v3;
      _2102_v3 = _module.D3.create_DC3((_this).f17);
      let _2103_v4;
      _2103_v4 = _dafny.Seq.of(_2102_v3);
      let _2104_v5;
      _2104_v5 = _module.D12.create_DC24(_this.f16, _2099_v0, _2101_v2, p2, _2103_v4);
      let _2105_v6;
      _2105_v6 = _module.D20.create_DC41(_dafny.MultiSet.fromElements(_2104_v5));
      let _2106_v7;
      _2106_v7 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(p1, (_dafny.ZERO).minus(_this.f14)),((_this.f16) ? (_2105_v6) : (_2105_v6)));
      let _2107_v8;
      let _nw405 = new _module.C6();
      _nw405.__ctor(_this.f14, (_this).f17, _this.f16, (_this).f17, p1);
      _2107_v8 = _nw405;
      let _2108_v9;
      _2108_v9 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_2107_v8);
      let _2109_v10;
      _2109_v10 = _module.D20.create_DC42(_module.__default.safeModuloInt(new BigNumber(393), p3), ((p2) ? (p2) : ((_module.D14.create_DC28(p3, p2, (((_2108_v9).contains(_this.f14)) ? ((_2108_v9).get(_this.f14)) : (_2107_v8)))).dtor_cf41)), _2107_v8);
      let _2110_v11;
      _2110_v11 = _dafny.Seq.of(_2106_v7);
      let _2111_v12;
      _2111_v12 = _dafny.MultiSet.fromElements(p0);
      let _2112_v13;
      _2112_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,(_2110_v11)[_module.__default.safeIndex(new BigNumber((_2111_v12).cardinality()), new BigNumber((_2110_v11).length))]);
      let _rhs383 = ((_2106_v7).Merge(_2106_v7)).Merge((((_2112_v13).contains(_this.f16)) ? ((_2112_v13).get(_this.f16)) : (_2106_v7)));
      let _rhs384 = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_this.f18, _this.f18), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.Concat(_this.f18, _this.f18)).length)), _module.__default.fm34(_2107_v8.f14, p1, _this.f16, p3, globalState)), _module.__default.safeIndex(_module.__default.fm0(true, globalState), new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_this.f18, _this.f18), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.Concat(_this.f18, _this.f18)).length)), _module.__default.fm34(_2107_v8.f14, p1, _this.f16, p3, globalState))).length)), p0);
      let _rhs385 = (_this).f17;
      let _rhs386 = function (_pat_let38_0) {
        return function (_2113_dt__update__tmp_h0) {
          return function (_pat_let39_0) {
            return function (_2114_dt__update_hcf65_h0) {
              return _module.D20.create_DC42(_2114_dt__update_hcf65_h0, (_2113_dt__update__tmp_h0).dtor_cf66, (_2113_dt__update__tmp_h0).dtor_cf67);
            }(_pat_let39_0);
          }(new BigNumber((_dafny.Seq.of((_this).f17, true, _this.f16)).length));
        }(_pat_let38_0);
      }(_2109_v10);
      let _lhs297 = _this;
      let _lhs298 = _this;
      _2106_v7 = _rhs383;
      _lhs297.f18 = _rhs384;
      _lhs298.f16 = _rhs385;
      _2109_v10 = _rhs386;
      (_this).f14 = _module.__default.safeModuloInt(p3, _this.f14);
      let _2115_i0;
      _2115_i0 = _dafny.ZERO;
      L12: {
        while (!(true) || (p2)) {
          C12: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2115_i0)) {
              break L12;
            }
            _2115_i0 = (_2115_i0).plus(_dafny.ONE);
            let _2116_v14;
            _2116_v14 = _dafny.Map.Empty.slice().updateUnsafe(_2107_v8.f14,_2107_v8.f14);
            let _2117_v15;
            _2117_v15 = _dafny.Set.fromElements(new BigNumber(-126), _2107_v8.f14, _2107_v8.f14);
            let _2118_v16;
            _2118_v16 = _dafny.Map.Empty.slice().updateUnsafe(_2117_v15,_this.f14);
            let _2119_v17;
            _2119_v17 = _dafny.Seq.of(p3);
            let _2120_v18;
            _2120_v18 = _dafny.Map.Empty.slice().updateUnsafe(_2119_v17,_dafny.Seq.of((_this).f17, (_this).f17));
            let _2121_v19;
            _2121_v19 = _dafny.Set.fromElements(_this.f16);
            let _2122_v20;
            _2122_v20 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,new BigNumber((_dafny.Set.fromElements(_this.f14)).length));
            let _2123_v21;
            let _nw406 = Array((new BigNumber(26)).toNumber());
            _nw406[(_dafny.ZERO).toNumber()] = (p3).minus(p3);
            _nw406[(_dafny.ONE).toNumber()] = _this.f14;
            _nw406[(new BigNumber(2)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_2116_v14).length), p3);
            _nw406[(new BigNumber(3)).toNumber()] = _module.__default.fm0(p2, globalState);
            _nw406[(new BigNumber(4)).toNumber()] = _2107_v8.f14;
            _nw406[(new BigNumber(5)).toNumber()] = p3;
            _nw406[(new BigNumber(6)).toNumber()] = p3;
            _nw406[(new BigNumber(7)).toNumber()] = p1;
            _nw406[(new BigNumber(8)).toNumber()] = p3;
            _nw406[(new BigNumber(9)).toNumber()] = p1;
            _nw406[(new BigNumber(10)).toNumber()] = (new BigNumber((_2118_v16).length)).minus((_this).fm12(p1, globalState));
            _nw406[(new BigNumber(11)).toNumber()] = _module.__default.safeDivisionInt(p1, p3);
            _nw406[(new BigNumber(12)).toNumber()] = _2107_v8.f14;
            _nw406[(new BigNumber(13)).toNumber()] = p3;
            _nw406[(new BigNumber(14)).toNumber()] = _this.f14;
            _nw406[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_2099_v0, (((_2120_v18).contains(_2119_v17)) ? ((_2120_v18).get(_2119_v17)) : (_2099_v0)))).length));
            _nw406[(new BigNumber(16)).toNumber()] = (p3).minus(_2107_v8.f14);
            _nw406[(new BigNumber(17)).toNumber()] = new BigNumber(((_2121_v19).Difference(_dafny.Set.fromElements(p2, (_this).f17))).length);
            _nw406[(new BigNumber(18)).toNumber()] = _this.f14;
            _nw406[(new BigNumber(19)).toNumber()] = new BigNumber((_2119_v17).length);
            _nw406[(new BigNumber(20)).toNumber()] = (p3).minus(_2107_v8.f14);
            _nw406[(new BigNumber(21)).toNumber()] = (_this.f14).minus(new BigNumber(728));
            _nw406[(new BigNumber(22)).toNumber()] = _module.__default.safeModuloInt(p3, new BigNumber((_2099_v0).length));
            _nw406[(new BigNumber(23)).toNumber()] = ((_dafny.ZERO).minus(_2107_v8.f14)).multipliedBy(new BigNumber((_2099_v0).length));
            _nw406[(new BigNumber(24)).toNumber()] = new BigNumber((_2122_v20).length);
            _nw406[(new BigNumber(25)).toNumber()] = _2107_v8.f14;
            _2123_v21 = _nw406;
            let _index344 = _module.__default.safeIndex(new BigNumber(490), new BigNumber((_2123_v21).length));
            (_2123_v21)[_index344] = _2107_v8.f14;
            let _index345 = _module.__default.safeIndex(new BigNumber(490), new BigNumber((_2123_v21).length));
            (_2123_v21)[_index345] = (((_this).f17) ? (new BigNumber((_this.f18).length)) : (_module.__default.fm0(!((_this).f17), globalState)));
            let _2124_v22;
            let _nw407 = Array((new BigNumber(9)).toNumber()).fill(_dafny.MultiSet.Empty);
            _2124_v22 = _nw407;
            let _2125_v23;
            let _init65 = ((_2126_p2) => function (_2127_i1) {
              return _2126_p2;
            })(p2);
            let _nw408 = Array((new BigNumber(27)).toNumber());
            for (let _i0_65 = 0; _i0_65 < new BigNumber(_nw408.length); _i0_65++) {
              _nw408[_i0_65] = _init65(new BigNumber(_i0_65));
            }
            _2125_v23 = _nw408;
            let _2128_v24;
            _2128_v24 = _dafny.Seq.of(_2124_v22);
            let _2129_v25;
            _2129_v25 = _dafny.MultiSet.fromElements(_this.f16, _this.f16, p2, (_this).f17, _module.__default.fm31(_this.f16, (_this).f17, globalState));
            let _2130_v26;
            let _nw409 = new _module.C2();
            _nw409.__ctor(_module.D3.create_DC3(p2), (_this).f17, _this.f16);
            _2130_v26 = _nw409;
            let _2131_v27;
            _2131_v27 = _dafny.Seq.of(_2130_v26, _2130_v26, _2130_v26, _2130_v26, _2130_v26);
            let _2132_v28;
            _2132_v28 = _dafny.Map.Empty.slice().updateUnsafe(_2131_v27,((_this.f16) ? (_2125_v23) : (_2125_v23)));
            let _rhs387 = (_2128_v24)[_module.__default.safeIndex(p1, new BigNumber((_2128_v24).length))];
            let _rhs388 = _module.__default.safeDivisionInt(((false) ? (new BigNumber((_2129_v25).cardinality())) : (_module.__default.fm0(_this.f16, globalState))), ((_2119_v17)[_module.__default.safeIndex(_2107_v8.f14, new BigNumber((_2119_v17).length))]).plus(_this.f14));
            let _rhs389 = (((_2132_v28).contains(_2131_v27)) ? ((_2132_v28).get(_2131_v27)) : (_2125_v23));
            let _lhs299 = globalState;
            _2124_v22 = _rhs387;
            _lhs299.f2 = _rhs388;
            _2125_v23 = _rhs389;
            if ((_this).f17) {
              let _2133_v29;
              _2133_v29 = _dafny.Map.Empty.slice().updateUnsafe(true,p2);
              _2116_v14 = (_2116_v14).update(new BigNumber(((_module.__default.fm40(p2, globalState)).Merge(_2133_v29)).length), new BigNumber((_2119_v17).length));
              let _2134_v30;
              _2134_v30 = _module.D16.create_DC33(_this.f16, _this.f16, (_2123_v21)[_module.__default.safeIndex(new BigNumber(490), new BigNumber((_2123_v21).length))], _this.f16);
              let _2135_v31;
              _2135_v31 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let40_0) {
                return function (_2136_dt__update__tmp_h1) {
                  return function (_pat_let41_0) {
                    return function (_2137_dt__update_hcf50_h0) {
                      return _module.D16.create_DC33((_2136_dt__update__tmp_h1).dtor_cf47, (_2136_dt__update__tmp_h1).dtor_cf48, (_2136_dt__update__tmp_h1).dtor_cf49, _2137_dt__update_hcf50_h0);
                    }(_pat_let41_0);
                  }(false);
                }(_pat_let40_0);
              }(_2134_v30),new BigNumber(970));
              _2135_v31 = (_2135_v31).update(_2134_v30, ((_dafny.ZERO).minus(p3)).plus(_this.f14));
              let _2138_v32;
              _2138_v32 = _module.D4.create_DC6(_dafny.Map.Empty.slice().updateUnsafe(p2,p2));
              let _2139_v33;
              let _nw410 = new _module.C4();
              _nw410.__ctor((_this).f17, ((_2138_v32).dtor_cf10).contains(_this.f16), (_this).f17);
              _2139_v33 = _nw410;
              let _index346 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_2125_v23).length));
              (_2125_v23)[_index346] = false;
              let _index347 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_2125_v23).length));
              (_2125_v23)[_index347] = _module.__default.fm17(false, _2107_v8.f14, p3, (_2123_v21)[_module.__default.safeIndex(new BigNumber(490), new BigNumber((_2123_v21).length))], globalState);
              (_this).f16 = (_2139_v33).f21;
            } else {
              let _2140_v34;
              _2140_v34 = _dafny.Seq.of(_2125_v23, _2125_v23);
              let _2141_v35;
              _2141_v35 = _dafny.Seq.of(_2140_v34, _2140_v34, _2140_v34);
              let _2142_v36;
              _2142_v36 = (_2141_v35)[_module.__default.safeIndex(p1, new BigNumber((_2141_v35).length))];
              _2140_v34 = (_2142_v36);
              let _2143_v37;
              let _nw411 = new _module.C2();
              _nw411.__ctor(_2130_v26.f24, _module.__default.fm43(_this.f16, globalState), !(p2));
              _2143_v37 = _nw411;
              let _2144_v38;
              _2144_v38 = _dafny.Map.Empty.slice().updateUnsafe(_2121_v19,(_this).f17);
              let _2145_v39;
              _2145_v39 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,true);
              _2144_v38 = (_2144_v38).update((((_this).f17) ? (_dafny.Set.fromElements(true, _module.__default.fm17(p2, p3, _2107_v8.f14, new BigNumber((_2099_v0).length), globalState), (((_2145_v39).contains(new BigNumber((_dafny.Seq.UnicodeFromString("lph")).length))) ? ((_2145_v39).get(new BigNumber((_dafny.Seq.UnicodeFromString("lph")).length))) : (_this.f16)))) : (_2121_v19)), (_this).f17);
              (globalState).f2 = new BigNumber((_this.f18).length);
              let _index348 = _module.__default.safeIndex(new BigNumber(490), new BigNumber((_2123_v21).length));
              (_2123_v21)[_index348] = new BigNumber(((((_2129_v25).IsDisjointFrom(_dafny.MultiSet.fromElements((_this).f17))) ? (_dafny.Seq.of(p0)) : (_this.f18))).length);
            }
            let _index349 = _module.__default.safeIndex(new BigNumber(211), new BigNumber((_2125_v23).length));
            (_2125_v23)[_index349] = _this.f16;
            let _2146_v40;
            _2146_v40 = _dafny.Seq.of(_2125_v23, _2125_v23, _2125_v23, _2125_v23, _2125_v23);
            let _2147_v41;
            _2147_v41 = _2146_v40;
            let _index350 = _module.__default.safeIndex(new BigNumber(211), new BigNumber((_2125_v23).length));
            let _rhs390 = false;
            let _rhs391 = _2147_v41;
            let _lhs300 = _2125_v23;
            let _lhs301 = _module.__default.safeIndex(new BigNumber(211), new BigNumber((_2125_v23).length));
            _lhs300[_lhs301] = _rhs390;
            _2147_v41 = _rhs391;
          }
        }
      }
      if (!(_this.f16)) {
        r0 = p3;
        let _2148_v42;
        let _nw412 = Array((new BigNumber(21)).toNumber()).fill(false);
        _2148_v42 = _nw412;
        let _2149_v43;
        _2149_v43 = _dafny.Map.Empty.slice().updateUnsafe(p2,_2148_v42);
        let _2150_v44;
        _2150_v44 = _dafny.Set.fromElements((((_2149_v43).contains((_this).f17)) ? ((_2149_v43).get((_this).f17)) : (_2148_v42)));
        (globalState).f2 = new BigNumber((_2150_v44).length);
        if (!((new BigNumber(-292)).isLessThanOrEqualTo(p3))) {
          let _2151_v45;
          let _nw413 = Array((new BigNumber(6)).toNumber()).fill(_module.D9.Default());
          _2151_v45 = _nw413;
          _2151_v45 = _2151_v45;
          let _2152_v47;
          let _init66 = ((_2153_v0) => function (_2154_i2) {
            return function () {
              let _coll76 = new _dafny.Set();
              for (const _compr_76 of (_dafny.Seq.of(_2153_v0)).Elements) {
                let _2155_v46 = _compr_76;
                if (_dafny.Seq.contains(_dafny.Seq.of(_2153_v0), _2155_v46)) {
                  _coll76.add(_2155_v46);
                }
              }
              return _coll76;
            }();
          })(_2099_v0);
          let _nw414 = Array((new BigNumber(28)).toNumber());
          for (let _i0_66 = 0; _i0_66 < new BigNumber(_nw414.length); _i0_66++) {
            _nw414[_i0_66] = _init66(new BigNumber(_i0_66));
          }
          _2152_v47 = _nw414;
          let _2156_v48;
          _2156_v48 = _dafny.Set.fromElements(_dafny.Seq.of(!(p2)), _2099_v0);
          let _index351 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_2152_v47).length));
          (_2152_v47)[_index351] = (_2156_v48).Union(_dafny.Set.fromElements(_dafny.Seq.of((_this).f17), _2099_v0, _2099_v0));
          let _index352 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_2152_v47).length));
          (_2152_v47)[_index352] = ((!(_2107_v8.f14).isEqualTo(_2107_v8.f14)) ? (_2156_v48) : (function () {
            let _coll77 = new _dafny.Set();
            for (const _compr_77 of (_dafny.MultiSet.fromElements(_dafny.Seq.of(p2), _2099_v0)).Elements) {
              let _2157_v49 = _compr_77;
              if ((_dafny.MultiSet.fromElements(_dafny.Seq.of(p2), _2099_v0)).contains(_2157_v49)) {
                _coll77.add(_2157_v49);
              }
            }
            return _coll77;
          }()));
          (_this).f16 = false;
          let _2158_v50;
          let _nw415 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
          _2158_v50 = _nw415;
          let _index353 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((_2158_v50).length));
          (_2158_v50)[_index353] = (_2107_v8.f14).plus(_2107_v8.f14);
          let _2159_v51;
          _2159_v51 = _dafny.Seq.of(_2158_v50, _2158_v50);
          let _index354 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((_2158_v50).length));
          let _rhs392 = _2107_v8.f14;
          let _rhs393 = _dafny.Seq.Concat(_this.f18, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lj"), _this.f18));
          let _rhs394 = (_2159_v51)[_module.__default.safeIndex(_2107_v8.f14, new BigNumber((_2159_v51).length))];
          let _rhs395 = (_this).f17;
          let _rhs396 = (_dafny.ZERO).minus((_this.f14).minus((_dafny.ZERO).minus(new BigNumber((((_this).fm15(globalState)).Intersect(_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(495)), ((_2160_p0) => function (_2161_i3) {
            return _2160_p0;
          })(p0)), _this.f18, _dafny.Seq.UnicodeFromString("ibuvhh")))).cardinality()))));
          let _lhs302 = _2107_v8;
          let _lhs303 = _this;
          let _lhs304 = _this;
          let _lhs305 = _2158_v50;
          let _lhs306 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((_2158_v50).length));
          _lhs302.f14 = _rhs392;
          _lhs303.f18 = _rhs393;
          _2158_v50 = _rhs394;
          _lhs304.f16 = _rhs395;
          _lhs305[_lhs306] = _rhs396;
          let _2162_v52;
          _2162_v52 = _module.D19.create_DC40(_2107_v8.f14, p2, (_this.f14).minus(_2107_v8.f14));
          _2162_v52 = _2162_v52;
        } else {
          (_this).f16 = !(true);
          (globalState).f2 = _2107_v8.f14;
          let _2163_v53;
          _2163_v53 = _module.D16.create_DC34(false, _2107_v8, p3);
          let _2164_v54;
          _2164_v54 = _module.D16.create_DC35(_2163_v53);
          let _2165_v55;
          _2165_v55 = _module.D16.create_DC35(_2163_v53);
          let _2166_v56;
          _2166_v56 = _module.D16.create_DC35(_2163_v53);
          let _2167_v57;
          _2167_v57 = _dafny.Seq.of(_2166_v56, _2166_v56);
          let _2168_v58;
          let _nw416 = new _module.C6();
          _nw416.__ctor(_2107_v8.f14, (_2107_v8.f14).isLessThanOrEqualTo(p1), _this.f16, p2, new BigNumber((_2167_v57).length));
          _2168_v58 = _nw416;
          _2168_v58 = _2168_v58;
          (_2168_v58).f16 = (_this).fm14(globalState);
          _2099_v0 = _2099_v0;
        }
        let _2169_v59;
        _2169_v59 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2148_v42);
        let _2170_v60;
        let _nw417 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.of());
        _2170_v60 = _nw417;
        let _index355 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_2170_v60).length));
        (_2170_v60)[_index355] = _dafny.Seq.Concat(_2099_v0, _2099_v0);
        let _index356 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_2170_v60).length));
        let _rhs397 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_2107_v8.f14, _module.__default.fm0(_this.f16, globalState)),_2148_v42);
        let _rhs398 = _2099_v0;
        let _lhs307 = _2170_v60;
        let _lhs308 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_2170_v60).length));
        _2169_v59 = _rhs397;
        _lhs307[_lhs308] = _rhs398;
        let _2171_v61;
        _2171_v61 = _module.D19.create_DC40(_2107_v8.f14, _this.f16, _2107_v8.f14);
        let _source33 = _2171_v61;
        if (_source33.is_DC40) {
          let _2172___mcc_h0 = (_source33).cf61;
          let _2173___mcc_h1 = (_source33).cf62;
          let _2174___mcc_h2 = (_source33).cf63;
          let _2175_cf63 = _2174___mcc_h2;
          let _2176_cf62 = _2173___mcc_h1;
          let _2177_cf61 = _2172___mcc_h0;
          r0 = p1;
          let _2178_v62;
          let _nw418 = Array((new BigNumber(4)).toNumber()).fill([]);
          _2178_v62 = _nw418;
          let _index357 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_2178_v62).length));
          (_2178_v62)[_index357] = _2148_v42;
          let _index358 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_2178_v62).length));
          (_2178_v62)[_index358] = _2148_v42;
          let _2179_v63;
          let _init67 = function (_2180_i4) {
            return _dafny.Set.fromElements(_this.f16, !((_this).f17));
          };
          let _nw419 = Array((new BigNumber(20)).toNumber());
          for (let _i0_67 = 0; _i0_67 < new BigNumber(_nw419.length); _i0_67++) {
            _nw419[_i0_67] = _init67(new BigNumber(_i0_67));
          }
          _2179_v63 = _nw419;
          let _2181_v64;
          let _nw420 = new _module.C0();
          _nw420.__ctor(_2179_v63, new BigNumber((_dafny.Seq.Concat(_this.f18, _this.f18)).length));
          _2181_v64 = _nw420;
          let _2182_v65;
          _2182_v65 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f16) || (_this.f16),((_this.f16) ? ((_2170_v60)[_module.__default.safeIndex(new BigNumber(509), new BigNumber((_2170_v60).length))]) : (_2099_v0)));
          _2182_v65 = (_2182_v65).update(!((_this).f17) || (_this.f16), _dafny.Seq.Concat((_2099_v0), _2099_v0));
        } else {
          let _2183___mcc_h3 = (_source33).cf60;
          let _2184_cf60 = _2183___mcc_h3;
          let _2185_v66;
          _2185_v66 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,(_this).f17);
          let _2186_v67;
          _2186_v67 = _dafny.MultiSet.fromElements(_this.f16);
          let _2187_v68;
          let _nw421 = new _module.C1();
          _nw421.__ctor((((_2185_v66).contains(new BigNumber((_dafny.Set.fromElements(p1, new BigNumber((_2186_v67).cardinality()), p1)).length))) ? ((_2185_v66).get(new BigNumber((_dafny.Set.fromElements(p1, new BigNumber((_2186_v67).cardinality()), p1)).length))) : (false)), (_this).fm12(_2107_v8.f14, globalState));
          _2187_v68 = _nw421;
          let _2188_v69;
          _2188_v69 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,(_2186_v67).update((_this).f17, _module.__default.abs((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('c'.codePointAt(0)))).length)))));
          let _2189_v70;
          _2189_v70 = _dafny.Set.fromElements(_2187_v68.f25);
          let _2190_v71;
          _2190_v71 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_2189_v70).length));
          let _2191_v72;
          _2191_v72 = _dafny.Seq.of(p1);
          _2188_v69 = (_2188_v69).update(((_module.__default.fm17(_this.f16, new BigNumber((_2190_v71).length), new BigNumber(((_2191_v72)).length), _this.f14, globalState)) ? (_2187_v68.f25) : ((_this).f17)), _dafny.MultiSet.fromElements(_2187_v68.f25, !((_this).f17), _2187_v68.f25));
          let _2192_v73;
          let _init68 = function (_2193_i5) {
            return _this.f18;
          };
          let _nw422 = Array((new BigNumber(11)).toNumber());
          for (let _i0_68 = 0; _i0_68 < new BigNumber(_nw422.length); _i0_68++) {
            _nw422[_i0_68] = _init68(new BigNumber(_i0_68));
          }
          _2192_v73 = _nw422;
          let _2194_v74;
          let _nw423 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _2194_v74 = _nw423;
          let _index359 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_2194_v74).length));
          (_2194_v74)[_index359] = _2107_v8.f14;
          let _pat_let_tv33 = _2099_v0;
          let _pat_let_tv34 = _2101_v2;
          let _2195_v75;
          _2195_v75 = _dafny.MultiSet.fromElements(_2104_v5, _2104_v5, _2104_v5, function (_pat_let42_0) {
            return function (_2196_dt__update__tmp_h2) {
              return function (_pat_let43_0) {
                return function (_2197_dt__update_hcf34_h0) {
                  return function (_pat_let44_0) {
                    return function (_2198_dt__update_hcf35_h0) {
                      return _module.D12.create_DC24((_2196_dt__update__tmp_h2).dtor_cf33, _2197_dt__update_hcf34_h0, _2198_dt__update_hcf35_h0, (_2196_dt__update__tmp_h2).dtor_cf36, (_2196_dt__update__tmp_h2).dtor_cf37);
                    }(_pat_let44_0);
                  }(_pat_let_tv34);
                }(_pat_let43_0);
              }(_pat_let_tv33);
            }(_pat_let42_0);
          }(_2104_v5));
          let _2199_v76;
          _2199_v76 = _dafny.Map.Empty.slice().updateUnsafe(_2107_v8.f14,_2186_v67);
          let _index360 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_2194_v74).length));
          let _rhs399 = _2192_v73;
          let _rhs400 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(_2107_v8.f14, new BigNumber((_2199_v76).length)), _2107_v8.f14);
          let _rhs401 = (_2190_v71).Merge(_2190_v71);
          let _rhs402 = _dafny.MultiSet.fromElements(_module.D12.create_DC24(false, _dafny.Seq.of(_this.f16), _2101_v2, _this.f16, _2103_v4), _2104_v5, _2104_v5);
          let _lhs309 = _2194_v74;
          let _lhs310 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_2194_v74).length));
          _2192_v73 = _rhs399;
          _lhs309[_lhs310] = _rhs400;
          _2190_v71 = _rhs401;
          _2195_v75 = _rhs402;
          let _2200_v77;
          let _nw424 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.of());
          _2200_v77 = _nw424;
          let _index361 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_2194_v74).length));
          let _rhs403 = _dafny.Seq.Concat(_this.f18, _dafny.Seq.Concat(_this.f18, _this.f18));
          let _rhs404 = p1;
          let _rhs405 = _2200_v77;
          let _lhs311 = _this;
          let _lhs312 = _2194_v74;
          let _lhs313 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_2194_v74).length));
          _lhs311.f18 = _rhs403;
          _lhs312[_lhs313] = _rhs404;
          _2200_v77 = _rhs405;
        }
      } else {
        let _rhs406 = _dafny.Seq.UnicodeFromString("hsv");
        let _rhs407 = _2107_v8.f14;
        let _rhs408 = _2107_v8.f14;
        let _lhs314 = _this;
        let _lhs315 = _this;
        let _lhs316 = globalState;
        _lhs314.f18 = _rhs406;
        _lhs315.f14 = _rhs407;
        _lhs316.f2 = _rhs408;
        let _2201_v78;
        _2201_v78 = _dafny.MultiSet.fromElements(_2107_v8.f14);
        let _2202_v79;
        _2202_v79 = _module.D19.create_DC39(_2201_v78);
        let _2203_v80;
        _2203_v80 = _dafny.Set.fromElements(true, !(_this.f16));
        let _rhs409 = _module.__default.fm31((_2203_v80).contains(true), ((_this).f17) === ((_this).f17), globalState);
        let _rhs410 = _2104_v5;
        let _rhs411 = _module.__default.fm63(globalState);
        let _lhs317 = _this;
        _lhs317.f16 = _rhs409;
        _2104_v5 = _rhs410;
        _2202_v79 = _rhs411;
        (globalState).f2 = _2107_v8.f14;
        let _2204_v81;
        let _nw425 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _2204_v81 = _nw425;
        let _index362 = _module.__default.safeIndex(new BigNumber(912), new BigNumber((_2204_v81).length));
        (_2204_v81)[_index362] = _dafny.Seq.UnicodeFromString("sxckdbjch");
        let _index363 = _module.__default.safeIndex(new BigNumber(912), new BigNumber((_2204_v81).length));
        (_2204_v81)[_index363] = _this.f18;
        if (!(_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(377)), ((_2205_p0) => function (_2206_i6) {
          return _2205_p0;
        })(p0)), (_2204_v81)[_module.__default.safeIndex(new BigNumber(912), new BigNumber((_2204_v81).length))])) || (!(!(p2)))) {
          let _index364 = _module.__default.safeIndex(new BigNumber(912), new BigNumber((_2204_v81).length));
          (_2204_v81)[_index364] = (_2204_v81)[_module.__default.safeIndex(new BigNumber(912), new BigNumber((_2204_v81).length))];
          let _2207_v82;
          let _nw426 = Array((new BigNumber(20)).toNumber()).fill(_dafny.MultiSet.Empty);
          _2207_v82 = _nw426;
          let _2208_v83;
          _2208_v83 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_2201_v78);
          let _index365 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_2207_v82).length));
          (_2207_v82)[_index365] = (((_2208_v83).contains(_this.f16)) ? ((_2208_v83).get(_this.f16)) : (_2201_v78));
          let _index366 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_2207_v82).length));
          (_2207_v82)[_index366] = _2201_v78;
          (_this).f16 = p2;
          r0 = _2107_v8.f14;
          let _2209_v84;
          _2209_v84 = _dafny.Set.fromElements(_2107_v8.f14, _2107_v8.f14);
          let _2210_v85;
          _2210_v85 = _dafny.Seq.of(_2209_v84, _2209_v84, _2209_v84);
          let _2211_v86;
          _2211_v86 = _dafny.Seq.of(p3);
          let _2212_v87;
          _2212_v87 = _dafny.Seq.of(_module.__default.fm0(_module.__default.fm17(p2, new BigNumber((_dafny.Seq.of(p0)).length), _2107_v8.f14, new BigNumber((_2211_v86).length), globalState), globalState));
          let _2213_v88;
          _2213_v88 = _dafny.Map.Empty.slice().updateUnsafe(true,_2107_v8.f14);
          let _2214_v89;
          _2214_v89 = _dafny.Map.Empty.slice().updateUnsafe(_2107_v8.f14,_2107_v8.f14);
          let _2215_v90;
          _2215_v90 = _dafny.MultiSet.fromElements(!(_this.f16), p2);
          let _2216_v92;
          let _nw427 = Array((new BigNumber(25)).toNumber());
          _nw427[(_dafny.ZERO).toNumber()] = _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f16,(_2207_v82)[_module.__default.safeIndex(new BigNumber(203), new BigNumber((_2207_v82).length))])).length));
          _nw427[(_dafny.ONE).toNumber()] = _2209_v84;
          _nw427[(new BigNumber(2)).toNumber()] = (_2210_v85)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_2099_v0, _module.__default.safeIndex(new BigNumber((_2212_v87).length), new BigNumber((_2099_v0).length)), !(_this.f16))).length), new BigNumber((_2210_v85).length))];
          _nw427[(new BigNumber(3)).toNumber()] = _module.__default.fm49(p3, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_this.f16)).length), (_this).f17, new BigNumber(((_2204_v81)[_module.__default.safeIndex(new BigNumber(912), new BigNumber((_2204_v81).length))]).length), globalState);
          _nw427[(new BigNumber(4)).toNumber()] = _2209_v84;
          _nw427[(new BigNumber(5)).toNumber()] = _2209_v84;
          _nw427[(new BigNumber(6)).toNumber()] = _module.__default.fm49((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-407)), ((_2217_p0) => function (_2218_i7) {
            return _2217_p0;
          })(p0))).length), _2107_v8.f14, _this.f14)).length)), (((_2213_v88).contains(p2)) ? ((_2213_v88).get(p2)) : ((((_2214_v89).contains(new BigNumber((_2201_v78).cardinality()))) ? ((_2214_v89).get(new BigNumber((_2201_v78).cardinality()))) : (_2107_v8.f14)))), _this.f16, _2107_v8.f14, globalState);
          _nw427[(new BigNumber(7)).toNumber()] = _module.__default.fm49(_2107_v8.f14, _2107_v8.f14, false, _2107_v8.f14, globalState);
          _nw427[(new BigNumber(8)).toNumber()] = _dafny.Set.fromElements((((_2215_v90).contains(_this.f16)) ? ((_2215_v90).get(_this.f16)) : (_2107_v8.f14)));
          _nw427[(new BigNumber(9)).toNumber()] = _2209_v84;
          _nw427[(new BigNumber(10)).toNumber()] = _2209_v84;
          _nw427[(new BigNumber(11)).toNumber()] = _2209_v84;
          _nw427[(new BigNumber(12)).toNumber()] = _2209_v84;
          _nw427[(new BigNumber(13)).toNumber()] = _2209_v84;
          _nw427[(new BigNumber(14)).toNumber()] = _2209_v84;
          _nw427[(new BigNumber(15)).toNumber()] = (_2209_v84).Difference(_dafny.Set.fromElements(p1, new BigNumber(-151)));
          _nw427[(new BigNumber(16)).toNumber()] = _dafny.Set.fromElements(new BigNumber((_2203_v80).length));
          _nw427[(new BigNumber(17)).toNumber()] = ((_this.f16) ? (_dafny.Set.fromElements(_2107_v8.f14)) : (_2209_v84));
          _nw427[(new BigNumber(18)).toNumber()] = function () {
            let _coll78 = new _dafny.Set();
            for (const _compr_78 of _dafny.IntegerRange(new BigNumber(909), new BigNumber(310))) {
              let _2219_v91 = _compr_78;
              if (((new BigNumber(909)).isLessThanOrEqualTo(_2219_v91)) && ((_2219_v91).isLessThan(new BigNumber(310)))) {
                _coll78.add(_module.__default.safeDivisionInt(_2219_v91, new BigNumber(((_2204_v81)[_module.__default.safeIndex(new BigNumber(912), new BigNumber((_2204_v81).length))]).length)));
              }
            }
            return _coll78;
          }();
          _nw427[(new BigNumber(19)).toNumber()] = _dafny.Set.fromElements(p3, new BigNumber(134), new BigNumber((_this.f18).length));
          _nw427[(new BigNumber(20)).toNumber()] = (_2209_v84).Intersect(_2209_v84);
          _nw427[(new BigNumber(21)).toNumber()] = _2209_v84;
          _nw427[(new BigNumber(22)).toNumber()] = _2209_v84;
          _nw427[(new BigNumber(23)).toNumber()] = _module.__default.fm49(_2107_v8.f14, p1, p2, new BigNumber((_dafny.Set.fromElements((_this).f17, (_this).f17, _this.f16, p2)).length), globalState);
          _nw427[(new BigNumber(24)).toNumber()] = _2209_v84;
          _2216_v92 = _nw427;
          let _2220_v93;
          _2220_v93 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,p2);
          let _index367 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_2216_v92).length));
          (_2216_v92)[_index367] = _module.__default.fm49(new BigNumber((_2220_v93).length), new BigNumber((_dafny.Seq.of(p3)).length), _this.f16, p3, globalState);
          let _index368 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_2216_v92).length));
          (_2216_v92)[_index368] = (_2209_v84).Difference((_2209_v84).Union(_2209_v84));
        } else {
          let _rhs412 = ((_2107_v8.f14).minus(p3)).isLessThan(new BigNumber(301));
          let _rhs413 = _2099_v0;
          let _rhs414 = _2107_v8.f14;
          let _rhs415 = ((p3).multipliedBy(p3)).plus(_2107_v8.f14);
          let _lhs318 = _this;
          let _lhs319 = globalState;
          let _lhs320 = globalState;
          _lhs318.f16 = _rhs412;
          _2099_v0 = _rhs413;
          _lhs319.f2 = _rhs414;
          _lhs320.f7 = _rhs415;
          let _2221_v94;
          let _init69 = function (_2222_i8) {
            return _module.__default.safeModuloInt(_2222_i8, new BigNumber(544));
          };
          let _nw428 = Array((new BigNumber(6)).toNumber());
          for (let _i0_69 = 0; _i0_69 < new BigNumber(_nw428.length); _i0_69++) {
            _nw428[_i0_69] = _init69(new BigNumber(_i0_69));
          }
          _2221_v94 = _nw428;
          let _index369 = _module.__default.safeIndex(new BigNumber(142), new BigNumber((_2221_v94).length));
          (_2221_v94)[_index369] = (p3).multipliedBy(_2107_v8.f14);
          let _index370 = _module.__default.safeIndex(new BigNumber(142), new BigNumber((_2221_v94).length));
          (_2221_v94)[_index370] = _this.f14;
          let _2223_v95;
          _2223_v95 = _dafny.Map.Empty.slice().updateUnsafe(p1,_this.f16);
          let _2224_v96;
          _2224_v96 = _module.D12.create_DC23(p2, _2107_v8.f14, new BigNumber((_2223_v95).length));
          let _2225_v97;
          _2225_v97 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.update(_2099_v0, _module.__default.safeIndex((_2221_v94)[_module.__default.safeIndex(new BigNumber(142), new BigNumber((_2221_v94).length))], new BigNumber((_2099_v0).length)), _this.f16)).length),_2221_v94);
          let _2226_v99;
          _2226_v99 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_this.f16,(_this).f17));
          let _2227_v100;
          _2227_v100 = _module.D4.create_DC6(_module.__default.fm40((_this).f17, globalState));
          let _2228_v101;
          _2228_v101 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,false);
          let _2229_v102;
          _2229_v102 = _dafny.Map.Empty.slice().updateUnsafe(_2227_v100,_2228_v101);
          let _pat_let_tv35 = _2221_v94;
          let _pat_let_tv36 = _2221_v94;
          let _pat_let_tv37 = _2225_v97;
          let _2230_v103;
          let _nw429 = Array((new BigNumber(21)).toNumber());
          _nw429[(_dafny.ZERO).toNumber()] = (_this).f17;
          _nw429[(_dafny.ONE).toNumber()] = !((function (_pat_let45_0) {
            return function (_2231_dt__update__tmp_h3) {
              return function (_pat_let46_0) {
                return function (_2232_dt__update_hcf32_h0) {
                  return function (_pat_let47_0) {
                    return function (_2233_dt__update_hcf31_h0) {
                      return _module.D12.create_DC23((_2231_dt__update__tmp_h3).dtor_cf30, _2233_dt__update_hcf31_h0, _2232_dt__update_hcf32_h0);
                    }(_pat_let47_0);
                  }(new BigNumber((_pat_let_tv37).length));
                }(_pat_let46_0);
              }((_pat_let_tv36)[_module.__default.safeIndex(new BigNumber(142), new BigNumber((_pat_let_tv35).length))]);
            }(_pat_let45_0);
          }(_2224_v96)).dtor_cf30) || (_this.f16);
          _nw429[(new BigNumber(2)).toNumber()] = p2;
          _nw429[(new BigNumber(3)).toNumber()] = !(new BigNumber((function () {
            let _coll79 = new _dafny.Map();
            for (const _compr_79 of (_dafny.Map.Empty.slice().updateUnsafe(p0,(_dafny.ZERO).minus(p1))).Keys.Elements) {
              let _2234_v98 = _compr_79;
              if ((_dafny.Map.Empty.slice().updateUnsafe(p0,(_dafny.ZERO).minus(p1))).contains(_2234_v98)) {
                _coll79.push([_2234_v98,_2107_v8.f14]);
              }
            }
            return _coll79;
          }()).length)).isEqualTo(_module.__default.fm0(p2, globalState));
          _nw429[(new BigNumber(4)).toNumber()] = p2;
          _nw429[(new BigNumber(5)).toNumber()] = p2;
          _nw429[(new BigNumber(6)).toNumber()] = _this.f16;
          _nw429[(new BigNumber(7)).toNumber()] = (_this).f17;
          _nw429[(new BigNumber(8)).toNumber()] = !(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(293)), function (_2235_i9) {
            return _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f17);
          }), _dafny.Seq.update(_2226_v99, _module.__default.safeIndex(_2107_v8.f14, new BigNumber((_2226_v99).length)), ((((_2229_v102).contains(_module.D4.create_DC6(_2228_v101))) ? ((_2229_v102).get(_module.D4.create_DC6(_2228_v101))) : (_2228_v101))).update(_this.f16, p2))));
          _nw429[(new BigNumber(9)).toNumber()] = _this.f16;
          _nw429[(new BigNumber(10)).toNumber()] = (_2201_v78).IsSubsetOf(_2201_v78);
          _nw429[(new BigNumber(11)).toNumber()] = !_dafny.Seq.contains((_this).fm13(new BigNumber(-80), true, (_this).f17, globalState), _this.f16);
          _nw429[(new BigNumber(12)).toNumber()] = false;
          _nw429[(new BigNumber(13)).toNumber()] = p2;
          _nw429[(new BigNumber(14)).toNumber()] = _this.f16;
          _nw429[(new BigNumber(15)).toNumber()] = !((_this).f17);
          _nw429[(new BigNumber(16)).toNumber()] = _this.f16;
          _nw429[(new BigNumber(17)).toNumber()] = true;
          _nw429[(new BigNumber(18)).toNumber()] = _this.f16;
          _nw429[(new BigNumber(19)).toNumber()] = p2;
          _nw429[(new BigNumber(20)).toNumber()] = _this.f16;
          _2230_v103 = _nw429;
          let _index371 = _module.__default.safeIndex(new BigNumber(912), new BigNumber((_2204_v81).length));
          let _index372 = _module.__default.safeIndex(new BigNumber(142), new BigNumber((_2221_v94).length));
          let _rhs416 = _this.f18;
          let _rhs417 = _2107_v8.f14;
          let _rhs418 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f14));
          let _rhs419 = _2230_v103;
          let _rhs420 = ((p3).multipliedBy(_this.f14)).plus(_module.__default.fm0(false, globalState));
          let _lhs321 = _2204_v81;
          let _lhs322 = _module.__default.safeIndex(new BigNumber(912), new BigNumber((_2204_v81).length));
          let _lhs323 = _2107_v8;
          let _lhs324 = _this;
          let _lhs325 = _2221_v94;
          let _lhs326 = _module.__default.safeIndex(new BigNumber(142), new BigNumber((_2221_v94).length));
          _lhs321[_lhs322] = _rhs416;
          _lhs323.f14 = _rhs417;
          _lhs324.f14 = _rhs418;
          _2230_v103 = _rhs419;
          _lhs325[_lhs326] = _rhs420;
          let _2236_v104;
          _2236_v104 = _dafny.Seq.of(_2230_v103, _2230_v103);
          let _2237_v105;
          _2237_v105 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(743),((true) ? (p3) : (new BigNumber((_2236_v104).length))));
          _2237_v105 = _2237_v105;
          let _rhs421 = (_module.__default.fm0((_this).f17, globalState)).isEqualTo(p3);
          let _rhs422 = false;
          let _lhs327 = _this;
          let _lhs328 = _this;
          _lhs327.f16 = _rhs421;
          _lhs328.f16 = _rhs422;
        }
      }
      (_2107_v8).f14 = new BigNumber((_2111_v12).cardinality());
      _2099_v0 = _dafny.Seq.Concat(_2099_v0, _2099_v0);
      r0 = (_2107_v8.f14).minus(_2107_v8.f14);
      return r0;
    }
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.Seq.UnicodeFromString("");
      let r3 = false;
      let _2238_v0;
      _2238_v0 = _dafny.MultiSet.fromElements(p0);
      let _2239_v1;
      _2239_v1 = _dafny.Map.Empty.slice().updateUnsafe(p1,_this.f16);
      let _2240_v2;
      _2240_v2 = _module.D15.create_DC30(p3);
      let _2241_v3;
      _2241_v3 = _dafny.MultiSet.fromElements(_2240_v2);
      (_this).f14 = (((_2238_v0).contains((((_2239_v1).contains(p0)) ? ((_2239_v1).get(p0)) : (_this.f16)))) ? ((_2238_v0).get((((_2239_v1).contains(p0)) ? ((_2239_v1).get(p0)) : (_this.f16)))) : ((_dafny.ZERO).minus(new BigNumber((_2241_v3).cardinality()))));
      let _hi18 = _this.f14;
      for (let _2242_i0 = _module.__default.safeDivisionInt(p2, p2); _2242_i0.isLessThan(_hi18); _2242_i0 = _2242_i0.plus(_dafny.ONE)) {
        if (p1) {
          let _rhs423 = _module.__default.fm43((_this.f16) || (!(p0)), globalState);
          let _rhs424 = _dafny.Seq.Concat(_this.f18, _dafny.Seq.UnicodeFromString("bvioansl"));
          let _rhs425 = false;
          let _lhs329 = _this;
          let _lhs330 = _this;
          _lhs329.f16 = _rhs423;
          _lhs330.f18 = _rhs424;
          r1 = _rhs425;
          let _2243_v4;
          _2243_v4 = _dafny.Set.fromElements(p2);
          let _2244_v5;
          let _nw430 = Array((new BigNumber(10)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _2244_v5 = _nw430;
          let _2245_v6;
          _2245_v6 = _dafny.Map.Empty.slice().updateUnsafe(_2243_v4,(((_this).f17) ? (_2244_v5) : (_2244_v5)));
          _2245_v6 = (_2245_v6).update(_2243_v4, _2244_v5);
          (_this).f14 = _this.f14;
          (globalState).f2 = _2242_i0;
          let _2246_v7;
          _2246_v7 = _dafny.Seq.of((_this).f17, p0);
          r3 = !(!_dafny.Seq.contains(_2246_v7, _this.f16));
        } else {
          let _2247_v8;
          _2247_v8 = _dafny.Seq.of((_this).fm14(globalState));
          let _2248_v9;
          _2248_v9 = _dafny.Map.Empty.slice().updateUnsafe(_2247_v8,_2242_i0);
          let _2249_v11;
          _2249_v11 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll80 = new _dafny.Set();
            for (const _compr_80 of (_2248_v9).Keys.Elements) {
              let _2250_v10 = _compr_80;
              if ((_2248_v9).contains(_2250_v10)) {
                _coll80.add(_2250_v10);
              }
            }
            return _coll80;
          }(),_this.f14);
          _2249_v11 = (_2249_v11).update(_dafny.Set.fromElements(_dafny.Seq.of((_2247_v8)[_module.__default.safeIndex(_2242_i0, new BigNumber((_2247_v8).length))])), _2242_i0);
          let _2251_v12;
          _2251_v12 = new _dafny.CodePoint('p'.codePointAt(0));
          _2251_v12 = p3;
          let _2252_v13;
          _2252_v13 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2239_v1).length),p0);
          _2252_v13 = (_2252_v13).update(_this.f14, true);
          let _2253_v14;
          let _nw431 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.of());
          _2253_v14 = _nw431;
          let _nw432 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.of());
          _2253_v14 = _nw432;
          (globalState).f7 = (_2242_i0).minus(_module.__default.safeModuloInt(p2, _2242_i0));
        }
        let _2254_v15;
        let _nw433 = Array((new BigNumber(20)).toNumber()).fill(false);
        _2254_v15 = _nw433;
        let _2255_v16;
        _2255_v16 = _dafny.MultiSet.fromElements(_2254_v15);
        if (!(!(true) || ((_2255_v16).IsProperSubsetOf(_2255_v16)))) {
          let _2256_v17;
          _2256_v17 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_dafny.ZERO).minus(_this.f14));
          let _2257_v18;
          _2257_v18 = _dafny.Set.fromElements(_2256_v17, _2256_v17, (_dafny.Map.Empty.slice().updateUnsafe(p2,_2242_i0)).Merge(_2256_v17), _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_this.f14),(_dafny.ZERO).minus(_2242_i0)));
          _2257_v18 = _2257_v18;
          let _2258_v19;
          _2258_v19 = _module.D11.create_DC19(_2256_v17);
          let _2259_v20;
          _2259_v20 = _module.D11.create_DC21(_2258_v19);
          let _2260_v21;
          _2260_v21 = _module.D11.create_DC21(_2259_v20);
          let _2261_v22;
          _2261_v22 = _dafny.MultiSet.fromElements(_2260_v21, _2260_v21, _2260_v21, _2260_v21);
          let _2262_v23;
          let _nw434 = new _module.C7();
          _nw434.__ctor((_this).f17, (_2261_v22).IsProperSubsetOf(_2261_v22), _this.f14);
          _2262_v23 = _nw434;
          let _2263_v24;
          let _init70 = ((_2264_i0) => function (_2265_i1) {
            return (_2265_i1).plus(_2264_i0);
          })(_2242_i0);
          let _nw435 = Array((new BigNumber(10)).toNumber());
          for (let _i0_70 = 0; _i0_70 < new BigNumber(_nw435.length); _i0_70++) {
            _nw435[_i0_70] = _init70(new BigNumber(_i0_70));
          }
          _2263_v24 = _nw435;
          let _index373 = _module.__default.safeIndex(new BigNumber(97), new BigNumber((_2263_v24).length));
          (_2263_v24)[_index373] = _module.__default.safeDivisionInt(_this.f14, new BigNumber((_this.f18).length));
          let _2266_v25;
          _2266_v25 = _dafny.Seq.of(p2, p2, _2242_i0, (_this).fm12(_2242_i0, globalState), _2242_i0);
          let _2267_v26;
          let _nw436 = new _module.C7();
          _nw436.__ctor(p0, (_this).f17, new BigNumber((_dafny.Seq.of((_this).f17)).length));
          _2267_v26 = _nw436;
          let _2268_v27;
          _2268_v27 = _module.D20.create_DC42(new BigNumber((_dafny.Seq.of(p1)).length), false, _2267_v26);
          let _2269_v28;
          _2269_v28 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2266_v25).length),_module.__default.fm21(globalState));
          let _2270_v29;
          _2270_v29 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm64(new BigNumber(66), _2242_i0, _module.__default.fm21(globalState), new BigNumber(852), globalState)).length),p0), (_2269_v28).update(_2242_i0, (_this).f17));
          let _index374 = _module.__default.safeIndex(new BigNumber(97), new BigNumber((_2263_v24).length));
          let _rhs426 = (_this).fm12((_2266_v25)[_module.__default.safeIndex(_this.f14, new BigNumber((_2266_v25).length))], globalState);
          let _rhs427 = ((_2242_i0).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(631)), ((_2271_p3) => function (_2272_i2) {
            return _2271_p3;
          })(p3))).length))).isLessThanOrEqualTo((_this.f14).plus(p2));
          let _rhs428 = ((p0) ? (_module.__default.safeModuloInt(_2242_i0, _2242_i0)) : (((_2268_v27).dtor_cf65).plus(new BigNumber((_2270_v29).length))));
          let _lhs331 = globalState;
          let _lhs332 = _2263_v24;
          let _lhs333 = _module.__default.safeIndex(new BigNumber(97), new BigNumber((_2263_v24).length));
          _lhs331.f7 = _rhs426;
          r3 = _rhs427;
          _lhs332[_lhs333] = _rhs428;
          let _2273_v30;
          _2273_v30 = _dafny.Map.Empty.slice().updateUnsafe(p3,_module.D20.create_DC42(new BigNumber(-219), p1, _2267_v26));
          let _2274_v31;
          _2274_v31 = _module.D23.create_DC47(_2273_v30);
          let _pat_let_tv38 = _2267_v26;
          let _2275_v32;
          let _nw437 = Array((new BigNumber(17)).toNumber());
          _nw437[(_dafny.ZERO).toNumber()] = (_2273_v30).Merge(_2273_v30);
          _nw437[(_dafny.ONE).toNumber()] = (_2273_v30).Merge(_2273_v30);
          _nw437[(new BigNumber(2)).toNumber()] = _2273_v30;
          _nw437[(new BigNumber(3)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),_module.D20.create_DC42(_2267_v26.f14, (_this).f17, _2267_v26))).update(p3, _2268_v27);
          _nw437[(new BigNumber(4)).toNumber()] = (_2273_v30).update((_this.f18)[_module.__default.safeIndex(_this.f14, new BigNumber((_this.f18).length))], _2268_v27);
          _nw437[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p3,function (_pat_let48_0) {
            return function (_2276_dt__update__tmp_h0) {
              return function (_pat_let49_0) {
                return function (_2277_dt__update_hcf67_h0) {
                  return _module.D20.create_DC42((_2276_dt__update__tmp_h0).dtor_cf65, (_2276_dt__update__tmp_h0).dtor_cf66, _2277_dt__update_hcf67_h0);
                }(_pat_let49_0);
              }(_pat_let_tv38);
            }(_pat_let48_0);
          }(_2268_v27));
          _nw437[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p3,_2268_v27);
          _nw437[(new BigNumber(7)).toNumber()] = (_2273_v30).Merge((_2274_v31).dtor_cf73);
          _nw437[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p3,_2268_v27);
          _nw437[(new BigNumber(9)).toNumber()] = (_2273_v30).Merge(_2273_v30);
          _nw437[(new BigNumber(10)).toNumber()] = _2273_v30;
          _nw437[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p3,_2268_v27);
          _nw437[(new BigNumber(12)).toNumber()] = _2273_v30;
          _nw437[(new BigNumber(13)).toNumber()] = _2273_v30;
          _nw437[(new BigNumber(14)).toNumber()] = _2273_v30;
          _nw437[(new BigNumber(15)).toNumber()] = (_2273_v30).Merge(_2273_v30);
          _nw437[(new BigNumber(16)).toNumber()] = _2273_v30;
          _2275_v32 = _nw437;
          let _index375 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_2275_v32).length));
          (_2275_v32)[_index375] = _2273_v30;
          let _index376 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_2275_v32).length));
          let _rhs429 = ((_this).f17) === (p0);
          let _rhs430 = _2273_v30;
          let _rhs431 = (_module.__default.fm0((_this).f17, globalState)).plus(_2242_i0);
          let _lhs334 = _2275_v32;
          let _lhs335 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_2275_v32).length));
          let _lhs336 = globalState;
          r3 = _rhs429;
          _lhs334[_lhs335] = _rhs430;
          _lhs336.f2 = _rhs431;
          let _index377 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_2254_v15).length));
          (_2254_v15)[_index377] = p0;
          let _2278_v33;
          _2278_v33 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('b'.codePointAt(0)));
          let _index378 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_2254_v15).length));
          (_2254_v15)[_index378] = !((_2278_v33).IsSubsetOf(_2278_v33));
        } else {
          (_this).f16 = p1;
          (globalState).f2 = _2242_i0;
          let _2279_v34;
          _2279_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_this.f18).length),(_this).f17);
          let _2280_v35;
          _2280_v35 = _dafny.Seq.of(_2279_v34);
          _2280_v35 = _dafny.Seq.update(_2280_v35, _module.__default.safeIndex(p2, new BigNumber((_2280_v35).length)), _2279_v34);
          (_this).f14 = _2242_i0;
          (globalState).f7 = p2;
        }
        (globalState).f2 = p2;
        let _2281_v36;
        _2281_v36 = _dafny.Map.Empty.slice().updateUnsafe(p0,_module.__default.fm0(_this.f16, globalState));
        let _2282_v37;
        _2282_v37 = _dafny.MultiSet.fromElements((((_2281_v36).contains(_this.f16)) ? ((_2281_v36).get(_this.f16)) : ((_dafny.ZERO).minus(p2))), (_2242_i0).multipliedBy(p2));
        (_this).f14 = (((_2282_v37).contains((_dafny.ZERO).minus(p2))) ? ((_2282_v37).get((_dafny.ZERO).minus(p2))) : (_2242_i0));
      }
      (globalState).f2 = _this.f14;
      r3 = false;
      let _2283_i3;
      _2283_i3 = _dafny.ZERO;
      L13: {
        while (((p2).multipliedBy(_module.__default.fm0(false, globalState))).isLessThanOrEqualTo(_this.f14)) {
          C13: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2283_i3)) {
              break L13;
            }
            _2283_i3 = (_2283_i3).plus(_dafny.ONE);
            let _2284_v38;
            _2284_v38 = _dafny.MultiSet.fromElements(p2, p2);
            let _2285_v39;
            _2285_v39 = _dafny.Set.fromElements(_this.f14, _this.f14, _this.f14);
            let _2286_v40;
            let _nw438 = Array((new BigNumber(20)).toNumber());
            _nw438[(_dafny.ZERO).toNumber()] = p0;
            _nw438[(_dafny.ONE).toNumber()] = p1;
            _nw438[(new BigNumber(2)).toNumber()] = (_this.f14).isLessThan(p2);
            _nw438[(new BigNumber(3)).toNumber()] = _this.f16;
            _nw438[(new BigNumber(4)).toNumber()] = (_2284_v38).equals(_dafny.MultiSet.fromElements(_this.f14, new BigNumber(799)));
            _nw438[(new BigNumber(5)).toNumber()] = false;
            _nw438[(new BigNumber(6)).toNumber()] = (_this).f17;
            _nw438[(new BigNumber(7)).toNumber()] = false;
            _nw438[(new BigNumber(8)).toNumber()] = false;
            _nw438[(new BigNumber(9)).toNumber()] = (_this).f17;
            _nw438[(new BigNumber(10)).toNumber()] = _this.f16;
            _nw438[(new BigNumber(11)).toNumber()] = p1;
            _nw438[(new BigNumber(12)).toNumber()] = p0;
            _nw438[(new BigNumber(13)).toNumber()] = p0;
            _nw438[(new BigNumber(14)).toNumber()] = (new BigNumber((_2285_v39).length)).isLessThanOrEqualTo(p2);
            _nw438[(new BigNumber(15)).toNumber()] = _this.f16;
            _nw438[(new BigNumber(16)).toNumber()] = p1;
            _nw438[(new BigNumber(17)).toNumber()] = _module.__default.fm21(globalState);
            _nw438[(new BigNumber(18)).toNumber()] = p1;
            _nw438[(new BigNumber(19)).toNumber()] = !((false) === ((_this).f17));
            _2286_v40 = _nw438;
            let _index379 = _module.__default.safeIndex(new BigNumber(576), new BigNumber((_2286_v40).length));
            (_2286_v40)[_index379] = !(p1) || (p1);
            let _index380 = _module.__default.safeIndex(new BigNumber(576), new BigNumber((_2286_v40).length));
            (_2286_v40)[_index380] = _this.f16;
            let _2287_v41;
            let _nw439 = Array((new BigNumber(13)).toNumber()).fill(_module.D9.Default());
            _2287_v41 = _nw439;
            let _2288_v42;
            let _init71 = ((_2289_p3) => function (_2290_i4) {
              return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-351)), ((_2291_p3) => function (_2292_i5) {
                return _2291_p3;
              })(_2289_p3));
            })(p3);
            let _nw440 = Array((new BigNumber(13)).toNumber());
            for (let _i0_71 = 0; _i0_71 < new BigNumber(_nw440.length); _i0_71++) {
              _nw440[_i0_71] = _init71(new BigNumber(_i0_71));
            }
            _2288_v42 = _nw440;
            let _index381 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_2288_v42).length));
            (_2288_v42)[_index381] = _dafny.Seq.UnicodeFromString("xugnav");
            let _index382 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_2288_v42).length));
            let _rhs432 = _2287_v41;
            let _rhs433 = _module.__default.safeModuloInt(_this.f14, (((_2286_v40)[_module.__default.safeIndex(new BigNumber(576), new BigNumber((_2286_v40).length))]) ? (p2) : ((_dafny.ZERO).minus(p2))));
            let _rhs434 = _this.f18;
            let _lhs337 = globalState;
            let _lhs338 = _2288_v42;
            let _lhs339 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_2288_v42).length));
            _2287_v41 = _rhs432;
            _lhs337.f2 = _rhs433;
            _lhs338[_lhs339] = _rhs434;
            (_this).f16 = (_this).f17;
            r0 = ((_this.f16) ? (new BigNumber((_2238_v0).cardinality())) : (new BigNumber(338)));
          }
        }
      }
      let _hi19 = _this.f14;
      for (let _2293_i6 = p2; _2293_i6.isLessThan(_hi19); _2293_i6 = _2293_i6.plus(_dafny.ONE)) {
        let _2294_v43;
        let _nw441 = new _module.C1();
        _nw441.__ctor(!(p0), _2293_i6);
        _2294_v43 = _nw441;
        let _2295_v44;
        _2295_v44 = _dafny.Seq.of(p0, _2294_v43.f25, _this.f16);
        let _2296_v45;
        let _nw442 = new _module.C3();
        _nw442.__ctor(_this.f14, _dafny.Seq.Concat(_module.__default.fm42(_2294_v43.f25, _this.f14, new BigNumber((_dafny.Seq.of(!(_2294_v43.f25))).length), false, globalState), _dafny.Seq.update(_this.f18, _module.__default.safeIndex(new BigNumber((_2295_v44).length), new BigNumber((_this.f18).length)), p3)), !(((_this.f16) ? (p0) : (true))), _this.f16);
        _2296_v45 = _nw442;
        let _2297_v46;
        _2297_v46 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,p2);
        let _2298_v47;
        _2298_v47 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_2297_v46).length));
        let _2299_v48;
        _2299_v48 = _dafny.MultiSet.fromElements(_2298_v47);
        let _2300_v49;
        _2300_v49 = _dafny.Map.Empty.slice().updateUnsafe(_2293_i6,(_2299_v48).Union(_2299_v48));
        _2300_v49 = (_2300_v49).update(p2, _2299_v48);
        let _source34 = _module.__default.fm65((_2296_v45).f22, globalState);
        if (_source34.is_DC4) {
          let _2301___mcc_h0 = (_source34).cf4;
          let _2302___mcc_h1 = (_source34).cf5;
          let _2303___mcc_h2 = (_source34).cf6;
          let _2304___mcc_h3 = (_source34).cf7;
          let _2305___mcc_h4 = (_source34).cf8;
          let _2306_cf8 = _2305___mcc_h4;
          let _2307_cf7 = _2304___mcc_h3;
          let _2308_cf6 = _2303___mcc_h2;
          let _2309_cf5 = _2302___mcc_h1;
          let _2310_cf4 = _2301___mcc_h0;
          r0 = (_2296_v45).f22;
          let _2311_v50;
          let _nw443 = Array((new BigNumber(19)).toNumber()).fill([]);
          _2311_v50 = _nw443;
          let _2312_v51;
          let _init72 = function (_2313_i7) {
            return (_2313_i7).multipliedBy(new BigNumber(-822));
          };
          let _nw444 = Array((new BigNumber(16)).toNumber());
          for (let _i0_72 = 0; _i0_72 < new BigNumber(_nw444.length); _i0_72++) {
            _nw444[_i0_72] = _init72(new BigNumber(_i0_72));
          }
          _2312_v51 = _nw444;
          let _index383 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_2311_v50).length));
          (_2311_v50)[_index383] = _2312_v51;
          let _index384 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_2311_v50).length));
          let _nw445 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          (_2311_v50)[_index384] = _nw445;
          let _2314_v52;
          _2314_v52 = _dafny.MultiSet.fromElements(p3);
          let _2315_v53;
          _2315_v53 = _dafny.Map.Empty.slice().updateUnsafe(p2,_2306_cf8);
          let _2316_v54;
          _2316_v54 = _dafny.MultiSet.fromElements(_2315_v53, _2315_v53);
          let _arr0 = (_2311_v50)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_2311_v50).length))];
          let _index385 = _module.__default.safeIndex(new BigNumber(240), new BigNumber(((_2311_v50)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_2311_v50).length))]).length));
          _arr0[_index385] = (((_2314_v52).contains(new _dafny.CodePoint('u'.codePointAt(0)))) ? ((_2314_v52).get(new _dafny.CodePoint('u'.codePointAt(0)))) : (new BigNumber((_2316_v54).cardinality())));
          let _2317_v55;
          _2317_v55 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_2296_v45).f22);
          let _arr1 = (_2311_v50)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_2311_v50).length))];
          let _index386 = _module.__default.safeIndex(new BigNumber(240), new BigNumber(((_2311_v50)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_2311_v50).length))]).length));
          _arr1[_index386] = _module.__default.fm0((_2317_v55).contains(new _dafny.CodePoint('k'.codePointAt(0))), globalState);
          let _2318_v56;
          _2318_v56 = _dafny.Seq.of((_2296_v45).f22, ((_2311_v50)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_2311_v50).length))])[_module.__default.safeIndex(new BigNumber(240), new BigNumber(((_2311_v50)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_2311_v50).length))]).length))], (_2296_v45).f22);
          let _2319_v57;
          let _nw446 = Array((new BigNumber(24)).toNumber());
          _nw446[(_dafny.ZERO).toNumber()] = _module.__default.fm21(globalState);
          _nw446[(_dafny.ONE).toNumber()] = true;
          _nw446[(new BigNumber(2)).toNumber()] = (_this).f17;
          _nw446[(new BigNumber(3)).toNumber()] = (_2294_v43.f25) && (p1);
          _nw446[(new BigNumber(4)).toNumber()] = _module.__default.fm27(new BigNumber((_2315_v53).length), new BigNumber(((_2296_v45).f23).length), _2293_i6, _2297_v46, globalState);
          _nw446[(new BigNumber(5)).toNumber()] = !_dafny.Seq.contains(_2318_v56, ((_2311_v50)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_2311_v50).length))])[_module.__default.safeIndex(new BigNumber(240), new BigNumber(((_2311_v50)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_2311_v50).length))]).length))]);
          _nw446[(new BigNumber(6)).toNumber()] = !(_2238_v0).contains(p1);
          _nw446[(new BigNumber(7)).toNumber()] = _2307_cf7;
          _nw446[(new BigNumber(8)).toNumber()] = _module.__default.fm43((((_2239_v1).contains((_this).fm14(globalState))) ? ((_2239_v1).get((_this).fm14(globalState))) : (_2310_cf4)), globalState);
          _nw446[(new BigNumber(9)).toNumber()] = _this.f16;
          _nw446[(new BigNumber(10)).toNumber()] = (_this).f17;
          _nw446[(new BigNumber(11)).toNumber()] = (_this).f17;
          _nw446[(new BigNumber(12)).toNumber()] = p1;
          _nw446[(new BigNumber(13)).toNumber()] = _2306_cf8;
          _nw446[(new BigNumber(14)).toNumber()] = (_2295_v44)[_module.__default.safeIndex((_2296_v45).f22, new BigNumber((_2295_v44).length))];
          _nw446[(new BigNumber(15)).toNumber()] = _2306_cf8;
          _nw446[(new BigNumber(16)).toNumber()] = _2294_v43.f25;
          _nw446[(new BigNumber(17)).toNumber()] = p0;
          _nw446[(new BigNumber(18)).toNumber()] = !_dafny.areEqual(_2318_v56, _2318_v56);
          _nw446[(new BigNumber(19)).toNumber()] = (_this).f17;
          _nw446[(new BigNumber(20)).toNumber()] = _2294_v43.f25;
          _nw446[(new BigNumber(21)).toNumber()] = _2307_cf7;
          _nw446[(new BigNumber(22)).toNumber()] = true;
          _nw446[(new BigNumber(23)).toNumber()] = _dafny.Seq.IsPrefixOf((_2296_v45).f23, _dafny.Seq.UnicodeFromString("awdr"));
          _2319_v57 = _nw446;
          let _index387 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_2319_v57).length));
          (_2319_v57)[_index387] = _this.f16;
          let _index388 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_2319_v57).length));
          let _rhs435 = _2307_cf7;
          let _rhs436 = (new BigNumber(((_2296_v45).f23).length)).isLessThanOrEqualTo((_2296_v45).f22);
          let _lhs340 = _2319_v57;
          let _lhs341 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_2319_v57).length));
          _lhs340[_lhs341] = _rhs435;
          _2307_cf7 = _rhs436;
        } else if (_source34.is_DC3) {
          let _2320___mcc_h5 = (_source34).cf3;
          let _2321_cf3 = _2320___mcc_h5;
          let _2322_v58;
          let _init73 = ((_2323_cf3, _2324_p1, _2325_p0) => function (_2326_i8) {
            return (_dafny.Map.Empty.slice().updateUnsafe(_2324_p1,_2325_p0)).contains(_2323_cf3);
          })(_2321_cf3, p1, p0);
          let _nw447 = Array((new BigNumber(23)).toNumber());
          for (let _i0_73 = 0; _i0_73 < new BigNumber(_nw447.length); _i0_73++) {
            _nw447[_i0_73] = _init73(new BigNumber(_i0_73));
          }
          _2322_v58 = _nw447;
          let _index389 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_2322_v58).length));
          (_2322_v58)[_index389] = p1;
          let _2327_v59;
          let _nw448 = Array((new BigNumber(17)).toNumber());
          _nw448[(_dafny.ZERO).toNumber()] = p3;
          _nw448[(_dafny.ONE).toNumber()] = p3;
          _nw448[(new BigNumber(2)).toNumber()] = p3;
          _nw448[(new BigNumber(3)).toNumber()] = p3;
          _nw448[(new BigNumber(4)).toNumber()] = p3;
          _nw448[(new BigNumber(5)).toNumber()] = p3;
          _nw448[(new BigNumber(6)).toNumber()] = p3;
          _nw448[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('g'.codePointAt(0));
          _nw448[(new BigNumber(8)).toNumber()] = p3;
          _nw448[(new BigNumber(9)).toNumber()] = p3;
          _nw448[(new BigNumber(10)).toNumber()] = p3;
          _nw448[(new BigNumber(11)).toNumber()] = p3;
          _nw448[(new BigNumber(12)).toNumber()] = p3;
          _nw448[(new BigNumber(13)).toNumber()] = new _dafny.CodePoint('m'.codePointAt(0));
          _nw448[(new BigNumber(14)).toNumber()] = p3;
          _nw448[(new BigNumber(15)).toNumber()] = p3;
          _nw448[(new BigNumber(16)).toNumber()] = p3;
          _2327_v59 = _nw448;
          let _2328_v60;
          _2328_v60 = _dafny.Map.Empty.slice().updateUnsafe((_2296_v45).f22,_2327_v59);
          let _2329_v61;
          _2329_v61 = _dafny.Seq.of((new BigNumber((_2328_v60).length)).multipliedBy(p2), _2293_i6, _module.__default.fm0(p0, globalState), (_2296_v45).f22, _module.__default.fm0(_this.f16, globalState));
          let _index390 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_2322_v58).length));
          let _rhs437 = ((new BigNumber(((_2296_v45).f23).length)).minus(new BigNumber(-659))).multipliedBy(new BigNumber((_dafny.Seq.Concat(_this.f18, (_2296_v45).f23)).length));
          let _rhs438 = !(_2294_v43.f25);
          let _rhs439 = !((_this).f17);
          let _rhs440 = (_2329_v61)[_module.__default.safeIndex(_2293_i6, new BigNumber((_2329_v61).length))];
          let _lhs342 = _2294_v43;
          let _lhs343 = _2322_v58;
          let _lhs344 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_2322_v58).length));
          let _lhs345 = _this;
          r0 = _rhs437;
          _lhs342.f25 = _rhs438;
          _lhs343[_lhs344] = _rhs439;
          _lhs345.f14 = _rhs440;
          (_2294_v43).f25 = _module.__default.fm43(false, globalState);
          let _2330_v62;
          _2330_v62 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(p3, p3),(_dafny.ZERO).minus(_this.f14));
          _2330_v62 = (_2330_v62).update((_2296_v45).f23, new BigNumber(271));
          let _2331_v63;
          let _nw449 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2331_v63 = _nw449;
          let _2332_v64;
          let _nw450 = Array((new BigNumber(22)).toNumber());
          _nw450[(_dafny.ZERO).toNumber()] = _2331_v63;
          _nw450[(_dafny.ONE).toNumber()] = _2331_v63;
          _nw450[(new BigNumber(2)).toNumber()] = _2331_v63;
          _nw450[(new BigNumber(3)).toNumber()] = _2331_v63;
          _nw450[(new BigNumber(4)).toNumber()] = _2331_v63;
          _nw450[(new BigNumber(5)).toNumber()] = _2331_v63;
          _nw450[(new BigNumber(6)).toNumber()] = _2331_v63;
          _nw450[(new BigNumber(7)).toNumber()] = _2331_v63;
          _nw450[(new BigNumber(8)).toNumber()] = _2331_v63;
          _nw450[(new BigNumber(9)).toNumber()] = _2331_v63;
          _nw450[(new BigNumber(10)).toNumber()] = _2331_v63;
          _nw450[(new BigNumber(11)).toNumber()] = _2331_v63;
          _nw450[(new BigNumber(12)).toNumber()] = ((true) ? (_2331_v63) : (_2331_v63));
          _nw450[(new BigNumber(13)).toNumber()] = _2331_v63;
          _nw450[(new BigNumber(14)).toNumber()] = _2331_v63;
          _nw450[(new BigNumber(15)).toNumber()] = _2331_v63;
          _nw450[(new BigNumber(16)).toNumber()] = _2331_v63;
          _nw450[(new BigNumber(17)).toNumber()] = _2331_v63;
          _nw450[(new BigNumber(18)).toNumber()] = _2331_v63;
          _nw450[(new BigNumber(19)).toNumber()] = _2331_v63;
          _nw450[(new BigNumber(20)).toNumber()] = _2331_v63;
          _nw450[(new BigNumber(21)).toNumber()] = _2331_v63;
          _2332_v64 = _nw450;
          let _index391 = _module.__default.safeIndex(new BigNumber(40), new BigNumber((_2332_v64).length));
          (_2332_v64)[_index391] = _2331_v63;
          let _index392 = _module.__default.safeIndex(new BigNumber(40), new BigNumber((_2332_v64).length));
          (_2332_v64)[_index392] = _2331_v63;
        } else {
          let _2333___mcc_h6 = (_source34).cf9;
          let _2334_cf9 = _2333___mcc_h6;
          let _rhs441 = _this.f14;
          let _rhs442 = _2294_v43.f25;
          let _lhs346 = globalState;
          _lhs346.f2 = _rhs441;
          r3 = _rhs442;
          let _2335_v65;
          _2335_v65 = _dafny.Seq.of(_dafny.MultiSet.fromElements((_2296_v45).f22));
          let _2336_v66;
          _2336_v66 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.Concat(_2335_v65, _2335_v65));
          (globalState).f7 = (_dafny.ZERO).minus(new BigNumber(((((_2336_v66).contains(_2294_v43.f25)) ? ((_2336_v66).get(_2294_v43.f25)) : (_2335_v65))).length));
          (_2294_v43).f25 = ((p0) ? ((_this).f17) : ((false) === (!(_this.f16))));
          let _2337_v67;
          let _nw451 = Array((new BigNumber(7)).toNumber());
          _nw451[(_dafny.ZERO).toNumber()] = _this.f14;
          _nw451[(_dafny.ONE).toNumber()] = new BigNumber(613);
          _nw451[(new BigNumber(2)).toNumber()] = _this.f14;
          _nw451[(new BigNumber(3)).toNumber()] = (_2296_v45).f22;
          _nw451[(new BigNumber(4)).toNumber()] = (_2296_v45).f22;
          _nw451[(new BigNumber(5)).toNumber()] = p2;
          _nw451[(new BigNumber(6)).toNumber()] = (_2296_v45).f22;
          _2337_v67 = _nw451;
          let _2338_v68;
          _2338_v68 = _dafny.Seq.of(_2337_v67);
          let _2339_v69;
          _2339_v69 = _dafny.Set.fromElements(p0);
          let _2340_v70;
          _2340_v70 = _module.D15.create_DC29(_2339_v69);
          let _2341_v71;
          _2341_v71 = _dafny.Map.Empty.slice().updateUnsafe(_2338_v68,_2340_v70);
          _2341_v71 = (_2341_v71).update(_2338_v68, _2340_v70);
        }
      }
      r0 = p2;
      let _2342_v72;
      _2342_v72 = _dafny.Set.fromElements(_this.f14);
      let _2343_v73;
      _2343_v73 = _dafny.Map.Empty.slice().updateUnsafe(_2342_v72,_this.f14);
      let _2344_v74;
      _2344_v74 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,p2);
      let _2345_v76;
      _2345_v76 = _dafny.Set.fromElements(_2344_v74, function () {
        let _coll81 = new _dafny.Map();
        for (const _compr_81 of _dafny.IntegerRange(new BigNumber(117), new BigNumber(-228))) {
          let _2346_v75 = _compr_81;
          if (((new BigNumber(117)).isLessThanOrEqualTo(_2346_v75)) && ((_2346_v75).isLessThan(new BigNumber(-228)))) {
            _coll81.push([(_2346_v75).minus(_this.f14),_this.f14]);
          }
        }
        return _coll81;
      }(), _2344_v74, _2344_v74, _2344_v74);
      r1 = !(!((_2345_v76).IsProperSubsetOf(_module.__default.fm66(p2, _2343_v73, p2, p0, globalState))));
      let _2347_v77;
      _2347_v77 = _module.D13.create_DC25(_dafny.Map.Empty.slice().updateUnsafe(_this.f18,p2));
      let _2348_v78;
      _2348_v78 = _dafny.Map.Empty.slice().updateUnsafe(_this.f18,p2);
      let _pat_let_tv39 = _2348_v78;
      r2 = function (_source35) {
        if (_source35.is_DC26) {
          return _dafny.Seq.Concat(_this.f18, _this.f18);
        } else {
          let _2349___mcc_h7 = (_source35).cf38;
          let _2350_cf38 = _2349___mcc_h7;
          return _this.f18;
        }
      }(_module.D13.create_DC25((function (_pat_let50_0) {
  return function (_2351_dt__update__tmp_h1) {
    return function (_pat_let51_0) {
      return function (_2352_dt__update_hcf38_h0) {
        return _module.D13.create_DC25(_2352_dt__update_hcf38_h0);
      }(_pat_let51_0);
    }(_pat_let_tv39);
  }(_pat_let50_0);
}(_2347_v77)).dtor_cf38));
      r3 = p1;
      return [r0, r1, r2, r3];
    }
    m7(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let _2353_v0;
      _2353_v0 = _dafny.Seq.of(p0);
      let _2354_v1;
      _2354_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,_dafny.MultiSet.FromArray(_2353_v0));
      let _2355_v2;
      _2355_v2 = _dafny.MultiSet.fromElements(_this.f16);
      let _2356_v3;
      _2356_v3 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(565),p0);
      let _2357_v4;
      _2357_v4 = _module.D16.create_DC33(_this.f16, (_this).f17, (((_2355_v2).contains(_this.f16)) ? ((_2355_v2).get(_this.f16)) : (new BigNumber((_2356_v3).length))), (_this).f17);
      let _2358_v5;
      _2358_v5 = _module.D11.create_DC19(_2356_v3);
      let _2359_v6;
      _2359_v6 = _dafny.Seq.of(_2358_v5, _2358_v5);
      let _2360_v7;
      _2360_v7 = _dafny.MultiSet.fromElements(_this.f14, new BigNumber((_dafny.Seq.update(_2359_v6, _module.__default.safeIndex(_this.f14, new BigNumber((_2359_v6).length)), _2358_v5)).length));
      let _2361_i0;
      _2361_i0 = _dafny.ZERO;
      L14: {
        while (_module.__default.fm27(new BigNumber(((((_2354_v1).contains((_2357_v4).dtor_cf50)) ? ((_2354_v1).get((_2357_v4).dtor_cf50)) : (_2360_v7))).cardinality()), (_dafny.ZERO).minus(new BigNumber((_this.f18).length)), _this.f14, _2356_v3, globalState)) {
          C14: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2361_i0)) {
              break L14;
            }
            _2361_i0 = (_2361_i0).plus(_dafny.ONE);
            let _2362_v8;
            _2362_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_this.f18).length),(_this).f17);
            (globalState).f7 = (_dafny.ZERO).minus((new BigNumber((_2362_v8).length)).multipliedBy(new BigNumber((_this.f18).length)));
            (globalState).f7 = _this.f14;
            let _2363_v9;
            let _init74 = function (_2364_i1) {
              return (_2364_i1).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f14)));
            };
            let _nw452 = Array((new BigNumber(21)).toNumber());
            for (let _i0_74 = 0; _i0_74 < new BigNumber(_nw452.length); _i0_74++) {
              _nw452[_i0_74] = _init74(new BigNumber(_i0_74));
            }
            _2363_v9 = _nw452;
            let _2365_v10;
            _2365_v10 = _dafny.MultiSet.fromElements(_2363_v9, _2363_v9);
            (globalState).f7 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((((_2365_v10).contains(_2363_v9)) ? ((_2365_v10).get(_2363_v9)) : (_this.f14)),((_this.f16) ? (_this.f14) : (new BigNumber((_this.f18).length))))).length);
            let _2366_v11;
            let _nw453 = new _module.C4();
            _nw453.__ctor(!(_this.f16) || (_this.f16), (_this).f17, (_2365_v10).equals(_2365_v10));
            _2366_v11 = _nw453;
          }
        }
      }
      if (false) {
        (_this).f14 = _this.f14;
        let _2367_v12;
        _2367_v12 = _module.D21.create_DC44(_this.f14, (_this).f17);
        let _2368_v13;
        _2368_v13 = _dafny.Seq.of((_this).f17, (_2367_v12).dtor_cf70, _this.f16);
        let _2369_v14;
        let _nw454 = new _module.C6();
        _nw454.__ctor(p0, _dafny.Seq.IsPrefixOf(_2368_v13, _2368_v13), _this.f16, (_this).f17, (p0).multipliedBy(_this.f14));
        _2369_v14 = _nw454;
        let _2370_v15;
        let _nw455 = new _module.C3();
        _nw455.__ctor(p0, _this.f18, _this.f16, _this.f16);
        _2370_v15 = _nw455;
        let _2371_v16;
        _2371_v16 = _dafny.Map.Empty.slice().updateUnsafe(_2370_v15,(_this).f17);
        (globalState).f7 = new BigNumber(((_2371_v16).Merge(_2371_v16)).length);
        let _2372_v17;
        let _nw456 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
        _2372_v17 = _nw456;
        let _index393 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_2372_v17).length));
        (_2372_v17)[_index393] = (((_2369_v14).f20) ? (new BigNumber(12)) : (p0));
        let _index394 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_2372_v17).length));
        (_2372_v17)[_index394] = (p0).plus(new BigNumber(640));
        let _2373_v18;
        _2373_v18 = _module.D12.create_DC22(_dafny.Seq.UnicodeFromString("tw"));
        let _source36 = _2373_v18;
        if (_source36.is_DC23) {
          let _2374___mcc_h0 = (_source36).cf30;
          let _2375___mcc_h1 = (_source36).cf31;
          let _2376___mcc_h2 = (_source36).cf32;
          let _2377_cf32 = _2376___mcc_h2;
          let _2378_cf31 = _2375___mcc_h1;
          let _2379_cf30 = _2374___mcc_h0;
          let _2380_v19;
          let _nw457 = Array((new BigNumber(18)).toNumber()).fill(false);
          _2380_v19 = _nw457;
          let _2381_v20;
          _2381_v20 = _2380_v19;
          _2380_v19 = (_2381_v20);
          let _2382_v21;
          let _nw458 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Map.Empty);
          _2382_v21 = _nw458;
          _2382_v21 = _2382_v21;
          let _index395 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((_2380_v19).length));
          (_2380_v19)[_index395] = (_this).f17;
          let _index396 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((_2380_v19).length));
          (_2380_v19)[_index396] = _2379_cf30;
          let _2383_v22;
          _2383_v22 = new _dafny.CodePoint('k'.codePointAt(0));
          _2383_v22 = _2383_v22;
        } else if (_source36.is_DC24) {
          let _2384___mcc_h3 = (_source36).cf33;
          let _2385___mcc_h4 = (_source36).cf34;
          let _2386___mcc_h5 = (_source36).cf35;
          let _2387___mcc_h6 = (_source36).cf36;
          let _2388___mcc_h7 = (_source36).cf37;
          let _2389_cf37 = _2388___mcc_h7;
          let _2390_cf36 = _2387___mcc_h6;
          let _2391_cf35 = _2386___mcc_h5;
          let _2392_cf34 = _2385___mcc_h4;
          let _2393_cf33 = _2384___mcc_h3;
          let _2394_v23;
          _2394_v23 = _dafny.Map.Empty.slice().updateUnsafe(_2393_cf33,(_2369_v14).f20);
          _2394_v23 = _2394_v23;
          let _2395_v24;
          _2395_v24 = _module.D7.create_DC12();
          let _2396_v25;
          _2396_v25 = _dafny.Map.Empty.slice().updateUnsafe((_2369_v14).f20,(_2370_v15).f22);
          let _2397_v26;
          _2397_v26 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_2396_v25);
          _2395_v24 = _module.__default.fm67((_2397_v26).Merge(_2397_v26), globalState);
          _2372_v17 = _2372_v17;
          let _2398_v27;
          let _init75 = ((_2399_v15) => function (_2400_i2) {
            return (_2399_v15).f23;
          })(_2370_v15);
          let _nw459 = Array((new BigNumber(17)).toNumber());
          for (let _i0_75 = 0; _i0_75 < new BigNumber(_nw459.length); _i0_75++) {
            _nw459[_i0_75] = _init75(new BigNumber(_i0_75));
          }
          _2398_v27 = _nw459;
          let _2401_v28;
          let _nw460 = Array((new BigNumber(18)).toNumber());
          _nw460[(_dafny.ZERO).toNumber()] = _2398_v27;
          _nw460[(_dafny.ONE).toNumber()] = _2398_v27;
          _nw460[(new BigNumber(2)).toNumber()] = _2398_v27;
          _nw460[(new BigNumber(3)).toNumber()] = _2398_v27;
          _nw460[(new BigNumber(4)).toNumber()] = _2398_v27;
          _nw460[(new BigNumber(5)).toNumber()] = _2398_v27;
          _nw460[(new BigNumber(6)).toNumber()] = _2398_v27;
          _nw460[(new BigNumber(7)).toNumber()] = _2398_v27;
          _nw460[(new BigNumber(8)).toNumber()] = _2398_v27;
          _nw460[(new BigNumber(9)).toNumber()] = _2398_v27;
          _nw460[(new BigNumber(10)).toNumber()] = _2398_v27;
          _nw460[(new BigNumber(11)).toNumber()] = _2398_v27;
          _nw460[(new BigNumber(12)).toNumber()] = _2398_v27;
          _nw460[(new BigNumber(13)).toNumber()] = _2398_v27;
          _nw460[(new BigNumber(14)).toNumber()] = _2398_v27;
          _nw460[(new BigNumber(15)).toNumber()] = ((_2390_cf36) ? (_2398_v27) : (_2398_v27));
          _nw460[(new BigNumber(16)).toNumber()] = _2398_v27;
          _nw460[(new BigNumber(17)).toNumber()] = _2398_v27;
          _2401_v28 = _nw460;
          let _index397 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_2401_v28).length));
          (_2401_v28)[_index397] = _2398_v27;
          let _2402_v29;
          _2402_v29 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_2369_v14).f19, new BigNumber((_2353_v0).length)),_2398_v27);
          let _index398 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_2401_v28).length));
          (_2401_v28)[_index398] = (((_2402_v29).contains((_2372_v17)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_2372_v17).length))])) ? ((_2402_v29).get((_2372_v17)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_2372_v17).length))])) : (_2398_v27));
        } else {
          let _2403___mcc_h8 = (_source36).cf29;
          let _2404_cf29 = _2403___mcc_h8;
          let _index399 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_2372_v17).length));
          (_2372_v17)[_index399] = (_dafny.ZERO).minus(p0);
          let _2405_v30;
          let _nw461 = Array((new BigNumber(13)).toNumber()).fill(_module.D9.Default());
          _2405_v30 = _nw461;
          let _2406_v31;
          _2406_v31 = _dafny.Seq.of(_2405_v30);
          let _2407_v32;
          _2407_v32 = _module.D3.create_DC3((_this).f17);
          let _2408_v33;
          _2408_v33 = _dafny.Seq.of(_module.D3.create_DC3((_this).f17), _2407_v32);
          let _2409_v34;
          _2409_v34 = _module.D12.create_DC24((_2369_v14).f20, _2368_v13, _2406_v31, (_2369_v14).f20, _2408_v33);
          let _2410_v35;
          let _nw462 = Array((new BigNumber(18)).toNumber());
          _nw462[(_dafny.ZERO).toNumber()] = _this.f16;
          _nw462[(_dafny.ONE).toNumber()] = (_this).f17;
          _nw462[(new BigNumber(2)).toNumber()] = (_this).f17;
          _nw462[(new BigNumber(3)).toNumber()] = (_this).f17;
          _nw462[(new BigNumber(4)).toNumber()] = (_2369_v14).f20;
          _nw462[(new BigNumber(5)).toNumber()] = (_2369_v14).f20;
          _nw462[(new BigNumber(6)).toNumber()] = (_2369_v14).f20;
          _nw462[(new BigNumber(7)).toNumber()] = (_2369_v14).f20;
          _nw462[(new BigNumber(8)).toNumber()] = (_2369_v14).f20;
          _nw462[(new BigNumber(9)).toNumber()] = (_this).f17;
          _nw462[(new BigNumber(10)).toNumber()] = (_this).f17;
          _nw462[(new BigNumber(11)).toNumber()] = _this.f16;
          _nw462[(new BigNumber(12)).toNumber()] = (_this).f17;
          _nw462[(new BigNumber(13)).toNumber()] = (_2409_v34).dtor_cf36;
          _nw462[(new BigNumber(14)).toNumber()] = (_this).f17;
          _nw462[(new BigNumber(15)).toNumber()] = (_2369_v14).f20;
          _nw462[(new BigNumber(16)).toNumber()] = (_2369_v14).f20;
          _nw462[(new BigNumber(17)).toNumber()] = _this.f16;
          _2410_v35 = _nw462;
          let _2411_v36;
          _2411_v36 = _dafny.Seq.of(_2410_v35, _2410_v35, _2410_v35, _2410_v35, _2410_v35);
          let _2412_v37;
          _2412_v37 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(904)).minus((_2370_v15).f22),_dafny.Seq.update(_dafny.Seq.Concat(_2411_v36, _2411_v36), _module.__default.safeIndex((_2370_v15).f22, new BigNumber((_dafny.Seq.Concat(_2411_v36, _2411_v36)).length)), _2410_v35));
          _2412_v37 = (_2412_v37).update((_2369_v14).f19, _dafny.Seq.Concat(_2411_v36, _2411_v36));
          let _index400 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_2372_v17).length));
          (_2372_v17)[_index400] = _this.f14;
          let _2413_v38;
          _2413_v38 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),_dafny.Seq.IsProperPrefixOf(_2353_v0, _2353_v0));
          let _2414_v39;
          _2414_v39 = new _dafny.CodePoint('l'.codePointAt(0));
          _2413_v38 = (_2413_v38).update(_2414_v39, (_2409_v34).dtor_cf36);
        }
      } else {
        let _2415_v40;
        _2415_v40 = new _dafny.CodePoint('y'.codePointAt(0));
        let _2416_v41;
        _2416_v41 = _module.D21.create_DC43(_dafny.MultiSet.fromElements(_2415_v40, _2415_v40));
        let _2417_v42;
        _2417_v42 = _dafny.Seq.of(_2353_v0);
        let _2418_v43;
        _2418_v43 = _dafny.Seq.of(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(877)), ((_2419_p0) => function (_2420_i3) {
          return _2419_p0;
        })(p0))), _2417_v42);
        let _2421_v44;
        _2421_v44 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,(_2418_v43)[_module.__default.safeIndex(_this.f14, new BigNumber((_2418_v43).length))]);
        let _rhs443 = new BigNumber((_this.f18).length);
        let _rhs444 = _2416_v41;
        let _rhs445 = _dafny.Seq.Concat(_dafny.Seq.update((((_2421_v44).contains(true)) ? ((_2421_v44).get(true)) : (_dafny.Seq.of(_dafny.Seq.of(p0, _this.f14, _this.f14)))), _module.__default.safeIndex(_this.f14, new BigNumber(((((_2421_v44).contains(true)) ? ((_2421_v44).get(true)) : (_dafny.Seq.of(_dafny.Seq.of(p0, _this.f14, _this.f14))))).length)), _2353_v0), _2417_v42);
        let _rhs446 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lbxjip"), _this.f18)).length);
        let _lhs347 = globalState;
        let _lhs348 = globalState;
        _lhs347.f7 = _rhs443;
        _2416_v41 = _rhs444;
        _2417_v42 = _rhs445;
        _lhs348.f2 = _rhs446;
        let _2422_v45;
        let _init76 = ((_2423_v40) => function (_2424_i4) {
          return _2423_v40;
        })(_2415_v40);
        let _nw463 = Array((new BigNumber(25)).toNumber());
        for (let _i0_76 = 0; _i0_76 < new BigNumber(_nw463.length); _i0_76++) {
          _nw463[_i0_76] = _init76(new BigNumber(_i0_76));
        }
        _2422_v45 = _nw463;
        let _2425_v46;
        _2425_v46 = _dafny.Map.Empty.slice().updateUnsafe(_2422_v45,(_this).f17);
        (globalState).f2 = new BigNumber((_2425_v46).length);
        let _2426_v47;
        _2426_v47 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,(_this).f17);
        let _2427_v48;
        _2427_v48 = _dafny.Map.Empty.slice().updateUnsafe(_2426_v47,_module.__default.safeModuloInt(p0, _this.f14));
        _2427_v48 = (_2427_v48).update(_2426_v47, new BigNumber((_2356_v3).length));
        if (true) {
          let _2428_v49;
          _2428_v49 = _dafny.Seq.of((_this).f17);
          let _2429_v50;
          _2429_v50 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_2428_v49);
          let _2430_v51;
          _2430_v51 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,(_this).f17);
          let _2431_v52;
          _2431_v52 = _dafny.Map.Empty.slice().updateUnsafe(_2430_v51,_2428_v49);
          _2428_v49 = _dafny.Seq.Concat(_2428_v49, _dafny.Seq.Concat(_2428_v49, (((_2429_v50).contains(_this.f14)) ? ((_2429_v50).get(_this.f14)) : ((((_2431_v52).contains(_dafny.Map.Empty.slice().updateUnsafe(_this.f16,(_this).f17))) ? ((_2431_v52).get(_dafny.Map.Empty.slice().updateUnsafe(_this.f16,(_this).f17))) : (_dafny.Seq.of(_this.f16, true)))))));
          let _2432_v53;
          _2432_v53 = _module.D18.create_DC38(_this.f14, _this.f16, new BigNumber(-384));
          let _2433_v54;
          _2433_v54 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f17,(_this).f17));
          let _2434_v55;
          let _nw464 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _2434_v55 = _nw464;
          let _2435_v56;
          _2435_v56 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-583)),_2434_v55);
          let _2436_v57;
          let _nw465 = Array((new BigNumber(20)).toNumber());
          _nw465[(_dafny.ZERO).toNumber()] = _this.f14;
          _nw465[(_dafny.ONE).toNumber()] = _this.f14;
          _nw465[(new BigNumber(2)).toNumber()] = _this.f14;
          _nw465[(new BigNumber(3)).toNumber()] = _this.f14;
          _nw465[(new BigNumber(4)).toNumber()] = new BigNumber((_2430_v51).length);
          _nw465[(new BigNumber(5)).toNumber()] = p0;
          _nw465[(new BigNumber(6)).toNumber()] = new BigNumber(946);
          _nw465[(new BigNumber(7)).toNumber()] = new BigNumber(-631);
          _nw465[(new BigNumber(8)).toNumber()] = (new BigNumber((_2360_v7).cardinality())).multipliedBy((_dafny.ZERO).minus(_this.f14));
          _nw465[(new BigNumber(9)).toNumber()] = new BigNumber(672);
          _nw465[(new BigNumber(10)).toNumber()] = (_2432_v53).dtor_cf57;
          _nw465[(new BigNumber(11)).toNumber()] = p0;
          _nw465[(new BigNumber(12)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_this.f18, _this.f18)).length);
          _nw465[(new BigNumber(13)).toNumber()] = _module.__default.fm0((_this).f17, globalState);
          _nw465[(new BigNumber(14)).toNumber()] = new BigNumber(711);
          _nw465[(new BigNumber(15)).toNumber()] = (new BigNumber((_2433_v54).length)).multipliedBy(new BigNumber((_2435_v56).length));
          _nw465[(new BigNumber(16)).toNumber()] = (p0).plus(_this.f14);
          _nw465[(new BigNumber(17)).toNumber()] = p0;
          _nw465[(new BigNumber(18)).toNumber()] = _module.__default.safeDivisionInt(_this.f14, new BigNumber((_2353_v0).length));
          _nw465[(new BigNumber(19)).toNumber()] = _module.__default.safeDivisionInt(_this.f14, p0);
          _2436_v57 = _nw465;
          let _index401 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_2436_v57).length));
          (_2436_v57)[_index401] = p0;
          let _2437_v58;
          let _nw466 = new _module.C6();
          _nw466.__ctor(_this.f14, !(_this.f16), _this.f16, !(_this.f16), (_dafny.ZERO).minus(p0));
          _2437_v58 = _nw466;
          let _2438_v59;
          _2438_v59 = _module.D20.create_DC42(p0, _this.f16, _2437_v58);
          let _index402 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_2436_v57).length));
          (_2436_v57)[_index402] = (p0).multipliedBy(((_2438_v59).dtor_cf65).minus(_this.f14));
          let _2439_v60;
          let _nw467 = Array((_dafny.ONE).toNumber()).fill(_dafny.Map.Empty);
          _2439_v60 = _nw467;
          let _2440_v61;
          _2440_v61 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,(_2436_v57)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_2436_v57).length))]);
          let _rhs447 = _2439_v60;
          let _rhs448 = _module.__default.fm31(_dafny.Seq.contains(_2353_v0, new BigNumber((_2440_v61).length)), true, globalState);
          let _rhs449 = new BigNumber(696);
          let _rhs450 = _dafny.Seq.Concat(_2428_v49, _2428_v49);
          let _lhs349 = _this;
          let _lhs350 = globalState;
          _2439_v60 = _rhs447;
          _lhs349.f16 = _rhs448;
          _lhs350.f2 = _rhs449;
          _2428_v49 = _rhs450;
          _2430_v51 = (_2430_v51).update(!(_dafny.Seq.contains(_this.f18, _2415_v40)), (_this).f17);
          let _2441_v62;
          _2441_v62 = _module.D14.create_DC28(_2437_v58.f14, (_this).f17, _2437_v58);
          let _2442_v63;
          let _nw468 = Array((new BigNumber(26)).toNumber());
          _nw468[(_dafny.ZERO).toNumber()] = _this.f16;
          _nw468[(_dafny.ONE).toNumber()] = _this.f16;
          _nw468[(new BigNumber(2)).toNumber()] = _this.f16;
          _nw468[(new BigNumber(3)).toNumber()] = _this.f16;
          _nw468[(new BigNumber(4)).toNumber()] = _this.f16;
          _nw468[(new BigNumber(5)).toNumber()] = (_this).f17;
          _nw468[(new BigNumber(6)).toNumber()] = true;
          _nw468[(new BigNumber(7)).toNumber()] = (_this).f17;
          _nw468[(new BigNumber(8)).toNumber()] = _this.f16;
          _nw468[(new BigNumber(9)).toNumber()] = _this.f16;
          _nw468[(new BigNumber(10)).toNumber()] = _this.f16;
          _nw468[(new BigNumber(11)).toNumber()] = _this.f16;
          _nw468[(new BigNumber(12)).toNumber()] = (_this).f17;
          _nw468[(new BigNumber(13)).toNumber()] = (_this).f17;
          _nw468[(new BigNumber(14)).toNumber()] = true;
          _nw468[(new BigNumber(15)).toNumber()] = _this.f16;
          _nw468[(new BigNumber(16)).toNumber()] = (_this).f17;
          _nw468[(new BigNumber(17)).toNumber()] = _this.f16;
          _nw468[(new BigNumber(18)).toNumber()] = true;
          _nw468[(new BigNumber(19)).toNumber()] = _this.f16;
          _nw468[(new BigNumber(20)).toNumber()] = (_this).f17;
          _nw468[(new BigNumber(21)).toNumber()] = (_this).f17;
          _nw468[(new BigNumber(22)).toNumber()] = (_2441_v62).dtor_cf41;
          _nw468[(new BigNumber(23)).toNumber()] = _this.f16;
          _nw468[(new BigNumber(24)).toNumber()] = (_this).f17;
          _nw468[(new BigNumber(25)).toNumber()] = _this.f16;
          _2442_v63 = _nw468;
          let _2443_v64;
          _2443_v64 = _dafny.MultiSet.fromElements(_2442_v63);
          let _nw469 = Array((new BigNumber(2)).toNumber());
          _nw469[(_dafny.ZERO).toNumber()] = _this.f14;
          _nw469[(_dafny.ONE).toNumber()] = (((_2443_v64).contains(_2442_v63)) ? ((_2443_v64).get(_2442_v63)) : (_this.f14));
          _2434_v55 = _nw469;
        } else {
          let _2444_v65;
          let _nw470 = Array((new BigNumber(15)).toNumber()).fill([]);
          _2444_v65 = _nw470;
          let _2445_v67;
          let _init77 = function (_2446_i5) {
            return function () {
              let _coll82 = new _dafny.Map();
              for (const _compr_82 of _dafny.IntegerRange(new BigNumber(785), new BigNumber(-560))) {
                let _2447_v66 = _compr_82;
                if (((new BigNumber(785)).isLessThanOrEqualTo(_2447_v66)) && ((_2447_v66).isLessThan(new BigNumber(-560)))) {
                  _coll82.push([(_2447_v66).minus(_this.f14),(_this).f17]);
                }
              }
              return _coll82;
            }();
          };
          let _nw471 = Array((new BigNumber(10)).toNumber());
          for (let _i0_77 = 0; _i0_77 < new BigNumber(_nw471.length); _i0_77++) {
            _nw471[_i0_77] = _init77(new BigNumber(_i0_77));
          }
          _2445_v67 = _nw471;
          let _index403 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_2444_v65).length));
          (_2444_v65)[_index403] = _2445_v67;
          let _index404 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_2444_v65).length));
          let _nw472 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
          (_2444_v65)[_index404] = _nw472;
          let _2448_v68;
          let _init78 = function (_2449_i6) {
            return _this.f16;
          };
          let _nw473 = Array((new BigNumber(18)).toNumber());
          for (let _i0_78 = 0; _i0_78 < new BigNumber(_nw473.length); _i0_78++) {
            _nw473[_i0_78] = _init78(new BigNumber(_i0_78));
          }
          _2448_v68 = _nw473;
          let _index405 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_2448_v68).length));
          (_2448_v68)[_index405] = (!(_this.f16)) && (!(_module.__default.fm31(_this.f16, false, globalState)));
          let _2450_v69;
          _2450_v69 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-978)), function (_2451_i7) {
            return new _dafny.CodePoint('b'.codePointAt(0));
          }));
          let _index406 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_2448_v68).length));
          (_2448_v68)[_index406] = ((_2450_v69).Difference(_2450_v69)).IsDisjointFrom((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("jm"))).Intersect(_2450_v69));
          let _2452_v70;
          _2452_v70 = _dafny.Seq.of((_this).f17);
          let _2453_v71;
          _2453_v71 = _dafny.Map.Empty.slice().updateUnsafe(_2448_v68,(_2452_v70)[_module.__default.safeIndex(_this.f14, new BigNumber((_2452_v70).length))]);
          (_this).f16 = (((_2453_v71).contains(_2448_v68)) ? ((_2453_v71).get(_2448_v68)) : ((((_2448_v68)[_module.__default.safeIndex(new BigNumber(611), new BigNumber((_2448_v68).length))]) ? ((_2448_v68)[_module.__default.safeIndex(new BigNumber(611), new BigNumber((_2448_v68).length))]) : (_this.f16))));
          let _2454_v72;
          _2454_v72 = _dafny.Seq.of(_2360_v7);
          let _2455_v73;
          let _nw474 = new _module.C1();
          _nw474.__ctor((new BigNumber((_2454_v72).length)).isEqualTo(_this.f14), (p0).multipliedBy(_module.__default.fm0(false, globalState)));
          _2455_v73 = _nw474;
          let _2456_v74;
          _2456_v74 = _dafny.Map.Empty.slice().updateUnsafe(_2455_v73.f25,_this.f16);
          let _2457_v75;
          _2457_v75 = _dafny.Seq.of(_2456_v74, _2456_v74);
          (_2455_v73).f25 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(659)), function (_2458_i8) {
            return _dafny.Map.Empty.slice().updateUnsafe(false,false);
          }), _2457_v75), _2457_v75);
        }
        let _2459_v76;
        _2459_v76 = _module.D19.create_DC40(p0, _this.f16, p0);
        let _2460_v77;
        let _nw475 = Array((new BigNumber(20)).toNumber());
        _nw475[(_dafny.ZERO).toNumber()] = _2353_v0;
        _nw475[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_2353_v0, _2353_v0);
        _nw475[(new BigNumber(2)).toNumber()] = _2353_v0;
        _nw475[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_this.f14, p0, new BigNumber((_2360_v7).cardinality()), new BigNumber((_2353_v0).length), (_2459_v76).dtor_cf61), _2353_v0);
        _nw475[(new BigNumber(4)).toNumber()] = _2353_v0;
        _nw475[(new BigNumber(5)).toNumber()] = _2353_v0;
        _nw475[(new BigNumber(6)).toNumber()] = _2353_v0;
        _nw475[(new BigNumber(7)).toNumber()] = _2353_v0;
        _nw475[(new BigNumber(8)).toNumber()] = _2353_v0;
        _nw475[(new BigNumber(9)).toNumber()] = _2353_v0;
        _nw475[(new BigNumber(10)).toNumber()] = _2353_v0;
        _nw475[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(302)), function (_2461_i9) {
          return _this.f14;
        });
        _nw475[(new BigNumber(12)).toNumber()] = _2353_v0;
        _nw475[(new BigNumber(13)).toNumber()] = _2353_v0;
        _nw475[(new BigNumber(14)).toNumber()] = _2353_v0;
        _nw475[(new BigNumber(15)).toNumber()] = (((_this).f17) ? (_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("hhr")).length)))) : (_2353_v0));
        _nw475[(new BigNumber(16)).toNumber()] = _2353_v0;
        _nw475[(new BigNumber(17)).toNumber()] = _dafny.Seq.of(new BigNumber((_2426_v47).length));
        _nw475[(new BigNumber(18)).toNumber()] = _2353_v0;
        _nw475[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(131)), _dafny.Seq.of(p0));
        _2460_v77 = _nw475;
        _2460_v77 = _2460_v77;
      }
      (globalState).f7 = ((_this.f16) ? ((_dafny.ZERO).minus(_this.f14)) : (_this.f14));
      let _2462_v78;
      _2462_v78 = _dafny.Seq.of((_this).f17, (_this).f17, _module.__default.fm17((_this).f17, _this.f14, p0, new BigNumber((_2360_v7).cardinality()), globalState));
      let _hi20 = p0;
      for (let _2463_i10 = new BigNumber((_2462_v78).length); _2463_i10.isLessThan(_hi20); _2463_i10 = _2463_i10.plus(_dafny.ONE)) {
        let _2464_v79;
        _2464_v79 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,(_this).f17);
        let _2465_v80;
        let _init79 = function (_2466_i11) {
          return _this.f16;
        };
        let _nw476 = Array((new BigNumber(13)).toNumber());
        for (let _i0_79 = 0; _i0_79 < new BigNumber(_nw476.length); _i0_79++) {
          _nw476[_i0_79] = _init79(new BigNumber(_i0_79));
        }
        _2465_v80 = _nw476;
        let _2467_v81;
        _2467_v81 = _dafny.Map.Empty.slice().updateUnsafe((((_2464_v79).contains(_this.f16)) ? ((_2464_v79).get(_this.f16)) : ((_this).f17)),_2465_v80);
        (globalState).f2 = (p0).multipliedBy(new BigNumber((_2467_v81).length));
        let _2468_v82;
        _2468_v82 = new _dafny.CodePoint('l'.codePointAt(0));
        let _2469_v83;
        let _init80 = ((_2470_p0) => function (_2471_i12) {
          return (_dafny.Map.Empty.slice().updateUnsafe(_this.f14,(_this).f17)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2470_p0,_this.f16));
        })(p0);
        let _nw477 = Array((new BigNumber(8)).toNumber());
        for (let _i0_80 = 0; _i0_80 < new BigNumber(_nw477.length); _i0_80++) {
          _nw477[_i0_80] = _init80(new BigNumber(_i0_80));
        }
        _2469_v83 = _nw477;
        let _2472_v84;
        _2472_v84 = _dafny.Map.Empty.slice().updateUnsafe(_2463_i10,(_this).f17);
        let _index407 = _module.__default.safeIndex(new BigNumber(490), new BigNumber((_2469_v83).length));
        (_2469_v83)[_index407] = (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(244),false)).Merge(_2472_v84);
        let _index408 = _module.__default.safeIndex(new BigNumber(490), new BigNumber((_2469_v83).length));
        let _rhs451 = !(((_2463_i10).multipliedBy((_dafny.ZERO).minus((_this).fm12(_2463_i10, globalState)))).isLessThanOrEqualTo(_2463_i10));
        let _rhs452 = ((_dafny.Seq.contains(_this.f18, _2468_v82)) ? (_2468_v82) : (_2468_v82));
        let _rhs453 = _2472_v84;
        let _rhs454 = p0;
        let _lhs351 = _this;
        let _lhs352 = _2469_v83;
        let _lhs353 = _module.__default.safeIndex(new BigNumber(490), new BigNumber((_2469_v83).length));
        let _lhs354 = globalState;
        _lhs351.f16 = _rhs451;
        _2468_v82 = _rhs452;
        _lhs352[_lhs353] = _rhs453;
        _lhs354.f7 = _rhs454;
        if ((_module.__default.safeModuloInt(_2463_i10, _2463_i10)).isLessThan(_module.__default.safeDivisionInt(_this.f14, p0))) {
          let _2473_v85;
          let _nw478 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2473_v85 = _nw478;
          let _index409 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_2473_v85).length));
          (_2473_v85)[_index409] = _this.f18;
          let _2474_v86;
          _2474_v86 = _module.D12.create_DC22(_this.f18);
          let _index410 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_2473_v85).length));
          (_2473_v85)[_index410] = (_2474_v86).dtor_cf29;
          let _2475_v87;
          _2475_v87 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_2469_v83)[_module.__default.safeIndex(new BigNumber(490), new BigNumber((_2469_v83).length))]);
          let _2476_v88;
          let _nw479 = new _module.C7();
          _nw479.__ctor(_this.f16, _this.f16, _module.__default.fm0(_this.f16, globalState));
          _2476_v88 = _nw479;
          let _2477_v89;
          _2477_v89 = _module.D14.create_DC28(new BigNumber(647), _this.f16, _2476_v88);
          _2475_v87 = (_2475_v87).update(((_2477_v89).dtor_cf40).minus(_this.f14), ((_2469_v83)[_module.__default.safeIndex(new BigNumber(490), new BigNumber((_2469_v83).length))]).Merge(_2472_v84));
          let _2478_v90;
          let _init81 = ((_2479_v86) => function (_2480_i13) {
            return _module.__default.safeModuloInt(_2480_i13, new BigNumber(((_2479_v86).dtor_cf29).length));
          })(_2474_v86);
          let _nw480 = Array((new BigNumber(8)).toNumber());
          for (let _i0_81 = 0; _i0_81 < new BigNumber(_nw480.length); _i0_81++) {
            _nw480[_i0_81] = _init81(new BigNumber(_i0_81));
          }
          _2478_v90 = _nw480;
          let _2481_v91;
          _2481_v91 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_2478_v90);
          let _2482_v92;
          _2482_v92 = _dafny.Set.fromElements(_this.f14, _this.f14);
          let _rhs455 = (((_2481_v91).contains((_dafny.ZERO).minus(_2476_v88.f14))) ? ((_2481_v91).get((_dafny.ZERO).minus(_2476_v88.f14))) : (_2478_v90));
          let _rhs456 = (_2482_v92).IsProperSubsetOf(_2482_v92);
          let _rhs457 = _2476_v88.f14;
          let _lhs355 = _this;
          _2478_v90 = _rhs455;
          r0 = _rhs456;
          _lhs355.f14 = _rhs457;
          let _2483_v93;
          _2483_v93 = _module.D23.create_DC47(_dafny.Map.Empty.slice().updateUnsafe(_2468_v82,_module.D20.create_DC42(_2463_i10, true, _2476_v88)));
          _2483_v93 = _2483_v93;
          let _2484_v94;
          _2484_v94 = _module.D9.create_DC15(_this.f16, _this.f16, _dafny.Set.fromElements(_2463_i10, new BigNumber((_2464_v79).length), new BigNumber(281), _2463_i10), _2478_v90, _2465_v80);
          let _2485_v96;
          let _nw481 = Array((new BigNumber(20)).toNumber());
          _nw481[(_dafny.ZERO).toNumber()] = _2484_v94;
          _nw481[(_dafny.ONE).toNumber()] = _2484_v94;
          _nw481[(new BigNumber(2)).toNumber()] = _2484_v94;
          _nw481[(new BigNumber(3)).toNumber()] = _2484_v94;
          _nw481[(new BigNumber(4)).toNumber()] = _module.D9.create_DC15((_this).f17, _this.f16, _2482_v92, _2478_v90, _2465_v80);
          _nw481[(new BigNumber(5)).toNumber()] = _2484_v94;
          _nw481[(new BigNumber(6)).toNumber()] = _2484_v94;
          _nw481[(new BigNumber(7)).toNumber()] = _2484_v94;
          _nw481[(new BigNumber(8)).toNumber()] = _2484_v94;
          _nw481[(new BigNumber(9)).toNumber()] = _2484_v94;
          _nw481[(new BigNumber(10)).toNumber()] = _2484_v94;
          _nw481[(new BigNumber(11)).toNumber()] = _2484_v94;
          _nw481[(new BigNumber(12)).toNumber()] = _2484_v94;
          _nw481[(new BigNumber(13)).toNumber()] = _2484_v94;
          _nw481[(new BigNumber(14)).toNumber()] = _module.D9.create_DC15((_this).fm14(globalState), (_this).f17, _dafny.Set.fromElements(new BigNumber(-27)), _2478_v90, _2465_v80);
          _nw481[(new BigNumber(15)).toNumber()] = _2484_v94;
          _nw481[(new BigNumber(16)).toNumber()] = _2484_v94;
          _nw481[(new BigNumber(17)).toNumber()] = _2484_v94;
          _nw481[(new BigNumber(18)).toNumber()] = _module.D9.create_DC15((_this).f17, (_this).f17, function () {
  let _coll83 = new _dafny.Set();
  for (const _compr_83 of _dafny.IntegerRange(new BigNumber(129), new BigNumber(847))) {
    let _2486_v95 = _compr_83;
    if (((new BigNumber(129)).isLessThanOrEqualTo(_2486_v95)) && ((_2486_v95).isLessThan(new BigNumber(847)))) {
      _coll83.add(_module.__default.safeModuloInt(_2486_v95, p0));
    }
  }
  return _coll83;
}(), _2478_v90, _2465_v80);
          _nw481[(new BigNumber(19)).toNumber()] = _2484_v94;
          _2485_v96 = _nw481;
          let _2487_v97;
          _2487_v97 = _dafny.Seq.of(_2485_v96);
          let _2488_v98;
          _2488_v98 = _module.D3.create_DC3(_this.f16);
          let _2489_v99;
          _2489_v99 = _dafny.Seq.of(_2488_v98);
          let _2490_v100;
          _2490_v100 = _module.D12.create_DC24(_this.f16, _2462_v78, _2487_v97, _this.f16, _2489_v99);
          let _2491_v101;
          let _nw482 = new _module.C5();
          _nw482.__ctor((_this).f17, (_2490_v100).dtor_cf36);
          _2491_v101 = _nw482;
          let _2492_v102;
          _2492_v102 = _dafny.Map.Empty.slice().updateUnsafe(_2491_v101,_2491_v101.f16);
          let _2493_v104;
          _2493_v104 = _dafny.MultiSet.fromElements(_this.f18);
          let _2494_v105;
          _2494_v105 = _dafny.Map.Empty.slice().updateUnsafe((_2473_v85)[_module.__default.safeIndex(new BigNumber(924), new BigNumber((_2473_v85).length))],_this.f16);
          let _2495_v106;
          _2495_v106 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber(769), _2476_v88.f14),(((((_2492_v102).contains(_2491_v101)) ? ((_2492_v102).get(_2491_v101)) : (_2491_v101.f16))) ? (function () {
            let _coll84 = new _dafny.Map();
            for (const _compr_84 of (_2493_v104).Elements) {
              let _2496_v103 = _compr_84;
              if ((_2493_v104).contains(_2496_v103)) {
                _coll84.push([_2496_v103,(_this).f17]);
              }
            }
            return _coll84;
          }()) : (_2494_v105)));
          _2495_v106 = (_2495_v106).update(new BigNumber(840), (_dafny.Map.Empty.slice().updateUnsafe((_2473_v85)[_module.__default.safeIndex(new BigNumber(924), new BigNumber((_2473_v85).length))],false)).Merge(_2494_v105));
        } else {
          let _rhs458 = (p0).minus(_2463_i10);
          let _rhs459 = _this.f14;
          let _rhs460 = p0;
          let _rhs461 = _2468_v82;
          let _lhs356 = globalState;
          let _lhs357 = _this;
          let _lhs358 = _this;
          _lhs356.f2 = _rhs458;
          _lhs357.f14 = _rhs459;
          _lhs358.f14 = _rhs460;
          _2468_v82 = _rhs461;
          let _2497_v107;
          _2497_v107 = _dafny.Set.fromElements(_this.f14);
          _2497_v107 = (_2497_v107).Difference(_2497_v107);
          _2497_v107 = _2497_v107;
          (globalState).f7 = p0;
          let _2498_v108;
          _2498_v108 = _dafny.Set.fromElements((_this).f17);
          (_this).f16 = !(!(!((_2498_v108).IsSubsetOf((_2498_v108).Union(_2498_v108)))));
        }
        let _2499_v109;
        let _nw483 = new _module.C1();
        _nw483.__ctor(_this.f16, _2463_i10);
        _2499_v109 = _nw483;
      }
      let _2500_v110;
      _2500_v110 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2355_v2).cardinality()),_module.__default.fm44(false, (_this).f17, globalState));
      let _2501_v111;
      _2501_v111 = _module.D18.create_DC38(_this.f14, _this.f16, new BigNumber(((_2500_v110).Merge(_2500_v110)).length));
      _2501_v111 = _2501_v111;
      let _2502_v112;
      _2502_v112 = _dafny.Set.fromElements(new BigNumber(618));
      let _2503_v113;
      let _init82 = function (_2504_i14) {
        return (_2504_i14).minus(new BigNumber((_this.f18).length));
      };
      let _nw484 = Array((new BigNumber(16)).toNumber());
      for (let _i0_82 = 0; _i0_82 < new BigNumber(_nw484.length); _i0_82++) {
        _nw484[_i0_82] = _init82(new BigNumber(_i0_82));
      }
      _2503_v113 = _nw484;
      let _2505_v114;
      let _nw485 = Array((new BigNumber(28)).toNumber()).fill(false);
      _2505_v114 = _nw485;
      let _2506_v115;
      _2506_v115 = _module.D9.create_DC15((_this).f17, (_this).f17, _2502_v112, _2503_v113, _2505_v114);
      let _2507_v116;
      let _nw486 = Array((new BigNumber(6)).toNumber());
      _nw486[(_dafny.ZERO).toNumber()] = (_2506_v115).dtor_cf22;
      _nw486[(_dafny.ONE).toNumber()] = _2505_v114;
      _nw486[(new BigNumber(2)).toNumber()] = _2505_v114;
      _nw486[(new BigNumber(3)).toNumber()] = _2505_v114;
      _nw486[(new BigNumber(4)).toNumber()] = _2505_v114;
      _nw486[(new BigNumber(5)).toNumber()] = _2505_v114;
      _2507_v116 = _nw486;
      let _2508_v117;
      _2508_v117 = _dafny.Map.Empty.slice().updateUnsafe(_2507_v116,(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2356_v3).length),(_this).f17)).update(_this.f14, (_this).f17));
      let _2509_v118;
      _2509_v118 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f17);
      _2508_v117 = (_2508_v117).update(_2507_v116, _2509_v118);
      r0 = (_this).f17;
      return r0;
    }
    m8(p0, p1, p2, p3, globalState) {
      let _this = this;
      if ((_this.f16) && ((_this).f17)) {
        let _2510_v0;
        let _nw487 = new _module.C6();
        _nw487.__ctor(_this.f14, _this.f16, _this.f16, (((_this).f17) ? (_module.__default.fm43((_this).f17, globalState)) : (!(false))), p2);
        _2510_v0 = _nw487;
        let _2511_v1;
        _2511_v1 = _dafny.Set.fromElements(p2);
        let _2512_v2;
        _2512_v2 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_2510_v0).f19);
        let _rhs462 = !(!((_2511_v1).IsSubsetOf(_2511_v1)));
        let _rhs463 = (_2510_v0).fm12(_module.__default.fm0(!((_2510_v0).f20), globalState), globalState);
        let _rhs464 = _module.__default.fm0(_module.__default.fm27(_this.f14, p3, _this.f14, _2512_v2, globalState), globalState);
        let _lhs359 = _this;
        let _lhs360 = globalState;
        let _lhs361 = globalState;
        _lhs359.f16 = _rhs462;
        _lhs360.f2 = _rhs463;
        _lhs361.f7 = _rhs464;
        let _2513_v3;
        let _nw488 = new _module.C4();
        _nw488.__ctor(true, (_2510_v0).f20, !(false));
        _2513_v3 = _nw488;
        let _2514_v4;
        _2514_v4 = new _dafny.CodePoint('r'.codePointAt(0));
        let _2515_v5;
        let _init83 = ((_2516_v3) => function (_2517_i0) {
          return (_2516_v3).f21;
        })(_2513_v3);
        let _nw489 = Array((new BigNumber(4)).toNumber());
        for (let _i0_83 = 0; _i0_83 < new BigNumber(_nw489.length); _i0_83++) {
          _nw489[_i0_83] = _init83(new BigNumber(_i0_83));
        }
        _2515_v5 = _nw489;
        let _2518_v6;
        _2518_v6 = _dafny.MultiSet.fromElements(_2515_v5, _2515_v5);
        let _2519_v7;
        _2519_v7 = _dafny.Map.Empty.slice().updateUnsafe((_2510_v0).f20,_this.f16);
        let _2520_v8;
        _2520_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2518_v6).cardinality()),_2519_v7);
        let _2521_v9;
        let _out35;
        _out35 = (_2510_v0).m3(_2514_v4, new BigNumber(((_2520_v8).Merge(_2520_v8)).length), (_2513_v3).f21, (_2510_v0).f19, globalState);
        _2521_v9 = _out35;
        let _2522_v10;
        _2522_v10 = _dafny.MultiSet.fromElements(_this.f18);
        (_this).f16 = ((_this).f17) || ((_2522_v10).IsProperSubsetOf(_2522_v10));
      } else {
        (_this).f16 = false;
        (_this).f14 = p3;
        (_this).f16 = (_this).f17;
        let _2523_v11;
        let _init84 = function (_2524_i1) {
          return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),(_this).f17);
        };
        let _nw490 = Array((new BigNumber(12)).toNumber());
        for (let _i0_84 = 0; _i0_84 < new BigNumber(_nw490.length); _i0_84++) {
          _nw490[_i0_84] = _init84(new BigNumber(_i0_84));
        }
        _2523_v11 = _nw490;
        let _2525_v12;
        _2525_v12 = _module.D11.create_DC20();
        let _2526_v13;
        _2526_v13 = _dafny.Map.Empty.slice().updateUnsafe((_module.__default.fm68(_module.D11.create_DC21(_2525_v12), (_this).f17, globalState)).dtor_cf14,false);
        let _index411 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_2523_v11).length));
        (_2523_v11)[_index411] = (_2526_v13).Merge(_2526_v13);
        let _index412 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_2523_v11).length));
        (_2523_v11)[_index412] = _2526_v13;
        let _2527_v14;
        let _init85 = function (_2528_i2) {
          return (_dafny.MultiSet.fromElements((_this).f17, (_this).f17, false)).Difference(_dafny.MultiSet.fromElements(_this.f16));
        };
        let _nw491 = Array((new BigNumber(11)).toNumber());
        for (let _i0_85 = 0; _i0_85 < new BigNumber(_nw491.length); _i0_85++) {
          _nw491[_i0_85] = _init85(new BigNumber(_i0_85));
        }
        _2527_v14 = _nw491;
        let _2529_v15;
        _2529_v15 = _dafny.MultiSet.fromElements((_this).f17);
        let _index413 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_2527_v14).length));
        (_2527_v14)[_index413] = _2529_v15;
        let _index414 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_2527_v14).length));
        (_2527_v14)[_index414] = ((_2529_v15).Difference(_dafny.MultiSet.fromElements(_this.f16, _this.f16, (_this).f17, false, _this.f16))).Difference(_2529_v15);
      }
      let _2530_v16;
      let _nw492 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
      _2530_v16 = _nw492;
      let _2531_v17;
      _2531_v17 = _module.D9.create_DC14(_2530_v16);
      let _2532_v18;
      _2532_v18 = _dafny.Set.fromElements(p2, new BigNumber((_this.f18).length), p3);
      let _2533_v19;
      let _nw493 = Array((new BigNumber(8)).toNumber()).fill(false);
      _2533_v19 = _nw493;
      let _2534_v20;
      _2534_v20 = _module.D9.create_DC15(_this.f16, _this.f16, _2532_v18, _2530_v16, _2533_v19);
      let _2535_v21;
      let _nw494 = Array((new BigNumber(7)).toNumber());
      _2535_v21 = _nw494;
      let _2536_v22;
      _2536_v22 = _dafny.Map.Empty.slice().updateUnsafe(_2535_v21,_2530_v16);
      let _2537_v23;
      let _nw495 = Array((new BigNumber(26)).toNumber());
      _nw495[(_dafny.ZERO).toNumber()] = _2530_v16;
      _nw495[(_dafny.ONE).toNumber()] = _2530_v16;
      _nw495[(new BigNumber(2)).toNumber()] = _2530_v16;
      _nw495[(new BigNumber(3)).toNumber()] = _2530_v16;
      _nw495[(new BigNumber(4)).toNumber()] = _2530_v16;
      _nw495[(new BigNumber(5)).toNumber()] = (_2531_v17).dtor_cf17;
      _nw495[(new BigNumber(6)).toNumber()] = _2530_v16;
      _nw495[(new BigNumber(7)).toNumber()] = ((_this.f16) ? (_2530_v16) : (_2530_v16));
      _nw495[(new BigNumber(8)).toNumber()] = _2530_v16;
      _nw495[(new BigNumber(9)).toNumber()] = _2530_v16;
      _nw495[(new BigNumber(10)).toNumber()] = (_2534_v20).dtor_cf21;
      _nw495[(new BigNumber(11)).toNumber()] = _2530_v16;
      _nw495[(new BigNumber(12)).toNumber()] = _2530_v16;
      _nw495[(new BigNumber(13)).toNumber()] = _2530_v16;
      _nw495[(new BigNumber(14)).toNumber()] = _2530_v16;
      _nw495[(new BigNumber(15)).toNumber()] = (((_2536_v22).contains(_2535_v21)) ? ((_2536_v22).get(_2535_v21)) : (_2530_v16));
      _nw495[(new BigNumber(16)).toNumber()] = (_2534_v20).dtor_cf21;
      _nw495[(new BigNumber(17)).toNumber()] = _2530_v16;
      _nw495[(new BigNumber(18)).toNumber()] = _2530_v16;
      _nw495[(new BigNumber(19)).toNumber()] = _2530_v16;
      _nw495[(new BigNumber(20)).toNumber()] = _2530_v16;
      _nw495[(new BigNumber(21)).toNumber()] = _2530_v16;
      _nw495[(new BigNumber(22)).toNumber()] = _2530_v16;
      _nw495[(new BigNumber(23)).toNumber()] = _2530_v16;
      _nw495[(new BigNumber(24)).toNumber()] = _2530_v16;
      _nw495[(new BigNumber(25)).toNumber()] = _2530_v16;
      _2537_v23 = _nw495;
      let _index415 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_2537_v23).length));
      (_2537_v23)[_index415] = _2530_v16;
      let _index416 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_2537_v23).length));
      (_2537_v23)[_index416] = _2530_v16;
      (_this).f16 = (_this).f17;
      let _2538_v24;
      _2538_v24 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber(175));
      let _2539_v25;
      _2539_v25 = _module.D11.create_DC19(_2538_v24);
      let _2540_i3;
      _2540_i3 = _dafny.ZERO;
      L15: {
        let _pat_let_tv40 = p3;
        let _pat_let_tv41 = p3;
        while (!(function (_source38) {
          if (_source38.is_DC20) {
            return _this.f16;
          } else if (_source38.is_DC19) {
            let _2567___mcc_h7 = (_source38).cf27;
            let _2568_cf27 = _2567___mcc_h7;
            return (_pat_let_tv40).isLessThan(_pat_let_tv41);
          } else {
            let _2569___mcc_h8 = (_source38).cf28;
            let _2570_cf28 = _2569___mcc_h8;
            return (_this).f17;
          }
        }(_2539_v25))) {
          C15: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2540_i3)) {
              break L15;
            }
            _2540_i3 = (_2540_i3).plus(_dafny.ONE);
            let _2541_v26;
            _2541_v26 = _module.D3.create_DC4(true, p2, p3, (_this).f17, !(_this.f16));
            let _source37 = _2541_v26;
            if (_source37.is_DC4) {
              let _2542___mcc_h0 = (_source37).cf4;
              let _2543___mcc_h1 = (_source37).cf5;
              let _2544___mcc_h2 = (_source37).cf6;
              let _2545___mcc_h3 = (_source37).cf7;
              let _2546___mcc_h4 = (_source37).cf8;
              let _2547_cf8 = _2546___mcc_h4;
              let _2548_cf7 = _2545___mcc_h3;
              let _2549_cf6 = _2544___mcc_h2;
              let _2550_cf5 = _2543___mcc_h1;
              let _2551_cf4 = _2542___mcc_h0;
              (_this).f16 = _2548_cf7;
              let _index417 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_2533_v19).length));
              (_2533_v19)[_index417] = _2547_cf8;
              let _index418 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_2533_v19).length));
              (_2533_v19)[_index418] = (new BigNumber(((function () {
                let _coll85 = new _dafny.Map();
                for (const _compr_85 of _dafny.IntegerRange(new BigNumber(348), new BigNumber(586))) {
                  let _2552_v27 = _compr_85;
                  if (((new BigNumber(348)).isLessThanOrEqualTo(_2552_v27)) && ((_2552_v27).isLessThan(new BigNumber(586)))) {
                    _coll85.push([_module.__default.safeDivisionInt(_2552_v27, _2550_cf5),p3]);
                  }
                }
                return _coll85;
              }()).Merge(_2538_v24)).length)).isEqualTo(new BigNumber(257));
              _2549_cf6 = ((_module.__default.fm17(!(_2547_cf8), (_dafny.ZERO).minus(_this.f14), new BigNumber((_this.f18).length), new BigNumber(867), globalState)) ? (_2550_cf5) : (_module.__default.safeModuloInt(_2549_cf6, p3)));
              let _2553_v28;
              let _nw496 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Set.Empty);
              _2553_v28 = _nw496;
              let _2554_v29;
              let _nw497 = new _module.C0();
              _nw497.__ctor(_2553_v28, _2550_cf5);
              _2554_v29 = _nw497;
            } else if (_source37.is_DC3) {
              let _2555___mcc_h5 = (_source37).cf3;
              let _2556_cf3 = _2555___mcc_h5;
              (globalState).f7 = p2;
              (_this).f16 = _this.f16;
              let _2557_v30;
              _2557_v30 = new _dafny.CodePoint('m'.codePointAt(0));
              _2556_cf3 = !_dafny.Seq.contains(_dafny.Seq.Concat(_this.f18, _this.f18), _2557_v30);
              (globalState).f2 = _this.f14;
            } else {
              let _2558___mcc_h6 = (_source37).cf9;
              let _2559_cf9 = _2558___mcc_h6;
              let _2560_v31;
              _2560_v31 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_this.f16);
              let _2561_v32;
              let _nw498 = new _module.C5();
              _nw498.__ctor(false, (((_2560_v31).contains(_this.f16)) ? ((_2560_v31).get(_this.f16)) : (_this.f16)));
              _2561_v32 = _nw498;
              let _2562_v33;
              _2562_v33 = _module.D3.create_DC3((_this).f17);
              let _2563_v34;
              let _nw499 = new _module.C2();
              _nw499.__ctor(_2562_v33, (_this).f17, (_this).f17);
              _2563_v34 = _nw499;
              (globalState).f7 = (_dafny.ZERO).minus((p2).multipliedBy((_module.__default.fm0(_this.f16, globalState)).multipliedBy(p3)));
              let _2564_v35;
              let _nw500 = Array((new BigNumber(28)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
              _2564_v35 = _nw500;
              let _2565_v36;
              _2565_v36 = new _dafny.CodePoint('a'.codePointAt(0));
              let _index419 = _module.__default.safeIndex(new BigNumber(83), new BigNumber((_2564_v35).length));
              (_2564_v35)[_index419] = _2565_v36;
              let _index420 = _module.__default.safeIndex(new BigNumber(83), new BigNumber((_2564_v35).length));
              (_2564_v35)[_index420] = _2565_v36;
            }
            let _index421 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_2530_v16).length));
            (_2530_v16)[_index421] = p3;
            let _index422 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_2530_v16).length));
            (_2530_v16)[_index422] = (new BigNumber(715)).multipliedBy(p3);
            let _2566_v37;
            _2566_v37 = _dafny.Set.fromElements((_this).f17, (_this).f17, (_this).f17);
            _2566_v37 = ((_2566_v37).Intersect(_2566_v37)).Intersect((_2566_v37).Intersect(_2566_v37));
            (globalState).f7 = _module.__default.fm0(_this.f16, globalState);
          }
        }
      }
      (_this).f18 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(543)), function (_2571_i4) {
        return ((_this.f16) ? (new _dafny.CodePoint('u'.codePointAt(0))) : (new _dafny.CodePoint('n'.codePointAt(0))));
      });
      let _2572_v38;
      _2572_v38 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_2533_v19);
      let _2573_v39;
      _2573_v39 = _dafny.Set.fromElements(_2530_v16);
      _2572_v38 = (_2572_v38).update(new BigNumber((_2573_v39).length), _2533_v19);
      return;
    }
  };

  $module.C9 = class C9 {
    constructor () {
      this._tname = "_module.C9";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    m4(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let _2574_v0;
      _2574_v0 = _dafny.Seq.UnicodeFromString("qnopbue");
      let _2575_v1;
      _2575_v1 = _dafny.Set.fromElements(false, true, p0);
      let _2576_v2;
      let _nw501 = new _module.C8();
      _nw501.__ctor(_2574_v0, !(_2575_v1).contains(p0), ((false) ? (p0) : (p0)), (_dafny.ZERO).minus((new BigNumber(488)).multipliedBy(new BigNumber(179))));
      _2576_v2 = _nw501;
      let _hi21 = p3;
      for (let _2577_i0 = ((p0) ? (p1) : (p1)); _2577_i0.isLessThan(_hi21); _2577_i0 = _2577_i0.plus(_dafny.ONE)) {
        (globalState).f2 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("nno")).length));
        let _2578_v3;
        let _init86 = ((_2579_p1) => function (_2580_i1) {
          return (_2580_i1).multipliedBy(_2579_p1);
        })(p1);
        let _nw502 = Array((new BigNumber(22)).toNumber());
        for (let _i0_86 = 0; _i0_86 < new BigNumber(_nw502.length); _i0_86++) {
          _nw502[_i0_86] = _init86(new BigNumber(_i0_86));
        }
        _2578_v3 = _nw502;
        let _index423 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_2578_v3).length));
        (_2578_v3)[_index423] = new BigNumber((_2576_v2.f18).length);
        let _index424 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_2578_v3).length));
        (_2578_v3)[_index424] = p2;
        let _2581_v4;
        _2581_v4 = new _dafny.CodePoint('g'.codePointAt(0));
        let _2582_v5;
        _2582_v5 = _dafny.MultiSet.fromElements(_2581_v4, _2581_v4, _2581_v4, _2581_v4, _2581_v4);
        let _2583_v6;
        _2583_v6 = _dafny.Seq.of(_2582_v5);
        let _index425 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_2578_v3).length));
        let _index426 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_2578_v3).length));
        let _rhs465 = p2;
        let _rhs466 = new BigNumber(((_2583_v6)[_module.__default.safeIndex(new BigNumber((_2576_v2.f18).length), new BigNumber((_2583_v6).length))]).cardinality());
        let _rhs467 = p3;
        let _rhs468 = (_2575_v1).Union(_2575_v1);
        let _rhs469 = p2;
        let _lhs362 = _2578_v3;
        let _lhs363 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_2578_v3).length));
        let _lhs364 = globalState;
        let _lhs365 = _2578_v3;
        let _lhs366 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_2578_v3).length));
        let _lhs367 = globalState;
        _lhs362[_lhs363] = _rhs465;
        _lhs364.f2 = _rhs466;
        _lhs365[_lhs366] = _rhs467;
        _2575_v1 = _rhs468;
        _lhs367.f7 = _rhs469;
        let _2584_v7;
        let _nw503 = Array((new BigNumber(5)).toNumber());
        _2584_v7 = _nw503;
        let _2585_v8;
        let _nw504 = new _module.C6();
        _nw504.__ctor(new BigNumber(297), p0, p0, !(p0), p1);
        _2585_v8 = _nw504;
        let _index427 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_2584_v7).length));
        (_2584_v7)[_index427] = _2585_v8;
        let _index428 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_2584_v7).length));
        (_2584_v7)[_index428] = _2585_v8;
      }
      let _2586_i2;
      _2586_i2 = _dafny.ZERO;
      L16: {
        while (false) {
          C16: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2586_i2)) {
              break L16;
            }
            _2586_i2 = (_2586_i2).plus(_dafny.ONE);
            (globalState).f2 = new BigNumber(820);
            let _2587_v9;
            let _init87 = ((_2588_p0) => function (_2589_i3) {
              return _2588_p0;
            })(p0);
            let _nw505 = Array((new BigNumber(19)).toNumber());
            for (let _i0_87 = 0; _i0_87 < new BigNumber(_nw505.length); _i0_87++) {
              _nw505[_i0_87] = _init87(new BigNumber(_i0_87));
            }
            _2587_v9 = _nw505;
            let _index429 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_2587_v9).length));
            (_2587_v9)[_index429] = p0;
            let _index430 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_2587_v9).length));
            (_2587_v9)[_index430] = p0;
            if (_module.__default.fm31((_2587_v9)[_module.__default.safeIndex(new BigNumber(756), new BigNumber((_2587_v9).length))], p0, globalState)) {
              let _index431 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_2587_v9).length));
              (_2587_v9)[_index431] = _module.__default.fm17((_2587_v9)[_module.__default.safeIndex(new BigNumber(756), new BigNumber((_2587_v9).length))], new BigNumber(843), (new BigNumber(637)).minus((_dafny.ZERO).minus(p1)), _module.__default.safeDivisionInt(p2, p3), globalState);
              let _2590_v10;
              let _nw506 = Array((new BigNumber(6)).toNumber());
              _nw506[(_dafny.ZERO).toNumber()] = p2;
              _nw506[(_dafny.ONE).toNumber()] = p3;
              _nw506[(new BigNumber(2)).toNumber()] = p1;
              _nw506[(new BigNumber(3)).toNumber()] = new BigNumber(352);
              _nw506[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_2574_v0).length)));
              _nw506[(new BigNumber(5)).toNumber()] = p1;
              _2590_v10 = _nw506;
              let _index432 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_2590_v10).length));
              (_2590_v10)[_index432] = (new BigNumber((_2576_v2.f18).length)).multipliedBy(p2);
              let _index433 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_2590_v10).length));
              (_2590_v10)[_index433] = p3;
              let _2591_v11;
              _2591_v11 = _dafny.Seq.of(p3, (_2590_v10)[_module.__default.safeIndex(new BigNumber(36), new BigNumber((_2590_v10).length))], p3, ((p0) ? (new BigNumber((_dafny.Seq.UnicodeFromString("eiimxauy")).length)) : ((_2590_v10)[_module.__default.safeIndex(new BigNumber(36), new BigNumber((_2590_v10).length))])));
              let _2592_v12;
              _2592_v12 = _2591_v11;
              _2591_v11 = _dafny.Seq.Concat(_dafny.Seq.update((_2592_v12), _module.__default.safeIndex(new BigNumber(-484), new BigNumber(((_2592_v12)).length)), p3), _2591_v11);
              _2587_v9 = _2587_v9;
              let _index434 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_2587_v9).length));
              (_2587_v9)[_index434] = (_2587_v9)[_module.__default.safeIndex(new BigNumber(756), new BigNumber((_2587_v9).length))];
            } else {
              let _2593_v13;
              _2593_v13 = new _dafny.CodePoint('b'.codePointAt(0));
              let _2594_v14;
              _2594_v14 = _module.D15.create_DC30(_2593_v13);
              let _2595_v15;
              _2595_v15 = _dafny.Map.Empty.slice().updateUnsafe((_2594_v14).dtor_cf44,true);
              let _2596_v16;
              _2596_v16 = _dafny.Map.Empty.slice().updateUnsafe((_2595_v15).contains(_2593_v13),p3);
              _2596_v16 = (_2596_v16).update(p0, p3);
              let _index435 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_2587_v9).length));
              (_2587_v9)[_index435] = (_2575_v1).equals(_2575_v1);
              let _2597_v17;
              let _nw507 = new _module.C7();
              _nw507.__ctor(p0, p0, ((p0) ? (p3) : (p2)));
              _2597_v17 = _nw507;
              let _index436 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_2587_v9).length));
              (_2587_v9)[_index436] = p0;
              (globalState).f7 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(378)), ((_2598_v13) => function (_2599_i4) {
                return _2598_v13;
              })(_2593_v13))).length));
            }
            let _2600_v18;
            _2600_v18 = new _dafny.CodePoint('o'.codePointAt(0));
            let _2601_v19;
            _2601_v19 = _dafny.Map.Empty.slice().updateUnsafe(_2600_v18,p0);
            let _index437 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_2587_v9).length));
            (_2587_v9)[_index437] = !((((_2601_v19).contains(_2600_v18)) ? ((_2601_v19).get(_2600_v18)) : (false)));
          }
        }
      }
      let _hi22 = p3;
      for (let _2602_i5 = p1; _2602_i5.isLessThan(_hi22); _2602_i5 = _2602_i5.plus(_dafny.ONE)) {
        (globalState).f7 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber(-213)), p2));
        r0 = _module.__default.safeDivisionInt(p1, p1);
        let _2603_v20;
        _2603_v20 = new _dafny.CodePoint('g'.codePointAt(0));
        _2603_v20 = new _dafny.CodePoint('w'.codePointAt(0));
        let _2604_v21;
        let _init88 = function (_2605_i6) {
          return !(true);
        };
        let _nw508 = Array((new BigNumber(7)).toNumber());
        for (let _i0_88 = 0; _i0_88 < new BigNumber(_nw508.length); _i0_88++) {
          _nw508[_i0_88] = _init88(new BigNumber(_i0_88));
        }
        _2604_v21 = _nw508;
        let _index438 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_2604_v21).length));
        (_2604_v21)[_index438] = p0;
        let _index439 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_2604_v21).length));
        (_2604_v21)[_index439] = p0;
      }
      let _2606_v22;
      _2606_v22 = _dafny.Seq.of(p0);
      let _2607_v23;
      _2607_v23 = _2606_v22;
      let _2608_v24;
      _2608_v24 = _dafny.MultiSet.fromElements((_2607_v23), _2606_v22);
      let _hi23 = p3;
      for (let _2609_i7 = (_2576_v2).fm12(new BigNumber(((_2608_v24).update(_2606_v22, _module.__default.abs((_dafny.ZERO).minus(p2)))).cardinality()), globalState); _2609_i7.isLessThan(_hi23); _2609_i7 = _2609_i7.plus(_dafny.ONE)) {
        (globalState).f7 = ((p0) ? ((p3).plus(_2609_i7)) : (p1));
        let _2610_v25;
        _2610_v25 = new _dafny.CodePoint('u'.codePointAt(0));
        let _2611_v26;
        let _nw509 = new _module.C8();
        _nw509.__ctor(_dafny.Seq.update(_2576_v2.f18, _module.__default.safeIndex(p2, new BigNumber((_2576_v2.f18).length)), _2610_v25), p0, p0, (_dafny.ZERO).minus(_2609_i7));
        _2611_v26 = _nw509;
        let _2612_v27;
        _2612_v27 = _module.D16.create_DC34((_2606_v22)[_module.__default.safeIndex(p3, new BigNumber((_2606_v22).length))], _2611_v26, p1);
        let _source39 = _2612_v27;
        if (_source39.is_DC33) {
          let _2613___mcc_h0 = (_source39).cf47;
          let _2614___mcc_h1 = (_source39).cf48;
          let _2615___mcc_h2 = (_source39).cf49;
          let _2616___mcc_h3 = (_source39).cf50;
          let _2617_cf50 = _2616___mcc_h3;
          let _2618_cf49 = _2615___mcc_h2;
          let _2619_cf48 = _2614___mcc_h1;
          let _2620_cf47 = _2613___mcc_h0;
          let _2621_v28;
          let _nw510 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _2621_v28 = _nw510;
          let _index440 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_2621_v28).length));
          (_2621_v28)[_index440] = (_dafny.ZERO).minus(p3);
          let _2622_v29;
          _2622_v29 = _dafny.MultiSet.fromElements(p3);
          let _2623_v30;
          _2623_v30 = _module.D16.create_DC33(false, !(p0), _2609_i7, _2617_cf50);
          let _pat_let_tv42 = _2619_cf48;
          let _index441 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_2621_v28).length));
          (_2621_v28)[_index441] = _module.__default.safeModuloInt(new BigNumber((_2622_v29).cardinality()), (function (_pat_let52_0) {
            return function (_2624_dt__update__tmp_h0) {
              return function (_pat_let53_0) {
                return function (_2625_dt__update_hcf48_h0) {
                  return function (_pat_let54_0) {
                    return function (_2626_dt__update_hcf47_h0) {
                      return _module.D16.create_DC33(_2626_dt__update_hcf47_h0, _2625_dt__update_hcf48_h0, (_2624_dt__update__tmp_h0).dtor_cf49, (_2624_dt__update__tmp_h0).dtor_cf50);
                    }(_pat_let54_0);
                  }(_pat_let_tv42);
                }(_pat_let53_0);
              }(true);
            }(_pat_let52_0);
          }(_2623_v30)).dtor_cf49);
          let _index442 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_2621_v28).length));
          (_2621_v28)[_index442] = (_2621_v28)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_2621_v28).length))];
          let _2627_v31;
          _2627_v31 = _dafny.MultiSet.fromElements(_2610_v25);
          let _2628_v32;
          _2628_v32 = _module.D21.create_DC43(_2627_v31);
          let _2629_v33;
          _2629_v33 = _module.D21.create_DC45(_2628_v32);
          _2629_v33 = _module.D21.create_DC45(_2628_v32);
          _2610_v25 = _2610_v25;
        } else if (_source39.is_DC34) {
          let _2630___mcc_h4 = (_source39).cf51;
          let _2631___mcc_h5 = (_source39).cf52;
          let _2632___mcc_h6 = (_source39).cf53;
          let _2633_cf53 = _2632___mcc_h6;
          let _2634_cf52 = _2631___mcc_h5;
          let _2635_cf51 = _2630___mcc_h4;
          let _2636_v34;
          let _nw511 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Seq.of());
          _2636_v34 = _nw511;
          _2636_v34 = _2636_v34;
          _2635_cf51 = _dafny.Seq.IsPrefixOf(_2606_v22, _dafny.Seq.update(_2606_v22, _module.__default.safeIndex(new BigNumber((_2575_v1).length), new BigNumber((_2606_v22).length)), _2635_cf51));
          _2635_cf51 = p0;
          let _2637_v36;
          let _init89 = ((_2638_i7, _2639_v26, _2640_cf52, _2641_p1, _2642_p3) => function (_2643_i8) {
            return _dafny.Seq.of(function () {
              let _coll86 = new _dafny.Set();
              for (const _compr_86 of _dafny.IntegerRange(new BigNumber(980), new BigNumber(226))) {
                let _2644_v35 = _compr_86;
                if (((new BigNumber(980)).isLessThanOrEqualTo(_2644_v35)) && ((_2644_v35).isLessThan(new BigNumber(226)))) {
                  _coll86.add(_module.__default.safeModuloInt(_2644_v35, (_dafny.ZERO).minus(_2639_v26.f14)));
                }
              }
              return _coll86;
            }(), _dafny.Set.fromElements(_2638_i7, new BigNumber(458), _2640_cf52.f14, new BigNumber((_dafny.Set.fromElements(_2641_p1, _2642_p3, _2641_p1)).length), _2639_v26.f14));
          })(_2609_i7, _2611_v26, _2634_cf52, p1, p3);
          let _nw512 = Array((new BigNumber(28)).toNumber());
          for (let _i0_89 = 0; _i0_89 < new BigNumber(_nw512.length); _i0_89++) {
            _nw512[_i0_89] = _init89(new BigNumber(_i0_89));
          }
          _2637_v36 = _nw512;
          let _2645_v37;
          _2645_v37 = _dafny.Set.fromElements(p1);
          let _2646_v38;
          _2646_v38 = _dafny.Seq.of(_2645_v37, _2645_v37, _2645_v37);
          let _2647_v39;
          _2647_v39 = _2646_v38;
          let _index443 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_2637_v36).length));
          (_2637_v36)[_index443] = _2647_v39;
          let _index444 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_2637_v36).length));
          (_2637_v36)[_index444] = ((((_dafny.ZERO).minus(_2634_cf52.f14)).isLessThanOrEqualTo(p2)) ? (_2647_v39) : (_2647_v39));
        } else if (_source39.is_DC32) {
          let _2648___mcc_h7 = (_source39).cf46;
          let _2649_cf46 = _2648___mcc_h7;
          let _2650_v40;
          let _init90 = ((_2651_p0) => function (_2652_i9) {
            return (_2652_i9).minus(new BigNumber((_dafny.Seq.of(_2651_p0)).length));
          })(p0);
          let _nw513 = Array((new BigNumber(20)).toNumber());
          for (let _i0_90 = 0; _i0_90 < new BigNumber(_nw513.length); _i0_90++) {
            _nw513[_i0_90] = _init90(new BigNumber(_i0_90));
          }
          _2650_v40 = _nw513;
          let _2653_v41;
          let _nw514 = Array((new BigNumber(14)).toNumber()).fill(false);
          _2653_v41 = _nw514;
          let _2654_v42;
          _2654_v42 = _module.D9.create_DC15(p0, p0, _module.__default.fm49(p1, p2, p0, _2611_v26.f14, globalState), _2650_v40, _2653_v41);
          let _2655_v43;
          _2655_v43 = _dafny.Seq.of(_2650_v40);
          let _2656_v44;
          _2656_v44 = _dafny.Set.fromElements(_2650_v40, _2650_v40, (_2655_v43)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_2606_v22, _module.__default.safeIndex(p3, new BigNumber((_2606_v22).length)), p0)).length), new BigNumber((_2655_v43).length))], _2650_v40);
          let _2657_v45;
          let _nw515 = Array((new BigNumber(3)).toNumber());
          _nw515[(_dafny.ZERO).toNumber()] = (_2654_v42).dtor_cf18;
          _nw515[(_dafny.ONE).toNumber()] = !((_dafny.Set.fromElements(_2650_v40, _2650_v40)).IsDisjointFrom(_2656_v44));
          _nw515[(new BigNumber(2)).toNumber()] = p0;
          _2657_v45 = _nw515;
          let _index445 = _module.__default.safeIndex(new BigNumber(519), new BigNumber((_2653_v41).length));
          (_2653_v41)[_index445] = (_2611_v26.f14).isLessThan(_2609_i7);
          let _index446 = _module.__default.safeIndex(new BigNumber(519), new BigNumber((_2653_v41).length));
          (_2653_v41)[_index446] = p0;
          (globalState).f2 = (_dafny.ZERO).minus(_2611_v26.f14);
          let _2658_v46;
          let _nw516 = new _module.C3();
          _nw516.__ctor(p2, _dafny.Seq.Concat(_2574_v0, _dafny.Seq.UnicodeFromString("bomh")), (_2653_v41)[_module.__default.safeIndex(new BigNumber(519), new BigNumber((_2653_v41).length))], (_2653_v41)[_module.__default.safeIndex(new BigNumber(519), new BigNumber((_2653_v41).length))]);
          _2658_v46 = _nw516;
          let _index447 = _module.__default.safeIndex(new BigNumber(519), new BigNumber((_2653_v41).length));
          (_2653_v41)[_index447] = ((false) ? (_module.__default.fm31(p0, false, globalState)) : (p0));
        } else {
          let _2659___mcc_h8 = (_source39).cf54;
          let _2660_cf54 = _2659___mcc_h8;
          let _2661_v47;
          let _2662_v48;
          let _out36;
          let _out37;
          let _outcollector11 = (_this).m5(globalState);
          _out36 = _outcollector11[0];
          _out37 = _outcollector11[1];
          _2661_v47 = _out36;
          _2662_v48 = _out37;
          let _2663_v49;
          let _init91 = ((_2664_p1) => function (_2665_i10) {
            return _module.__default.safeModuloInt(_2665_i10, _2664_p1);
          })(p1);
          let _nw517 = Array((new BigNumber(22)).toNumber());
          for (let _i0_91 = 0; _i0_91 < new BigNumber(_nw517.length); _i0_91++) {
            _nw517[_i0_91] = _init91(new BigNumber(_i0_91));
          }
          _2663_v49 = _nw517;
          let _index448 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_2663_v49).length));
          (_2663_v49)[_index448] = p2;
          let _index449 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_2663_v49).length));
          (_2663_v49)[_index449] = _2609_i7;
          let _2666_v50;
          _2666_v50 = _dafny.MultiSet.fromElements(_2611_v26.f14, p1);
          _2666_v50 = ((_2666_v50).update((_2663_v49)[_module.__default.safeIndex(new BigNumber(740), new BigNumber((_2663_v49).length))], _module.__default.abs(_2611_v26.f14))).Intersect(_2666_v50);
          (globalState).f2 = (((_2666_v50).contains(_2611_v26.f14)) ? ((_2666_v50).get(_2611_v26.f14)) : (_2611_v26.f14));
        }
        let _2667_v51;
        _2667_v51 = false;
        let _2668_v52;
        _2668_v52 = _dafny.MultiSet.fromElements(_2609_i7, _2609_i7, _2609_i7, new BigNumber((_dafny.Seq.update(_2576_v2.f18, _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(46)), ((_2669_v25) => function (_2670_i11) {
          return _2669_v25;
        })(_2610_v25))).length), new BigNumber((_2576_v2.f18).length)), _2610_v25)).length), p2);
        let _rhs470 = ((p1).multipliedBy(_2611_v26.f14)).isEqualTo((p1).plus(p2));
        let _rhs471 = !(!((_2668_v52).IsDisjointFrom(_2668_v52)));
        let _rhs472 = p0;
        _2667_v51 = _rhs470;
        _2667_v51 = _rhs471;
        _2667_v51 = _rhs472;
        if (_2667_v51) {
          let _2671_v53;
          _2671_v53 = _module.D3.create_DC4(p0, p2, p2, p0, _2667_v51);
          let _2672_v54;
          _2672_v54 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2667_v51);
          let _2673_v55;
          _2673_v55 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_2576_v2).fm12(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-952),_2611_v26.f14)).length), globalState)),_2576_v2.f18);
          let _2674_v56;
          _2674_v56 = _dafny.Map.Empty.slice().updateUnsafe(_2611_v26.f14,_2611_v26.f14);
          let _2675_v57;
          _2675_v57 = _dafny.Seq.of(_2574_v0);
          let _2676_v58;
          let _nw518 = Array((new BigNumber(19)).toNumber());
          _nw518[(_dafny.ZERO).toNumber()] = _2576_v2.f18;
          _nw518[(_dafny.ONE).toNumber()] = _module.__default.fm28(p0, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("faf")).length),p3), _2671_v53, _2672_v54, globalState);
          _nw518[(new BigNumber(2)).toNumber()] = _2576_v2.f18;
          _nw518[(new BigNumber(3)).toNumber()] = _2574_v0;
          _nw518[(new BigNumber(4)).toNumber()] = _2576_v2.f18;
          _nw518[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_2576_v2.f18, _module.__default.safeIndex(new BigNumber((_dafny.Seq.of(new BigNumber((_2668_v52).cardinality()))).length), new BigNumber((_2576_v2.f18).length)), new _dafny.CodePoint('c'.codePointAt(0))), _2576_v2.f18);
          _nw518[(new BigNumber(6)).toNumber()] = _2576_v2.f18;
          _nw518[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_2576_v2.f18, _2574_v0);
          _nw518[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("ecxp"), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.UnicodeFromString("ecxp")).length)), _2610_v25), _dafny.Seq.UnicodeFromString("qoltf"));
          _nw518[(new BigNumber(9)).toNumber()] = _2574_v0;
          _nw518[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_2576_v2.f18, (((_2673_v55).contains(p1)) ? ((_2673_v55).get(p1)) : (_dafny.Seq.update(_2576_v2.f18, _module.__default.safeIndex(p1, new BigNumber((_2576_v2.f18).length)), _2610_v25))));
          _nw518[(new BigNumber(11)).toNumber()] = _2574_v0;
          _nw518[(new BigNumber(12)).toNumber()] = _2576_v2.f18;
          _nw518[(new BigNumber(13)).toNumber()] = _2576_v2.f18;
          _nw518[(new BigNumber(14)).toNumber()] = _module.__default.fm28(p0, _2674_v56, _2671_v53, _2672_v54, globalState);
          _nw518[(new BigNumber(15)).toNumber()] = _2576_v2.f18;
          _nw518[(new BigNumber(16)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-673)), ((_2677_v25) => function (_2678_i12) {
            return _2677_v25;
          })(_2610_v25));
          _nw518[(new BigNumber(17)).toNumber()] = _2576_v2.f18;
          _nw518[(new BigNumber(18)).toNumber()] = (_2675_v57)[_module.__default.safeIndex(_2609_i7, new BigNumber((_2675_v57).length))];
          _2676_v58 = _nw518;
          let _2679_v59;
          _2679_v59 = _dafny.Map.Empty.slice().updateUnsafe(!(_2667_v51),_2672_v54);
          let _2680_v60;
          _2680_v60 = _dafny.Map.Empty.slice().updateUnsafe(_2667_v51,_2679_v59);
          let _2681_v61;
          _2681_v61 = _dafny.Seq.of(_2680_v60);
          let _2682_v62;
          _2682_v62 = _dafny.Map.Empty.slice().updateUnsafe((_2681_v61)[_module.__default.safeIndex((_dafny.ZERO).minus((_2576_v2).fm12(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(528)), ((_2683_v25) => function (_2684_i13) {
            return _2683_v25;
          })(_2610_v25))).length), globalState)), new BigNumber((_2681_v61).length))],_2609_i7);
          let _2685_v63;
          _2685_v63 = _module.D14.create_DC27(_2682_v62);
          let _rhs473 = _2676_v58;
          let _rhs474 = _2685_v63;
          _2676_v58 = _rhs473;
          _2685_v63 = _rhs474;
          let _2686_v64;
          let _init92 = ((_2687_v1) => function (_2688_i14) {
            return _2687_v1;
          })(_2575_v1);
          let _nw519 = Array((new BigNumber(2)).toNumber());
          for (let _i0_92 = 0; _i0_92 < new BigNumber(_nw519.length); _i0_92++) {
            _nw519[_i0_92] = _init92(new BigNumber(_i0_92));
          }
          _2686_v64 = _nw519;
          let _2689_v65;
          let _nw520 = new _module.C0();
          _nw520.__ctor(_2686_v64, new BigNumber(-759));
          _2689_v65 = _nw520;
          let _nw521 = new _module.C8();
          _nw521.__ctor((_2675_v57)[_module.__default.safeIndex(_2611_v26.f14, new BigNumber((_2675_v57).length))], (_2611_v26.f14).isEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(p3))), p0, _2609_i7);
          _2576_v2 = _nw521;
          (_2611_v26).f14 = p2;
          _2610_v25 = new _dafny.CodePoint('t'.codePointAt(0));
        } else {
          let _2690_v66;
          _2690_v66 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
          let _2691_v67;
          _2691_v67 = _module.D18.create_DC38(new BigNumber((_2690_v66).length), _2667_v51, _2609_i7);
          let _2692_v68;
          let _nw522 = new _module.C5();
          _nw522.__ctor((_2691_v67).dtor_cf58, _2667_v51);
          _2692_v68 = _nw522;
          let _2693_v69;
          let _nw523 = new _module.C8();
          _nw523.__ctor(_2576_v2.f18, !(_2667_v51), p0, p3);
          _2693_v69 = _nw523;
          let _2694_v70;
          _2694_v70 = _dafny.Map.Empty.slice().updateUnsafe(_2612_v27,_2611_v26.f14);
          _2694_v70 = _2694_v70;
          _2667_v51 = !(false);
          let _2695_v71;
          let _out38;
          _out38 = (_2576_v2).m3(_2610_v25, new BigNumber((_2606_v22).length), false, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("iir"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(422)), ((_2696_v25) => function (_2697_i15) {
            return _2696_v25;
          })(_2610_v25)))).length), globalState);
          _2695_v71 = _out38;
        }
      }
      let _2698_v72;
      _2698_v72 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-978),p2);
      let _2699_v73;
      _2699_v73 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _2700_i16;
      _2700_i16 = _dafny.ZERO;
      L17: {
        while (!(p3).isEqualTo(new BigNumber(((_2698_v72).update(new BigNumber((_2699_v73).length), p3)).length))) {
          C17: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2700_i16)) {
              break L17;
            }
            _2700_i16 = (_2700_i16).plus(_dafny.ONE);
            let _2701_v74;
            _2701_v74 = _module.D15.create_DC29(_2575_v1);
            let _2702_v75;
            let _nw524 = new _module.C6();
            _nw524.__ctor(new BigNumber((_2576_v2.f18).length), p0, p0, p0, new BigNumber(((_2701_v74).dtor_cf43).length));
            _2702_v75 = _nw524;
            (globalState).f2 = (_2702_v75).f19;
            let _2703_v76;
            let _nw525 = Array((new BigNumber(13)).toNumber()).fill([]);
            _2703_v76 = _nw525;
            _2703_v76 = _2703_v76;
            let _2704_v77;
            _2704_v77 = _dafny.MultiSet.fromElements((_2702_v75).f20, (_2702_v75).f20, (_2702_v75).f20);
            let _2705_v78;
            let _nw526 = Array((new BigNumber(5)).toNumber());
            _nw526[(_dafny.ZERO).toNumber()] = ((_2702_v75).f19).isLessThan(new BigNumber((_2606_v22).length));
            _nw526[(_dafny.ONE).toNumber()] = p0;
            _nw526[(new BigNumber(2)).toNumber()] = !((_2702_v75).f20) || (_module.__default.fm17(p0, (_2702_v75).f19, p3, (_dafny.ZERO).minus(p1), globalState));
            _nw526[(new BigNumber(3)).toNumber()] = p0;
            _nw526[(new BigNumber(4)).toNumber()] = (_2704_v77).IsSubsetOf(_2704_v77);
            _2705_v78 = _nw526;
            let _index450 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_2705_v78).length));
            (_2705_v78)[_index450] = p0;
            let _2706_v79;
            _2706_v79 = _dafny.Seq.of(p1, new BigNumber((_2704_v77).cardinality()), p1);
            let _index451 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_2705_v78).length));
            (_2705_v78)[_index451] = _module.__default.fm27(new BigNumber((_2706_v79).length), p3, p2, _2698_v72, globalState);
          }
        }
      }
      let _2707_v81;
      _2707_v81 = _dafny.Map.Empty.slice().updateUnsafe(p3,p0);
      let _2708_v83;
      _2708_v83 = _dafny.Seq.of(new BigNumber((function () {
        let _coll87 = new _dafny.Set();
        for (const _compr_87 of (_2707_v81).Keys.Elements) {
          let _2709_v82 = _compr_87;
          if ((_2707_v81).contains(_2709_v82)) {
            _coll87.add((_2709_v82).plus(new BigNumber((_dafny.Seq.of(new BigNumber(-594), new BigNumber(445))).length)));
          }
        }
        return _coll87;
      }()).length), p3);
      r0 = ((!(function () {
        let _coll88 = new _dafny.Map();
        for (const _compr_88 of (_2708_v83).Elements) {
          let _2710_v80 = _compr_88;
          if (_dafny.Seq.contains(_2708_v83, _2710_v80)) {
            _coll88.push([(_2710_v80).minus(new BigNumber((_dafny.Set.fromElements(p2)).length)),(_dafny.ZERO).minus(p1)]);
          }
        }
        return _coll88;
      }()).equals(_2698_v72)) ? (p1) : (p3));
      r1 = _module.__default.safeModuloInt(new BigNumber((_2708_v83).length), p2);
      return [r0, r1];
    }
    m5(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let _2711_v0;
      _2711_v0 = new BigNumber(205);
      r0 = _2711_v0;
      let _hi24 = new BigNumber(-132);
      for (let _2712_i0 = _2711_v0; _2712_i0.isLessThan(_hi24); _2712_i0 = _2712_i0.plus(_dafny.ONE)) {
        let _2713_v1;
        let _init93 = ((_2714_v0) => function (_2715_i1) {
          return (_dafny.Set.fromElements(new BigNumber(454))).IsDisjointFrom(_dafny.Set.fromElements(_2714_v0));
        })(_2711_v0);
        let _nw527 = Array((new BigNumber(23)).toNumber());
        for (let _i0_93 = 0; _i0_93 < new BigNumber(_nw527.length); _i0_93++) {
          _nw527[_i0_93] = _init93(new BigNumber(_i0_93));
        }
        _2713_v1 = _nw527;
        let _2716_v2;
        _2716_v2 = false;
        let _index452 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_2713_v1).length));
        (_2713_v1)[_index452] = _2716_v2;
        let _2717_v3;
        _2717_v3 = _dafny.Seq.UnicodeFromString("wluvi");
        let _2718_v4;
        _2718_v4 = _dafny.Map.Empty.slice().updateUnsafe(_2717_v3,new BigNumber(-752));
        let _2719_v6;
        _2719_v6 = _dafny.Seq.of(_2717_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(511)), function (_2720_i2) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        }));
        let _2721_v9;
        _2721_v9 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ugrshctcj"),false);
        let _2722_v10;
        _2722_v10 = _dafny.Map.Empty.slice().updateUnsafe(_2716_v2,_2716_v2);
        let _2723_v11;
        _2723_v11 = _dafny.Map.Empty.slice().updateUnsafe(_2722_v10,_2718_v4);
        let _2724_v12;
        _2724_v12 = _dafny.Seq.of(_2718_v4);
        let _2725_v13;
        let _nw528 = Array((new BigNumber(28)).toNumber());
        _nw528[(_dafny.ZERO).toNumber()] = _2718_v4;
        _nw528[(_dafny.ONE).toNumber()] = _2718_v4;
        _nw528[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2717_v3,_2712_i0);
        _nw528[(new BigNumber(3)).toNumber()] = (_2718_v4).Merge(_2718_v4);
        _nw528[(new BigNumber(4)).toNumber()] = _2718_v4;
        _nw528[(new BigNumber(5)).toNumber()] = _2718_v4;
        _nw528[(new BigNumber(6)).toNumber()] = _2718_v4;
        _nw528[(new BigNumber(7)).toNumber()] = _2718_v4;
        _nw528[(new BigNumber(8)).toNumber()] = (_2718_v4).Merge(function () {
          let _coll89 = new _dafny.Map();
          for (const _compr_89 of (_2719_v6).Elements) {
            let _2726_v5 = _compr_89;
            if (_dafny.Seq.contains(_2719_v6, _2726_v5)) {
              _coll89.push([_2726_v5,new BigNumber((_dafny.Seq.UnicodeFromString("dwkk")).length)]);
            }
          }
          return _coll89;
        }());
        _nw528[(new BigNumber(9)).toNumber()] = (_2718_v4).update(_2717_v3, _2712_i0);
        _nw528[(new BigNumber(10)).toNumber()] = (_2718_v4).Merge(_2718_v4);
        _nw528[(new BigNumber(11)).toNumber()] = (_2718_v4).Merge((_2718_v4).update(_2717_v3, _2711_v0));
        _nw528[(new BigNumber(12)).toNumber()] = (function () {
          let _coll90 = new _dafny.Map();
          for (const _compr_90 of (_2719_v6).Elements) {
            let _2727_v7 = _compr_90;
            if (_dafny.Seq.contains(_2719_v6, _2727_v7)) {
              _coll90.push([_2727_v7,new BigNumber((_dafny.MultiSet.fromElements(_2712_i0, new BigNumber((_2717_v3).length))).cardinality())]);
            }
          }
          return _coll90;
        }()).Merge(function () {
          let _coll91 = new _dafny.Map();
          for (const _compr_91 of (_2721_v9).Keys.Elements) {
            let _2728_v8 = _compr_91;
            if ((_2721_v9).contains(_2728_v8)) {
              _coll91.push([_2728_v8,_2712_i0]);
            }
          }
          return _coll91;
        }());
        _nw528[(new BigNumber(13)).toNumber()] = _module.__default.fm69(_2711_v0, globalState);
        _nw528[(new BigNumber(14)).toNumber()] = _2718_v4;
        _nw528[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2717_v3,_module.__default.fm0(_2716_v2, globalState));
        _nw528[(new BigNumber(16)).toNumber()] = (_2718_v4).Merge(_2718_v4);
        _nw528[(new BigNumber(17)).toNumber()] = (_2718_v4).Merge(_2718_v4);
        _nw528[(new BigNumber(18)).toNumber()] = ((_2716_v2) ? (_2718_v4) : (_2718_v4));
        _nw528[(new BigNumber(19)).toNumber()] = (_2718_v4).Merge(_2718_v4);
        _nw528[(new BigNumber(20)).toNumber()] = _2718_v4;
        _nw528[(new BigNumber(21)).toNumber()] = _2718_v4;
        _nw528[(new BigNumber(22)).toNumber()] = _2718_v4;
        _nw528[(new BigNumber(23)).toNumber()] = _2718_v4;
        _nw528[(new BigNumber(24)).toNumber()] = (((_2723_v11).contains(_dafny.Map.Empty.slice().updateUnsafe(_2716_v2,_2716_v2))) ? ((_2723_v11).get(_dafny.Map.Empty.slice().updateUnsafe(_2716_v2,_2716_v2))) : (_2718_v4));
        _nw528[(new BigNumber(25)).toNumber()] = _2718_v4;
        _nw528[(new BigNumber(26)).toNumber()] = _2718_v4;
        _nw528[(new BigNumber(27)).toNumber()] = (_2724_v12)[_module.__default.safeIndex(new BigNumber(-890), new BigNumber((_2724_v12).length))];
        _2725_v13 = _nw528;
        let _index453 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_2725_v13).length));
        (_2725_v13)[_index453] = (_2724_v12)[_module.__default.safeIndex(_2712_i0, new BigNumber((_2724_v12).length))];
        let _2729_v14;
        _2729_v14 = _dafny.MultiSet.fromElements(_2716_v2);
        let _2730_v15;
        _2730_v15 = _dafny.Set.fromElements(_2712_i0, new BigNumber(929));
        let _2731_v16;
        _2731_v16 = _dafny.Seq.of(_2711_v0, new BigNumber((_2730_v15).length));
        let _2732_v17;
        _2732_v17 = _dafny.Seq.of(false, true);
        let _2733_v18;
        _2733_v18 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(!(_2716_v2), globalState),_2711_v0);
        let _2734_v19;
        let _nw529 = Array((new BigNumber(24)).toNumber());
        _nw529[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_2711_v0);
        _nw529[(_dafny.ONE).toNumber()] = (((_2729_v14).contains(_2716_v2)) ? ((_2729_v14).get(_2716_v2)) : (_2711_v0));
        _nw529[(new BigNumber(2)).toNumber()] = _2711_v0;
        _nw529[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(_2712_i0, _2711_v0)).cardinality());
        _nw529[(new BigNumber(4)).toNumber()] = new BigNumber(181);
        _nw529[(new BigNumber(5)).toNumber()] = new BigNumber((_2729_v14).cardinality());
        _nw529[(new BigNumber(6)).toNumber()] = _2712_i0;
        _nw529[(new BigNumber(7)).toNumber()] = _2712_i0;
        _nw529[(new BigNumber(8)).toNumber()] = _2711_v0;
        _nw529[(new BigNumber(9)).toNumber()] = (_2731_v16)[_module.__default.safeIndex(_2711_v0, new BigNumber((_2731_v16).length))];
        _nw529[(new BigNumber(10)).toNumber()] = new BigNumber((_2732_v17).length);
        _nw529[(new BigNumber(11)).toNumber()] = _2711_v0;
        _nw529[(new BigNumber(12)).toNumber()] = _2711_v0;
        _nw529[(new BigNumber(13)).toNumber()] = new BigNumber((_2733_v18).length);
        _nw529[(new BigNumber(14)).toNumber()] = new BigNumber((_2717_v3).length);
        _nw529[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_2731_v16).length));
        _nw529[(new BigNumber(16)).toNumber()] = new BigNumber(568);
        _nw529[(new BigNumber(17)).toNumber()] = new BigNumber(188);
        _nw529[(new BigNumber(18)).toNumber()] = _2712_i0;
        _nw529[(new BigNumber(19)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("k")).length);
        _nw529[(new BigNumber(20)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2712_i0,(_dafny.ZERO).minus(_2711_v0))).length);
        _nw529[(new BigNumber(21)).toNumber()] = _2711_v0;
        _nw529[(new BigNumber(22)).toNumber()] = (_2731_v16)[_module.__default.safeIndex(_2711_v0, new BigNumber((_2731_v16).length))];
        _nw529[(new BigNumber(23)).toNumber()] = (_dafny.ZERO).minus(_2711_v0);
        _2734_v19 = _nw529;
        let _2735_v20;
        _2735_v20 = _dafny.Map.Empty.slice().updateUnsafe(_2734_v19,_dafny.Seq.Concat(_2732_v17, _dafny.Seq.update(_module.__default.fm52(true, _2711_v0, _2711_v0, new BigNumber(229), globalState), _module.__default.safeIndex(_2712_i0, new BigNumber((_module.__default.fm52(true, _2711_v0, _2711_v0, new BigNumber(229), globalState)).length)), true)));
        let _index454 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_2713_v1).length));
        let _index455 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_2725_v13).length));
        let _rhs475 = (_dafny.MultiSet.fromElements(_2716_v2)).IsSubsetOf(_2729_v14);
        let _rhs476 = _2712_i0;
        let _rhs477 = (_2724_v12)[_module.__default.safeIndex(_2712_i0, new BigNumber((_2724_v12).length))];
        let _rhs478 = (_2712_i0).plus(_2712_i0);
        let _rhs479 = _2735_v20;
        let _lhs368 = _2713_v1;
        let _lhs369 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_2713_v1).length));
        let _lhs370 = globalState;
        let _lhs371 = _2725_v13;
        let _lhs372 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_2725_v13).length));
        _lhs368[_lhs369] = _rhs475;
        _lhs370.f7 = _rhs476;
        _lhs371[_lhs372] = _rhs477;
        _2711_v0 = _rhs478;
        _2735_v20 = _rhs479;
        let _2736_v21;
        let _init94 = ((_2737_v3) => function (_2738_i3) {
          return _2737_v3;
        })(_2717_v3);
        let _nw530 = Array((new BigNumber(3)).toNumber());
        for (let _i0_94 = 0; _i0_94 < new BigNumber(_nw530.length); _i0_94++) {
          _nw530[_i0_94] = _init94(new BigNumber(_i0_94));
        }
        _2736_v21 = _nw530;
        let _2739_v22;
        _2739_v22 = _dafny.Map.Empty.slice().updateUnsafe(_2736_v21,(new BigNumber((_2732_v17).length)).minus(_2712_i0));
        _2739_v22 = (_2739_v22).update(_2736_v21, _module.__default.safeModuloInt((_dafny.ZERO).minus(_2711_v0), _2711_v0));
        let _index456 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_2713_v1).length));
        (_2713_v1)[_index456] = !((_2713_v1)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_2713_v1).length))]);
        if ((_2713_v1)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_2713_v1).length))]) {
          let _2740_v23;
          _2740_v23 = _module.D19.create_DC39(_dafny.MultiSet.fromElements(_2711_v0));
          let _2741_v24;
          _2741_v24 = _dafny.MultiSet.fromElements(_2712_i0);
          let _2742_v25;
          _2742_v25 = _dafny.Map.Empty.slice().updateUnsafe(_2712_i0,(_2713_v1)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_2713_v1).length))]);
          let _2743_v26;
          _2743_v26 = _module.D7.create_DC11(_2742_v25);
          let _pat_let_tv43 = _2741_v24;
          let _2744_v27;
          _2744_v27 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let55_0) {
            return function (_2745_dt__update__tmp_h0) {
              return function (_pat_let56_0) {
                return function (_2746_dt__update_hcf60_h0) {
                  return _module.D19.create_DC39(_2746_dt__update_hcf60_h0);
                }(_pat_let56_0);
              }(_pat_let_tv43);
            }(_pat_let55_0);
          }(_2740_v23),_2743_v26);
          let _pat_let_tv44 = _2741_v24;
          _2744_v27 = (_2744_v27).update(function (_pat_let57_0) {
            return function (_2747_dt__update__tmp_h1) {
              return function (_pat_let58_0) {
                return function (_2748_dt__update_hcf60_h1) {
                  return _module.D19.create_DC39(_2748_dt__update_hcf60_h1);
                }(_pat_let58_0);
              }(_pat_let_tv44);
            }(_pat_let57_0);
          }(_2740_v23), _2743_v26);
          r0 = ((_dafny.ZERO).minus(new BigNumber((_2729_v14).cardinality()))).plus(_2711_v0);
          (globalState).f7 = _2711_v0;
          let _2749_v28;
          _2749_v28 = _module.D3.create_DC4((_2713_v1)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_2713_v1).length))], _2711_v0, _2711_v0, _2716_v2, (_2713_v1)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_2713_v1).length))]);
          _2717_v3 = _module.__default.fm28(true, _2733_v18, _2749_v28, _2722_v10, globalState);
          (globalState).f7 = _module.__default.safeModuloInt(new BigNumber((_2729_v14).cardinality()), _module.__default.fm0(_2716_v2, globalState));
        } else {
          let _index457 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_2734_v19).length));
          (_2734_v19)[_index457] = _2712_i0;
          let _index458 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_2734_v19).length));
          (_2734_v19)[_index458] = new BigNumber((_2717_v3).length);
          r1 = (_2713_v1)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_2713_v1).length))];
          let _2750_v29;
          let _nw531 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Set.Empty);
          _2750_v29 = _nw531;
          let _2751_v30;
          _2751_v30 = _dafny.Set.fromElements((_2732_v17)[_module.__default.safeIndex((_2734_v19)[_module.__default.safeIndex(new BigNumber(971), new BigNumber((_2734_v19).length))], new BigNumber((_2732_v17).length))], (_2713_v1)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_2713_v1).length))], _2716_v2, _2716_v2, !((_2713_v1)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_2713_v1).length))]));
          let _index459 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_2750_v29).length));
          (_2750_v29)[_index459] = _2751_v30;
          let _index460 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_2750_v29).length));
          (_2750_v29)[_index460] = _2751_v30;
          let _index461 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_2734_v19).length));
          let _rhs480 = (_2713_v1)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_2713_v1).length))];
          let _rhs481 = (_2732_v17)[_module.__default.safeIndex(_module.__default.safeDivisionInt(new BigNumber((_2730_v15).length), (_2734_v19)[_module.__default.safeIndex(new BigNumber(971), new BigNumber((_2734_v19).length))]), new BigNumber((_2732_v17).length))];
          let _rhs482 = !_dafny.areEqual(_2717_v3, _2717_v3);
          let _rhs483 = _module.__default.fm0(((true) ? (_2716_v2) : ((_2713_v1)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_2713_v1).length))])), globalState);
          let _lhs373 = _2734_v19;
          let _lhs374 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_2734_v19).length));
          r1 = _rhs480;
          _2716_v2 = _rhs481;
          _2716_v2 = _rhs482;
          _lhs373[_lhs374] = _rhs483;
          (globalState).f7 = (_dafny.ZERO).minus(_2711_v0);
        }
      }
      let _2752_v31;
      _2752_v31 = new _dafny.CodePoint('t'.codePointAt(0));
      let _2753_v32;
      _2753_v32 = _dafny.Map.Empty.slice().updateUnsafe(_2752_v31,_2711_v0);
      let _2754_v33;
      _2754_v33 = false;
      let _2755_v34;
      _2755_v34 = _dafny.Seq.UnicodeFromString("jpujsfs");
      let _2756_v35;
      _2756_v35 = _dafny.Map.Empty.slice().updateUnsafe(_2711_v0,new BigNumber((_2755_v34).length));
      let _2757_v36;
      let _nw532 = new _module.C1();
      _nw532.__ctor(_2754_v33, _2711_v0);
      _2757_v36 = _nw532;
      let _2758_v37;
      _2758_v37 = _module.D20.create_DC42(new BigNumber((_dafny.Seq.of(false)).length), _2754_v33, _2757_v36);
      let _2759_v38;
      let _nw533 = new _module.C6();
      _nw533.__ctor(_2711_v0, _2754_v33, (_2758_v37).dtor_cf66, _2754_v33, _2757_v36.f14);
      _2759_v38 = _nw533;
      let _2760_v39;
      _2760_v39 = _dafny.Seq.of(_2759_v38);
      let _2761_v40;
      _2761_v40 = _dafny.Map.Empty.slice().updateUnsafe(_2759_v38.f16,_2754_v33);
      _2753_v32 = (_2753_v32).update(_2752_v31, new BigNumber((_module.__default.fm28(_2754_v33, _2756_v35, _module.D3.create_DC4(true, new BigNumber((_2760_v39).length), _2711_v0, _2754_v33, _2759_v38.f16), _2761_v40, globalState)).length));
      let _2762_v42;
      let _init95 = ((_2763_v38) => function (_2764_i4) {
        return _module.__default.safeModuloInt(_2764_i4, new BigNumber((function () {
          let _coll92 = new _dafny.Set();
          for (const _compr_92 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_2763_v38.f16, true),(_2763_v38).f17)).Keys.Elements) {
            let _2765_v41 = _compr_92;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_2763_v38.f16, true),(_2763_v38).f17)).contains(_2765_v41)) {
              _coll92.add(_2765_v41);
            }
          }
          return _coll92;
        }()).length));
      })(_2759_v38);
      let _nw534 = Array((new BigNumber(20)).toNumber());
      for (let _i0_95 = 0; _i0_95 < new BigNumber(_nw534.length); _i0_95++) {
        _nw534[_i0_95] = _init95(new BigNumber(_i0_95));
      }
      _2762_v42 = _nw534;
      _2762_v42 = _2762_v42;
      let _2766_v43;
      let _nw535 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.of());
      _2766_v43 = _nw535;
      let _2767_v44;
      _2767_v44 = _dafny.Seq.of(true);
      let _index462 = _module.__default.safeIndex(new BigNumber(754), new BigNumber((_2766_v43).length));
      (_2766_v43)[_index462] = _2767_v44;
      let _index463 = _module.__default.safeIndex(new BigNumber(754), new BigNumber((_2766_v43).length));
      (_2766_v43)[_index463] = _2767_v44;
      let _hi25 = _2757_v36.f14;
      for (let _2768_i5 = _2757_v36.f14; _2768_i5.isLessThan(_hi25); _2768_i5 = _2768_i5.plus(_dafny.ONE)) {
        let _2769_v45;
        let _nw536 = Array((new BigNumber(9)).toNumber());
        _nw536[(_dafny.ZERO).toNumber()] = !((_2759_v38).f17);
        _nw536[(_dafny.ONE).toNumber()] = (_2759_v38).f17;
        _nw536[(new BigNumber(2)).toNumber()] = _2759_v38.f16;
        _nw536[(new BigNumber(3)).toNumber()] = false;
        _nw536[(new BigNumber(4)).toNumber()] = !(!(_2759_v38.f16));
        _nw536[(new BigNumber(5)).toNumber()] = (_2757_v36.f14).isLessThan(_2711_v0);
        _nw536[(new BigNumber(6)).toNumber()] = !_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("kgmbqxi"), _2752_v31);
        _nw536[(new BigNumber(7)).toNumber()] = _2759_v38.f16;
        _nw536[(new BigNumber(8)).toNumber()] = (_2759_v38.f16) && (_2754_v33);
        _2769_v45 = _nw536;
        _2769_v45 = _2769_v45;
        let _2770_v46;
        _2770_v46 = _dafny.Set.fromElements(_2759_v38.f16);
        let _2771_v47;
        _2771_v47 = _dafny.Map.Empty.slice().updateUnsafe(_2768_i5,_2770_v46);
        let _2772_v48;
        let _nw537 = Array((new BigNumber(20)).toNumber());
        _nw537[(_dafny.ZERO).toNumber()] = _2770_v46;
        _nw537[(_dafny.ONE).toNumber()] = (((_2771_v47).contains(_2768_i5)) ? ((_2771_v47).get(_2768_i5)) : (_2770_v46));
        _nw537[(new BigNumber(2)).toNumber()] = _2770_v46;
        _nw537[(new BigNumber(3)).toNumber()] = _2770_v46;
        _nw537[(new BigNumber(4)).toNumber()] = _module.__default.fm44((_2759_v38).f17, _2754_v33, globalState);
        _nw537[(new BigNumber(5)).toNumber()] = _module.__default.fm44(!(_2759_v38.f16), _2754_v33, globalState);
        _nw537[(new BigNumber(6)).toNumber()] = _2770_v46;
        _nw537[(new BigNumber(7)).toNumber()] = _2770_v46;
        _nw537[(new BigNumber(8)).toNumber()] = _2770_v46;
        _nw537[(new BigNumber(9)).toNumber()] = _2770_v46;
        _nw537[(new BigNumber(10)).toNumber()] = _dafny.Set.fromElements(_2754_v33, !((_2759_v38).f17), _2754_v33);
        _nw537[(new BigNumber(11)).toNumber()] = _dafny.Set.fromElements((_2759_v38).f17, true);
        _nw537[(new BigNumber(12)).toNumber()] = _2770_v46;
        _nw537[(new BigNumber(13)).toNumber()] = _2770_v46;
        _nw537[(new BigNumber(14)).toNumber()] = _2770_v46;
        _nw537[(new BigNumber(15)).toNumber()] = _dafny.Set.fromElements(_2759_v38.f16, _2759_v38.f16, (_2759_v38).f17);
        _nw537[(new BigNumber(16)).toNumber()] = _2770_v46;
        _nw537[(new BigNumber(17)).toNumber()] = _2770_v46;
        _nw537[(new BigNumber(18)).toNumber()] = _2770_v46;
        _nw537[(new BigNumber(19)).toNumber()] = _dafny.Set.fromElements(false, (_2759_v38).f17);
        _2772_v48 = _nw537;
        let _2773_v49;
        let _nw538 = new _module.C0();
        _nw538.__ctor(_2772_v48, _2768_i5);
        _2773_v49 = _nw538;
        (_2757_v36).f14 = new BigNumber(101);
        r1 = !(_2759_v38.f16) || (!((_2757_v36.f14).isLessThanOrEqualTo(_2768_i5)));
      }
      r0 = (_2711_v0).multipliedBy(_2757_v36.f14);
      r1 = ((_2766_v43)[_module.__default.safeIndex(new BigNumber(754), new BigNumber((_2766_v43).length))])[_module.__default.safeIndex(_2757_v36.f14, new BigNumber(((_2766_v43)[_module.__default.safeIndex(new BigNumber(754), new BigNumber((_2766_v43).length))]).length))];
      return [r0, r1];
    }
  };

  $module.C10 = class C10 {
    constructor () {
      this._tname = "_module.C10";
      this._f16 = false;
      this._f14 = _dafny.ZERO;
      this._f17 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f16() {
      let _this = this;
      return _this._f16;
    };
    set f16(value) {
      let _this = this;
      _this._f16 = value;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    set f14(value) {
      let _this = this;
      _this._f14 = value;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    __ctor(f16, f17, f14) {
      let _this = this;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f14 = f14;
      return;
    }
    fm3(globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe(_this.f14,new BigNumber((_dafny.Seq.UnicodeFromString("kgtxp")).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f14,_this.f14))).Merge((((_this).f17) ? (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f14)),_this.f14)) : (_dafny.Map.Empty.slice().updateUnsafe(_this.f14,new BigNumber((_dafny.Seq.of((_this).f17)).length)))));
    };
    fm9(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(((_this.f16) ? (_dafny.Seq.UnicodeFromString("nsvqcphrb")) : (_dafny.Seq.UnicodeFromString("txojgxj"))), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-432)), function (_2774_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(348)), function (_2775_i1) {
        return new _dafny.CodePoint('i'.codePointAt(0));
      })));
    };
    fm10(p0, p1, p2, globalState) {
      let _this = this;
      return _this.f16;
    };
    m2(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      if ((_this).f17) {
        let _2776_v0;
        _2776_v0 = new _dafny.CodePoint('h'.codePointAt(0));
        let _2777_v1;
        _2777_v1 = _dafny.MultiSet.fromElements(_2776_v0);
        let _2778_v2;
        _2778_v2 = _dafny.Seq.of(_2777_v1);
        _2778_v2 = _dafny.Seq.update(_dafny.Seq.update(_2778_v2, _module.__default.safeIndex(p0, new BigNumber((_2778_v2).length)), _2777_v1), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update(_2778_v2, _module.__default.safeIndex(p0, new BigNumber((_2778_v2).length)), _2777_v1)).length)), (_2778_v2)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(688)), ((_2779_v0) => function (_2780_i0) {
          return _2779_v0;
        })(_2776_v0))).length), new BigNumber((_2778_v2).length))]);
        let _2781_v3;
        _2781_v3 = _dafny.Seq.UnicodeFromString("bdwbsv");
        let _2782_v4;
        _2782_v4 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("scqd"), _2781_v3));
        _2782_v4 = _2782_v4;
        let _2783_v5;
        _2783_v5 = _this.f14;
        let _2784_v6;
        _2784_v6 = _dafny.Map.Empty.slice().updateUnsafe(false,p0);
        let _2785_v7;
        _2785_v7 = _dafny.MultiSet.fromElements(new BigNumber((_2784_v6).length));
        let _2786_v8;
        _2786_v8 = _dafny.Set.fromElements(p1, (_this).f17);
        let _2787_v9;
        _2787_v9 = _dafny.Seq.of((_2783_v5), _this.f14, p0, (_dafny.ZERO).minus((((_2785_v7).contains(new BigNumber((_2786_v8).length))) ? ((_2785_v7).get(new BigNumber((_2786_v8).length))) : (p0))), new BigNumber(446));
        (globalState).f7 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_2787_v9, _2787_v9), _2787_v9)).length);
        let _2788_v10;
        _2788_v10 = _dafny.Seq.of(_this.f16);
        let _2789_v11;
        _2789_v11 = _2788_v10;
        let _2790_v12;
        _2790_v12 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,(_this).f17);
        let _2791_v13;
        let _nw539 = Array((_dafny.ONE).toNumber());
        _nw539[(_dafny.ZERO).toNumber()] = _2784_v6;
        _2791_v13 = _nw539;
        let _2792_v14;
        _2792_v14 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).fm10(_2789_v11, (_this).fm9(_this.f14, _module.__default.fm11(_this.f14, (((_2790_v12).contains(_this.f16)) ? ((_2790_v12).get(_this.f16)) : ((_this).f17)), new BigNumber((_dafny.Set.fromElements(_this.f14, p0)).length), globalState), _2776_v0, p0, globalState), _2776_v0, globalState)),_2791_v13);
        _2792_v14 = (_2792_v14).update(p1, _2791_v13);
        r0 = _this.f14;
      } else {
        let _2793_v15;
        _2793_v15 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this.f14).isLessThanOrEqualTo(_this.f14));
        let _2794_v16;
        _2794_v16 = _module.D3.create_DC3(p1);
        _2793_v15 = (_2793_v15).update((p1) === ((_2794_v16).dtor_cf3), p1);
        if (_this.f16) {
          (_this).f16 = ((_this).f17) === (true);
          (_this).f16 = !(_this.f16);
          let _2795_v17;
          let _nw540 = Array((new BigNumber(6)).toNumber()).fill(false);
          _2795_v17 = _nw540;
          let _2796_v18;
          _2796_v18 = _dafny.Seq.UnicodeFromString("nolroa");
          let _rhs484 = _this.f16;
          let _rhs485 = _2795_v17;
          let _rhs486 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2796_v18, _2796_v18), _dafny.Seq.UnicodeFromString("xkamp"));
          let _lhs375 = _this;
          _lhs375.f16 = _rhs484;
          _2795_v17 = _rhs485;
          _2796_v18 = _rhs486;
          let _2797_v19;
          _2797_v19 = new _dafny.CodePoint('d'.codePointAt(0));
          _2797_v19 = new _dafny.CodePoint('h'.codePointAt(0));
          let _2798_v20;
          _2798_v20 = _dafny.MultiSet.fromElements(p1, (_this).f17, _this.f16);
          (globalState).f7 = (((_2798_v20).contains(_this.f16)) ? ((_2798_v20).get(_this.f16)) : (_this.f14));
        } else {
          let _2799_v21;
          let _init96 = ((_2800_v15) => function (_2801_i1) {
            return (_module.D4.create_DC6(_2800_v15)).dtor_cf10;
          })(_2793_v15);
          let _nw541 = Array((new BigNumber(4)).toNumber());
          for (let _i0_96 = 0; _i0_96 < new BigNumber(_nw541.length); _i0_96++) {
            _nw541[_i0_96] = _init96(new BigNumber(_i0_96));
          }
          _2799_v21 = _nw541;
          let _2802_v22;
          _2802_v22 = _dafny.Seq.of(_2793_v15);
          let _index464 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_2799_v21).length));
          (_2799_v21)[_index464] = ((_2802_v22)[_module.__default.safeIndex(p0, new BigNumber((_2802_v22).length))]).update(_this.f16, p1);
          let _index465 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_2799_v21).length));
          (_2799_v21)[_index465] = (_2793_v15).Merge(_2793_v15);
          let _2803_v23;
          let _nw542 = Array((new BigNumber(4)).toNumber()).fill([]);
          _2803_v23 = _nw542;
          let _2804_v24;
          let _init97 = function (_2805_i2) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          };
          let _nw543 = Array((new BigNumber(9)).toNumber());
          for (let _i0_97 = 0; _i0_97 < new BigNumber(_nw543.length); _i0_97++) {
            _nw543[_i0_97] = _init97(new BigNumber(_i0_97));
          }
          _2804_v24 = _nw543;
          let _index466 = _module.__default.safeIndex(new BigNumber(345), new BigNumber((_2803_v23).length));
          (_2803_v23)[_index466] = _2804_v24;
          let _index467 = _module.__default.safeIndex(new BigNumber(345), new BigNumber((_2803_v23).length));
          (_2803_v23)[_index467] = _2804_v24;
          let _2806_v25;
          let _init98 = function (_2807_i3) {
            return (_this).f17;
          };
          let _nw544 = Array((new BigNumber(22)).toNumber());
          for (let _i0_98 = 0; _i0_98 < new BigNumber(_nw544.length); _i0_98++) {
            _nw544[_i0_98] = _init98(new BigNumber(_i0_98));
          }
          _2806_v25 = _nw544;
          let _index468 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_2806_v25).length));
          (_2806_v25)[_index468] = (_this).f17;
          let _index469 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_2806_v25).length));
          (_2806_v25)[_index469] = (_this).f17;
          (_this).f16 = _this.f16;
          let _index470 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_2806_v25).length));
          (_2806_v25)[_index470] = p1;
        }
        let _2808_v26;
        _2808_v26 = _dafny.MultiSet.fromElements(_this.f14, (_dafny.ZERO).minus(_this.f14));
        let _2809_v27;
        _2809_v27 = _dafny.Seq.of(_this.f16);
        let _2810_v28;
        _2810_v28 = _2809_v27;
        let _2811_v29;
        _2811_v29 = _dafny.Seq.UnicodeFromString("qxnnlkg");
        let _2812_v30;
        _2812_v30 = new _dafny.CodePoint('c'.codePointAt(0));
        let _2813_v31;
        _2813_v31 = new BigNumber(593);
        let _2814_v32;
        _2814_v32 = _dafny.MultiSet.fromElements(p0);
        let _2815_v33;
        let _nw545 = Array((new BigNumber(13)).toNumber());
        _nw545[(_dafny.ZERO).toNumber()] = p1;
        _nw545[(_dafny.ONE).toNumber()] = false;
        _nw545[(new BigNumber(2)).toNumber()] = _this.f16;
        _nw545[(new BigNumber(3)).toNumber()] = (_2808_v26).IsDisjointFrom(_2808_v26);
        _nw545[(new BigNumber(4)).toNumber()] = _this.f16;
        _nw545[(new BigNumber(5)).toNumber()] = !(p1) || (_this.f16);
        _nw545[(new BigNumber(6)).toNumber()] = (_this).fm10(_2809_v27, _2811_v29, _2812_v30, globalState);
        _nw545[(new BigNumber(7)).toNumber()] = !(false);
        _nw545[(new BigNumber(8)).toNumber()] = !(p1);
        _nw545[(new BigNumber(9)).toNumber()] = !(!(_2808_v26).equals(_2808_v26));
        _nw545[(new BigNumber(10)).toNumber()] = (_this).f17;
        _nw545[(new BigNumber(11)).toNumber()] = (((_this).f17) ? ((_this).f17) : (_this.f16));
        _nw545[(new BigNumber(12)).toNumber()] = (_2809_v27)[_module.__default.safeIndex(new BigNumber((_2814_v32).cardinality()), new BigNumber((_2809_v27).length))];
        _2815_v33 = _nw545;
        let _index471 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_2815_v33).length));
        (_2815_v33)[_index471] = (_this).f17;
        let _index472 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_2815_v33).length));
        (_2815_v33)[_index472] = p1;
        let _2816_v34;
        let _nw546 = new _module.C9();
        _nw546.__ctor();
        _2816_v34 = _nw546;
        let _2817_v35;
        _2817_v35 = _dafny.Seq.of(_2811_v29);
        let _2818_v36;
        _2818_v36 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-751)), ((_2819_p0, _2820_v26) => function (_2821_i4) {
          return _dafny.Set.fromElements(_2819_p0, new BigNumber((_2820_v26).cardinality()));
        })(p0, _2808_v26));
        _2817_v35 = _dafny.Seq.update(_2817_v35, _module.__default.safeIndex(_this.f14, new BigNumber((_2817_v35).length)), (_this).fm9(p0, (_2818_v36), _2812_v30, _this.f14, globalState));
      }
      if (!(_this.f16)) {
        let _2822_v37;
        _2822_v37 = _dafny.Seq.UnicodeFromString("ujipqu");
        let _2823_v38;
        let _nw547 = Array((new BigNumber(8)).toNumber()).fill(false);
        _2823_v38 = _nw547;
        let _2824_v39;
        _2824_v39 = _dafny.Seq.of(_2823_v38, _2823_v38);
        let _2825_v40;
        _2825_v40 = _dafny.Seq.of(true);
        let _2826_v41;
        _2826_v41 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_this.f14);
        let _2827_v42;
        _2827_v42 = _dafny.Seq.of(_this.f14, new BigNumber((_2824_v39).length), new BigNumber((_2825_v40).length), new BigNumber((_2826_v41).length));
        let _2828_v43;
        let _nw548 = new _module.C8();
        _nw548.__ctor(_2822_v37, true, !((_this).f17), (_dafny.ZERO).minus((_dafny.ZERO).minus((_2827_v42)[_module.__default.safeIndex(new BigNumber(-8), new BigNumber((_2827_v42).length))])));
        _2828_v43 = _nw548;
        let _2829_v44;
        _2829_v44 = new _dafny.CodePoint('e'.codePointAt(0));
        _2829_v44 = new _dafny.CodePoint('w'.codePointAt(0));
        let _2830_v45;
        let _nw549 = Array((new BigNumber(11)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _2830_v45 = _nw549;
        let _2831_v46;
        _2831_v46 = _module.D12.create_DC23(_this.f16, _this.f14, p0);
        let _pat_let_tv45 = p1;
        let _pat_let_tv46 = p0;
        let _rhs487 = _module.__default.fm0((function (_pat_let59_0) {
          return function (_2832_dt__update__tmp_h2) {
            return function (_pat_let60_0) {
              return function (_2833_dt__update_hcf30_h0) {
                return function (_pat_let61_0) {
                  return function (_2834_dt__update_hcf31_h0) {
                    return _module.D12.create_DC23(_2833_dt__update_hcf30_h0, _2834_dt__update_hcf31_h0, (_2832_dt__update__tmp_h2).dtor_cf32);
                  }(_pat_let61_0);
                }(_pat_let_tv46);
              }(_pat_let60_0);
            }(_pat_let_tv45);
          }(_pat_let59_0);
        }(_2831_v46)).dtor_cf30, globalState);
        let _rhs488 = (_this.f14).plus(((_dafny.ZERO).minus(_this.f14)).minus(_this.f14));
        let _rhs489 = _2830_v45;
        let _rhs490 = p1;
        let _lhs376 = globalState;
        let _lhs377 = globalState;
        let _lhs378 = _this;
        _lhs376.f7 = _rhs487;
        _lhs377.f2 = _rhs488;
        _2830_v45 = _rhs489;
        _lhs378.f16 = _rhs490;
        let _2835_v47;
        _2835_v47 = _dafny.MultiSet.fromElements((_2825_v40)[_module.__default.safeIndex(p0, new BigNumber((_2825_v40).length))]);
        let _2836_v48;
        let _nw550 = new _module.C4();
        _nw550.__ctor(p1, (_2835_v47).IsProperSubsetOf(_2835_v47), p1);
        _2836_v48 = _nw550;
        let _2837_v49;
        _2837_v49 = _dafny.Seq.of(_2825_v40, _2825_v40, _2825_v40);
        let _2838_v50;
        let _out39;
        _out39 = (_2828_v43).m7(new BigNumber((_2835_v47).cardinality()), _2837_v49, globalState);
        _2838_v50 = _out39;
      } else {
        let _2839_v51;
        _2839_v51 = _dafny.MultiSet.fromElements(p1);
        (_this).f16 = (_dafny.MultiSet.fromElements(_this.f16, p1, _this.f16)).IsProperSubsetOf(((false) ? (_dafny.MultiSet.fromElements(false, _this.f16, p1)) : (_2839_v51)));
        let _2840_v52;
        _2840_v52 = _dafny.Seq.UnicodeFromString("rsbh");
        let _2841_v53;
        _2841_v53 = _dafny.Set.fromElements(p1);
        let _2842_v54;
        _2842_v54 = _dafny.Seq.of(_dafny.MultiSet.fromElements(p0));
        let _2843_v55;
        _2843_v55 = _module.D23.create_DC48(_this.f16, _2842_v54);
        let _pat_let_tv47 = p1;
        let _2844_v56;
        let _nw551 = new _module.C3();
        _nw551.__ctor(new BigNumber(-458), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qaguuedgl"), _2840_v52), !(_2841_v53).equals(_2841_v53), (function (_pat_let62_0) {
          return function (_2845_dt__update__tmp_h3) {
            return function (_pat_let63_0) {
              return function (_2846_dt__update_hcf74_h0) {
                return _module.D23.create_DC48(_2846_dt__update_hcf74_h0, (_2845_dt__update__tmp_h3).dtor_cf75);
              }(_pat_let63_0);
            }(!(_pat_let_tv47));
          }(_pat_let62_0);
        }(_2843_v55)).dtor_cf74);
        _2844_v56 = _nw551;
        let _2847_v57;
        _2847_v57 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f17) || (_this.f16),((_this).f17) === (_this.f16));
        _2847_v57 = (_2847_v57).update((_this).f17, ((_this).f17) || (_this.f16));
        let _2848_v58;
        let _init99 = ((_2849_p0) => function (_2850_i5) {
          return _module.__default.safeDivisionInt(_2850_i5, _2849_p0);
        })(p0);
        let _nw552 = Array((new BigNumber(28)).toNumber());
        for (let _i0_99 = 0; _i0_99 < new BigNumber(_nw552.length); _i0_99++) {
          _nw552[_i0_99] = _init99(new BigNumber(_i0_99));
        }
        _2848_v58 = _nw552;
        (_this).f16 = (_dafny.MultiSet.fromElements(_2848_v58)).contains(_2848_v58);
        (_this).f16 = (_this).f17;
      }
      (_this).f16 = !((p0).isLessThan(_this.f14));
      let _2851_v59;
      let _init100 = ((_2852_p0) => function (_2853_i6) {
        return (_2853_i6).minus(_2852_p0);
      })(p0);
      let _nw553 = Array((new BigNumber(24)).toNumber());
      for (let _i0_100 = 0; _i0_100 < new BigNumber(_nw553.length); _i0_100++) {
        _nw553[_i0_100] = _init100(new BigNumber(_i0_100));
      }
      _2851_v59 = _nw553;
      let _2854_v60;
      _2854_v60 = _dafny.MultiSet.fromElements(_2851_v59, _2851_v59);
      if ((_2854_v60).IsProperSubsetOf(_2854_v60)) {
        if ((_this.f14).isLessThanOrEqualTo(_module.__default.safeDivisionInt(p0, (_dafny.ZERO).minus(p0)))) {
          (_this).f16 = (_this).f17;
          (globalState).f2 = p0;
          let _2855_v61;
          _2855_v61 = _dafny.MultiSet.fromElements(p0);
          let _index473 = _module.__default.safeIndex(new BigNumber(433), new BigNumber((_2851_v59).length));
          (_2851_v59)[_index473] = _module.__default.safeDivisionInt((((_2855_v61).contains(_this.f14)) ? ((_2855_v61).get(_this.f14)) : (_this.f14)), p0);
          let _2856_v62;
          _2856_v62 = _dafny.Map.Empty.slice().updateUnsafe(_2851_v59,p0);
          let _index474 = _module.__default.safeIndex(new BigNumber(433), new BigNumber((_2851_v59).length));
          (_2851_v59)[_index474] = new BigNumber((((_2856_v62).Merge(_2856_v62)).Merge(_2856_v62)).length);
          (_this).f16 = !(!(p1)) || (true);
          let _2857_v63;
          _2857_v63 = _dafny.Seq.UnicodeFromString("qdlxlyqfn");
          _2857_v63 = _dafny.Seq.Concat(_2857_v63, _2857_v63);
        } else {
          let _2858_v64;
          _2858_v64 = _dafny.Seq.UnicodeFromString("rjmveunxy");
          let _2859_v65;
          _2859_v65 = _module.D12.create_DC22(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("giwttqei"), _2858_v64));
          _2859_v65 = _2859_v65;
          (_this).f16 = !(p1);
          (_this).f16 = (_this).f17;
          (_this).f16 = _this.f16;
          let _2860_v66;
          let _nw554 = Array((new BigNumber(15)).toNumber()).fill(false);
          _2860_v66 = _nw554;
          let _index475 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_2860_v66).length));
          (_2860_v66)[_index475] = (_this.f16) === (_this.f16);
          let _index476 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_2860_v66).length));
          (_2860_v66)[_index476] = _dafny.Seq.contains(_2858_v64, new _dafny.CodePoint('b'.codePointAt(0)));
        }
        let _2861_v67;
        let _nw555 = new _module.C9();
        _nw555.__ctor();
        _2861_v67 = _nw555;
        let _index477 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_2851_v59).length));
        (_2851_v59)[_index477] = _this.f14;
        let _index478 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_2851_v59).length));
        (_2851_v59)[_index478] = _this.f14;
        let _2862_v68;
        _2862_v68 = _dafny.Seq.UnicodeFromString("q");
        (_this).f16 = _dafny.Seq.IsProperPrefixOf(_2862_v68, _2862_v68);
        (_this).f16 = p1;
      } else {
        let _2863_v69;
        _2863_v69 = _dafny.Seq.of(_this.f14);
        let _2864_v70;
        _2864_v70 = _dafny.Set.fromElements(_2863_v69, _2863_v69, _2863_v69, _dafny.Seq.Create(_module.__default.abs(new BigNumber(385)), function (_2865_i7) {
          return new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("y")))).cardinality());
        }));
        if ((_2864_v70).IsProperSubsetOf(_2864_v70)) {
          let _2866_v71;
          let _nw556 = new _module.C3();
          _nw556.__ctor(new BigNumber(472), _dafny.Seq.UnicodeFromString("hkqdyi"), (_this.f14).isLessThan(p0), (p1) || (p1));
          _2866_v71 = _nw556;
          let _index479 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_2851_v59).length));
          (_2851_v59)[_index479] = (_dafny.ZERO).minus((_this.f14).minus(p0));
          let _2867_v72;
          _2867_v72 = _dafny.Seq.UnicodeFromString("gwnvcomc");
          let _index480 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_2851_v59).length));
          let _rhs491 = _this.f14;
          let _rhs492 = new BigNumber(433);
          let _rhs493 = _module.__default.safeModuloInt(p0, _this.f14);
          let _rhs494 = _dafny.Seq.Concat(((!(true)) ? (_2867_v72) : ((_2866_v71).f23)), _2867_v72);
          let _lhs379 = _2851_v59;
          let _lhs380 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_2851_v59).length));
          let _lhs381 = globalState;
          let _lhs382 = globalState;
          _lhs379[_lhs380] = _rhs491;
          _lhs381.f7 = _rhs492;
          _lhs382.f7 = _rhs493;
          _2867_v72 = _rhs494;
          let _2868_v73;
          let _nw557 = Array((new BigNumber(26)).toNumber()).fill(false);
          _2868_v73 = _nw557;
          let _2869_v74;
          _2869_v74 = _dafny.Map.Empty.slice().updateUnsafe((_2851_v59)[_module.__default.safeIndex(new BigNumber(125), new BigNumber((_2851_v59).length))],p1);
          let _index481 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_2868_v73).length));
          (_2868_v73)[_index481] = (((_2869_v74).contains(p0)) ? ((_2869_v74).get(p0)) : (_this.f16));
          let _index482 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_2868_v73).length));
          (_2868_v73)[_index482] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_2863_v69, _2863_v69), _2863_v69);
          let _2870_v75;
          let _nw558 = new _module.C9();
          _nw558.__ctor();
          _2870_v75 = _nw558;
          let _2871_v77;
          _2871_v77 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll93 = new _dafny.Map();
            for (const _compr_93 of _dafny.IntegerRange(new BigNumber(320), new BigNumber(498))) {
              let _2872_v76 = _compr_93;
              if (((new BigNumber(320)).isLessThanOrEqualTo(_2872_v76)) && ((_2872_v76).isLessThan(new BigNumber(498)))) {
                _coll93.push([(_2872_v76).minus(new BigNumber(-33)),(_2851_v59)[_module.__default.safeIndex(new BigNumber(125), new BigNumber((_2851_v59).length))]]);
              }
            }
            return _coll93;
          }(),(_2868_v73)[_module.__default.safeIndex(new BigNumber(908), new BigNumber((_2868_v73).length))]);
          let _2873_v78;
          _2873_v78 = _dafny.MultiSet.fromElements(_this.f14, (_2851_v59)[_module.__default.safeIndex(new BigNumber(125), new BigNumber((_2851_v59).length))], _this.f14);
          let _2874_v79;
          _2874_v79 = _dafny.Map.Empty.slice().updateUnsafe((_2851_v59)[_module.__default.safeIndex(new BigNumber(125), new BigNumber((_2851_v59).length))],new BigNumber((_2873_v78).cardinality()));
          (globalState).f2 = new BigNumber((((_2871_v77).update(_2874_v79, _this.f16)).update(_2874_v79, !(_this.f16) || ((_2868_v73)[_module.__default.safeIndex(new BigNumber(908), new BigNumber((_2868_v73).length))]))).length);
        } else {
          let _2875_v80;
          _2875_v80 = _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(p1, _this.f16)).length));
          let _2876_v81;
          let _nw559 = Array((new BigNumber(2)).toNumber()).fill(false);
          _2876_v81 = _nw559;
          let _2877_v82;
          _2877_v82 = _module.D9.create_DC15(p1, p1, _2875_v80, _2851_v59, _2876_v81);
          let _2878_v83;
          _2878_v83 = _dafny.Set.fromElements(false, (_this).f17, (_2877_v82).dtor_cf19, (_this).f17, p1);
          let _2879_v84;
          _2879_v84 = _module.D15.create_DC29(_2878_v83);
          let _2880_v85;
          _2880_v85 = _dafny.Seq.of(p1, _this.f16);
          let _2881_v86;
          let _nw560 = Array((new BigNumber(20)).toNumber());
          _nw560[(_dafny.ZERO).toNumber()] = _dafny.Set.fromElements(false, (_this).f17, _this.f16);
          _nw560[(_dafny.ONE).toNumber()] = _2878_v83;
          _nw560[(new BigNumber(2)).toNumber()] = _2878_v83;
          _nw560[(new BigNumber(3)).toNumber()] = _2878_v83;
          _nw560[(new BigNumber(4)).toNumber()] = _2878_v83;
          _nw560[(new BigNumber(5)).toNumber()] = _2878_v83;
          _nw560[(new BigNumber(6)).toNumber()] = _2878_v83;
          _nw560[(new BigNumber(7)).toNumber()] = _2878_v83;
          _nw560[(new BigNumber(8)).toNumber()] = _module.__default.fm44(_this.f16, p1, globalState);
          _nw560[(new BigNumber(9)).toNumber()] = _2878_v83;
          _nw560[(new BigNumber(10)).toNumber()] = _2878_v83;
          _nw560[(new BigNumber(11)).toNumber()] = _2878_v83;
          _nw560[(new BigNumber(12)).toNumber()] = (_2879_v84).dtor_cf43;
          _nw560[(new BigNumber(13)).toNumber()] = _2878_v83;
          _nw560[(new BigNumber(14)).toNumber()] = _2878_v83;
          _nw560[(new BigNumber(15)).toNumber()] = _2878_v83;
          _nw560[(new BigNumber(16)).toNumber()] = _dafny.Set.fromElements((_2880_v85)[_module.__default.safeIndex(new BigNumber(920), new BigNumber((_2880_v85).length))], _this.f16, (_2880_v85)[_module.__default.safeIndex(_this.f14, new BigNumber((_2880_v85).length))], !((_this).f17), !(true));
          _nw560[(new BigNumber(17)).toNumber()] = _module.__default.fm44(_module.__default.fm17((_this).f17, p0, new BigNumber(831), _this.f14, globalState), p1, globalState);
          _nw560[(new BigNumber(18)).toNumber()] = _module.__default.fm44((_this).f17, (_this).f17, globalState);
          _nw560[(new BigNumber(19)).toNumber()] = _module.__default.fm44((_this).f17, false, globalState);
          _2881_v86 = _nw560;
          let _2882_v87;
          _2882_v87 = _dafny.Seq.UnicodeFromString("hubavys");
          let _2883_v88;
          let _nw561 = new _module.C0();
          _nw561.__ctor(_2881_v86, new BigNumber((_2882_v87).length));
          _2883_v88 = _nw561;
          (_this).f16 = !(_this.f16);
          (_this).f16 = !(_2883_v88.f27).isEqualTo(p0);
          let _2884_v89;
          _2884_v89 = new _dafny.CodePoint('x'.codePointAt(0));
          let _2885_v90;
          _2885_v90 = _dafny.Map.Empty.slice().updateUnsafe(_2884_v89,_this.f14);
          let _2886_v91;
          _2886_v91 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2885_v90);
          _2886_v91 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(762),_dafny.Map.Empty.slice().updateUnsafe(_2884_v89,(_dafny.ZERO).minus(p0)));
          (_2883_v88).f27 = _2883_v88.f27;
        }
        let _2887_v92;
        let _nw562 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Set.Empty);
        _2887_v92 = _nw562;
        let _2888_v93;
        _2888_v93 = _dafny.Map.Empty.slice().updateUnsafe(true,_this.f14);
        let _2889_v94;
        _2889_v94 = _dafny.Set.fromElements(_2888_v93);
        let _index483 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_2887_v92).length));
        (_2887_v92)[_index483] = _2889_v94;
        let _2890_v95;
        _2890_v95 = _dafny.MultiSet.fromElements(_this.f16, _this.f16, p1);
        let _2891_v96;
        _2891_v96 = _dafny.Seq.of(_this.f16);
        let _2892_v97;
        _2892_v97 = _module.__default.fm52((_this).f17, _this.f14, new BigNumber((_dafny.Seq.of(_2890_v95, _dafny.MultiSet.fromElements(_this.f16), _dafny.MultiSet.FromArray(_2891_v96), _2890_v95)).length), (((_2888_v93).contains(p1)) ? ((_2888_v93).get(p1)) : (_this.f14)), globalState);
        let _2893_v98;
        _2893_v98 = new _dafny.CodePoint('v'.codePointAt(0));
        let _index484 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_2887_v92).length));
        let _rhs495 = (_this).fm10(_2892_v97, _dafny.Seq.Create(_module.__default.abs(new BigNumber(462)), function (_2894_i8) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        }), _2893_v98, globalState);
        let _rhs496 = (_2889_v94).Union(_2889_v94);
        let _lhs383 = _this;
        let _lhs384 = _2887_v92;
        let _lhs385 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_2887_v92).length));
        _lhs383.f16 = _rhs495;
        _lhs384[_lhs385] = _rhs496;
        let _2895_v99;
        _2895_v99 = _dafny.Seq.UnicodeFromString("qhplwtr");
        let _rhs497 = (_dafny.ZERO).minus(new BigNumber((_2895_v99).length));
        let _rhs498 = _module.__default.safeModuloInt(_this.f14, (_dafny.ZERO).minus(p0));
        let _lhs386 = _this;
        let _lhs387 = globalState;
        _lhs386.f14 = _rhs497;
        _lhs387.f7 = _rhs498;
        r0 = (new BigNumber(721)).minus(_this.f14);
        (_this).f16 = !((_module.D9.create_DC16(_this.f14, false)).dtor_cf24);
      }
      let _2896_v100;
      _2896_v100 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-935),_this.f14);
      let _2897_v101;
      _2897_v101 = _dafny.Map.Empty.slice().updateUnsafe(!(p1),(_dafny.Map.Empty.slice().updateUnsafe(p0,_module.__default.fm0(_this.f16, globalState))).Merge(_2896_v100));
      _2897_v101 = (_2897_v101).update((_this).f17, (_this).fm3(globalState));
      let _2898_v102;
      _2898_v102 = _module.D7.create_DC12();
      let _2899_v103;
      _2899_v103 = _dafny.Seq.of(p1);
      let _2900_v104;
      _2900_v104 = _dafny.Map.Empty.slice().updateUnsafe(_2898_v102,_2899_v103);
      _2900_v104 = _2900_v104;
      r0 = p0;
      return r0;
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _2901_v0;
      _2901_v0 = new _dafny.CodePoint('k'.codePointAt(0));
      _2901_v0 = _2901_v0;
      if (!((_this).f17)) {
        let _2902_v1;
        _2902_v1 = _dafny.Set.fromElements((_this).f17);
        let _2903_v2;
        _2903_v2 = _dafny.Seq.UnicodeFromString("mqllnxkei");
        let _2904_v3;
        _2904_v3 = _dafny.Map.Empty.slice().updateUnsafe(_2902_v1,_2903_v2);
        _2904_v3 = (_2904_v3).update((_2902_v1).Union(_2902_v1), _2903_v2);
        let _2905_v4;
        let _nw563 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
        _2905_v4 = _nw563;
        let _index485 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_2905_v4).length));
        (_2905_v4)[_index485] = (p1).multipliedBy(_this.f14);
        let _index486 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_2905_v4).length));
        (_2905_v4)[_index486] = _this.f14;
        let _2906_v5;
        _2906_v5 = _dafny.Set.fromElements(p0, new _dafny.CodePoint('h'.codePointAt(0)), _2901_v0);
        let _2907_v6;
        _2907_v6 = _dafny.Set.fromElements(_2906_v5);
        _2907_v6 = _2907_v6;
        _2905_v4 = _2905_v4;
        let _2908_v7;
        _2908_v7 = _module.D9.create_DC16(_this.f14, p2);
        let _2909_v8;
        _2909_v8 = _module.D9.create_DC17(_2908_v7);
        let _pat_let_tv48 = p1;
        let _2910_v9;
        let _nw564 = Array((_dafny.ONE).toNumber());
        _nw564[(_dafny.ZERO).toNumber()] = function (_pat_let64_0) {
          return function (_2911_dt__update__tmp_h0) {
            return function (_pat_let65_0) {
              return function (_2912_dt__update_hcf25_h0) {
                return _module.D9.create_DC17(_2912_dt__update_hcf25_h0);
              }(_pat_let65_0);
            }(_module.D9.create_DC16(_pat_let_tv48, (_this).f17));
          }(_pat_let64_0);
        }(_2909_v8);
        _2910_v9 = _nw564;
        let _index487 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_2910_v9).length));
        (_2910_v9)[_index487] = _2909_v8;
        let _index488 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_2910_v9).length));
        (_2910_v9)[_index488] = _2909_v8;
      } else {
        let _2913_v10;
        _2913_v10 = _dafny.MultiSet.fromElements(_this.f14);
        let _2914_v11;
        let _nw565 = Array((new BigNumber(8)).toNumber());
        _nw565[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(p1)).cardinality());
        _nw565[(_dafny.ONE).toNumber()] = (_this.f14).multipliedBy(_this.f14);
        _nw565[(new BigNumber(2)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_2913_v10).cardinality()), new BigNumber(-184));
        _nw565[(new BigNumber(3)).toNumber()] = _this.f14;
        _nw565[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt(_this.f14, p1);
        _nw565[(new BigNumber(5)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("ripqy")).length), _this.f14);
        _nw565[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("fes")).length));
        _nw565[(new BigNumber(7)).toNumber()] = _module.__default.safeModuloInt(p1, p1);
        _2914_v11 = _nw565;
        let _index489 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_2914_v11).length));
        (_2914_v11)[_index489] = p1;
        let _index490 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_2914_v11).length));
        (_2914_v11)[_index490] = (_dafny.ZERO).minus(p1);
        let _2915_v12;
        let _nw566 = new _module.C7();
        _nw566.__ctor(_this.f16, (_this).f17, _this.f14);
        _2915_v12 = _nw566;
        let _index491 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_2914_v11).length));
        (_2914_v11)[_index491] = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_dafny.ZERO).minus(p3), new BigNumber((_dafny.Seq.of(_2915_v12)).length)));
        let _index492 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_2914_v11).length));
        (_2914_v11)[_index492] = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(new BigNumber(-631), _this.f14), new BigNumber(128));
        let _2916_v13;
        _2916_v13 = _dafny.Seq.UnicodeFromString("ip");
        let _2917_v14;
        _2917_v14 = _dafny.Seq.of(_this.f16);
        let _2918_v15;
        let _nw567 = new _module.C8();
        _nw567.__ctor(_dafny.Seq.Concat(_2916_v13, _2916_v13), !(p2) || ((_this).fm10(_2917_v14, _dafny.Seq.UnicodeFromString("x"), _2901_v0, globalState)), !(_this.f16) || (p2), (new BigNumber((_2916_v13).length)).multipliedBy(_this.f14));
        _2918_v15 = _nw567;
        let _2919_v16;
        _2919_v16 = _dafny.Map.Empty.slice().updateUnsafe(_2914_v11,_2914_v11);
        let _2920_v17;
        _2920_v17 = _dafny.Map.Empty.slice().updateUnsafe((((_2919_v16).contains(_2914_v11)) ? ((_2919_v16).get(_2914_v11)) : (_2914_v11)),_2901_v0);
        _2920_v17 = (_2920_v17).update(_2914_v11, p0);
      }
      let _2921_v18;
      _2921_v18 = _module.D12.create_DC23((_this).f17, p3, p1);
      let _2922_v19;
      _2922_v19 = _dafny.Seq.UnicodeFromString("lneeohgn");
      (_this).f16 = (_module.__default.safeDivisionInt((_2921_v18).dtor_cf32, new BigNumber((_2922_v19).length))).isLessThanOrEqualTo(p3);
      let _2923_v20;
      let _nw568 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
      _2923_v20 = _nw568;
      let _index493 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_2923_v20).length));
      (_2923_v20)[_index493] = p1;
      let _index494 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_2923_v20).length));
      (_2923_v20)[_index494] = (_dafny.ZERO).minus(_this.f14);
      let _2924_v21;
      let _nw569 = Array((new BigNumber(13)).toNumber());
      _nw569[(_dafny.ZERO).toNumber()] = p0;
      _nw569[(_dafny.ONE).toNumber()] = _2901_v0;
      _nw569[(new BigNumber(2)).toNumber()] = p0;
      _nw569[(new BigNumber(3)).toNumber()] = p0;
      _nw569[(new BigNumber(4)).toNumber()] = p0;
      _nw569[(new BigNumber(5)).toNumber()] = _2901_v0;
      _nw569[(new BigNumber(6)).toNumber()] = _2901_v0;
      _nw569[(new BigNumber(7)).toNumber()] = _module.__default.fm39(globalState);
      _nw569[(new BigNumber(8)).toNumber()] = _2901_v0;
      _nw569[(new BigNumber(9)).toNumber()] = _2901_v0;
      _nw569[(new BigNumber(10)).toNumber()] = p0;
      _nw569[(new BigNumber(11)).toNumber()] = ((p2) ? (_2901_v0) : (p0));
      _nw569[(new BigNumber(12)).toNumber()] = p0;
      _2924_v21 = _nw569;
      let _index495 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_2924_v21).length));
      (_2924_v21)[_index495] = (_2922_v19)[_module.__default.safeIndex(new BigNumber(782), new BigNumber((_2922_v19).length))];
      let _2925_v22;
      _2925_v22 = _dafny.Seq.of(p1, ((_this.f16) ? ((_2923_v20)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_2923_v20).length))]) : ((_2923_v20)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_2923_v20).length))])), (_2923_v20)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_2923_v20).length))]);
      let _2926_v23;
      _2926_v23 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm21(globalState),_this.f14);
      let _2927_v24;
      _2927_v24 = _2926_v23;
      let _index496 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_2924_v21).length));
      let _index497 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_2923_v20).length));
      let _rhs499 = _2901_v0;
      let _rhs500 = p1;
      let _rhs501 = _module.__default.fm34((_2921_v18).dtor_cf32, (_2923_v20)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_2923_v20).length))], _this.f16, ((p2) ? (p1) : (p1)), globalState);
      let _rhs502 = _2925_v22;
      let _rhs503 = _module.__default.safeModuloInt(new BigNumber(((_2927_v24)).length), (_dafny.ZERO).minus((p3).plus(p3)));
      let _lhs388 = _2924_v21;
      let _lhs389 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_2924_v21).length));
      let _lhs390 = _2923_v20;
      let _lhs391 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_2923_v20).length));
      _lhs388[_lhs389] = _rhs499;
      _lhs390[_lhs391] = _rhs500;
      _2901_v0 = _rhs501;
      _2925_v22 = _rhs502;
      r0 = _rhs503;
      let _index498 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_2924_v21).length));
      (_2924_v21)[_index498] = new _dafny.CodePoint('w'.codePointAt(0));
      let _2928_v25;
      _2928_v25 = _dafny.Set.fromElements(p3, (_2925_v22)[_module.__default.safeIndex(p1, new BigNumber((_2925_v22).length))], new BigNumber((_2922_v19).length), (_2923_v20)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_2923_v20).length))]);
      let _2929_v26;
      _2929_v26 = _dafny.Map.Empty.slice().updateUnsafe(_2928_v25,p1);
      r0 = new BigNumber((_module.__default.fm66(p1, _2929_v26, (new BigNumber((_2922_v19).length)).multipliedBy(p1), p2, globalState)).length);
      return r0;
    }
  };

  $module.C11 = class C11 {
    constructor () {
      this._tname = "_module.C11";
      this._f14 = _dafny.ZERO;
      this._f16 = false;
      this._f17 = false;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f14() {
      let _this = this;
      return _this._f14;
    };
    set f14(value) {
      let _this = this;
      _this._f14 = value;
    };
    get f16() {
      let _this = this;
      return _this._f16;
    };
    set f16(value) {
      let _this = this;
      _this._f16 = value;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    __ctor(f14, f16, f17) {
      let _this = this;
      (_this)._f14 = f14;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      return;
    }
    fm3(globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_dafny.Seq.of(_this.f16, !(_this.f16)))).length)),(_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(false, _this.f16)).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f14,_this.f14));
    };
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return _this.f16;
    };
    fm5(p0, p1, globalState) {
      let _this = this;
      return _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("owwe"));
    };
    m2(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _2930_i0;
      _2930_i0 = _dafny.ZERO;
      L18: {
        while (p1) {
          C18: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2930_i0)) {
              break L18;
            }
            _2930_i0 = (_2930_i0).plus(_dafny.ONE);
            let _2931_v0;
            let _nw570 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
            _2931_v0 = _nw570;
            let _index499 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_2931_v0).length));
            (_2931_v0)[_index499] = p0;
            let _2932_v1;
            _2932_v1 = new _dafny.CodePoint('f'.codePointAt(0));
            let _2933_v2;
            _2933_v2 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2932_v1);
            let _index500 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_2931_v0).length));
            (_2931_v0)[_index500] = new BigNumber(((_2933_v2).update(_this.f16, _2932_v1)).length);
            let _2934_v3;
            _2934_v3 = _dafny.Seq.UnicodeFromString("lcloiqko");
            let _2935_v4;
            _2935_v4 = _dafny.Seq.of((_this).f17);
            let _2936_v5;
            _2936_v5 = _dafny.Map.Empty.slice().updateUnsafe(_2935_v4,_module.__default.fm6(new BigNumber((_module.__default.fm7(p1, (_this).f17, p1, globalState)).length), _this.f16, !(false), p1, globalState));
            let _2937_v6;
            _2937_v6 = _dafny.Seq.of(_2934_v3);
            _2934_v3 = (((_2936_v5).contains(_2935_v4)) ? ((_2936_v5).get(_2935_v4)) : (_dafny.Seq.Concat((_2937_v6)[_module.__default.safeIndex((_2931_v0)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_2931_v0).length))], new BigNumber((_2937_v6).length))], _dafny.Seq.UnicodeFromString("chchihtm"))));
            if (p1) {
              let _2938_v7;
              _2938_v7 = _dafny.Seq.of(p1);
              _2935_v4 = _dafny.Seq.of(_this.f16);
              let _2939_v8;
              let _nw571 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
              _2939_v8 = _nw571;
              let _2940_v10;
              _2940_v10 = _dafny.Map.Empty.slice().updateUnsafe(function () {
                let _coll94 = new _dafny.Set();
                for (const _compr_94 of _dafny.IntegerRange(new BigNumber(350), new BigNumber(-32))) {
                  let _2941_v9 = _compr_94;
                  if (((new BigNumber(350)).isLessThanOrEqualTo(_2941_v9)) && ((_2941_v9).isLessThan(new BigNumber(-32)))) {
                    _coll94.add((_2941_v9).minus(p0));
                  }
                }
                return _coll94;
              }(),_this.f14);
              let _index501 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_2939_v8).length));
              (_2939_v8)[_index501] = _2940_v10;
              let _2942_v11;
              let _nw572 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _2942_v11 = _nw572;
              let _index502 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_2942_v11).length));
              (_2942_v11)[_index502] = _2934_v3;
              let _2943_v12;
              let _nw573 = Array((new BigNumber(23)).toNumber()).fill(false);
              _2943_v12 = _nw573;
              let _2944_v13;
              _2944_v13 = _2943_v12;
              let _2945_v14;
              let _nw574 = Array((new BigNumber(25)).toNumber());
              _nw574[(_dafny.ZERO).toNumber()] = _2943_v12;
              _nw574[(_dafny.ONE).toNumber()] = _2943_v12;
              _nw574[(new BigNumber(2)).toNumber()] = _2943_v12;
              _nw574[(new BigNumber(3)).toNumber()] = _2943_v12;
              _nw574[(new BigNumber(4)).toNumber()] = _2943_v12;
              _nw574[(new BigNumber(5)).toNumber()] = _2943_v12;
              _nw574[(new BigNumber(6)).toNumber()] = (_2944_v13);
              _nw574[(new BigNumber(7)).toNumber()] = _2943_v12;
              _nw574[(new BigNumber(8)).toNumber()] = _2943_v12;
              _nw574[(new BigNumber(9)).toNumber()] = ((_this.f16) ? (_2943_v12) : (_2943_v12));
              _nw574[(new BigNumber(10)).toNumber()] = _2943_v12;
              _nw574[(new BigNumber(11)).toNumber()] = _2943_v12;
              _nw574[(new BigNumber(12)).toNumber()] = ((_this.f16) ? (_2943_v12) : (_2943_v12));
              _nw574[(new BigNumber(13)).toNumber()] = _2943_v12;
              _nw574[(new BigNumber(14)).toNumber()] = _2943_v12;
              _nw574[(new BigNumber(15)).toNumber()] = _2943_v12;
              _nw574[(new BigNumber(16)).toNumber()] = (((_this).f17) ? (_2943_v12) : (_2943_v12));
              _nw574[(new BigNumber(17)).toNumber()] = _2943_v12;
              _nw574[(new BigNumber(18)).toNumber()] = _2943_v12;
              _nw574[(new BigNumber(19)).toNumber()] = _2943_v12;
              _nw574[(new BigNumber(20)).toNumber()] = _2943_v12;
              _nw574[(new BigNumber(21)).toNumber()] = _2943_v12;
              _nw574[(new BigNumber(22)).toNumber()] = _2943_v12;
              _nw574[(new BigNumber(23)).toNumber()] = _2943_v12;
              _nw574[(new BigNumber(24)).toNumber()] = _2943_v12;
              _2945_v14 = _nw574;
              let _index503 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_2945_v14).length));
              (_2945_v14)[_index503] = _2943_v12;
              let _2946_v15;
              _2946_v15 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_this.f14), _this.f14);
              let _index504 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_2939_v8).length));
              let _index505 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_2942_v11).length));
              let _index506 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_2945_v14).length));
              let _rhs504 = _this.f16;
              let _rhs505 = _2940_v10;
              let _rhs506 = _2934_v3;
              let _rhs507 = _2943_v12;
              let _rhs508 = new BigNumber(((_module.__default.fm8(!(p1), (_this).f17, globalState)).update(_2946_v15, ((!(p1)) ? (_2934_v3) : (_2934_v3)))).length);
              let _lhs392 = _this;
              let _lhs393 = _2939_v8;
              let _lhs394 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_2939_v8).length));
              let _lhs395 = _2942_v11;
              let _lhs396 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_2942_v11).length));
              let _lhs397 = _2945_v14;
              let _lhs398 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_2945_v14).length));
              let _lhs399 = globalState;
              _lhs392.f16 = _rhs504;
              _lhs393[_lhs394] = _rhs505;
              _lhs395[_lhs396] = _rhs506;
              _lhs397[_lhs398] = _rhs507;
              _lhs399.f7 = _rhs508;
              let _2947_v16;
              _2947_v16 = _dafny.Seq.of(_2943_v12);
              let _2948_v17;
              let _nw575 = Array((new BigNumber(27)).toNumber());
              _nw575[(_dafny.ZERO).toNumber()] = _2944_v13;
              _nw575[(_dafny.ONE).toNumber()] = _2944_v13;
              _nw575[(new BigNumber(2)).toNumber()] = _2944_v13;
              _nw575[(new BigNumber(3)).toNumber()] = _2944_v13;
              _nw575[(new BigNumber(4)).toNumber()] = _2944_v13;
              _nw575[(new BigNumber(5)).toNumber()] = _2944_v13;
              _nw575[(new BigNumber(6)).toNumber()] = (_2945_v14)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_2945_v14).length))];
              _nw575[(new BigNumber(7)).toNumber()] = _2944_v13;
              _nw575[(new BigNumber(8)).toNumber()] = _2944_v13;
              _nw575[(new BigNumber(9)).toNumber()] = ((_this.f16) ? (_2944_v13) : (_2944_v13));
              _nw575[(new BigNumber(10)).toNumber()] = ((p1) ? (_2944_v13) : (_2944_v13));
              _nw575[(new BigNumber(11)).toNumber()] = _2944_v13;
              _nw575[(new BigNumber(12)).toNumber()] = _2944_v13;
              _nw575[(new BigNumber(13)).toNumber()] = _2944_v13;
              _nw575[(new BigNumber(14)).toNumber()] = _2944_v13;
              _nw575[(new BigNumber(15)).toNumber()] = (_2945_v14)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_2945_v14).length))];
              _nw575[(new BigNumber(16)).toNumber()] = _2944_v13;
              _nw575[(new BigNumber(17)).toNumber()] = _2944_v13;
              _nw575[(new BigNumber(18)).toNumber()] = _2944_v13;
              _nw575[(new BigNumber(19)).toNumber()] = (_2945_v14)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_2945_v14).length))];
              _nw575[(new BigNumber(20)).toNumber()] = _2944_v13;
              _nw575[(new BigNumber(21)).toNumber()] = (_2945_v14)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_2945_v14).length))];
              _nw575[(new BigNumber(22)).toNumber()] = _2943_v12;
              _nw575[(new BigNumber(23)).toNumber()] = _2944_v13;
              _nw575[(new BigNumber(24)).toNumber()] = _2944_v13;
              _nw575[(new BigNumber(25)).toNumber()] = (_2947_v16)[_module.__default.safeIndex((_2931_v0)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_2931_v0).length))], new BigNumber((_2947_v16).length))];
              _nw575[(new BigNumber(26)).toNumber()] = _2944_v13;
              _2948_v17 = _nw575;
              let _index507 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_2948_v17).length));
              (_2948_v17)[_index507] = _2944_v13;
              let _index508 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_2948_v17).length));
              (_2948_v17)[_index508] = _2944_v13;
              let _2949_v18;
              _2949_v18 = _dafny.Map.Empty.slice().updateUnsafe((_2931_v0)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_2931_v0).length))],_this.f14);
              let _2950_v19;
              _2950_v19 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,new BigNumber((_2949_v18).length));
              let _2951_v20;
              _2951_v20 = (((_2950_v19).contains((_this).f17)) ? ((_2950_v19).get((_this).f17)) : (_this.f14));
              let _index509 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_2931_v0).length));
              (_2931_v0)[_index509] = ((_2931_v0)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_2931_v0).length))]);
              (_this).f16 = (_this).f17;
            } else {
              let _2952_v21;
              let _nw576 = Array((new BigNumber(5)).toNumber()).fill(false);
              _2952_v21 = _nw576;
              let _2953_v22;
              _2953_v22 = _2952_v21;
              let _2954_v23;
              _2954_v23 = _dafny.MultiSet.fromElements(_2953_v22);
              (_this).f16 = !((_2954_v23).IsProperSubsetOf(_2954_v23));
              let _2955_v24;
              _2955_v24 = _dafny.MultiSet.fromElements(_this.f14);
              let _2956_v25;
              let _nw577 = new _module.C6();
              _nw577.__ctor(_this.f14, (_this).f17, _this.f16, _this.f16, new BigNumber((_dafny.Seq.UnicodeFromString("afxfh")).length));
              _2956_v25 = _nw577;
              let _2957_v26;
              _2957_v26 = _module.D20.create_DC42(new BigNumber((_2934_v3).length), _module.__default.fm27(new BigNumber(759), _this.f14, new BigNumber((_2955_v24).cardinality()), _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(613)), globalState), _2956_v25);
              let _2958_v27;
              let _nw578 = new _module.C1();
              _nw578.__ctor((_this).f17, _module.__default.fm0((_2957_v26).dtor_cf66, globalState));
              _2958_v27 = _nw578;
              let _2959_v28;
              let _nw579 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.of());
              _2959_v28 = _nw579;
              let _2960_v29;
              let _nw580 = new _module.C9();
              _nw580.__ctor();
              _2960_v29 = _nw580;
              let _2961_v30;
              _2961_v30 = _dafny.Seq.of(_2960_v29);
              let _index510 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_2959_v28).length));
              (_2959_v28)[_index510] = _dafny.Seq.Concat(_2961_v30, _2961_v30);
              let _index511 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_2959_v28).length));
              (_2959_v28)[_index511] = _2961_v30;
              let _2962_v31;
              _2962_v31 = _dafny.Set.fromElements(_2934_v3);
              _2962_v31 = _2962_v31;
              let _2963_v32;
              _2963_v32 = _module.D12.create_DC22(_dafny.Seq.UnicodeFromString("ocj"));
              let _2964_v33;
              _2964_v33 = _dafny.Set.fromElements(_this.f14, p0);
              _2934_v3 = _dafny.Seq.Concat(_2934_v3, _dafny.Seq.Concat(_dafny.Seq.update((_2963_v32).dtor_cf29, _module.__default.safeIndex(new BigNumber((_2964_v33).length), new BigNumber(((_2963_v32).dtor_cf29).length)), _2932_v1), _2934_v3));
            }
            (_this).f16 = _this.f16;
          }
        }
      }
      let _2965_v34;
      _2965_v34 = _dafny.MultiSet.fromElements(_this.f14, _this.f14);
      let _2966_v35;
      _2966_v35 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,p0);
      let _2967_v36;
      _2967_v36 = _module.D19.create_DC39(_dafny.MultiSet.fromElements(p0, new BigNumber((_2966_v35).length)));
      (_this).f16 = (_2965_v34).IsDisjointFrom((_2967_v36).dtor_cf60);
      let _2968_v37;
      _2968_v37 = _module.D6.create_DC10();
      r0 = (_dafny.ZERO).minus(new BigNumber((function (_source40) {
        if (_source40.is_DC10) {
          return _dafny.Seq.UnicodeFromString("uhsmv");
        } else {
          let _2969___mcc_h0 = (_source40).cf14;
          let _2970_cf14 = _2969___mcc_h0;
          return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("krbh"), _dafny.Seq.UnicodeFromString("vcvjhvxyw"));
        }
      }(_2968_v37)).length));
      let _2971_v38;
      _2971_v38 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm17((_this).f17, _this.f14, p0, _this.f14, globalState),!(true));
      if ((((_2971_v38).contains(_module.__default.fm21(globalState))) ? ((_2971_v38).get(_module.__default.fm21(globalState))) : (_this.f16))) {
        let _2972_v39;
        _2972_v39 = _module.D3.create_DC4(p1, new BigNumber(344), p0, _this.f16, p1);
        if ((_2972_v39).dtor_cf7) {
          let _2973_v40;
          _2973_v40 = _dafny.Seq.UnicodeFromString("eea");
          let _2974_v41;
          _2974_v41 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_this.f14, p0),_2973_v40);
          let _2975_v42;
          let _nw581 = Array((new BigNumber(9)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _2975_v42 = _nw581;
          let _2976_v43;
          _2976_v43 = new _dafny.CodePoint('f'.codePointAt(0));
          let _2977_v44;
          _2977_v44 = _module.D15.create_DC30(_2976_v43);
          let _index512 = _module.__default.safeIndex(new BigNumber(709), new BigNumber((_2975_v42).length));
          (_2975_v42)[_index512] = (_2977_v44).dtor_cf44;
          let _index513 = _module.__default.safeIndex(new BigNumber(709), new BigNumber((_2975_v42).length));
          let _rhs509 = _2974_v41;
          let _rhs510 = _2976_v43;
          let _lhs400 = _2975_v42;
          let _lhs401 = _module.__default.safeIndex(new BigNumber(709), new BigNumber((_2975_v42).length));
          _2974_v41 = _rhs509;
          _lhs400[_lhs401] = _rhs510;
          let _2978_v45;
          _2978_v45 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2973_v40).length),_module.__default.fm28(_module.__default.fm21(globalState), (_this).fm3(globalState), _module.D3.create_DC4(p1, new BigNumber((_2973_v40).length), p0, _this.f16, (_this).f17), _2971_v38, globalState));
          let _rhs511 = _2978_v45;
          let _rhs512 = p0;
          _2978_v45 = _rhs511;
          r0 = _rhs512;
          let _2979_v46;
          _2979_v46 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,(_2975_v42)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_2975_v42).length))]);
          let _2980_v47;
          _2980_v47 = _dafny.Map.Empty.slice().updateUnsafe((((_2979_v46).contains(p1)) ? ((_2979_v46).get(p1)) : (_2976_v43)),_2976_v43);
          let _2981_v48;
          _2981_v48 = _module.D16.create_DC33((_this).f17, p1, _this.f14, (_this).f17);
          _2980_v47 = (_2980_v47).update(_2976_v43, (_2973_v40)[_module.__default.safeIndex((_2981_v48).dtor_cf49, new BigNumber((_2973_v40).length))]);
          let _2982_v49;
          let _nw582 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Set.Empty);
          _2982_v49 = _nw582;
          let _2983_v50;
          let _nw583 = new _module.C0();
          _nw583.__ctor(_2982_v49, (_dafny.ZERO).minus((_dafny.ZERO).minus(p0)));
          _2983_v50 = _nw583;
          let _2984_v51;
          let _nw584 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _2984_v51 = _nw584;
          let _index514 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_2984_v51).length));
          (_2984_v51)[_index514] = p0;
          let _index515 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_2984_v51).length));
          (_2984_v51)[_index515] = p0;
        } else {
          (_this).f16 = (_this).f17;
          let _2985_v52;
          let _nw585 = new _module.C4();
          _nw585.__ctor(true, (_this).f17, (_this).f17);
          _2985_v52 = _nw585;
          let _2986_v53;
          _2986_v53 = _dafny.Seq.of(_2985_v52);
          let _2987_v54;
          _2987_v54 = _dafny.MultiSet.fromElements(_2985_v52, _2985_v52);
          let _2988_v55;
          _2988_v55 = _dafny.Seq.of(_dafny.MultiSet.fromElements((_2986_v53)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(524)), function (_2989_i1) {
            return new _dafny.CodePoint('x'.codePointAt(0));
          })).length))).length), new BigNumber((_2986_v53).length))]), _2987_v54);
          let _2990_v56;
          _2990_v56 = _dafny.Seq.of((_this).f17, false);
          let _2991_v57;
          let _nw586 = new _module.C7();
          _nw586.__ctor(p1, (_this).f17, p0);
          _2991_v57 = _nw586;
          let _2992_v58;
          _2992_v58 = _module.D14.create_DC28(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(477)), function (_2993_i2) {
  return new _dafny.CodePoint('l'.codePointAt(0));
})).length), (_this).f17, _2991_v57);
          let _2994_v59;
          _2994_v59 = _module.D20.create_DC42((_2992_v58).dtor_cf40, _this.f16, _2991_v57);
          let _2995_v61;
          _2995_v61 = _dafny.Set.fromElements((_2985_v52).f21);
          let _2996_v62;
          _2996_v62 = _dafny.Seq.UnicodeFromString("whx");
          let _2997_v63;
          _2997_v63 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2996_v62).length),p1);
          let _2998_v65;
          _2998_v65 = _dafny.Seq.of(new BigNumber((function () {
            let _coll95 = new _dafny.Set();
            for (const _compr_95 of ((function () {
              let _coll96 = new _dafny.Map();
              for (const _compr_96 of (_dafny.MultiSet.fromElements(_2995_v61)).Elements) {
                let _2999_v60 = _compr_96;
                if ((_dafny.MultiSet.fromElements(_2995_v61)).contains(_2999_v60)) {
                  _coll96.push([_2999_v60,_dafny.Map.Empty.slice().updateUnsafe(_2991_v57.f14,(_this).f17)]);
                }
              }
              return _coll96;
            }()).update(_2995_v61, _2997_v63)).Keys.Elements) {
              let _3000_v64 = _compr_95;
              if (((function () {
                let _coll97 = new _dafny.Map();
                for (const _compr_97 of (_dafny.MultiSet.fromElements(_2995_v61)).Elements) {
                  let _2999_v60 = _compr_97;
                  if ((_dafny.MultiSet.fromElements(_2995_v61)).contains(_2999_v60)) {
                    _coll97.push([_2999_v60,_dafny.Map.Empty.slice().updateUnsafe(_2991_v57.f14,(_this).f17)]);
                  }
                }
                return _coll97;
              }()).update(_2995_v61, _2997_v63)).contains(_3000_v64)) {
                _coll95.add(_3000_v64);
              }
            }
            return _coll95;
          }()).length));
          let _3001_v66;
          _3001_v66 = new _dafny.CodePoint('u'.codePointAt(0));
          let _3002_v67;
          _3002_v67 = _module.D15.create_DC30(_3001_v66);
          let _3003_v68;
          _3003_v68 = _dafny.Set.fromElements(_3002_v67);
          let _3004_v70;
          _3004_v70 = _dafny.MultiSet.fromElements(_3001_v66);
          let _3005_v72;
          let _nw587 = Array((new BigNumber(24)).toNumber());
          _nw587[(_dafny.ZERO).toNumber()] = _this.f14;
          _nw587[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((new BigNumber(((_2988_v55)[_module.__default.safeIndex(p0, new BigNumber((_2988_v55).length))]).cardinality())).minus(_this.f14));
          _nw587[(new BigNumber(2)).toNumber()] = p0;
          _nw587[(new BigNumber(3)).toNumber()] = (_this.f14).minus(p0);
          _nw587[(new BigNumber(4)).toNumber()] = (_this.f14).minus(_this.f14);
          _nw587[(new BigNumber(5)).toNumber()] = _module.__default.safeDivisionInt(p0, new BigNumber((_2990_v56).length));
          _nw587[(new BigNumber(6)).toNumber()] = (_2985_v52).fm12(p0, globalState);
          _nw587[(new BigNumber(7)).toNumber()] = p0;
          _nw587[(new BigNumber(8)).toNumber()] = (_2994_v59).dtor_cf65;
          _nw587[(new BigNumber(9)).toNumber()] = _module.__default.safeModuloInt(_2991_v57.f14, p0);
          _nw587[(new BigNumber(10)).toNumber()] = (_2972_v39).dtor_cf6;
          _nw587[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus((_module.__default.fm0((_this).f17, globalState)).minus((_dafny.ZERO).minus(_this.f14)));
          _nw587[(new BigNumber(12)).toNumber()] = _module.__default.safeDivisionInt(p0, _2991_v57.f14);
          _nw587[(new BigNumber(13)).toNumber()] = (_2998_v65)[_module.__default.safeIndex(new BigNumber((_3003_v68).length), new BigNumber((_2998_v65).length))];
          _nw587[(new BigNumber(14)).toNumber()] = p0;
          _nw587[(new BigNumber(15)).toNumber()] = (p0).multipliedBy(_this.f14);
          _nw587[(new BigNumber(16)).toNumber()] = (new BigNumber((function () {
            let _coll98 = new _dafny.Set();
            for (const _compr_98 of _dafny.IntegerRange(new BigNumber(698), new BigNumber(631))) {
              let _3006_v69 = _compr_98;
              if (((new BigNumber(698)).isLessThanOrEqualTo(_3006_v69)) && ((_3006_v69).isLessThan(new BigNumber(631)))) {
                _coll98.add((_3006_v69).minus(new BigNumber((_dafny.MultiSet.FromArray(_2990_v56)).cardinality())));
              }
            }
            return _coll98;
          }()).length)).minus(_this.f14);
          _nw587[(new BigNumber(17)).toNumber()] = p0;
          _nw587[(new BigNumber(18)).toNumber()] = (new BigNumber(509)).minus(p0);
          _nw587[(new BigNumber(19)).toNumber()] = new BigNumber(-150);
          _nw587[(new BigNumber(20)).toNumber()] = (new BigNumber((function () {
            let _coll99 = new _dafny.Set();
            for (const _compr_99 of (_3004_v70).Elements) {
              let _3007_v71 = _compr_99;
              if ((_3004_v70).contains(_3007_v71)) {
                _coll99.add(_3007_v71);
              }
            }
            return _coll99;
          }()).length)).minus((_dafny.ZERO).minus(_this.f14));
          _nw587[(new BigNumber(21)).toNumber()] = _2991_v57.f14;
          _nw587[(new BigNumber(22)).toNumber()] = _this.f14;
          _nw587[(new BigNumber(23)).toNumber()] = p0;
          _3005_v72 = _nw587;
          _3005_v72 = _3005_v72;
          (globalState).f7 = new BigNumber(-954);
          let _3008_v73;
          let _init101 = ((_3009_v57, _3010_v52) => function (_3011_i3) {
            return _module.D21.create_DC44(_3009_v57.f14, (_3010_v52).f21);
          })(_2991_v57, _2985_v52);
          let _nw588 = Array((new BigNumber(7)).toNumber());
          for (let _i0_101 = 0; _i0_101 < new BigNumber(_nw588.length); _i0_101++) {
            _nw588[_i0_101] = _init101(new BigNumber(_i0_101));
          }
          _3008_v73 = _nw588;
          let _3012_v74;
          _3012_v74 = _module.D21.create_DC44(_this.f14, p1);
          let _index516 = _module.__default.safeIndex(new BigNumber(996), new BigNumber((_3008_v73).length));
          (_3008_v73)[_index516] = _3012_v74;
          let _index517 = _module.__default.safeIndex(new BigNumber(996), new BigNumber((_3008_v73).length));
          (_3008_v73)[_index517] = ((true) ? (_3012_v74) : (_3012_v74));
          let _3013_v75;
          _3013_v75 = _module.D15.create_DC30(new _dafny.CodePoint('j'.codePointAt(0)));
          let _3014_v76;
          _3014_v76 = _module.D15.create_DC31(_3013_v75);
          let _3015_v77;
          _3015_v77 = _module.D19.create_DC40(new BigNumber(793), !(p1), _2991_v57.f14);
          let _rhs513 = _3014_v76;
          let _rhs514 = p0;
          let _rhs515 = !((_3015_v77).dtor_cf62) || (p1);
          let _rhs516 = !(_this.f14).isEqualTo(new BigNumber((_2990_v56).length));
          let _rhs517 = _2985_v52;
          let _lhs402 = globalState;
          let _lhs403 = _this;
          let _lhs404 = _this;
          _3014_v76 = _rhs513;
          _lhs402.f7 = _rhs514;
          _lhs403.f16 = _rhs515;
          _lhs404.f16 = _rhs516;
          _2985_v52 = _rhs517;
        }
        let _3016_v78;
        let _init102 = function (_3017_i4) {
          return _this.f16;
        };
        let _nw589 = Array((new BigNumber(13)).toNumber());
        for (let _i0_102 = 0; _i0_102 < new BigNumber(_nw589.length); _i0_102++) {
          _nw589[_i0_102] = _init102(new BigNumber(_i0_102));
        }
        _3016_v78 = _nw589;
        let _index518 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_3016_v78).length));
        (_3016_v78)[_index518] = (p1) === (_this.f16);
        let _3018_v79;
        _3018_v79 = _module.D3.create_DC3(p1);
        let _3019_v80;
        _3019_v80 = _module.D18.create_DC38(_this.f14, (_3018_v79).dtor_cf3, _this.f14);
        let _index519 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_3016_v78).length));
        (_3016_v78)[_index519] = ((_this).f17) && ((_3019_v80).dtor_cf58);
        let _index520 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_3016_v78).length));
        (_3016_v78)[_index520] = (_3016_v78)[_module.__default.safeIndex(new BigNumber(525), new BigNumber((_3016_v78).length))];
        let _3020_v81;
        _3020_v81 = _module.D11.create_DC20();
        let _source41 = _module.D11.create_DC21(_3020_v81);
        if (_source41.is_DC20) {
          let _3021_v82;
          _3021_v82 = _dafny.MultiSet.fromElements(false);
          let _3022_v83;
          _3022_v83 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(333)), function (_3023_i5) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          }),_3021_v82);
          r0 = _module.__default.safeDivisionInt(_module.__default.fm0(p1, globalState), new BigNumber((_3022_v83).length));
          let _index521 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_3016_v78).length));
          (_3016_v78)[_index521] = !(false);
          let _3024_v84;
          _3024_v84 = _dafny.Seq.UnicodeFromString("fycuspo");
          let _3025_v85;
          let _nw590 = new _module.C3();
          _nw590.__ctor(new BigNumber(452), _3024_v84, _this.f16, (_this).f17);
          _3025_v85 = _nw590;
          let _3026_v86;
          _3026_v86 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_3025_v85);
          _3026_v86 = (_3026_v86).update((_this).f17, _3025_v85);
          let _3027_v87;
          _3027_v87 = _dafny.MultiSet.fromElements((_3025_v85).f23);
          let _3028_v88;
          _3028_v88 = _module.D21.create_DC44(new BigNumber((_3027_v87).cardinality()), p1);
          (_this).f16 = (!(_this.f16)) && ((_3028_v88).dtor_cf70);
        } else if (_source41.is_DC19) {
          let _3029___mcc_h1 = (_source41).cf27;
          let _3030_cf27 = _3029___mcc_h1;
          let _3031_v89;
          _3031_v89 = _dafny.Seq.UnicodeFromString("wtco");
          let _3032_v90;
          let _nw591 = new _module.C3();
          _nw591.__ctor(_this.f14, _3031_v89, _this.f16, true);
          _3032_v90 = _nw591;
          _3032_v90 = _3032_v90;
          let _3033_v91;
          _3033_v91 = (_dafny.ZERO).minus(new BigNumber((_3030_cf27).length));
          _3033_v91 = _3033_v91;
          let _3034_v92;
          _3034_v92 = _dafny.Seq.of(p1, (_this).f17);
          let _3035_v93;
          let _nw592 = Array((new BigNumber(13)).toNumber()).fill(_module.D9.Default());
          _3035_v93 = _nw592;
          let _3036_v94;
          _3036_v94 = _dafny.Seq.of(_3035_v93);
          let _3037_v95;
          _3037_v95 = _dafny.Seq.of(_3018_v79, _3018_v79);
          let _3038_v96;
          let _nw593 = new _module.C2();
          _nw593.__ctor(_3018_v79, (_3018_v79).dtor_cf3, !((_module.D12.create_DC24((_this).f17, _3034_v92, _dafny.Seq.update(_3036_v94, _module.__default.safeIndex((_3032_v90).f22, new BigNumber((_3036_v94).length)), _3035_v93), (_this).f17, _dafny.Seq.update(_3037_v95, _module.__default.safeIndex(_this.f14, new BigNumber((_3037_v95).length)), _module.D3.create_DC3((_3016_v78)[_module.__default.safeIndex(new BigNumber(525), new BigNumber((_3016_v78).length))])))).dtor_cf36));
          _3038_v96 = _nw593;
          let _3039_v97;
          _3039_v97 = _module.D12.create_DC22(_dafny.Seq.UnicodeFromString("efq"));
          let _3040_v98;
          _3040_v98 = new _dafny.CodePoint('l'.codePointAt(0));
          let _3041_v99;
          _3041_v99 = _dafny.MultiSet.fromElements(_3040_v98);
          let _3042_v100;
          _3042_v100 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("xsup"));
          let _3043_v101;
          _3043_v101 = _dafny.Set.fromElements((_3032_v90).f22, new BigNumber((_3031_v89).length));
          let _3044_v102;
          _3044_v102 = _dafny.MultiSet.fromElements(p1);
          let _3045_v103;
          let _nw594 = Array((new BigNumber(29)).toNumber());
          _nw594[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_this.f14);
          _nw594[(_dafny.ONE).toNumber()] = p0;
          _nw594[(new BigNumber(2)).toNumber()] = p0;
          _nw594[(new BigNumber(3)).toNumber()] = new BigNumber(((_3039_v97).dtor_cf29).length);
          _nw594[(new BigNumber(4)).toNumber()] = _this.f14;
          _nw594[(new BigNumber(5)).toNumber()] = (_3032_v90).f22;
          _nw594[(new BigNumber(6)).toNumber()] = _this.f14;
          _nw594[(new BigNumber(7)).toNumber()] = (_3019_v80).dtor_cf57;
          _nw594[(new BigNumber(8)).toNumber()] = _this.f14;
          _nw594[(new BigNumber(9)).toNumber()] = (_3032_v90).f22;
          _nw594[(new BigNumber(10)).toNumber()] = (((_3041_v99).contains(new _dafny.CodePoint('u'.codePointAt(0)))) ? ((_3041_v99).get(new _dafny.CodePoint('u'.codePointAt(0)))) : (new BigNumber((_3042_v100).cardinality())));
          _nw594[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(_module.__default.fm0(p1, globalState));
          _nw594[(new BigNumber(12)).toNumber()] = p0;
          _nw594[(new BigNumber(13)).toNumber()] = _this.f14;
          _nw594[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(_3043_v101, _3043_v101)).cardinality());
          _nw594[(new BigNumber(15)).toNumber()] = ((_3032_v90).f22).multipliedBy(p0);
          _nw594[(new BigNumber(16)).toNumber()] = new BigNumber(29);
          _nw594[(new BigNumber(17)).toNumber()] = ((((_2965_v34).contains(p0)) ? ((_2965_v34).get(p0)) : (new BigNumber((_dafny.Seq.of(new BigNumber((_3044_v102).cardinality()))).length)))).minus(p0);
          _nw594[(new BigNumber(18)).toNumber()] = (_3032_v90).f22;
          _nw594[(new BigNumber(19)).toNumber()] = (_dafny.ZERO).minus((_this.f14).multipliedBy(p0));
          _nw594[(new BigNumber(20)).toNumber()] = _this.f14;
          _nw594[(new BigNumber(21)).toNumber()] = (_3032_v90).f22;
          _nw594[(new BigNumber(22)).toNumber()] = (((_2965_v34).contains(p0)) ? ((_2965_v34).get(p0)) : (new BigNumber(915)));
          _nw594[(new BigNumber(23)).toNumber()] = new BigNumber(872);
          _nw594[(new BigNumber(24)).toNumber()] = (p0).plus(_this.f14);
          _nw594[(new BigNumber(25)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("byv"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(418)), ((_3046_v98) => function (_3047_i6) {
            return _3046_v98;
          })(_3040_v98)))).length);
          _nw594[(new BigNumber(26)).toNumber()] = _this.f14;
          _nw594[(new BigNumber(27)).toNumber()] = (_this.f14).multipliedBy(_this.f14);
          _nw594[(new BigNumber(28)).toNumber()] = _module.__default.safeDivisionInt(_this.f14, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_3044_v102)).length));
          _3045_v103 = _nw594;
          let _index522 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_3045_v103).length));
          (_3045_v103)[_index522] = ((_3032_v90).f22).multipliedBy(new BigNumber(((_3032_v90).f23).length));
          let _index523 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_3045_v103).length));
          (_3045_v103)[_index523] = (_3032_v90).f22;
        } else {
          let _3048___mcc_h2 = (_source41).cf28;
          let _3049_cf28 = _3048___mcc_h2;
          _2971_v38 = (_2971_v38).update((_3016_v78)[_module.__default.safeIndex(new BigNumber(525), new BigNumber((_3016_v78).length))], (_3016_v78)[_module.__default.safeIndex(new BigNumber(525), new BigNumber((_3016_v78).length))]);
          let _3050_v104;
          _3050_v104 = _dafny.Seq.of((_this.f14).plus(p0));
          let _3051_v105;
          let _nw595 = Array((new BigNumber(24)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _3051_v105 = _nw595;
          let _3052_v106;
          _3052_v106 = _dafny.Seq.of(_3051_v105);
          let _rhs518 = _3050_v104;
          let _rhs519 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.update(_3052_v106, _module.__default.safeIndex(_this.f14, new BigNumber((_3052_v106).length)), _3051_v105), _3052_v106), _dafny.Seq.update(_dafny.Seq.Concat(_3052_v106, _3052_v106), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_3052_v106, _3052_v106)).length)), _3051_v105));
          _3050_v104 = _rhs518;
          _3052_v106 = _rhs519;
          let _3053_v107;
          let _nw596 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
          _3053_v107 = _nw596;
          let _index524 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_3053_v107).length));
          (_3053_v107)[_index524] = _module.__default.safeModuloInt(_this.f14, _module.__default.fm0((_this).f17, globalState));
          let _index525 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_3053_v107).length));
          (_3053_v107)[_index525] = _this.f14;
          let _3054_v108;
          _3054_v108 = _dafny.Seq.UnicodeFromString("etgkflwva");
          let _3055_v109;
          _3055_v109 = new _dafny.CodePoint('i'.codePointAt(0));
          let _3056_v110;
          _3056_v110 = _dafny.MultiSet.fromElements(((!((_this).fm4(new BigNumber((_3050_v104).length), (_this).f17, (_3053_v107)[_module.__default.safeIndex(new BigNumber(505), new BigNumber((_3053_v107).length))], globalState))) ? (_3054_v108) : (_3054_v108)), _dafny.Seq.update(_3054_v108, _module.__default.safeIndex(_this.f14, new BigNumber((_3054_v108).length)), _3055_v109), _3054_v108, _dafny.Seq.Create(_module.__default.abs(new BigNumber(845)), ((_3057_v109) => function (_3058_i7) {
            return _3057_v109;
          })(_3055_v109)));
          _3056_v110 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(806)), function (_3059_i8) {
            return new _dafny.CodePoint('i'.codePointAt(0));
          }), _module.__default.safeIndex(_this.f14, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(806)), function (_3060_i8) {
            return new _dafny.CodePoint('i'.codePointAt(0));
          })).length)), _3055_v109), _module.__default.safeIndex(_this.f14, new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(806)), function (_3061_i8) {
            return new _dafny.CodePoint('i'.codePointAt(0));
          }), _module.__default.safeIndex(_this.f14, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(806)), function (_3062_i8) {
            return new _dafny.CodePoint('i'.codePointAt(0));
          })).length)), _3055_v109)).length)), _3055_v109), _3054_v108), _dafny.Seq.Concat(_3054_v108, _3054_v108), _dafny.Seq.Create(_module.__default.abs(new BigNumber(967)), ((_3063_v109) => function (_3064_i9) {
            return _3063_v109;
          })(_3055_v109)), _3054_v108);
        }
        let _3065_v111;
        _3065_v111 = new _dafny.CodePoint('l'.codePointAt(0));
        let _3066_v112;
        _3066_v112 = _dafny.Seq.of(p0);
        let _3067_v113;
        _3067_v113 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_3066_v112).length),_3065_v111);
        _3065_v111 = (_module.D6.create_DC9((((_3067_v113).contains(_this.f14)) ? ((_3067_v113).get(_this.f14)) : (_3065_v111)))).dtor_cf14;
      } else {
        (_this).f16 = _this.f16;
        let _3068_v114;
        _3068_v114 = _dafny.Set.fromElements(_module.__default.fm1(globalState));
        _3068_v114 = _3068_v114;
        let _3069_v115;
        let _nw597 = Array((new BigNumber(26)).toNumber()).fill(false);
        _3069_v115 = _nw597;
        let _3070_v116;
        _3070_v116 = _dafny.Seq.of(_3069_v115, _3069_v115, _3069_v115, _3069_v115);
        let _3071_v117;
        let _nw598 = new _module.C1();
        _nw598.__ctor(_this.f16, (_dafny.ZERO).minus(p0));
        _3071_v117 = _nw598;
        let _3072_v118;
        _3072_v118 = _module.D20.create_DC42(_this.f14, _this.f16, _3071_v117);
        let _3073_v119;
        _3073_v119 = _dafny.Set.fromElements(_this.f14);
        let _3074_v120;
        let _nw599 = new _module.C10();
        _nw599.__ctor(p1, (_this).f17, new BigNumber((_3073_v119).length));
        _3074_v120 = _nw599;
        let _3075_v121;
        _3075_v121 = _dafny.Set.fromElements(_3074_v120);
        let _3076_v122;
        let _nw600 = Array((new BigNumber(9)).toNumber());
        _nw600[(_dafny.ZERO).toNumber()] = _this.f14;
        _nw600[(_dafny.ONE).toNumber()] = _module.__default.fm0(_this.f16, globalState);
        _nw600[(new BigNumber(2)).toNumber()] = new BigNumber((_module.__default.fm6((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f14))), false, p1, (_3072_v118).dtor_cf66, globalState)).length);
        _nw600[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Seq.of(new BigNumber((_3075_v121).length))).length);
        _nw600[(new BigNumber(4)).toNumber()] = (p0).plus(_3071_v117.f14);
        _nw600[(new BigNumber(5)).toNumber()] = (_this.f14).plus(_3071_v117.f14);
        _nw600[(new BigNumber(6)).toNumber()] = p0;
        _nw600[(new BigNumber(7)).toNumber()] = p0;
        _nw600[(new BigNumber(8)).toNumber()] = p0;
        _3076_v122 = _nw600;
        let _index526 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((_3076_v122).length));
        (_3076_v122)[_index526] = new BigNumber(387);
        let _3077_v123;
        _3077_v123 = _dafny.Seq.UnicodeFromString("piced");
        let _3078_v124;
        _3078_v124 = _dafny.Seq.of(_3071_v117.f14);
        let _3079_v126;
        _3079_v126 = _dafny.Map.Empty.slice().updateUnsafe(false,_2971_v38);
        let _3080_v127;
        _3080_v127 = (_3079_v126).update(_this.f16, _2971_v38);
        let _3081_v128;
        _3081_v128 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,_3080_v127);
        let _3082_v129;
        _3082_v129 = _dafny.Map.Empty.slice().updateUnsafe(_module.D14.create_DC27(_dafny.Map.Empty.slice().updateUnsafe(_3081_v128,_this.f14)),new BigNumber(-956));
        let _index527 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((_3076_v122).length));
        let _rhs520 = _3070_v116;
        let _rhs521 = new BigNumber(837);
        let _rhs522 = _module.__default.fm27(new BigNumber((_3073_v119).length), new BigNumber((_3077_v123).length), (_3078_v124)[_module.__default.safeIndex(_3071_v117.f14, new BigNumber((_3078_v124).length))], _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_this.f14), globalState);
        let _rhs523 = (_this).f17;
        let _rhs524 = ((_this.f14).multipliedBy((_dafny.ZERO).minus(new BigNumber((function () {
          let _coll100 = new _dafny.Map();
          for (const _compr_100 of (_3082_v129).Keys.Elements) {
            let _3083_v125 = _compr_100;
            if ((_3082_v129).contains(_3083_v125)) {
              _coll100.push([_3083_v125,(_this).f17]);
            }
          }
          return _coll100;
        }()).length)))).multipliedBy(new BigNumber(-76));
        let _lhs405 = globalState;
        let _lhs406 = _this;
        let _lhs407 = _this;
        let _lhs408 = _3076_v122;
        let _lhs409 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((_3076_v122).length));
        _3070_v116 = _rhs520;
        _lhs405.f7 = _rhs521;
        _lhs406.f16 = _rhs522;
        _lhs407.f16 = _rhs523;
        _lhs408[_lhs409] = _rhs524;
        let _index528 = _module.__default.safeIndex(new BigNumber(141), new BigNumber((_3069_v115).length));
        (_3069_v115)[_index528] = (_this).f17;
        let _index529 = _module.__default.safeIndex(new BigNumber(141), new BigNumber((_3069_v115).length));
        (_3069_v115)[_index529] = _this.f16;
        _3078_v124 = _dafny.Seq.Concat(_3078_v124, ((_this.f16) ? (_3078_v124) : (_3078_v124)));
      }
      let _3084_v130;
      _3084_v130 = _dafny.Seq.UnicodeFromString("ttsybhds");
      let _3085_v131;
      let _nw601 = new _module.C3();
      _nw601.__ctor(p0, _3084_v130, _this.f16, p1);
      _3085_v131 = _nw601;
      let _3086_v132;
      _3086_v132 = _dafny.Seq.of((_3085_v131).f22, (_3085_v131).f22);
      let _hi26 = _module.__default.safeDivisionInt(new BigNumber(486), new BigNumber((_3086_v132).length));
      for (let _3087_i10 = (_dafny.ZERO).minus(p0); _3087_i10.isLessThan(_hi26); _3087_i10 = _3087_i10.plus(_dafny.ONE)) {
        let _3088_v133;
        let _init103 = function (_3089_i11) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        };
        let _nw602 = Array((_dafny.ONE).toNumber());
        for (let _i0_103 = 0; _i0_103 < new BigNumber(_nw602.length); _i0_103++) {
          _nw602[_i0_103] = _init103(new BigNumber(_i0_103));
        }
        _3088_v133 = _nw602;
        let _init104 = function (_3090_i12) {
          return new _dafny.CodePoint('n'.codePointAt(0));
        };
        let _nw603 = Array((new BigNumber(14)).toNumber());
        for (let _i0_104 = 0; _i0_104 < new BigNumber(_nw603.length); _i0_104++) {
          _nw603[_i0_104] = _init104(new BigNumber(_i0_104));
        }
        _3088_v133 = _nw603;
        let _3091_v134;
        _3091_v134 = _dafny.Set.fromElements(_3087_i10);
        _3091_v134 = _3091_v134;
        (globalState).f2 = (_this.f14).plus(_module.__default.safeDivisionInt((_3085_v131).f22, new BigNumber(((_3085_v131).f23).length)));
        _2965_v34 = (_2965_v34).Difference(_2965_v34);
      }
      r0 = (p0).multipliedBy(new BigNumber(176));
      return r0;
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _3092_v0;
      _3092_v0 = _dafny.Map.Empty.slice().updateUnsafe(p3,(p1).minus(p3));
      let _3093_v1;
      _3093_v1 = _dafny.Seq.of(p1);
      _3092_v0 = (_3092_v0).update((_this.f14).minus(new BigNumber((_3092_v0).length)), (_3093_v1)[_module.__default.safeIndex(_this.f14, new BigNumber((_3093_v1).length))]);
      let _3094_v2;
      _3094_v2 = _dafny.Seq.of((_this).f17, _this.f16);
      let _3095_v3;
      _3095_v3 = _dafny.Seq.of(_3094_v2, _3094_v2, _3094_v2, _3094_v2, _3094_v2);
      (_this).f16 = !((_dafny.ZERO).minus(new BigNumber((_3095_v3).length))).isEqualTo((_dafny.ZERO).minus(p1));
      let _3096_v4;
      _3096_v4 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,p0);
      _3096_v4 = (_3096_v4).update(p3, new _dafny.CodePoint('a'.codePointAt(0)));
      let _3097_v6;
      let _init105 = ((_3098_p0, _3099_p1) => function (_3100_i1) {
        return (function () {
          let _coll101 = new _dafny.Map();
          for (const _compr_101 of _dafny.IntegerRange(new BigNumber(939), new BigNumber(588))) {
            let _3101_v5 = _compr_101;
            if (((new BigNumber(939)).isLessThanOrEqualTo(_3101_v5)) && ((_3101_v5).isLessThan(new BigNumber(588)))) {
              _coll101.push([(_3101_v5).minus((_dafny.ZERO).minus(_3099_p1)),(_this).f17]);
            }
          }
          return _coll101;
        }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(333)), ((_3102_p0) => function (_3103_i2) {
          return _3102_p0;
        })(_3098_p0))).length),(_this).f17));
      })(p0, p1);
      let _nw604 = Array((new BigNumber(11)).toNumber());
      for (let _i0_105 = 0; _i0_105 < new BigNumber(_nw604.length); _i0_105++) {
        _nw604[_i0_105] = _init105(new BigNumber(_i0_105));
      }
      _3097_v6 = _nw604;
      for (const _guard_loop_14 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_3097_v6).length))) {
        let _3104_i0 = _guard_loop_14;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_3104_i0)) && ((_3104_i0).isLessThan(new BigNumber((_3097_v6).length))))) {
          (_3097_v6)[(_3104_i0)] = (_dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f17)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f14,!(!((_this).f17))));
        }
      }
      (_this).f16 = _this.f16;
      (_this).f16 = ((_module.__default.fm0(p2, globalState)).multipliedBy(new BigNumber((_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(-949)))).length))).isLessThan(p1);
      r0 = _this.f14;
      return r0;
    }
  };

  $module.C12 = class C12 {
    constructor () {
      this._tname = "_module.C12";
      this._f14 = _dafny.ZERO;
      this._f15 = false;
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
    __ctor(f15, f14) {
      let _this = this;
      (_this)._f15 = f15;
      (_this)._f14 = f14;
      return;
    }
    fm2(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _this.f14;
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = [];
      let r2 = _dafny.ZERO;
      let _3105_v0;
      _3105_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15);
      let _hi27 = new BigNumber((_3105_v0).length);
      for (let _3106_i0 = _this.f14; _3106_i0.isLessThan(_hi27); _3106_i0 = _3106_i0.plus(_dafny.ONE)) {
        let _3107_v1;
        let _init106 = ((_3108_v0) => function (_3109_i1) {
          return ((true) ? ((_this).f15) : ((((_3108_v0).contains((_this).f15)) ? ((_3108_v0).get((_this).f15)) : ((_this).f15))));
        })(_3105_v0);
        let _nw605 = Array((new BigNumber(22)).toNumber());
        for (let _i0_106 = 0; _i0_106 < new BigNumber(_nw605.length); _i0_106++) {
          _nw605[_i0_106] = _init106(new BigNumber(_i0_106));
        }
        _3107_v1 = _nw605;
        let _index530 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length));
        (_3107_v1)[_index530] = p0;
        let _index531 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length));
        (_3107_v1)[_index531] = p0;
        let _3110_v2;
        _3110_v2 = _dafny.Seq.of((_3107_v1)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length))], p0);
        let _3111_v3;
        _3111_v3 = _3110_v2;
        let _3112_v4;
        _3112_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_3111_v3)).length),new BigNumber((_3110_v2).length));
        let _3113_v5;
        _3113_v5 = _dafny.Map.Empty.slice().updateUnsafe(_3112_v4,_dafny.Seq.UnicodeFromString("vbsvqm"));
        r2 = (new BigNumber((_3113_v5).length)).plus(_3106_i0);
        if (p0) {
          let _3114_v6;
          let _init107 = ((_3115_i0) => function (_3116_i2) {
            return (_3116_i2).plus(_3115_i0);
          })(_3106_i0);
          let _nw606 = Array((new BigNumber(6)).toNumber());
          for (let _i0_107 = 0; _i0_107 < new BigNumber(_nw606.length); _i0_107++) {
            _nw606[_i0_107] = _init107(new BigNumber(_i0_107));
          }
          _3114_v6 = _nw606;
          let _index532 = _module.__default.safeIndex(new BigNumber(486), new BigNumber((_3114_v6).length));
          (_3114_v6)[_index532] = _3106_i0;
          let _index533 = _module.__default.safeIndex(new BigNumber(486), new BigNumber((_3114_v6).length));
          (_3114_v6)[_index533] = _this.f14;
          let _index534 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length));
          (_3107_v1)[_index534] = p0;
          let _3117_v7;
          let _nw607 = Array((_dafny.ONE).toNumber());
          _nw607[(_dafny.ZERO).toNumber()] = _3114_v6;
          _3117_v7 = _nw607;
          let _index535 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_3117_v7).length));
          (_3117_v7)[_index535] = _3114_v6;
          let _index536 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_3117_v7).length));
          (_3117_v7)[_index536] = _3114_v6;
          let _3118_v8;
          _3118_v8 = _dafny.Seq.UnicodeFromString("goxhiusc");
          let _3119_v9;
          _3119_v9 = _dafny.Set.fromElements((_3107_v1)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length))], false, (_3107_v1)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length))], (_3107_v1)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length))], (_this).f15);
          r2 = ((_3114_v6)[_module.__default.safeIndex(new BigNumber(486), new BigNumber((_3114_v6).length))]).plus((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_3118_v8).length), new BigNumber((_3119_v9).length))));
          let _index537 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length));
          (_3107_v1)[_index537] = !((_this).f15) || (!(p0) || ((_3107_v1)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length))]));
        } else {
          (globalState).f7 = _3106_i0;
          let _3120_v10;
          let _nw608 = new _module.C10();
          _nw608.__ctor(p0, (_this).f15, _3106_i0);
          _3120_v10 = _nw608;
          let _3121_v11;
          _3121_v11 = _dafny.Seq.of(_3120_v10);
          (globalState).f2 = _module.__default.safeDivisionInt(_3106_i0, (_this).fm2(new BigNumber((_3121_v11).length), _3120_v10.f14, _module.__default.fm21(globalState), _3120_v10.f14, globalState));
          let _index538 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length));
          (_3107_v1)[_index538] = (_3120_v10.f14).isEqualTo(_this.f14);
          let _3122_v12;
          let _nw609 = new _module.C4();
          _nw609.__ctor(!((!((_this).f15)) === ((_3107_v1)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length))])), false, (_this).f15);
          _3122_v12 = _nw609;
          let _index539 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length));
          (_3107_v1)[_index539] = !(false);
        }
        if ((_this).f15) {
          let _3123_v13;
          _3123_v13 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_this).f15, globalState),p0);
          _3123_v13 = _module.__default.fm22((_this).f15, globalState);
          let _3124_v14;
          let _nw610 = Array((new BigNumber(29)).toNumber()).fill(_dafny.MultiSet.Empty);
          _3124_v14 = _nw610;
          let _3125_v15;
          _3125_v15 = _dafny.MultiSet.fromElements((((_3107_v1)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length))]) ? ((_3107_v1)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length))]) : ((_3107_v1)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length))])));
          let _index540 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length));
          let _rhs525 = (((_3125_v15).contains(false)) ? ((_3125_v15).get(false)) : (_3106_i0));
          let _rhs526 = (_3107_v1)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length))];
          let _rhs527 = _3124_v14;
          let _rhs528 = _this.f14;
          let _lhs410 = globalState;
          let _lhs411 = _3107_v1;
          let _lhs412 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length));
          let _lhs413 = globalState;
          _lhs410.f2 = _rhs525;
          _lhs411[_lhs412] = _rhs526;
          _3124_v14 = _rhs527;
          _lhs413.f2 = _rhs528;
          let _3126_v16;
          let _init108 = function (_3127_i3) {
            return _dafny.Seq.of((_this).f15, (_this).f15);
          };
          let _nw611 = Array((new BigNumber(16)).toNumber());
          for (let _i0_108 = 0; _i0_108 < new BigNumber(_nw611.length); _i0_108++) {
            _nw611[_i0_108] = _init108(new BigNumber(_i0_108));
          }
          _3126_v16 = _nw611;
          let _index541 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_3126_v16).length));
          (_3126_v16)[_index541] = (((_3107_v1)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length))]) ? (_3110_v2) : (_3110_v2));
          let _3128_v17;
          _3128_v17 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_3110_v2);
          let _3129_v18;
          _3129_v18 = _dafny.Map.Empty.slice().updateUnsafe(_3110_v2,_dafny.Seq.of((_3107_v1)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length))]));
          let _3130_v19;
          _3130_v19 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_3110_v2);
          let _index542 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_3126_v16).length));
          (_3126_v16)[_index542] = (((_3128_v17).contains(_3106_i0)) ? ((_3128_v17).get(_3106_i0)) : (_dafny.Seq.Concat((((_3129_v18).contains((((_3130_v19).contains((_3107_v1)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length))])) ? ((_3130_v19).get((_3107_v1)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length))])) : (_3110_v2)))) ? ((_3129_v18).get((((_3130_v19).contains((_3107_v1)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length))])) ? ((_3130_v19).get((_3107_v1)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length))])) : (_3110_v2)))) : (_dafny.Seq.update(_3110_v2, _module.__default.safeIndex(new BigNumber(767), new BigNumber((_3110_v2).length)), (_this).f15))), _3110_v2)));
          let _3131_v20;
          let _nw612 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
          _3131_v20 = _nw612;
          let _index543 = _module.__default.safeIndex(new BigNumber(914), new BigNumber((_3131_v20).length));
          (_3131_v20)[_index543] = _this.f14;
          let _index544 = _module.__default.safeIndex(new BigNumber(168), new BigNumber((_3131_v20).length));
          (_3131_v20)[_index544] = _3106_i0;
          let _index545 = _module.__default.safeIndex(new BigNumber(914), new BigNumber((_3131_v20).length));
          let _index546 = _module.__default.safeIndex(new BigNumber(168), new BigNumber((_3131_v20).length));
          let _rhs529 = _this.f14;
          let _rhs530 = _3106_i0;
          let _lhs414 = _3131_v20;
          let _lhs415 = _module.__default.safeIndex(new BigNumber(914), new BigNumber((_3131_v20).length));
          let _lhs416 = _3131_v20;
          let _lhs417 = _module.__default.safeIndex(new BigNumber(168), new BigNumber((_3131_v20).length));
          _lhs414[_lhs415] = _rhs529;
          _lhs416[_lhs417] = _rhs530;
          let _index547 = _module.__default.safeIndex(new BigNumber(914), new BigNumber((_3131_v20).length));
          (_3131_v20)[_index547] = _3106_i0;
        } else {
          let _3132_v21;
          _3132_v21 = _dafny.Set.fromElements((_3107_v1)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length))]);
          let _3133_v22;
          _3133_v22 = _dafny.MultiSet.fromElements((_this).f15, (_this).f15, (_this).f15, p0);
          let _3134_v23;
          _3134_v23 = _dafny.Seq.UnicodeFromString("skoh");
          let _3135_v24;
          _3135_v24 = new _dafny.CodePoint('w'.codePointAt(0));
          let _3136_v25;
          let _nw613 = Array((new BigNumber(11)).toNumber());
          _nw613[(_dafny.ZERO).toNumber()] = (((_3133_v22).contains((_3107_v1)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length))])) ? ((_3133_v22).get((_3107_v1)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length))])) : (_3106_i0));
          _nw613[(_dafny.ONE).toNumber()] = _3106_i0;
          _nw613[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_3106_i0,(_3107_v1)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length))])).length);
          _nw613[(new BigNumber(3)).toNumber()] = _3106_i0;
          _nw613[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.update(_3134_v23, _module.__default.safeIndex(_this.f14, new BigNumber((_3134_v23).length)), _3135_v24)).length);
          _nw613[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(15)), ((_3137_i0) => function (_3138_i4) {
            return _3137_i0;
          })(_3106_i0))).length);
          _nw613[(new BigNumber(6)).toNumber()] = _this.f14;
          _nw613[(new BigNumber(7)).toNumber()] = _this.f14;
          _nw613[(new BigNumber(8)).toNumber()] = _3106_i0;
          _nw613[(new BigNumber(9)).toNumber()] = _3106_i0;
          _nw613[(new BigNumber(10)).toNumber()] = new BigNumber(-349);
          _3136_v25 = _nw613;
          let _3139_v26;
          _3139_v26 = _dafny.Map.Empty.slice().updateUnsafe(_3132_v21,_3136_v25);
          let _index548 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length));
          (_3107_v1)[_index548] = _module.__default.fm27(_3106_i0, _3106_i0, _module.__default.fm0(!((_3107_v1)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length))]), globalState), _dafny.Map.Empty.slice().updateUnsafe(_this.f14,new BigNumber((_3139_v26).length)), globalState);
          (globalState).f7 = _3106_i0;
          _3132_v21 = (_3132_v21).Union(_3132_v21);
          let _index549 = _module.__default.safeIndex(new BigNumber(486), new BigNumber((_3136_v25).length));
          (_3136_v25)[_index549] = _3106_i0;
          let _index550 = _module.__default.safeIndex(new BigNumber(486), new BigNumber((_3136_v25).length));
          let _index551 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length));
          let _index552 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length));
          let _rhs531 = new BigNumber(709);
          let _rhs532 = !(_module.__default.fm43((_3133_v22).IsProperSubsetOf(_module.__default.fm1(globalState)), globalState));
          let _rhs533 = ((_3107_v1)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length))]) && (!(_module.__default.fm21(globalState)));
          let _lhs418 = _3136_v25;
          let _lhs419 = _module.__default.safeIndex(new BigNumber(486), new BigNumber((_3136_v25).length));
          let _lhs420 = _3107_v1;
          let _lhs421 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length));
          let _lhs422 = _3107_v1;
          let _lhs423 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length));
          _lhs418[_lhs419] = _rhs531;
          _lhs420[_lhs421] = _rhs532;
          _lhs422[_lhs423] = _rhs533;
          let _index553 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_3107_v1).length));
          (_3107_v1)[_index553] = (((_module.__default.fm21(globalState)) ? (_3106_i0) : ((_3136_v25)[_module.__default.safeIndex(new BigNumber(486), new BigNumber((_3136_v25).length))]))).isLessThanOrEqualTo((_3136_v25)[_module.__default.safeIndex(new BigNumber(486), new BigNumber((_3136_v25).length))]);
        }
      }
      (globalState).f7 = _module.__default.safeModuloInt(_this.f14, _this.f14);
      let _3140_v27;
      _3140_v27 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm17(!(!((_this).f15)), (_dafny.ZERO).minus(_this.f14), _this.f14, _this.f14, globalState),_this.f14);
      let _3141_v28;
      _3141_v28 = _dafny.Seq.UnicodeFromString("p");
      let _3142_v29;
      _3142_v29 = _dafny.Seq.of((_this).f15);
      let _3143_v30;
      _3143_v30 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_3141_v28).length)),_3142_v29);
      _3140_v27 = (_3140_v27).update((_this).f15, (((_this).f15) ? (new BigNumber((_3143_v30).length)) : (_this.f14)));
      let _3144_v31;
      _3144_v31 = _dafny.Seq.of((_dafny.ZERO).minus(_this.f14), _this.f14, _this.f14);
      let _nw614 = Array((new BigNumber(3)).toNumber());
      _nw614[(_dafny.ZERO).toNumber()] = !_dafny.areEqual(_3144_v31, _3144_v31);
      _nw614[(_dafny.ONE).toNumber()] = (_this).f15;
      _nw614[(new BigNumber(2)).toNumber()] = (((_this).f15) ? ((_this).f15) : (p0));
      r1 = _nw614;
      (globalState).f7 = ((_this.f14).minus(new BigNumber(926))).minus(new BigNumber(405));
      let _3145_v32;
      _3145_v32 = true;
      let _3146_v33;
      _3146_v33 = new _dafny.CodePoint('c'.codePointAt(0));
      let _3147_v34;
      _3147_v34 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_this.f14);
      let _3148_v35;
      _3148_v35 = _module.D11.create_DC19(_3147_v34);
      let _3149_v36;
      _3149_v36 = _module.D11.create_DC21(_3148_v35);
      let _pat_let_tv49 = _3145_v32;
      let _rhs534 = (_this).f15;
      let _rhs535 = function (_source42) {
        if (_source42.is_DC20) {
          return _this.f14;
        } else if (_source42.is_DC19) {
          let _3150___mcc_h0 = (_source42).cf27;
          let _3151_cf27 = _3150___mcc_h0;
          return (_module.D18.create_DC38(_this.f14, _pat_let_tv49, new BigNumber((_3151_cf27).length))).dtor_cf59;
        } else {
          let _3152___mcc_h1 = (_source42).cf28;
          let _3153_cf28 = _3152___mcc_h1;
          return _this.f14;
        }
      }(_3149_v36);
      let _rhs536 = (_this.f14).plus(_this.f14);
      let _rhs537 = _3146_v33;
      let _rhs538 = _3141_v28;
      let _lhs424 = _this;
      _3145_v32 = _rhs534;
      r2 = _rhs535;
      _lhs424.f14 = _rhs536;
      _3146_v33 = _rhs537;
      _3141_v28 = _rhs538;
      r0 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_3144_v31, _3144_v31), _module.__default.safeIndex(_this.f14, new BigNumber((_dafny.Seq.Concat(_3144_v31, _3144_v31)).length)), new BigNumber((_3141_v28).length)), _3144_v31);
      let _nw615 = Array((_dafny.ONE).toNumber());
      _nw615[(_dafny.ZERO).toNumber()] = p0;
      r1 = _nw615;
      r2 = _this.f14;
      return [r0, r1, r2];
    }
    get f15() {
      let _this = this;
      return _this._f15;
    };
  };

  $module.C13 = class C13 {
    constructor () {
      this._tname = "_module.C13";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    m0(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = false;
      let _3154_v0;
      _3154_v0 = new BigNumber(-653);
      (globalState).f7 = _module.__default.safeDivisionInt(_3154_v0, _3154_v0);
      if (p0) {
        let _3155_v1;
        _3155_v1 = _dafny.MultiSet.fromElements(p0);
        let _3156_v2;
        _3156_v2 = _dafny.Map.Empty.slice().updateUnsafe(p0,_3154_v0);
        let _3157_v3;
        _3157_v3 = _dafny.Map.Empty.slice().updateUnsafe(_3155_v1,new BigNumber(((_3156_v2).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_3154_v0))).length));
        _3157_v3 = (_3157_v3).update(_3155_v1, (_dafny.ZERO).minus(_module.__default.fm0(p0, globalState)));
        if (p0) {
          let _3158_v4;
          _3158_v4 = _dafny.Map.Empty.slice().updateUnsafe(_3154_v0,_3154_v0);
          let _3159_v5;
          _3159_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _3160_v6;
          _3160_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_3159_v5).length),p0);
          let _3161_v7;
          let _nw616 = Array((new BigNumber(8)).toNumber());
          _nw616[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_3158_v4).length),p0), (_3160_v6).update((_dafny.ZERO).minus(_3154_v0), p0), _dafny.Map.Empty.slice().updateUnsafe(_3154_v0,false), _dafny.Map.Empty.slice().updateUnsafe(_3154_v0,p0), _3160_v6)).cardinality());
          _nw616[(_dafny.ONE).toNumber()] = new BigNumber(761);
          _nw616[(new BigNumber(2)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(102), _3154_v0);
          _nw616[(new BigNumber(3)).toNumber()] = (_3154_v0).minus(_3154_v0);
          _nw616[(new BigNumber(4)).toNumber()] = _3154_v0;
          _nw616[(new BigNumber(5)).toNumber()] = (_3154_v0).minus(_3154_v0);
          _nw616[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_3154_v0);
          _nw616[(new BigNumber(7)).toNumber()] = _3154_v0;
          _3161_v7 = _nw616;
          _3161_v7 = ((false) ? (_3161_v7) : (_3161_v7));
          let _3162_v8;
          let _init109 = ((_3163_p0) => function (_3164_i0) {
            return (_dafny.Set.fromElements(_3163_p0)).IsProperSubsetOf(_dafny.Set.fromElements(_3163_p0));
          })(p0);
          let _nw617 = Array((new BigNumber(18)).toNumber());
          for (let _i0_109 = 0; _i0_109 < new BigNumber(_nw617.length); _i0_109++) {
            _nw617[_i0_109] = _init109(new BigNumber(_i0_109));
          }
          _3162_v8 = _nw617;
          let _index554 = _module.__default.safeIndex(new BigNumber(753), new BigNumber((_3162_v8).length));
          (_3162_v8)[_index554] = p0;
          let _3165_v9;
          _3165_v9 = _dafny.Seq.UnicodeFromString("ft");
          let _index555 = _module.__default.safeIndex(new BigNumber(753), new BigNumber((_3162_v8).length));
          let _rhs539 = _3154_v0;
          let _rhs540 = p0;
          let _rhs541 = _3161_v7;
          let _rhs542 = ((p0) ? (!(p0)) : (_dafny.Seq.IsProperPrefixOf(_3165_v9, _3165_v9)));
          let _lhs425 = _3162_v8;
          let _lhs426 = _module.__default.safeIndex(new BigNumber(753), new BigNumber((_3162_v8).length));
          _3154_v0 = _rhs539;
          _lhs425[_lhs426] = _rhs540;
          _3161_v7 = _rhs541;
          r1 = _rhs542;
          r1 = (_3162_v8)[_module.__default.safeIndex(new BigNumber(753), new BigNumber((_3162_v8).length))];
          let _3166_v10;
          let _nw618 = new _module.C6();
          _nw618.__ctor(_3154_v0, p0, _module.__default.fm17((_3162_v8)[_module.__default.safeIndex(new BigNumber(753), new BigNumber((_3162_v8).length))], _3154_v0, _3154_v0, _3154_v0, globalState), (_3162_v8)[_module.__default.safeIndex(new BigNumber(753), new BigNumber((_3162_v8).length))], _3154_v0);
          _3166_v10 = _nw618;
          (globalState).f2 = _3154_v0;
        } else {
          let _3167_v11;
          _3167_v11 = _dafny.Seq.UnicodeFromString("pkeaspy");
          let _3168_v12;
          _3168_v12 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_3167_v11, _3167_v11), _3167_v11);
          _3168_v12 = (_3168_v12).update(_3167_v11, _module.__default.abs(_3154_v0));
          let _3169_v13;
          _3169_v13 = _module.D12.create_DC22(_3167_v11);
          let _3170_v14;
          _3170_v14 = _dafny.Seq.of(p0);
          (globalState).f2 = (new BigNumber(-285)).multipliedBy((_dafny.ZERO).minus((new BigNumber(((_3169_v13).dtor_cf29).length)).plus(new BigNumber((_3170_v14).length))));
          (globalState).f7 = _3154_v0;
          let _3171_v15;
          let _nw619 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
          _3171_v15 = _nw619;
          let _index556 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_3171_v15).length));
          (_3171_v15)[_index556] = _3154_v0;
          let _index557 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_3171_v15).length));
          (_3171_v15)[_index557] = _3154_v0;
          _3170_v14 = _3170_v14;
        }
        let _3172_v16;
        _3172_v16 = _dafny.Seq.of(!(p0), p0);
        _3172_v16 = _3172_v16;
        let _3173_v17;
        _3173_v17 = new _dafny.CodePoint('s'.codePointAt(0));
        let _3174_v18;
        _3174_v18 = _module.D6.create_DC9(_3173_v17);
        _3174_v18 = _3174_v18;
        let _3175_v19;
        _3175_v19 = _dafny.Seq.of(new BigNumber(802));
        _3175_v19 = _dafny.Seq.of(_module.__default.fm0(p0, globalState));
      } else {
        r1 = p0;
        let _3176_v20;
        _3176_v20 = _dafny.Set.fromElements(false, p0);
        let _3177_v21;
        _3177_v21 = _module.D15.create_DC29(_module.__default.fm44(false, !(p0), globalState));
        let _3178_v22;
        let _nw620 = new _module.C7();
        _nw620.__ctor(p0, p0, new BigNumber((_module.__default.fm70(_3177_v21, p0, _3154_v0, globalState)).length));
        _3178_v22 = _nw620;
        let _3179_v23;
        _3179_v23 = _module.D14.create_DC28(new BigNumber((_3176_v20).length), !(!(p0)), _3178_v22);
        (globalState).f2 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_3179_v23).dtor_cf40));
        r2 = p0;
        let _3180_v24;
        _3180_v24 = _dafny.MultiSet.fromElements(_3154_v0);
        let _3181_v25;
        _3181_v25 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),true);
        let _3182_v26;
        let _nw621 = new _module.C5();
        _nw621.__ctor((_3180_v24).IsSubsetOf(_3180_v24), (p0) || ((((_3181_v25).contains(p0)) ? ((_3181_v25).get(p0)) : (p0))));
        _3182_v26 = _nw621;
        let _3183_v27;
        _3183_v27 = _dafny.Map.Empty.slice().updateUnsafe((((_3181_v25).contains(p0)) ? ((_3181_v25).get(p0)) : (p0)),_dafny.Map.Empty.slice().updateUnsafe(p0,_3181_v25));
        let _3184_v28;
        _3184_v28 = _dafny.Map.Empty.slice().updateUnsafe(_3183_v27,_3178_v22.f14);
        let _3185_v29;
        _3185_v29 = new _dafny.CodePoint('m'.codePointAt(0));
        let _3186_v30;
        _3186_v30 = _dafny.Map.Empty.slice().updateUnsafe(_module.D14.create_DC27(_3184_v28),_3185_v29);
        _3186_v30 = _3186_v30;
      }
      let _3187_v31;
      let _nw622 = new _module.C4();
      _nw622.__ctor(p0, p0, !(p0));
      _3187_v31 = _nw622;
      let _3188_v32;
      _3188_v32 = _dafny.Set.fromElements(_3187_v31);
      let _3189_v33;
      _3189_v33 = _module.D13.create_DC26();
      let _3190_v34;
      _3190_v34 = _dafny.Map.Empty.slice().updateUnsafe(_3188_v32,_3189_v33);
      let _3191_v35;
      _3191_v35 = _dafny.Seq.UnicodeFromString("tm");
      let _3192_v36;
      _3192_v36 = _dafny.Map.Empty.slice().updateUnsafe(_3154_v0,_3154_v0);
      let _3193_v37;
      _3193_v37 = _dafny.Seq.of(_3154_v0, _3154_v0);
      let _3194_v38;
      _3194_v38 = _dafny.Set.fromElements(_dafny.Seq.of(_3154_v0));
      let _3195_v39;
      _3195_v39 = _module.D12.create_DC23(p0, new BigNumber((_3194_v38).length), new BigNumber((_3193_v37).length));
      let _3196_v40;
      let _nw623 = Array((new BigNumber(16)).toNumber());
      _nw623[(_dafny.ZERO).toNumber()] = _module.__default.safeDivisionInt(_3154_v0, new BigNumber((_3190_v34).length));
      _nw623[(_dafny.ONE).toNumber()] = new BigNumber((_3191_v35).length);
      _nw623[(new BigNumber(2)).toNumber()] = new BigNumber(-201);
      _nw623[(new BigNumber(3)).toNumber()] = _3154_v0;
      _nw623[(new BigNumber(4)).toNumber()] = _3154_v0;
      _nw623[(new BigNumber(5)).toNumber()] = _module.__default.fm0(p0, globalState);
      _nw623[(new BigNumber(6)).toNumber()] = _3154_v0;
      _nw623[(new BigNumber(7)).toNumber()] = ((false) ? (_3154_v0) : (_3154_v0));
      _nw623[(new BigNumber(8)).toNumber()] = _3154_v0;
      _nw623[(new BigNumber(9)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(908), _3154_v0);
      _nw623[(new BigNumber(10)).toNumber()] = _3154_v0;
      _nw623[(new BigNumber(11)).toNumber()] = (new BigNumber(-981)).multipliedBy(_3154_v0);
      _nw623[(new BigNumber(12)).toNumber()] = new BigNumber(((_3192_v36).update(_3154_v0, _3154_v0)).length);
      _nw623[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus((_3154_v0).multipliedBy(new BigNumber((_3193_v37).length)));
      _nw623[(new BigNumber(14)).toNumber()] = (_3195_v39).dtor_cf31;
      _nw623[(new BigNumber(15)).toNumber()] = _3154_v0;
      _3196_v40 = _nw623;
      let _index558 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length));
      (_3196_v40)[_index558] = (_3154_v0).plus(_3154_v0);
      let _index559 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length));
      (_3196_v40)[_index559] = _module.__default.safeModuloInt((_dafny.ZERO).minus(_3154_v0), _3154_v0);
      let _3197_i1;
      _3197_i1 = _dafny.ZERO;
      L19: {
        while (p0) {
          C19: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_3197_i1)) {
              break L19;
            }
            _3197_i1 = (_3197_i1).plus(_dafny.ONE);
            let _3198_v41;
            _3198_v41 = _dafny.Seq.of((_3187_v31).f21);
            let _3199_v42;
            _3199_v42 = _module.D16.create_DC33((_3187_v31).f21, (_3187_v31).f21, new BigNumber(465), (_3187_v31).f21);
            let _3200_v43;
            _3200_v43 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsProperPrefixOf(_3198_v41, _dafny.Seq.of(p0, p0, p0)),_3199_v42);
            _3200_v43 = (_3200_v43).update((_3154_v0).isLessThan((_dafny.ZERO).minus((_3196_v40)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length))])), _module.D16.create_DC33(p0, (_3187_v31).f21, (_3196_v40)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length))], (_3187_v31).f21));
            let _3201_v44;
            _3201_v44 = _dafny.Map.Empty.slice().updateUnsafe((_3187_v31).f21,new BigNumber((_3191_v35).length));
            let _3202_v45;
            _3202_v45 = new _dafny.CodePoint('l'.codePointAt(0));
            let _3203_v46;
            _3203_v46 = _module.D15.create_DC30(_3202_v45);
            let _3204_v47;
            _3204_v47 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_3201_v44).length)), (_3196_v40)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length))]),_3203_v46);
            _3204_v47 = (_3204_v47).update(_3154_v0, _module.D15.create_DC30(_3202_v45));
            let _index560 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length));
            let _rhs543 = !((((_3187_v31).f21) ? (false) : ((_3187_v31).f21)));
            let _rhs544 = (_3187_v31).f21;
            let _rhs545 = (_3196_v40)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length))];
            let _rhs546 = !(((_3196_v40)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length))]).isLessThan(_3154_v0));
            let _rhs547 = _module.__default.safeModuloInt(((_3196_v40)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length))]).multipliedBy((_3196_v40)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length))]), (_3193_v37)[_module.__default.safeIndex(_3154_v0, new BigNumber((_3193_v37).length))]);
            let _lhs427 = _3196_v40;
            let _lhs428 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length));
            r1 = _rhs543;
            r2 = _rhs544;
            r0 = _rhs545;
            r2 = _rhs546;
            _lhs427[_lhs428] = _rhs547;
            if ((_3198_v41)[_module.__default.safeIndex(_3154_v0, new BigNumber((_3198_v41).length))]) {
              _3196_v40 = _3196_v40;
              let _3205_v48;
              _3205_v48 = _module.D3.create_DC3((_3198_v41)[_module.__default.safeIndex((_3196_v40)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length))], new BigNumber((_3198_v41).length))]);
              let _3206_v49;
              let _nw624 = new _module.C2();
              _nw624.__ctor(_3205_v48, (_3187_v31).f21, (_3187_v31).f21);
              _3206_v49 = _nw624;
              let _rhs548 = _3202_v45;
              let _rhs549 = _3206_v49;
              _3202_v45 = _rhs548;
              _3206_v49 = _rhs549;
              let _3207_v50;
              _3207_v50 = _dafny.Map.Empty.slice().updateUnsafe((_3187_v31).f21,_3202_v45);
              let _3208_v51;
              _3208_v51 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.areEqual(_dafny.Seq.UnicodeFromString("ec"), _3191_v35),(((_3207_v50).contains((_3187_v31).f21)) ? ((_3207_v50).get((_3187_v31).f21)) : ((_3203_v46).dtor_cf44)));
              _3207_v50 = (_3207_v50);
              r1 = (_3198_v41)[_module.__default.safeIndex((_3196_v40)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length))], new BigNumber((_3198_v41).length))];
              (globalState).f7 = (((_3187_v31).f21) ? (_3154_v0) : (_3154_v0));
            } else {
              _3154_v0 = (new BigNumber(-387)).plus((_3154_v0).plus((_3196_v40)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length))]));
              let _3209_v52;
              let _nw625 = new _module.C3();
              _nw625.__ctor(_module.__default.safeModuloInt(_3154_v0, new BigNumber((_3191_v35).length)), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("dtmbyeffl"), _3191_v35), (_3187_v31).f21, p0);
              _3209_v52 = _nw625;
              let _3210_v53;
              _3210_v53 = _dafny.Map.Empty.slice().updateUnsafe((_3196_v40)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length))],(_3187_v31).f21);
              let _index561 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length));
              (_3196_v40)[_index561] = (new BigNumber((_3210_v53).length)).minus(_module.__default.safeDivisionInt(_3154_v0, (_3196_v40)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length))]));
              let _3211_v54;
              let _init110 = ((_3212_v31) => function (_3213_i2) {
                return (_3212_v31).f21;
              })(_3187_v31);
              let _nw626 = Array((new BigNumber(13)).toNumber());
              for (let _i0_110 = 0; _i0_110 < new BigNumber(_nw626.length); _i0_110++) {
                _nw626[_i0_110] = _init110(new BigNumber(_i0_110));
              }
              _3211_v54 = _nw626;
              let _index562 = _module.__default.safeIndex(new BigNumber(982), new BigNumber((_3211_v54).length));
              (_3211_v54)[_index562] = (_3187_v31).f21;
              let _3214_v55;
              _3214_v55 = _module.D12.create_DC22(_dafny.Seq.UnicodeFromString("roekfqycj"));
              let _index563 = _module.__default.safeIndex(new BigNumber(982), new BigNumber((_3211_v54).length));
              let _index564 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length));
              let _rhs550 = (_3187_v31).f21;
              let _rhs551 = (_dafny.ZERO).minus((_module.__default.safeDivisionInt((_3196_v40)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length))], new BigNumber(((_3214_v55).dtor_cf29).length))).multipliedBy((_3196_v40)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length))]));
              let _rhs552 = (_3209_v52).f22;
              let _lhs429 = _3211_v54;
              let _lhs430 = _module.__default.safeIndex(new BigNumber(982), new BigNumber((_3211_v54).length));
              let _lhs431 = _3196_v40;
              let _lhs432 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length));
              _lhs429[_lhs430] = _rhs550;
              _3154_v0 = _rhs551;
              _lhs431[_lhs432] = _rhs552;
              let _3215_v56;
              _3215_v56 = _dafny.Map.Empty.slice().updateUnsafe(_3191_v35,(_3187_v31).f21);
              _3215_v56 = (_3215_v56).update(_dafny.Seq.Concat((_3209_v52).f23, _3191_v35), (_3187_v31).f21);
            }
          }
        }
      }
      let _3216_v57;
      _3216_v57 = _module.D21.create_DC44((_dafny.ZERO).minus((_3196_v40)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length))]), (_3187_v31).f21);
      let _3217_v58;
      _3217_v58 = _dafny.MultiSet.fromElements((_3216_v57).dtor_cf69);
      let _3218_v59;
      _3218_v59 = new _dafny.CodePoint('f'.codePointAt(0));
      let _3219_v60;
      _3219_v60 = _dafny.Set.fromElements((_3196_v40)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length))], _3154_v0);
      let _3220_v61;
      let _nw627 = Array((new BigNumber(26)).toNumber());
      _nw627[(_dafny.ZERO).toNumber()] = _3193_v37;
      _nw627[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_3193_v37, _3193_v37);
      _nw627[(new BigNumber(2)).toNumber()] = _3193_v37;
      _nw627[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_3193_v37, _3193_v37);
      _nw627[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_3193_v37, _3193_v37);
      _nw627[(new BigNumber(5)).toNumber()] = _dafny.Seq.of(_3154_v0, (_3196_v40)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length))]);
      _nw627[(new BigNumber(6)).toNumber()] = _3193_v37;
      _nw627[(new BigNumber(7)).toNumber()] = _3193_v37;
      _nw627[(new BigNumber(8)).toNumber()] = _3193_v37;
      _nw627[(new BigNumber(9)).toNumber()] = _3193_v37;
      _nw627[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_3193_v37, _3193_v37);
      _nw627[(new BigNumber(11)).toNumber()] = _3193_v37;
      _nw627[(new BigNumber(12)).toNumber()] = _dafny.Seq.update(_3193_v37, _module.__default.safeIndex((_3196_v40)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length))], new BigNumber((_3193_v37).length)), (_3196_v40)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length))]);
      _nw627[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_3193_v37, _3193_v37);
      _nw627[(new BigNumber(14)).toNumber()] = _module.__default.fm35(new BigNumber((_3217_v58).cardinality()), !(p0), (_3196_v40)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length))], _3218_v59, globalState);
      _nw627[(new BigNumber(15)).toNumber()] = _3193_v37;
      _nw627[(new BigNumber(16)).toNumber()] = _3193_v37;
      _nw627[(new BigNumber(17)).toNumber()] = _3193_v37;
      _nw627[(new BigNumber(18)).toNumber()] = _3193_v37;
      _nw627[(new BigNumber(19)).toNumber()] = _3193_v37;
      _nw627[(new BigNumber(20)).toNumber()] = _3193_v37;
      _nw627[(new BigNumber(21)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(440)), ((_3221_v40) => function (_3222_i3) {
        return (_3221_v40)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_3221_v40).length))];
      })(_3196_v40));
      _nw627[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_3219_v60).length), new BigNumber(925)), _3193_v37);
      _nw627[(new BigNumber(23)).toNumber()] = _dafny.Seq.Concat(_3193_v37, _3193_v37);
      _nw627[(new BigNumber(24)).toNumber()] = _3193_v37;
      _nw627[(new BigNumber(25)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-444)), ((_3223_v0) => function (_3224_i4) {
        return _3223_v0;
      })(_3154_v0));
      _3220_v61 = _nw627;
      let _index565 = _module.__default.safeIndex(new BigNumber(70), new BigNumber((_3220_v61).length));
      (_3220_v61)[_index565] = _3193_v37;
      let _index566 = _module.__default.safeIndex(new BigNumber(70), new BigNumber((_3220_v61).length));
      (_3220_v61)[_index566] = _3193_v37;
      r1 = (_3187_v31).f21;
      r0 = (_3196_v40)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_3196_v40).length))];
      r1 = !((_3187_v31).f21) || (p0);
      r2 = !(_module.__default.fm21(globalState));
      return [r0, r1, r2];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
